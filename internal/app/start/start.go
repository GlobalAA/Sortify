package start

import (
	"main/internal/app/config"
	"main/internal/app/sort"
)

type Start struct {
}

func New() *Start {
	start := &Start{}
	return start
}

func (s *Start) Start() {
	config := config.New()
	trueValues := config.GetOnlyTrue()

	sort := sort.New(config, trueValues)
	sort.Sort()
}
