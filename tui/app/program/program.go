package program

import (
	"fmt"
	"log/slog"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/config"
)

var teaProgram *tea.Program
var internalApp *app.App

func GetProgram() *tea.Program {
	return teaProgram
}

func SetProgram(p *tea.Program) {
	teaProgram = p
}

func GetApp() *app.App {
	return internalApp
}

func SetApp(a *app.App) {
	internalApp = a
}

func GetConfig() *config.Config {
	return internalApp.Config()
}

func Send[T tea.Msg](msg T) error {
	if teaProgram == nil {
		slog.Error("program is not set")

		return fmt.Errorf("program is not set")
	}

	go teaProgram.Send(msg)

	return nil
}

func Command[T tea.Msg](msg T) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}

func CommandSkip() tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}
