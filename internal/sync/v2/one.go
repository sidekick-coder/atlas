package sync

import (
	"github.com/sidekick-coder/atlas/internal/metadata"
	"github.com/sidekick-coder/atlas/internal/drive/v2"
)


func (s *Sync) OneByInfo(info * drive.EntryInfo) error {
	handlers := metadata.GetHandlers(info)

	data, err := metadata.Extract(info, handlers)

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


	return nil
}

func (s *Sync) One(filepath string) error {
	info, err := s.drive.Get(filepath)

	if err != nil {
		return err
	}

	return s.OneByInfo(info)
}

