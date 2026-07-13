package syncer

import (
	"log/slog"

	"github.com/sidekick-coder/atlas/internal/repository/entry"
	"github.com/sidekick-coder/atlas/internal/repository/entrymeta"
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

	entryRepo := entry.New(s.database)

	em, err := entryRepo.Upsert(e.Path)

	if err != nil {
		return err
	}

	repo := entrymeta.New(s.database)

	err = repo.DeleteByEntryID(em.ID)

	if err != nil {
		return err
	}

	err = repo.InsertMany(em.ID, e.Metas)

	if err != nil {
		return err
	}

	slog.Info("Syncer: One", "path", path, "entry", e)

	return nil
}
