package scraper

import (
	"log"
	"net/url"
	"sync"
	"time"
)

func (s Scraper) IndexPreload(index Index) {
	for _, r := range index.Links {
		if r.IsDir {
			go s.DoScrape(index.Path + r.Link)
		}
	}
}

func (s Scraper) IndexAll() {
	queue := make(chan IndexJob)
	var wg sync.WaitGroup

	// add root job
	wg.Add(1)
	go func() {
		queue <- IndexJob{"/", time.Now()}
	}()

	// start workers
	for range s.Config.Workers {
		go s.IndexWorker(&wg, queue, queue)
	}

	wg.Wait()
	close(queue)
}

func (s Scraper) IndexWorker(wg *sync.WaitGroup, qr <-chan IndexJob, qw chan<- IndexJob) {
	for job := range qr {
		s.IndexRecursive(wg, qw, job)
	}
}

func (s Scraper) IndexRecursive(wg *sync.WaitGroup, queue chan<- IndexJob, job IndexJob) {
	index, err := s.CachedIndexRead(job.Path)
	decoded, _ := url.PathUnescape(job.Path)

	// fetch index if missing or older
	if err != nil || index.UpdatedAt.Before(job.LatestAt) {
		log.Printf("index miss: %s", decoded)

		latest, err := s.ScrapePath(job.Path)
		if err != nil {
			log.Fatalf("failed to index path (%s): %v", job.Path, err)
		}

		s.CachedIndexWrite(latest)
		index = latest
	} else {
		log.Printf("index hit:  %s", decoded)
	}

	// iterate links
	for _, link := range index.Links {
		if link.IsDir {
			wg.Add(1)
			go func() {
				queue <- IndexJob{job.Path + link.Link, link.UpdatedAt}
			}()
		}
	}

	wg.Done()
}
