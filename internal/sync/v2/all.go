package sync

import (
	"sync"
	"time"

	"github.com/sidekick-coder/atlas/internal/models"
)

type AllResult struct {
	Microseconds int64
	Concurrency   int
}

type AllPayload struct {
	Concurrency   int
	OnError func(e models.EntryInfo, err error)
	OnSuccess func(e models.EntryInfo)
	OnComplete func(result AllResult)
}


func (s *Sync) All(payload ...AllPayload) (AllResult, error) {
	p := AllPayload{}

	if len(payload) > 0 {
		p = payload[0]
	}

	concrrency := 1

	if p.Concurrency > 0 {
		concrrency = p.Concurrency
	}

	jobs := make(chan models.EntryInfo, concrrency)
	var wg sync.WaitGroup

	for i := 0; i < concrrency; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for e := range jobs {
				err := s.OneByInfo(&e)

				if err != nil && p.OnError != nil {
					p.OnError(e, err)
				}

				if err == nil && p.OnSuccess != nil {
					p.OnSuccess(e)
				}
			}
		}()
	}

	start := time.Now()

	err := s.drive.ScanStream(func(e models.EntryInfo) error {
		jobs <- e
		return nil
	})

	close(jobs)
	wg.Wait()

	result := AllResult{
		Microseconds: time.Since(start).Microseconds(),
		Concurrency:  concrrency,
	}

	if p.OnComplete != nil {
		p.OnComplete(result)
	}

	if err != nil {
		return result, err
	}

	return result, nil
}

