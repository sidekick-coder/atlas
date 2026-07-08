package writter

import (
	"strings"
	"sync"
	"sync/atomic"

	"github.com/sidekick-coder/atlas/internal/database"
	"github.com/sidekick-coder/atlas/internal/syncer/extractor"
)

type Worker struct {
	database  *database.Database
	processed atomic.Int32
	batchSize int
	count	 atomic.Int32
	nextID    atomic.Int32
	onSuccess func(e extractor.ExtractEntry)
	onError   func(e extractor.ExtractEntry, err error)
}

type Batch struct {
	ID      int32
	Entries []extractor.ExtractEntry
}

func Create() *Worker {
	return &Worker{
		batchSize: 10,
		nextID:    atomic.Int32{},
	}
}

func (w *Worker) GetCount() int32 {
	return w.count.Load()
}

func (w *Worker) SetDatabase(db *database.Database) *Worker {
	w.database = db
	return w
}

func (w *Worker) SetBatchSize(size int) *Worker {
	w.batchSize = size
	return w
}

func (w *Worker) OnSuccess(cb func(e extractor.ExtractEntry)) *Worker {
	w.onSuccess = cb
	return w
}

func (w *Worker) OnSuccessPath(cb func(p string)) *Worker {
	w.onSuccess = func(e extractor.ExtractEntry) {
		cb(e.Path)
	}
	return w
}

func (w *Worker) OnError(cb func(e extractor.ExtractEntry, err error)) *Worker {
	w.onError = cb
	return w
}

func (w *Worker) OnErrorPath(cb func(p string, err error)) *Worker {
	w.onError = func(e extractor.ExtractEntry, err error) {
		cb(e.Path, err)
	}
	return w
}

func (w *Worker) GetProcessedCount() int32 {
	return w.processed.Load()
}

func (w *Worker) Batcher(in <-chan extractor.ExtractEntry, out chan<- Batch) {
	id := w.nextID.Add(1)

	batch := Batch{
		ID:      id,
		Entries: make([]extractor.ExtractEntry, 0, w.batchSize),
	}

	for e := range in {
		batch.Entries = append(batch.Entries, e)

		if len(batch.Entries) >= w.batchSize {
			out <- batch

			id := w.nextID.Add(1)

			batch = Batch{
				ID:      id,
				Entries: make([]extractor.ExtractEntry, 0, w.batchSize),
			}
		}
	}

	if len(batch.Entries) > 0 {
		out <- batch
	}

}

func (w *Worker) BatcherRun(in <-chan extractor.ExtractEntry, out chan<- Batch, concurrency int) {
	var wg sync.WaitGroup

	for range concurrency {
		wg.Go(func() {
			w.Batcher(in, out)
		})
	}

	wg.Wait()
}

func (w *Worker) ExecuteBatch(batch Batch) error {
	tx, err := w.database.Connection.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	smtmt := []string{}
	params := []any{}
	values := []string{}

	smtmt = append(smtmt, "INSERT INTO entries (id, path) VALUES")

	for _, be := range batch.Entries {
		values = append(values, "(?, ?)")
		params = append(params, be.ID, be.Path)
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

func (w *Worker) emitSuccess(batch Batch) {
	if w.onSuccess != nil {
		for _, e := range batch.Entries {
			w.processed.Add(1)
			w.onSuccess(e)
		}
	}
}

func (w *Worker) emitError(batch Batch, err error) {
	if w.onError != nil {
		for _, e := range batch.Entries {
			w.onError(e, err)
			w.processed.Add(1)
		}
	}
}

func (w *Worker) Process(in <-chan Batch) {
	for batch := range in {
		err := w.ExecuteBatch(batch)

		if w.onSuccess != nil && err == nil {
			w.emitSuccess(batch)
		}

		if w.onError != nil && err != nil {
			w.emitError(batch, err)
		}

		for range batch.Entries {
			w.count.Add(1)
		}

		if err != nil {
			continue
		}

	}
}

func (w *Worker) Run(in <-chan Batch, concurrency int) {
	var wg sync.WaitGroup

	for range concurrency {
		wg.Go(func() {
			w.Process((in))
		})
	}

	wg.Wait()
}
