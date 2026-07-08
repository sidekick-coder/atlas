package scanner

import (
	"sync"
	"sync/atomic"

	"github.com/sidekick-coder/atlas/internal/drive/v2"
	"github.com/sidekick-coder/atlas/internal/models"
)

type Worker struct {
	drive *drive.Drive
	count atomic.Int32
}

func Create() *Worker {
	return &Worker{
		count: atomic.Int32{},
	}
}

func (w *Worker) GetCount() int32 {
	return w.count.Load()
}

func (w *Worker) SetDrive(d *drive.Drive) *Worker {
	w.drive = d
	return w
}

func (w *Worker) Process(out chan<- models.EntryInfo) {
	w.drive.ScanStream(func(e models.EntryInfo) error {
		w.count.Add(1)

		out <- e

		return nil
	})
}

func (w *Worker) Run(out chan<- models.EntryInfo, concurrency int) {
	var wg sync.WaitGroup

	for range concurrency {
		wg.Go(func() {
			w.Process(out)
		})
	}

	wg.Wait()
}
