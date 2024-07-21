package main

import (
	"log"
	"os"

	"github.com/thtg88/1brc/internal/configs"
	"github.com/thtg88/1brc/internal/datagenerator"
)

func main() {
	logger := log.New(os.Stdout, "", log.Lshortfile|log.Ltime)
	config := configs.NewDefaultDataGeneratorConfig()
	dataGenerator := datagenerator.NewDataGenerator(logger, config)
	dataGenerator.Generate()
}
