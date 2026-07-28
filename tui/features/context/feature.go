package context

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/utils"
)

type Feature struct {
	id     string
	label  string
	parent []string
	active bool

	entries []Entry
}

type Entry struct {
	Key   string
	Value any
}

type GetEntryOptions struct {
	ExcludeParent bool
	ExcludeGlobal bool
}

func Create() *Feature {
	id, err := utils.CreateID()

	if err != nil {
		panic(err)
	}

	f := &Feature{
		id:      id,
		label:   "",
		parent:  []string{},
		entries: []Entry{},
		active:  false,
	}

	return f
}

func (f *Feature) SetParent(ctx *Feature) {
	f.parent = append(f.parent, ctx.id)
}

func (f *Feature) SetLabel(label string) {
	f.label = label
}

func (f *Feature) GetLabel() string {
	return f.label
}

func (f *Feature) GetID() string {
	return f.id
}

func (f *Feature) SetID(id string) {
	f.id = id
}

func (f *Feature) GetParent() []string {
	return f.parent
}

func (f *Feature) IsActive() bool {
	return f.active
}

func (f *Feature) SetActive(active bool) {
	f.active = active
}

func (f *Feature) Activate() tea.Cmd {
	f.active = true

	return nil
}

func (f *Feature) Deactivate() tea.Cmd {
	f.active = false

	return nil
}

func (f *Feature) Unset(key string) {
	for i, e := range f.entries {
		if e.Key == key {
			f.entries = append(f.entries[:i], f.entries[i+1:]...)
			break
		}
	}
}

func (f *Feature) Set(key string, value any) {
	f.Unset(key)

	f.entries = append(f.entries, Entry{
		Key:   key,
		Value: value,
	})
}

func (f *Feature) SetAll(entries map[string]any) {
	for k, v := range entries {
		f.Set(k, v)
	}
}

func (f *Feature) GetParentEntries() []Entry {
	entries := []Entry{}

	ids := f.parent

	for _, id := range ids {
		c, ok := GetById(id)
		pe := []Entry{}

		if ok {
			pe = c.GetEntries(GetEntryOptions{
				ExcludeParent: false,
				ExcludeGlobal: true,
			})
		}

		for _, e := range pe {
			key := fmt.Sprintf("%s.%s", c.label, e.Key)

			e := Entry{
				Key:   key,
				Value: e.Value,
			}

			entries = append(entries, e)
		}

	}

	return entries
}

func (f *Feature) GetGlobalEntries() []Entry {
	if f.id == "global" {
		return []Entry{}
	}

	entries := []Entry{}

	ctx, ok := GetById("global")

	if !ok {
		return entries
	}

	return ctx.GetEntries()
}

func (f *Feature) GetEntries(args ...GetEntryOptions) []Entry {
	options := GetEntryOptions{}

	if len(args) > 0 {
		options = args[0]
	}

	entries := f.entries

	if !options.ExcludeParent {
		entries = append(entries, f.GetParentEntries()...)
	}

	if !options.ExcludeGlobal {
		entries = append(entries, f.GetGlobalEntries()...)
	}

	return entries
}

func (f *Feature) GetEntriesFlat() [][]string {
	result := [][]string{}

	flat := utils.FlattenMap(f.GetEntriesMap(), "")

	for k, v := range flat {
		value := ""
		vs, ok := v.(string)

		if ok {
			value = vs
		}

		if !ok {
			value = fmt.Sprintf("%T", v)
		}

		result = append(result, []string{k, value})
	}

	return result
}

func (f *Feature) GetEntriesMap() map[string]any {
	result := map[string]any{}

	for _, e := range f.GetEntries() {
		result[e.Key] = e.Value
	}

	return result
}

func (f *Feature) Init() tea.Cmd {
	registry = append(registry, f)

	return nil
}

func (f *Feature) Update(msg tea.Msg) tea.Cmd {
	return nil
}

func (f *Feature) Dispose() tea.Cmd {
	for i, v := range registry {
		if v.id == f.id {
			registry = append(registry[:i], registry[i+1:]...)
			break
		}
	}

	return nil
}
