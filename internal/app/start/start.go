package start

import (
	"fmt"
	"log"
	"main/internal/app/config"
	"main/internal/app/sort"
	"os"
)

type Start struct {
}

func New() *Start {
	start := &Start{}
	return start
}

func (s *Start) Start() {
	s.Init()
	config := config.New()
	trueValues := config.GetOnlyTrue()

	sort := sort.New(config, trueValues)
	sort.Sort()
}

func (s *Start) Init() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	configPath := fmt.Sprintf("%s/Sortify/config.toml", dirname)
	if _, err = os.Stat(configPath); err != nil {
		if _, err := os.Stat(configPath); err != nil {
			os.Mkdir(fmt.Sprintf("%s/Sortify", dirname), os.ModePerm)
			os.Create(configPath)
		}
	}
}
