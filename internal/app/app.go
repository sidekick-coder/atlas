package app

import (
	"github.com/sidekick-coder/atlas/internal/actionmanager"
	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/database"
	"github.com/sidekick-coder/atlas/internal/drive/v2"
	"github.com/sidekick-coder/atlas/internal/repository/entry"
	"github.com/sidekick-coder/atlas/internal/repository/entrymeta"
	sync "github.com/sidekick-coder/atlas/internal/sync/v2"
)

type App struct {
	config   *config.Config
	drive    *drive.Drive
	database *database.Database

	entryRepo     *entry.Repository
	entryMetaRepo *entrymeta.Repository

	actionManager *actionmanager.ActionManager
	syncer        *sync.Sync
}

func Create() (*App, error) {
	config, err := config.Create()

	if err != nil {
		return nil, err
	}

	drive, err := drive.New(config.Get("workspace.path"))

	if err != nil {
		return nil, err
	}

	database, err := database.Create(config.Get("workspace.database_path"))

	if err != nil {
		return nil, err
	}

	actionManager, err := actionmanager.New(config)

	if err != nil {
		return nil, err
	}


	entryRepo := entry.New(database)
	entryMetaRepo := entrymeta.New(database)

	syncer := sync.Create(drive, entryRepo, entryMetaRepo)

	app := &App{
		config:   config,
		database: database,
		drive:    drive,

		actionManager: actionManager,

		entryRepo:     entryRepo,
		entryMetaRepo: entryMetaRepo,

		syncer: syncer,
	}

	return app, nil
}

func (a *App) WorkspacePath() string {
	return a.config.Get("workspace.path")
}

func (a *App) ActionManager() *actionmanager.ActionManager {
	return a.actionManager
}

func (a *App) Config() *config.Config {
	return a.config
}

func (a *App) Drive() *drive.Drive {
	return a.drive
}

func (a *App) Syncer() *sync.Sync {
	return sync.Create(a.drive, a.entryRepo, a.entryMetaRepo)
}

func (a *App) EntryRepo() *entry.Repository {
	return a.entryRepo
}

func (a *App) EntryMetaRepo() *entrymeta.Repository {
	return a.entryMetaRepo
}
