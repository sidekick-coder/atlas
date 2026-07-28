package contextdialog

import (
	"slices"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/dialog"
	"github.com/sidekick-coder/atlas/tui/components/keyvalue"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/context"
)

type Component struct {
	width  int
	label  string
	dialog *dialog.Component
	kv     *keyvalue.Component
}

func Create() *Component {
	return &Component{
		dialog: dialog.Create(),
		kv:     keyvalue.Create(),
	}
}

func (f *Component) Load() tea.Cmd {

	allItems := []keyvalue.Item{}
	all := context.GetRegistry()

	for _, c := range all {
		if !c.IsActive() {
			continue
		}

		entries := c.GetEntriesFlat()
		items := []keyvalue.Item{}

		if len(entries) == 0 {
			continue
		}

		items = append(items, keyvalue.Item{
			Key:    c.GetLabel(),
			Value:  "",
			Header: true,
		})

		for _, e := range entries {
			key := e[0]
			value := e[1]

			items = append(items, keyvalue.Item{
				Key:   key,
				Value: value,
			})
		}

		slices.SortFunc(items, func(a, b keyvalue.Item) int {
			deepA := strings.Count(a.Key, ".")
			deepB := strings.Count(b.Key, ".")

			if a.Header && !b.Header {
				return -1
			}

			if !a.Header && b.Header {
				return 1
			}

			if deepA == 0 && deepB > 0 {
				return -1
			}

			if deepA > 0 && deepB == 0 {
				return 1
			}

			return strings.Compare(a.Key, b.Key)
		})

		allItems = append(allItems, items...)
	}

	f.kv.SetItems(allItems)

	return nil
}

func (f *Component) Init() tea.Cmd {
	f.dialog.OnRender(f.render)
	f.dialog.SetTitle("Context")

	return chain.Init(f.LoadBindings, f.kv.Init, f.dialog.Init)
}
