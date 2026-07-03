package sync

import (
	"github.com/sidekick-coder/atlas/internal/models"
)


func (s *Sync) All() error {

	cb := func(e models.EntryInfo) error {
		return s.OneByInfo(&e)
	}

	err := s.drive.ScanStream(cb)

	if err != nil {
		return err
	}

	return nil
}

