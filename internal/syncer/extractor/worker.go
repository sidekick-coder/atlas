package extractor

import (
	"fmt"
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

	m.SetHandlersFromConfig(w.config)

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

		if err != nil {
			fmt.Println("error extracting metadata for path:", e.Path, "error:", err)
		}

		out <- ee

	}

	close(out)
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
