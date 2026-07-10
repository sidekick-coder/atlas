package entryloader

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/internal/repository/entry"
	"github.com/sidekick-coder/atlas/tui/messages"
)

type Feature struct {
	repository entry.Repository

	limit  int
	offset int
	count  int
	query  []string

	entries []models.Entry
}

func Create(repository entry.Repository) *Feature {
	return &Feature{
		repository: repository,
		limit:      10,
		offset:     0,
		query:      []string{},
		entries:    []models.Entry{},
	}
}

func (f *Feature) GetEntries() []models.Entry {
	return f.entries
}

func (f *Feature) Load() error {
	options := entry.ListOptions{
		Limit:  f.limit,
		Offset: f.offset,
		Query:  f.query,
		LoadMetas: true,
	}

	entries, err := f.repository.List(options)

	if err != nil {
		return err
	}

	count, err := f.repository.Count()

	if err != nil {
		return err
	}

	f.count = count
	f.entries = entries

	return nil
}

func (f *Feature) Init() tea.Cmd {
	err := f.Load()

	if (err != nil) {
		return messages.ToastErrorCmd(err.Error())
	}

	return  nil
}
