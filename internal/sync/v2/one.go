package sync

import (
	"fmt"

	// "github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/internal/metadata"
	// "github.com/sidekick-coder/atlas/internal/drive/v2"
)


func (s *Sync) One(filepath string) error {
	info, err := s.drive.Get(filepath)

	if err != nil {
		return err
	}

	handlers := metadata.GetHandlers(info)

	data, err := metadata.Extract(info, handlers)

	if err != nil {
		return err 
	}

	entry, err := s.entryRepo.Upsert(info.Path)

	if err != nil {
		return err
	}

	metas, err := s.entryMetaRepo.UpsertMany(entry.ID, data)

	fmt.Println(data)
	fmt.Println(entry)
	fmt.Println(metas)

	return nil
}

