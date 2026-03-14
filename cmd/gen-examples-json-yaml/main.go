package main

import (
	"fmt"
	"log"
	"os"

	"github.com/grokify/mogo/os/osutil"
	"github.com/grokify/slogo/examples"
)

func main() {
	examplesDir := osutil.MustUserHomeDir("go/src/github.com/grokify/slogo/examples")

	err := examples.WriteExampleJSONAndYAMLFiles(examplesDir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")
	os.Exit(0)
}
