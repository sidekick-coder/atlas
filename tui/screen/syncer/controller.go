package syncer

import (
	"time"

	"github.com/sidekick-coder/atlas/internal/syncer"
)

func (s *Screen) Sync() {
	s.Running = true
	s.Completed = false
	s.Entries = []Entry{}

	go func() {
		startTime := time.Now()

		onSuccess := func(path string) {
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

		onComplete := func(result syncer.Result) {
			s.Running = false
			s.Time = time.Since(startTime)

			s.Program.Send(Completed{
				TotalEntries: int(result.Written),
				Time:         s.Time,
			})
		}

		s.Running = true

		s := s.App.Syncer()

		s.OnComplete(onComplete).OnSuccess(onSuccess).OnError(onError) 

		s.All()

	}()
}
