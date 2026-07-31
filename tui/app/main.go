package app

import (
	"fmt"
	"log/slog"
	"path/filepath"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/app"
	"github.com/sidekick-coder/atlas/internal/logger"
	"github.com/sidekick-coder/atlas/tui/app/model"
	"github.com/sidekick-coder/atlas/tui/app/program"
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
	atlasPath, exists := a.App.Config().Get("workspace.atlas_path")

	if !exists {
		return fmt.Errorf("atlas_path not found in config")
	}

	f, err := logger.CreateFileTransport(filepath.Join(atlasPath, "app.log"))

	if err != nil {
		return fmt.Errorf("failed to create file transport: %w", err)
	}

	sl := f.GetLogger()

	slog.SetDefault(sl)

	l := logger.Create()

	l.AddTransport(f)

	logger.SetLogger(l)

	logger.Info("Logger initialized")

	return nil
}

func (a *App) LoadProgram() error {
	model := model.Create(a.App)

	p := tea.NewProgram(model)

	a.program = p

	program.SetProgram(p)
	program.SetApp(a.App)

	return nil
}

func (a *App) Run() error {

	err := a.Init()

	if err != nil {
		return fmt.Errorf("failed to initialize app: %w", err)
	}

	err = a.LoadLogger()

	if err != nil {
		return fmt.Errorf("failed to load logger: %w", err)
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
