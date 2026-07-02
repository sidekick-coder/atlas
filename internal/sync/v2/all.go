package sync

import (
	"github.com/sidekick-coder/atlas/internal/drive/v2"
)


func (s *Sync) All() error {

	cb := func(e drive.EntryInfo) error {
		return s.OneByInfo(&e)
	}

	err := s.drive.ScanStream(cb)

	if err != nil {
		return err
	}

	return nil
}

