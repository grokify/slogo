package main

import (
	"log"
	"os"

	"github.com/grokify/slogo/examples"
	"github.com/grokify/slogo/ontology"
)

func main() {
	filename := "METRICS.md"
	sloSetsByDirectory := examples.ExampleSLOsByDirectory()

	// Open output file
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = ontology.WriteSLOSetsReport(f, sloSetsByDirectory)
	if err != nil {
		log.Fatal(err)
	}
}
