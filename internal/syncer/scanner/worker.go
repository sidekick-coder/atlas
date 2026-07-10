package scanner

import (
	"fmt"
	// "sync"
	"sync/atomic"

	"github.com/sidekick-coder/atlas/internal/drive"
	"github.com/sidekick-coder/atlas/internal/models"
)

type Worker struct {
	drive      *drive.Drive
	count      atomic.Int32
	onComplete func(count int32)
}

func Create() *Worker {
	return &Worker{
		count: atomic.Int32{},
	}
}

func (w *Worker) OnComplete(cb func(count int32)) *Worker {
	w.onComplete = cb
	return w
}

func (w *Worker) GetCount() int32 {
	return w.count.Load()
}

func (w *Worker) SetDrive(d *drive.Drive) *Worker {
	w.drive = d
	return w
}

func (w *Worker) Process(out chan<- models.EntryInfo) {
	entries, err := w.drive.Scan()

	if err != nil {
		fmt.Println("error scanning drive:", err)
	}

	for _, e := range entries {
		out <- e
	}

	w.count.Store(int32(len(entries)))

	if w.onComplete != nil {
		w.onComplete(w.count.Load())
	}
}

func (w *Worker) Run(out chan<- models.EntryInfo, concurrency int) {
	// var wg sync.WaitGroup
	//
	// for range concurrency {
	// 	wg.Go(func() {
	// 		w.Process(out)
	// 	})
	// }
	//
	// wg.Wait()
	// close(out)
	defer close(out)
	w.Process(out)
}
