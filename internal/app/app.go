package app

import (
	"github.com/sidekick-coder/atlas/internal/action"
	"github.com/sidekick-coder/atlas/internal/actionmanager"
	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/database"
	"github.com/sidekick-coder/atlas/internal/drive"
	"github.com/sidekick-coder/atlas/internal/repository/entry"
	"github.com/sidekick-coder/atlas/internal/repository/entrymeta"
	"github.com/sidekick-coder/atlas/internal/syncer"
)

type App struct {
	config   *config.Config
	drive    *drive.Drive
	database *database.Database

	entryRepo     *entry.Repository
	entryMetaRepo *entrymeta.Repository

	actionManager *actionmanager.ActionManager
	syncer        *syncer.Syncer

	Action *action.Manager
}

func Create() (*App, error) {
	config, err := config.Create()

	if err != nil {
		return nil, err
	}

	drive, err := drive.CreateFromConfig(config)

	if err != nil {
		return nil, err
	}

	drive.SetConfig(config)

	database, err := database.CreateFromConfig(config)

	if err != nil {
		return nil, err
	}

	actionManager, err := actionmanager.New(config)

	if err != nil {
		return nil, err
	}

	entryRepo := entry.New(database)
	entryMetaRepo := entrymeta.New(database)

	s := syncer.Create().SetConfig(config).SetDrive(drive).SetDatabase(database)
	a := action.Create()

	a.LoadConfigActions(config)

	app := &App{
		config:   config,
		database: database,
		drive:    drive,

		actionManager: actionManager,

		entryRepo:     entryRepo,
		entryMetaRepo: entryMetaRepo,

		syncer: s,

		Action: a,
	}

	return app, nil
}

func (a *App) WorkspacePath() string {
	wp, ok := a.config.Get("workspace.path")

	if !ok {
		return ""
	}

	return wp
}

func (a *App) ActionManager() *actionmanager.ActionManager {
	return a.actionManager
}

func (a *App) Config() *config.Config {
	return a.config
}

func (a *App) Database() *database.Database {
	return a.database
}

func (a *App) Drive() *drive.Drive {
	return a.drive
}

func (a *App) Syncer() *syncer.Syncer {
	return a.syncer
}

func (a *App) EntryRepo() *entry.Repository {
	return a.entryRepo
}

func (a *App) EntryMetaRepo() *entrymeta.Repository {
	return a.entryMetaRepo
}
