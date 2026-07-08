package syncer

import (
	"time"

	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/database"
	"github.com/sidekick-coder/atlas/internal/drive/v2"
	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/internal/syncer/batcher"
	"github.com/sidekick-coder/atlas/internal/syncer/extractor"
	"github.com/sidekick-coder/atlas/internal/syncer/scanner"
	"github.com/sidekick-coder/atlas/internal/syncer/writter"
)

type Result struct {
	Concurrency int 
	BatchSize int
	Time time.Duration
	Scanned int32
	Extracted int32
	Written int32
	Batches int32
}

type Syncer struct {
	drive    *drive.Drive
	database *database.Database
	config   *config.Config

	concurrency int

	scanner *scanner.Worker
	extractor *extractor.Worker
	batcher *batcher.Worker
	writter *writter.Worker

	onComplete func(result Result)
}

func Create() *Syncer {
	s := scanner.Create()
	e := extractor.Create()
	b := batcher.Create()
	w := writter.Create()


	return &Syncer{
		scanner: s,
		extractor: e,
		batcher: b,
		writter: w,
	}
}

func (s *Syncer) SetConcurrency(c int) *Syncer {
	s.concurrency = c
	return s
}

func (s *Syncer) SetBatchSize(b int) *Syncer {
	s.batcher.SetBatchSize(b)
	return s
}

func (s *Syncer) SetConfig(c *config.Config) *Syncer {
	s.config = c
	return s
}

func (s *Syncer) SetDrive(d *drive.Drive) *Syncer {
	s.drive = d

	return s
}

func (s *Syncer) SetDatabase(db *database.Database) *Syncer {
	s.database = db
	return s
}

func (s *Syncer) OnSuccess(f func(path string)) *Syncer {
	s.writter.OnSuccess(func(e extractor.ExtractEntry) {
		f(e.Path)
	})
	return s
}

func (s *Syncer) OnError(f func(path string, err error)) *Syncer {
	s.writter.OnError(func(e extractor.ExtractEntry, err error) {
		f(e.Path, err)
	})

	s.extractor.OnError(func(e models.EntryInfo, err error) {
		f(e.Path, err)
	})

	return s
}

func (s *Syncer) OnComplete(f func(result Result)) *Syncer {
	s.onComplete = f
	return s
}

func (s *Syncer) OnBatchComplete(f func(batch batcher.Batch)) *Syncer {
	s.writter.OnBatchComplete(f)
	return s
}

func (s *Syncer) OnScanComplete(f func(count int32)) *Syncer {
	s.scanner.OnComplete(f)
	return s
}
