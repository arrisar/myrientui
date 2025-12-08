package config

import (
	"log"
	"os"
	"os/user"
	"path/filepath"

	"sigs.k8s.io/yaml"
)

type Config struct {
	Home string
	Data ConfigData
}

func New() Config {
	var err error
	c := Config{}

	u, err := user.Current()
	if err != nil {
		log.Fatal("[Config] Setup:", err)
	}

	c.Home = filepath.Join(u.HomeDir, ".myrientui")
	c.load()

	return c
}

/**
 * HANDLERS
 */

func (c *Config) create() {
	var err error

	err = os.MkdirAll(c.Home, os.ModePerm)
	if err != nil {
		log.Fatal("[Config] Create config dir: ", err)
	}

	c.Data = c.DefaultData()

	err = os.MkdirAll(c.Data.Storage.BiosDir, os.ModePerm)
	if err != nil {
		log.Fatal("[Config] Create bios dir: ", err)
	}

	err = os.MkdirAll(c.Data.Storage.BiosDir, os.ModePerm)
	if err != nil {
		log.Fatal("[Config] Create cache dir: ", err)
	}

	err = os.MkdirAll(c.Data.Storage.RomsDir, os.ModePerm)
	if err != nil {
		log.Fatal("[Config] Create roms dir: ", err)
	}

	c.save()
}

func (c *Config) load() {
	var err error
	config := filepath.Join(c.Home, "config")

	// check if file exists
	_, err = os.Stat(config)
	if err != nil {
		c.create()
		return
	}

	// read file
	content, err := os.ReadFile(config)
	if err != nil {
		log.Fatal("[Config] Read: ", err)
	}

	// unmarshal
	err = yaml.Unmarshal(content, &c.Data)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Config) save() {
	var err error

	y, err := yaml.Marshal(c.Data)
	if err != nil {
		log.Fatal("[Config] Marshalling: ", err)
	}

	config := filepath.Join(c.Home, "config")
	err = os.WriteFile(config, []byte(y), 0666)
	if err != nil {
		log.Fatal("[Config] Saving: ", err)
	}
}
