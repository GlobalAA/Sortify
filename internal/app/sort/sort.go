package sort

import (
	"fmt"
	"io/fs"
	"log"
	"main/internal/app/config"
	"os"
	"path/filepath"
	"strings"
)

type Sort struct {
	config     *config.Config
	trueValues []string
}

func New(config *config.Config, trueValues []string) *Sort {
	return &Sort{
		config:     config,
		trueValues: trueValues,
	}
}

func (s *Sort) Sort() {
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	root += "/_data"

	for _, value := range ALLOWED_IMAGES {
		var files []string
		err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
				return nil
			}

			if !info.IsDir() && info.Name() != value && filepath.Ext(path) == value {
				files = append(files, path)
			}

			return nil
		})

		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			current_file := strings.Split(file, "/")
			charsets := CharsetsNew(s.trueValues, current_file[len(current_file)-1])
			charsets.Manage()
			filename := charsets.filename
			os.Mkdir(filepath.Join(root, value[1:]), os.ModePerm)
			path := file
			newPath := filepath.Join(root, value[1:], filename)

			if err := os.Rename(path, newPath); err != nil {
				log.Fatal(err)
			}
		}
	}
}
