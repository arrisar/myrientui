package scraper

import (
	"fmt"
	"log"
	"net/url"
	"path/filepath"

	"github.com/arrisar/myrientui/internal/files"
)

func (s Scraper) CachedIndexWrite(index Index) {
	dir, err := url.PathUnescape(index.Path)
	if err != nil {
		log.Fatalf("failed to decode index cache dir path: %v", err)
	}

	dir = filepath.Join(s.Config.CacheDir, dir)
	if err := files.MkDir(dir); err != nil {
		log.Fatalf("failed to make index dir: %v", err)
	}

	file := filepath.Join(dir, "index.yaml")
	if err := files.WriteYAML(file, index); err != nil {
		log.Fatalf("failed to write index file (%s): %v", file, err)
	}
}

func (s Scraper) CachedIndexRead(path string) (index Index, err error) {
	sub, err := url.PathUnescape(path)
	if err != nil {
		log.Fatalf("failed to decode index check dir path: %v", err)
	}

	dir := filepath.Join(s.Config.CacheDir, sub)
	if err = files.StatDir(dir); err != nil {
		err = fmt.Errorf("failed to stat absolute path to index dir: %s", dir)
		return
	}

	file := filepath.Join(dir, "index.yaml")
	if err = files.StatFile(file); err != nil {
		err = fmt.Errorf("failed to stat absolute path to index file: %s", file)
		return
	}

	if err = files.ReadYAML(file, &index); err != nil {
		log.Fatalf("failed to read index file: %v", err)
	}

	return
}
