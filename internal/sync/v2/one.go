package sync

import (
	"github.com/sidekick-coder/atlas/internal/metadata"
	"github.com/sidekick-coder/atlas/internal/models"
)

func (s *Sync) OneByInfo(info *models.EntryInfo) error {
	m, err := metadata.Create(info)

	if err != nil {
		return err
	}

	data, err := m.ExtractMap()

	if err != nil {
		return err
	}

	entry, err := s.entryRepo.Upsert(info.Path)

	if err != nil {
		return err
	}

	err = s.entryMetaRepo.DeleteByEntryID(entry.ID)

	if err != nil {
		return err
	}

	_, err = s.entryMetaRepo.UpsertMany(entry.ID, data)


	if err != nil {
		return err
	}

	return nil
}

func (s *Sync) One(filepath string) error {
	info, err := s.drive.Get(filepath)

	if err != nil {
		return err
	}

	return s.OneByInfo(info)
}
