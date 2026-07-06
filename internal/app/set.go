package app

import (
	"fmt"

	"github.com/sidekick-coder/atlas/internal/metadata"
)

func (a *App) SetEntryMeta(path string, name string, value string) error {
	entry, err := a.entryRepo.GetByPath(path)

	if err != nil {
		return err
	}

	info, err := a.drive.Get(entry.Path)

	if err != nil {
		return err
	}

	handlers := metadata.GetHandlers(info)

	success, err := metadata.Set(info, name, value, handlers)

	if err != nil {
		return err
	}

	if !success {
		return fmt.Errorf("could not set value: %s", name)
	}

	err = a.syncer.One(path)

	if err != nil {
		return err
	}

	return nil
}
