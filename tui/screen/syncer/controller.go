package syncer

import (
	"time"

	"github.com/sidekick-coder/atlas/internal/sync/v2"
)

func (s *Screen) Sync() {
	s.Running = true
	s.Completed = false
	s.Entries = []Entry{}

	go func() {
		startTime := time.Now()

		onSuccess := func(path string, metas map[string]string) {
			if !s.ViewList {
				return
			}

			e := Entry{
				Path:    path,
				Success: true,
				Error:   nil,
			}

			s.Program.Send(EntryAdd{AddEntry: e})

			s.Time = time.Since(startTime)
		}

		onError := func(path string, err error) {
			if !s.ViewList {
				return
			}

			e := Entry{
				Path:    path,
				Success: false,
				Error:   err,
			}

			s.Program.Send(EntryAdd{AddEntry: e})
			s.Time = time.Since(startTime)
		}

		onComplete := func(result sync.AllResult) {
			s.Running = false
			s.Time = time.Since(startTime)

			s.Program.Send(Completed{
				TotalEntries: result.TotalEntries,
				Time:         s.Time,
			})
		}

		s.Running = true

		s.Syncer.All(sync.AllPayload{
			Concurrency: 1,
			OnError:     onError,
			OnSuccess:   onSuccess,
			OnComplete:  onComplete,
		})
	}()
}
