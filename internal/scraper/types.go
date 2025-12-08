package scraper

import "time"

type Scraper struct {
	Config Config
}

type Config struct {
	CacheDir string
	Workers  int
}

type IndexJob struct {
	Path     string
	LatestAt time.Time
}

type Index struct {
	Path      string    `yaml:"path"`
	UpdatedAt time.Time `yaml:"updated_at"`
	Links     []Link    `yaml:"links"`
}

type Link struct {
	IsDir     bool      `yaml:"isDir"`
	Link      string    `yaml:"link"`
	Label     string    `yaml:"label"`
	Size      string    `yaml:"size,omitempty"`
	UpdatedAt time.Time `yaml:"updatedAt"`
	IndexedAt time.Time `yaml:"indexedAt"`
}
