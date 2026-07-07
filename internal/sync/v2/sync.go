package sync

import (
	"sync/atomic"

	"github.com/sidekick-coder/atlas/internal/database"
	"github.com/sidekick-coder/atlas/internal/drive/v2"
	"github.com/sidekick-coder/atlas/internal/repository/entry"
	"github.com/sidekick-coder/atlas/internal/repository/entrymeta"
)

type Sync struct {
	drive      *drive.Drive
	Database   *database.Database
	entryRepo  *entry.Repository
	entryMetaRepo   *entrymeta.Repository
	EntryRepo  *entry.Repository
	EntryMetaRepo   *entrymeta.Repository

	NextID atomic.Int64
	TotalEntries atomic.Int64 
	TotalEntriesErrors atomic.Int64 
	TotalBatches atomic.Int64
	TotalBatchesErrors atomic.Int64
}

type SyncPayload struct {
	Drive      *drive.Drive
	Database   *database.Database
	EntryRepo  *entry.Repository
	EntryMetaRepo   *entrymeta.Repository
}

func Create(p *SyncPayload) *Sync {
	return &Sync{
		drive:      p.Drive,
		Database:   p.Database,
		entryRepo:  p.EntryRepo,
		entryMetaRepo:   p.EntryMetaRepo,
		EntryRepo:  p.EntryRepo,
		EntryMetaRepo:   p.EntryMetaRepo,
	}
}

