package syncer

import (
	"github.com/sidekick-coder/atlas/internal/syncer/batcher"
	"github.com/sidekick-coder/atlas/internal/syncer/extractor"
)

func (s *Syncer) One(path string) error {
	s.scanner.SetDrive(s.drive)
	s.extractor.SetConfig(s.config)
	s.writter.SetDatabase(s.database)

	i, err := s.drive.Get(path)

	if err != nil {
		return err
	}

	e, err := s.extractor.Extract(*i)

	if err != nil {
		return err
	}

	batch := batcher.Batch{
		ID:      1,
		Entries: []extractor.ExtractEntry{e},
	}

	err = s.writter.Execute(batch)

	if err != nil {
		return err 
	}

	return nil
}
