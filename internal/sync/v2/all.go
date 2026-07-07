package sync

import (
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/sidekick-coder/atlas/internal/metadata"
	"github.com/sidekick-coder/atlas/internal/models"
)

type SyncBatchEntry struct {
	ID    int
	Path  string
	Metas map[string]string
}

type SyncBatch struct {
	ID       int
	Executed bool
	Entries  map[int]SyncBatchEntry
}

type AllResult struct {
	Concurrency        int
	TotalEntries       int
	TotalBatches       int
	TotalEntriesErrors int
	TotalBatchesErrors int
	Time               time.Duration
}

type AllPayload struct {
	Concurrency int
	OnError     func(path string, err error)
	OnSuccess   func(path string, metas map[string]string)
	OnComplete  func(result AllResult)
}

func (s *Sync) AllWorkerExtract(e models.EntryInfo, bache SyncBatch, bacheCount int) error {
	m, err := metadata.Create(&e)

	if err != nil {
		return err
	}

	id := int(s.NextID.Add(1))

	metas, err := m.ExtractMap()

	be := SyncBatchEntry{
		ID:    id,
		Path:  e.Path,
		Metas: metas,
	}

	bache.Entries[id] = be

	return nil
}

func (s *Sync) AllExecuteBatch(batch SyncBatch) error {
	tx, err := s.Database.Connection.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	smtmt := []string{}
	params := []any{}
	values := []string{}

	smtmt = append(smtmt, "INSERT INTO entries (id, path) VALUES")

	for id, be := range batch.Entries {
		values = append(values, "(?, ?)")
		params = append(params, id, be.Path)
	}

	smtmt = append(smtmt, strings.Join(values, ",\n"))

	_, err = tx.Exec(strings.Join(smtmt, " "), params...)

	if err != nil {
		tx.Rollback()
		return err
	}

	smtmt = []string{}
	params = []any{}
	values = []string{}

	smtmt = append(smtmt, "INSERT INTO entry_metas (entry_id, name, value) VALUES")

	for _, be := range batch.Entries {
		for k, v := range be.Metas {
			values = append(values, "(?, ?, ?)")
			params = append(params, be.ID, k, v)
		}
	}

	smtmt = append(smtmt, strings.Join(values, ",\n"))
	smtmt = append(smtmt, ";")

	_, err = tx.Exec(strings.Join(smtmt, " "), params...)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func OnBatchError(batch SyncBatch, err error, p AllPayload) {
	for _, be := range batch.Entries {
		p.OnError(be.Path, err)
	}
}

func OnBatchSuccess(batch SyncBatch, p AllPayload) {
	for _, be := range batch.Entries {
		p.OnSuccess(be.Path, be.Metas)
	}
}

func (s *Sync) AllWorker(wg *sync.WaitGroup, jobs <-chan models.EntryInfo, batches []SyncBatch, p AllPayload, batchSize int) {
	defer wg.Done()

	if len(batches) == 0 {
		s.TotalBatches.Add(1)
		batches = append(batches, SyncBatch{
			ID:      0,
			Entries: make(map[int]SyncBatchEntry),
		})
	}

	batchId := 0

	for e := range jobs {
		s.TotalEntries.Add(1)
		batch := batches[batchId]

		err := s.AllWorkerExtract(e, batch, batchId)

		if err != nil {
			p.OnError(e.Path, err)
			s.TotalEntriesErrors.Add(1)
			continue
		}

		if len(batch.Entries) >= batchSize {
			err := s.AllExecuteBatch(batch)

			batch.Executed = true

			if err != nil {
				OnBatchError(batch, err, p)
				s.TotalBatchesErrors.Add(1)
			}

			if err == nil {
				OnBatchSuccess(batch, p)
			}

			batchId++

			batch = SyncBatch{
				ID:      batchId,
				Entries: make(map[int]SyncBatchEntry),
			}

			batches = append(batches, batch)
			s.TotalBatches.Add(1)
		}
	}
}

func (s *Sync) AllCleanup() error {
	err := s.entryMetaRepo.DeleteAll()

	if err != nil {
		return err
	}

	err = s.entryRepo.DeleteAll()

	if err != nil {
		return err
	}

	return nil
}

func (s *Sync) All(payload ...AllPayload) (AllResult, error) {
	s.TotalEntries = atomic.Int64{}
	s.TotalEntriesErrors = atomic.Int64{}
	s.TotalBatches = atomic.Int64{}
	s.TotalBatchesErrors = atomic.Int64{}

	p := AllPayload{}

	if len(payload) > 0 {
		p = payload[0]
	}

	if p.OnComplete == nil {
		p.OnComplete = func(result AllResult) {}
	}

	if p.OnError == nil {
		p.OnError = func(e string, err error) {}
	}

	if p.OnSuccess == nil {
		p.OnSuccess = func(path string, metas map[string]string) {}
	}

	concrrency := 1
	batchSize := 100

	if p.Concurrency > 0 {
		concrrency = p.Concurrency
	}

	jobs := make(chan models.EntryInfo, concrrency)
	batch := []SyncBatch{}

	result := AllResult{
		Concurrency:  concrrency,
		TotalEntries: 0,
		TotalBatches: 0,
		Time:         0,
	}

	var wg sync.WaitGroup

	for i := 0; i < concrrency; i++ {
		wg.Add(1)

		go s.AllWorker(&wg, jobs, batch, p, batchSize)
	}

	start := time.Now()

	err := s.AllCleanup()

	if err != nil {
		return AllResult{}, err
	}

	err = s.drive.ScanStream(func(e models.EntryInfo) error {
		jobs <- e
		return nil
	})

	close(jobs)

	wg.Wait()

	for _, b := range batch {
		result.TotalEntries += len(b.Entries)

		if !b.Executed {
			err := s.AllExecuteBatch(b)

			if err != nil {
				OnBatchError(b, err, p)
			}

			if err == nil {
				OnBatchSuccess(b, p)
			}
		}
	}

	result.Time = time.Since(start)
	result.TotalBatches = int(s.TotalBatches.Load())
	result.TotalEntriesErrors = int(s.TotalEntriesErrors.Load())
	result.TotalEntries = int(s.TotalEntries.Load())
	result.TotalBatchesErrors = int(s.TotalBatchesErrors.Load())

	if err != nil {
		return result, err
	}

	p.OnComplete(result)

	return result, nil
}
