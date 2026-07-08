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

	m, err := metadata.Create(info)

	if err != nil {
		return err
	}

	err = m.SetHandlersFromConfig(a.config)

	if err != nil {
		return err
	}

	success, err := m.Set(name, value)

	if err != nil {
		return err
	}

	if !success {
		return fmt.Errorf("no handler to set set value: %s", name)
	}

	err = a.syncer.One(path)

	if err != nil {
		return err
	}

	return nil
}

func (a *App) UnsetEntryMeta(path string, name string) error {
	entry, err := a.entryRepo.GetByPath(path)

	if err != nil {
		return err
	}

	info, err := a.drive.Get(entry.Path)

	if err != nil {
		return err
	}

	m, err := metadata.Create(info)

	if err != nil {
		return err
	}

	err = m.SetHandlersFromConfig(a.config)

	if err != nil {
		return err
	}

	success, err := m.Unset(name)

	if err != nil {
		return err
	}

	if !success {
		return fmt.Errorf("could not unset meta: %s", name)
	}

	err = a.syncer.One(path)

	if err != nil {
		return err
	}

	return nil
}
