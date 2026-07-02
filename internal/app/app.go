package app 

import (
	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/database"
	"github.com/sidekick-coder/atlas/internal/drive/v2"
	"github.com/sidekick-coder/atlas/internal/repository/entry"
	"github.com/sidekick-coder/atlas/internal/repository/entrymeta"
)

type App struct {
	config *config.Config
	drive *drive.Drive
	database *database.Database
	entryRepo *entry.Repository
	entryMetaRepo *entrymeta.Repository
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

	entryRepo := entry.New(database)
	entryMetaRepo := entrymeta.New(database)

	app := &App{
		config: config,
		database: database,
		drive: drive,
		entryRepo: entryRepo,
		entryMetaRepo: entryMetaRepo,
	}

	return app, nil
}

func (a *App) Drive() *drive.Drive {
    return a.drive
}

func (a *App) EntryRepo() *entry.Repository {
    return a.entryRepo
}

func (a *App) EntryMetaRepo() *entrymeta.Repository {
    return a.entryMetaRepo
}
