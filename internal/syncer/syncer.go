package syncer

import (
	"time"

	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/database"
	"github.com/sidekick-coder/atlas/internal/drive/v2"
	"github.com/sidekick-coder/atlas/internal/syncer/extractor"
	"github.com/sidekick-coder/atlas/internal/syncer/scanner"
	"github.com/sidekick-coder/atlas/internal/syncer/writter"
)

type Result struct {
	Time time.Duration
	Scanned int32
	Extracted int32
	Written int32
}

type Syncer struct {
	drive    *drive.Drive
	database *database.Database
	config   *config.Config

	concurrency int

	scanner *scanner.Worker
	extractor *extractor.Worker
	writter *writter.Worker

	onSuccess func(path string)
	onError   func(path string, err error)
	onComplete func(result Result)
}

func Create() *Syncer {
	s := scanner.Create()
	e := extractor.Create()
	w := writter.Create()


	return &Syncer{
		scanner: s,
		extractor: e,
		writter: w,
	}
}

func (s *Syncer) SetConcurrency(c int) *Syncer {
	s.concurrency = c
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
	s.onSuccess = f
	return s
}

func (s *Syncer) OnError(f func(path string, err error)) *Syncer {
	s.onError = f
	return s
}

func (s *Syncer) OnComplete(f func(result Result)) *Syncer {
	s.onComplete = f
	return s
}
