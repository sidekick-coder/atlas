package entry

import (
	"github.com/sidekick-coder/atlas/internal/repository/entry"
)

func (s *Screen) Load() error {
	repo := s.App.EntryRepo()

	offset := (s.CurrentPage - 1) * s.Limit

	options := entry.ListOptions{
		Limit: s.Limit,
		Offset: offset,
	}

	if s.Query != "" {
		options.Query = []string{s.Query}
	}

	entries, err := repo.List(options)

	if err != nil {
		return err
	}

	count, err := repo.Count()

	if err != nil {
		return err
	}

	s.Count = count
	s.TotalPages = (count + s.Limit - 1) / s.Limit
	s.CurrentPage = (offset / s.Limit) + 1

	s.List.SetEntries(entries)

	return nil
}

func (s *Screen) NextPage() error {
	nextPage := s.CurrentPage + 1 

	if nextPage > s.TotalPages {
		return nil
	}

	s.CurrentPage = nextPage

	return s.Load()
}

func (s *Screen) PreviousPage() error {

	page := s.CurrentPage - 1

	if page < 1 {
		return nil
	}

	s.CurrentPage = page

	return s.Load()
}

func (s *Screen) Reset() error {
	s.CurrentPage = 1
	return s.Load()
}

