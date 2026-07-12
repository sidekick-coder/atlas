package app

import (
	"fmt"
	"log/slog"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/tui/app/program"
	"github.com/sidekick-coder/atlas/tui/root"
)

type App struct {
	*app.App
	program *tea.Program
}

func Create() (*App, error) {
	app := &App{}

	return app, nil
}

func (a *App) Init() error {
	ia, err := app.Create()

	if err != nil {
		return fmt.Errorf("failed to create app: %w", err)
	}

	a.App = ia

	return nil
}

func (a *App) LoadLogger() error {
	file, err := os.OpenFile("tui.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	logger := slog.New(slog.NewJSONHandler(file, nil))

	slog.SetDefault(logger)

	slog.Info("Logger initialized")

	return nil
}

func (a *App) LoadProgram() error {
	model := root.New(a.App)

	p := tea.NewProgram(model)

	a.program = p

	program.SetProgram(p)

	return nil
}

func (a *App) Run() error {
	err := a.LoadLogger()

	if err != nil {
		return fmt.Errorf("failed to load logger: %w", err)
	}

	err = a.Init()

	if err != nil {
		return fmt.Errorf("failed to initialize app: %w", err)
	}

	err = a.LoadProgram()

	if err != nil {
		return fmt.Errorf("failed to load program: %w", err)
	}

	_, err = a.program.Run()

	if err != nil {
		return fmt.Errorf("failed to run program: %w", err)
	}

	return nil
}
