package entryloader

import (
	"fmt"

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

func (f *Feature) GetEntry(index int) (models.Entry, error) {
	if index < 0 || index >= len(f.entries) {
		return models.Entry{}, fmt.Errorf("index %d out of bounds for entries (length: %d)", index, len(f.entries))
	}

	return f.entries[index], nil
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
		return fmt.Errorf("failed to list entries: %w", err)
	}

	count, err := f.repository.Count(entry.CountOptions{
		Query: f.query,
	})

	if err != nil {
		return fmt.Errorf("failed to count entries: %w", err)
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

func (f *Feature) SetLimit(limit int) {
	f.limit = limit
}

func (f *Feature) SetOffset(offset int) {
	f.offset = offset
}

func (f *Feature) SetQuery(query []string) {
	f.query = query 
}

func (f *Feature) GetQuery() []string {
	return f.query
}

func (f *Feature) GetCount() int {
	return f.count
}

func (f *Feature) GetLimit() int {
	return f.limit
}

func (f *Feature) GetOffset() int {
	return f.offset
}

func (f *Feature) Next() {
	if f.offset + f.limit < f.count {
		f.offset += f.limit
	}
}

func (f *Feature) Prev() {
	if f.offset - f.limit >= 0 {
		f.offset -= f.limit
	} else {
		f.offset = 0
	}
}
