package entrysingle

import (
	"fmt"
	"slices"
	"strings"

	"github.com/sidekick-coder/atlas/internal/models"
)

func (s *Screen) Load() error {
	repo := s.App.EntryMetaRepo()

	metas, err := repo.ListByEntryPath(s.Path)

	if err != nil {
		return fmt.Errorf("failed to load metadata for entry %s: %w", s.Path, err)
	}

	// sort 
	slices.SortFunc(metas, func(a, b models.EntryMeta) int {
		if len(a.Name) != len(b.Name) {
			return len(a.Name) - len(b.Name)
		}

		return strings.Compare(a.Name, b.Name)
	})

	s.EntryMetaComponent.SetMetas(metas)

	return nil
}

func (s *Screen) UnsetMeta(name string) error {
	err := s.App.UnsetEntryMeta(s.Path, name)

	if err != nil {
		return err
	}

	return s.Load()	
}

func (s *Screen) UnsetMetaSelected() error {
	em, ok := s.EntryMetaComponent.GetSelected()

	if !ok {
		return fmt.Errorf("no metadata selected to set value")
	}

	return s.UnsetMeta(em.Name)
}

func (s *Screen) SetMeta(name string, value string) error {
	err := s.App.SetEntryMeta(s.Path, name, value)

	if err != nil {
		return err
	}

	return s.Load()	
}

func (s *Screen) SetValue(value string) error {
	em, ok := s.EntryMetaComponent.GetSelected()

	if !ok {
		return fmt.Errorf("no metadata selected to set value")
	}

	return s.SetMeta(em.Name, value)
}

func (s *Screen) Sync() error {
	err := s.App.Syncer().One(s.Path)

	if err != nil {
		return err
	}

	return s.Load()
}
