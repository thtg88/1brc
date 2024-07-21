package main

import (
	"log"
	"os"

	"github.com/thtg88/1brc/internal/datagenerator"
)

func main() {
	logger := log.New(os.Stdout, "", log.Lshortfile|log.Ltime)
	dataGenerator := datagenerator.NewDataGenerator(logger)
	dataGenerator.Generate()
}
