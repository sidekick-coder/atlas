package screen

import (
	"log/slog"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/app/program"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

func (f *Feature) HandleMessages(msg tea.Msg) tea.Cmd {
	if wm, ok := msg.(tea.WindowSizeMsg); ok {
		f.windowWidth = wm.Width
		f.windowHeight = wm.Height

		program.Send(f.Size())
	}

	if as, ok := msg.(AddMsg); ok {
		_, err := f.Add(as.Name, as.Options)

		if err != nil {
			return toast.Error(err.Error())
		}

		return f.Size
	}

	if rs, ok := msg.(RemoveMsg); ok {
		err := f.Remove(rs.Index)

		if err != nil {
			return toast.Error(err.Error())
		}

		return f.Size
	}

	if rp, ok := msg.(ReplaceMsg); ok {
		_, err := f.Replace(rp.Index, rp.Name, rp.Options)

		if err != nil {
			return toast.Error(err.Error())
		}

		return f.Size
	}

	if rc, ok := msg.(ReplaceCurrentMsg); ok {
		slog.Info("replacing current screen", slog.String("name", rc.Name), slog.Any("options", rc.Options))

		current := f.GetCurrentIndex()

		if current == -1 {
			return toast.Error("no current screen to replace")
		}

		_, err := f.Replace(current, rc.Name, rc.Options)

		if err != nil {
			return toast.Error(err.Error())
		}

		return f.Size
	}

	if s, ok := f.GetCurrent(); ok {
		return s.Update(msg)
	}

	return nil
}

func (f *Feature) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(
		msg,
		f.HandleMessages,
		chain.OnKey(f.HandleBinding),
	)
}
