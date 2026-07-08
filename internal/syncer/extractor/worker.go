package extractor

import (
	"sync"
	"sync/atomic"

	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/metadata"
	"github.com/sidekick-coder/atlas/internal/models"
)

type Worker struct {
	config *config.Config
	nextId atomic.Int32
	count  atomic.Int32
	onError func(e models.EntryInfo, err error)
	onSuccess func(e ExtractEntry)
	onComplete func(count int32)
}

type ExtractEntry struct {
	ID    int
	Path  string
	Metas map[string]string
}

func Create() *Worker {
	return &Worker{
		nextId: atomic.Int32{},
	}
}

func (w *Worker) OnError(cb func(e models.EntryInfo, err error)) *Worker {
	w.onError = cb
	return w
}

func (w *Worker) OnSuccess(cb func(e ExtractEntry)) *Worker {
	w.onSuccess = cb
	return w
}

func (w *Worker) OnComplete(cb func(count int32)) *Worker {
	w.onComplete = cb
	return w
}

func (w *Worker) GetCount() int32 {
	return w.count.Load()
}

func (w *Worker) SetConfig(c *config.Config) *Worker {
	w.config = c
	return w
}

func (w *Worker) Extract(e models.EntryInfo) (ExtractEntry, error) {
	m, err := metadata.Create(&e)

	if err != nil {
		return ExtractEntry{}, err
	}

	err = m.SetHandlersFromConfig(w.config)

	if err != nil {
		return ExtractEntry{}, err
	}

	id := int(w.nextId.Add(1))

	metas, err := m.ExtractMap()

	if err != nil {
		return ExtractEntry{}, err
	}

	ee := ExtractEntry{
		ID:    id,
		Path:  e.Path,
		Metas: metas,
	}

	return ee, nil
}

func (w *Worker) Proccess(in <-chan models.EntryInfo, out chan<- ExtractEntry) {

	for e := range in {
		w.count.Add(1)

		ee, err := w.Extract(e)

		if err != nil && w.onError != nil {
			w.onError(e, err)
		}

		if err == nil && w.onSuccess != nil {
			w.onSuccess(ee)
		}

		if err != nil {
			continue
		}

		out <- ee
	}

	close(out)

	if (w.onComplete != nil) {
		w.onComplete(w.count.Load())
	}
}

func (w *Worker) Run(in <-chan models.EntryInfo, out chan<- ExtractEntry, concurrency int) {
	var wg sync.WaitGroup

	for range concurrency {
		wg.Go(func() {
			w.Proccess(in, out)
		})
	}

	wg.Wait()

}
