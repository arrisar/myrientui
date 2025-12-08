package config

import "path/filepath"

type StorageData struct {
	CacheDir string `json:"cacheDir"`
	BiosDir  string `json:"biosDir"`
	RomsDir  string `json:"romsDir"`
}

func (c Config) DefaultStorageData() StorageData {
	d := StorageData{}

	d.CacheDir = filepath.Join(c.Home, "cache")
	d.BiosDir = filepath.Join(c.Home, "library/bios")
	d.RomsDir = filepath.Join(c.Home, "library/roms")

	return d
}
