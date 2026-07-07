package syncer

import (
	"fmt"

	"github.com/sidekick-coder/atlas/internal/sync/v2"
	"github.com/sidekick-coder/atlas/tui/messages"
)

func (s *Screen) Sync() {
	s.Running = true
	s.Entries = []Entry{}

	go func() {
		onSuccess := func(path string, metas map[string]string) {
			e := Entry{
				Path:    path,
				Success: true,
				Error:   nil,
			}

			s.Program.Send(EntryAdd{AddEntry: e})
		}

		onError := func(path string, err error) {
			e := Entry{
				Path:    path,
				Success: false,
				Error:   err,
			}

			s.Program.Send(EntryAdd{AddEntry: e})
		}

		onComplete := func(result sync.AllResult) {
			s.Running = false

			message := fmt.Sprintf("Sync completed: %d entries, %d batches, %d errors, %.2fs", result.TotalEntries, result.TotalBatches, result.TotalEntriesErrors, result.Time.Seconds()*1000)

			s.Program.Send(messages.ToastSuccessMessage(message))
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
