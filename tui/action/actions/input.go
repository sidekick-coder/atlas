package actions

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/inputdialog"
)

type InputActionMsg struct {
	Label  string
	Submit string
}

var dialog *inputdialog.Component

func InitInputDialog() tea.Cmd {
	dialog = inputdialog.Create()

	return dialog.Init()
}

func InputAction(ctx map[string]any) (map[string]any, error) {
	result := make(map[string]any)

	msg := InputActionMsg{
		Label:  "Input",
		Submit: "",
	}

	if l, ok := ctx["label"].(string); ok {
		msg.Label = l
	}

	if s, ok := ctx["submit"].(string); ok {
		msg.Submit = s
	}

	result["tea_message"] = msg

	return result, nil
}

func HandleInput(exec func(string, ...map[string]any) tea.Cmd, msg tea.Msg) tea.Cmd {
	if dialog.IsOpen() {
		return dialog.Update(msg)
	}

	if m, ok := msg.(InputActionMsg); ok {
		dialog.SetTitle(m.Label)
		dialog.SetContent("")
		dialog.Open()
		dialog.OnSubmit(func(value string) tea.Cmd {
			if m.Submit == "" {
				return nil
			}

			actionCtx := map[string]any{
				"value": value,
			}

			return exec(m.Submit, actionCtx)
		})

		return nil
	}

	return nil
}
