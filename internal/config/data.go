package config

type ConfigData struct {
	Storage   StorageData  `json:"storage"`
	Platforms PlatformData `json:"platforms"`
}

func (c Config) DefaultData() ConfigData {
	d := ConfigData{}
	d.Storage = c.DefaultStorageData()
	d.Platforms = c.DefaultPlatformData()
	return d
}
