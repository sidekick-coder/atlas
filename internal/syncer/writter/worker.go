package writter

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/sidekick-coder/atlas/internal/database"
	"github.com/sidekick-coder/atlas/internal/syncer/batcher"
	"github.com/sidekick-coder/atlas/internal/syncer/extractor"
)

type Worker struct {
	database  *database.Database
	processed atomic.Int32
	count	 atomic.Int32
	onSuccess func(e extractor.ExtractEntry)
	onBatchComplete func(batch batcher.Batch)
	onError   func(e extractor.ExtractEntry, err error)
}

func Create() *Worker {
	return &Worker{
		count:    atomic.Int32{},
	}
}

func (w *Worker) GetCount() int32 {
	return w.count.Load()
}

func (w *Worker) SetDatabase(db *database.Database) *Worker {
	w.database = db
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

func (w *Worker) OnBatchComplete(cb func(batch batcher.Batch)) *Worker {
	w.onBatchComplete = cb
	return w
}

func (w *Worker) GetProcessedCount() int32 {
	return w.processed.Load()
}

func (w *Worker) Execute(batch batcher.Batch) error {
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

func (w *Worker) emitSuccess(batch batcher.Batch) {
	if w.onSuccess != nil {
		for _, e := range batch.Entries {
			w.processed.Add(1)
			w.onSuccess(e)
		}
	}
}

func (w *Worker) emitError(batch batcher.Batch, err error) {
	if w.onError != nil {
		for _, e := range batch.Entries {
			w.onError(e, err)
			w.processed.Add(1)
		}
	}
}

func (w *Worker) Process(in <-chan batcher.Batch) {
	for batch := range in {
		err := w.Execute(batch)

		if w.onBatchComplete != nil {
			w.onBatchComplete(batch)
		}

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
			fmt.Printf("Error processing batch %d: %v\n", batch.ID, err)
			continue
		}
	}
}

func (w *Worker) Run(in <-chan batcher.Batch, concurrency int) {
	var wg sync.WaitGroup

	for range concurrency {
		wg.Go(func() {
			w.Process((in))
		})
	}

	wg.Wait()
}
