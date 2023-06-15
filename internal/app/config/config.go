package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

const config_file = "config.toml"

type Config struct {
	FolderAssociation FolderAssociation `toml:"foldersAssociation"`
	Charsets          Charsets          `toml:"charsets"`
}

type FolderAssociation struct {
	PNG  string `toml:"PNG"`
	JPG  string `toml:"JPG"`
	JPEG string `toml:"JPEG"`
	WEBP string `toml:"WEBP"`
	SVG  string `toml:"SVG"`
}

type Charsets struct {
	LowerEmphasis bool `toml:"lowerEmphasis"`
	Dash          bool `toml:"dash"`
	Number        bool `toml:"number"`
	OnlyLowerCase bool `toml:"onlyLowerCase"`
}

func New() *Config {
	return &Config{}
}

func (c *Config) Read() *Config {
	var conf Config
	if _, err := toml.DecodeFile(config_file, &conf); err != nil {
		log.Fatal(err)
	}

	return &conf
}

func (c *Config) GetOnlyTrue() []string {
	config := *c.Read()

	returnMap := []string{}

	charsetsMap := map[string]bool{
		"LowerEmphasis": config.Charsets.LowerEmphasis,
		"Dash":          config.Charsets.Dash,
		"Number":        config.Charsets.Number,
		"OnlyLowerCase": config.Charsets.OnlyLowerCase,
	}

	for key, value := range charsetsMap {
		if value {
			returnMap = append(returnMap, key)
		}
	}

	return returnMap
}
