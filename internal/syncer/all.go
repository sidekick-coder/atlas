package syncer

import (
	"sync"
	"time"

	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/internal/syncer/extractor"
	"github.com/sidekick-coder/atlas/internal/syncer/writter"
)

func (s *Syncer) AllWorker(wg *sync.WaitGroup) {
	defer wg.Done()

}

func (s *Syncer) All() {
	concurrency := s.concurrency
	start := time.Now()

	s.scanner.SetDrive(s.drive)
	s.extractor.SetConfig(s.config)
	s.writter.SetDatabase(s.database)

	if s.onSuccess != nil {
		s.writter.OnSuccessPath(s.onSuccess)
	}

	if s.onError != nil {
		s.writter.OnErrorPath(s.onError)
	}

	err := s.database.Cleanup()

	if err != nil {
		panic(err)
	}

	infos := make(chan models.EntryInfo)
	extractions := make(chan extractor.ExtractEntry)
	batches := make(chan writter.Batch)

	go s.writter.Run(batches, concurrency)
	go s.writter.BatcherRun(extractions, batches, concurrency)
	go s.extractor.Run(infos, extractions, concurrency)

	s.scanner.Run(infos, concurrency)

	timeTaken := time.Since(start)

	if s.onComplete != nil {
		s.onComplete(Result{
			Time: timeTaken,
			Scanned:   s.scanner.GetCount(),
			Extracted: s.extractor.GetCount(),
			Written:   s.writter.GetCount(),
		})
	}
}
