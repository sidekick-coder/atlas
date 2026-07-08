package syncer

import (
	"sync"
	"time"

	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/internal/syncer/batcher"
	"github.com/sidekick-coder/atlas/internal/syncer/extractor"
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
	batches := make(chan batcher.Batch)
	done := make(chan struct{})

	go func(){
		s.writter.Run(batches, concurrency)
		close(done)
	}()

	go s.batcher.Run(extractions, batches, concurrency)
	go s.extractor.Run(infos, extractions, concurrency)

	s.scanner.Run(infos, concurrency)

	<-done

	if s.onComplete != nil {
		s.onComplete(Result{
			Concurrency: concurrency,
			BatchSize:  s.batcher.GetBatchSize(),
			Time:      time.Since(start),
			Scanned:   s.scanner.GetCount(),
			Extracted: s.extractor.GetCount(),
			Written:   s.writter.GetCount(),
			Batches:   s.batcher.GetCount(),
		})
	}
}
