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

	// if path not exists assume it was deleted and remove it from the database 

	exists, err := s.drive.Exists(path)

	if err != nil {
		return err
	}

	if !exists {
		entryRepo := entry.New(s.database)

		err = entryRepo.DeleteByPath(path)

		if err != nil {
			return err
		}

		slog.Info("Syncer: One", "path", path, "deleted", true)

		return nil
	}

	if err != nil {
		return err
	}

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

	slog.Info("Syncer: One", "path", path, "entry_id", em.ID, "metas", len(e.Metas))

	return nil
}
