package userkeymaps

import (
	"log/slog"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/actionmanager"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/key"
)

type Feature struct {
	app      *app.App

	am       *actionmanager.ActionManager
	ctx      map[string]any

	keymaps  []config.Keymap
	bindings []key.Binding
	groups   []string
	beforeActions []func() error
}

func Create() *Feature {
	return &Feature{
		groups:  []string{},
		keymaps: []config.Keymap{},
		ctx:     map[string]any{},
		beforeActions: []func() error{},
	}
}

func (f *Feature) SetApp(app *app.App) *Feature {
	f.app = app
	f.am = app.ActionManager()
	return f
}

func (f *Feature) OnBeforeActions(actions ...func() error) *Feature {
	f.beforeActions = append(f.beforeActions, actions...)
	return f
}

func (f *Feature) SetContext(ctx map[string]any) *Feature {
	f.ctx = ctx 
	return f
}

func (f *Feature) AddContext(key string, value any) *Feature {
	f.ctx[key] = value

	return f
}

func (f *Feature) SetGroups(groups []string) *Feature {
	f.groups = groups
	return f
}

func (f *Feature) AddGroup(group string) *Feature {
	f.groups = append(f.groups, group)
	return f
}

func (f *Feature) LoadKeymaps() {
	bindings := []key.Binding{}
	keymaps := []config.Keymap{}

	for _, group := range f.groups {
		gk := f.app.Config().GetKeymapsByGroup(group)

		keymaps = append(keymaps, gk...)
	}

	for _, action := range keymaps {
		b := key.CreateBinding(action.Keys...).
			SetDescription(action.Description).
			SetTags("user").
			SetHelp(action.Keys[0]). 
			SetID(action.ID)

		bindings = append(bindings, b)
	}

	f.bindings = bindings
	f.keymaps = keymaps
}

func (f *Feature) Init() error {
	return nil
}

func (f *Feature) Load() {
	f.LoadKeymaps()
	key.Register(f.bindings...)
}

func (f *Feature) Unload() {
	key.Unregister(f.bindings...)
	slog.Info("unloaded user keymaps", "groups", f.groups, "bindings", len(f.bindings))
}

func (f *Feature) ExecuteKeymap(km config.Keymap) error {
	for _, action := range f.beforeActions {
		err := action()

		if err != nil {
			return err
		}
	}

	err := f.am.Execute(km.Action, f.ctx)

	if err != nil {
		return err
	}

	return nil
}

func (f *Feature) HandleBinding(msg tea.KeyMsg) tea.Cmd {
	if len(f.groups) == 0 {
		return nil
	}

	for index, b := range f.bindings {
		km := f.keymaps[index]

		if key.Matches(b) {
			err := f.ExecuteKeymap(km)

			if err != nil {
				return toast.Error(err.Error())
			}
		}
	}

	return nil
}

func (f *Feature) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(msg, chain.OnKey(f.HandleBinding))
}
