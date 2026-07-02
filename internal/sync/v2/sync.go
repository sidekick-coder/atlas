package sync 

import (
	"github.com/sidekick-coder/atlas/internal/drive/v2"
	"github.com/sidekick-coder/atlas/internal/repository/entry"
	"github.com/sidekick-coder/atlas/internal/repository/entrymeta"
)

type Sync struct {
	drive      *drive.Drive
	entryRepo  *entry.Repository
	entryMetaRepo   *entrymeta.Repository
}

func Create(drive *drive.Drive, entryRepo *entry.Repository, entryMetaRepo *entrymeta.Repository) *Sync {
	return &Sync{
		drive:     drive,
		entryRepo: entryRepo,
		entryMetaRepo:  entryMetaRepo,
	}
}

