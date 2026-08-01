package formdialog

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/tui/components/form"
	"github.com/sidekick-coder/atlas/tui/components/form/field"
)

type ActionMsg struct {
	Title  string
	Submit string
	Fields []form.FieldData
	Action config.Action
}

var adialog *Component = Create()

func Action(ctx map[string]any) (map[string]any, error) {
	result := make(map[string]any)

	fields, err := field.CreateDataList(ctx["fields"])

	if err != nil {
		return result, err
	}

	msg := ActionMsg{
		Title:  "Form",
		Submit: "",
		Fields: fields,
	}

	if t, ok := ctx["title"].(string); ok {
		msg.Title = t
	}

	if ctx["action"] == nil {
		return result, fmt.Errorf("form action is required")
	}


	action, err := config.ParseAction(ctx["action"])

	if err != nil {
		return result, err
	}

	msg.Action = action

	result["tea_message"] = msg

	return result, nil
}

func HandleAction(exec func(string, ...map[string]any) tea.Cmd, msg tea.Msg) tea.Cmd {
	if adialog.IsOpen() {
		return adialog.Update(msg)
	}

	if m, ok := msg.(ActionMsg); ok {
		adialog.Init()
		adialog.SetTitle(m.Title)
		adialog.SetFields(m.Fields)
		adialog.SetAction(m.Action)
		adialog.Open()
		// adialog.OnSubmit(func(value string) tea.Cmd {
		// 	if m.Submit == "" {
		// 		return nil
		// 	}
		//
		// 	actionCtx := map[string]any{
		// 		"value": value,
		// 	}
		//
		// 	return exec(m.Submit, actionCtx)
		// })

		return nil
	}

	return nil
}
