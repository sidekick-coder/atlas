package batcher

import (
	"sync"
	"sync/atomic"

	"github.com/sidekick-coder/atlas/internal/syncer/extractor"
)

type Worker struct {
	size   int
	count  atomic.Int32
	nextID atomic.Int32
}

type Batch struct {
	ID      int32
	Entries []extractor.ExtractEntry
}

func Create() *Worker {
	return &Worker{
		size:   100,
		nextID: atomic.Int32{},
	}
}

func (w *Worker) GetCount() int32 {
	return w.count.Load()
}

func (w *Worker) SetBatchSize(size int) *Worker {
	w.size = size
	return w
}

func (w *Worker) GetBatchSize() int {
	return w.size
}

func (w *Worker) Procces(in <-chan extractor.ExtractEntry, out chan<- Batch) {
	batch := Batch{
		ID:      w.nextID.Add(1),
		Entries: make([]extractor.ExtractEntry, 0, w.size),
	}

	for e := range in {

		batch.Entries = append(batch.Entries, e)

		if len(batch.Entries) >= w.size {
			out <- batch

			w.count.Add(1)

			batch = Batch{
				ID:      w.nextID.Add(1),
				Entries: make([]extractor.ExtractEntry, 0, w.size),
			}
		}
	}

	if len(batch.Entries) > 0 {
		out <- batch
		w.count.Add(1)
	}

}

func (w *Worker) Run(in <-chan extractor.ExtractEntry, out chan<- Batch, concurrency int) {
	var wg sync.WaitGroup

	for range concurrency {
		wg.Go(func() {
			w.Procces(in, out)
		})
	}

	wg.Wait()
	close(out)
}
