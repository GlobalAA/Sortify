package config

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

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
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	config_file := fmt.Sprintf("%s/Sortify/config.toml", dirname)
	var conf Config
	if _, err := toml.DecodeFile(config_file, &conf); err != nil {
		log.Fatal(err)
	}

	associationMap := map[string]string{
		"PNG":  conf.FolderAssociation.PNG,
		"JPG":  conf.FolderAssociation.JPG,
		"JPEG": conf.FolderAssociation.JPEG,
		"WEBP": conf.FolderAssociation.WEBP,
		"SVG":  conf.FolderAssociation.SVG,
	}

	for _, value := range associationMap {
		if len(value) <= 0 {
			fmt.Println("Your configuration is corrupted!")
			os.Exit(0)
		}
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
