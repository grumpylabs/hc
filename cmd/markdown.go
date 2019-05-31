// +build ignore

// Creates Markdown files to summarize all available HomeKit service and characteristic types.

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/grumpylabs/hcf/gen"
	"github.com/grumpylabs/hcf/gen/markdown"
)

// LibPath is the location of the top directory
var LibPath = os.ExpandEnv("$GOPATH/src/github.com/grumpylabs/hcf")

// MetadataPath is the relative to LibPath JSON file that contains the output from plutil.
var MetadataPath = filepath.Join(GenPath, "metadata.json")

// GenPath is relative to the LibPath and has the gen files
var GenPath = filepath.Join(LibPath, "gen")

// SvcFilePath is where we will plop the service markdown
var SvcFilePath = filepath.Join(LibPath, "service/README.md")

// AccFilePath is where we will plop the accessory markdown
var AccFilePath = filepath.Join(LibPath, "accessory/README.md")

func main() {

	log.Println("Import data from", MetadataPath)

	// Open metadata file
	f, err := os.Open(MetadataPath)
	if err != nil {
		log.Fatal(err)
	}

	// Read content
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	// Import json
	metadata := gen.Metadata{}
	err = json.Unmarshal(b, &metadata)
	if err != nil {
		log.Fatal(err)
	}

	if b, err := markdown.CategoriesCode(&metadata); err != nil {
		log.Fatal(err)
	} else {
		filePath := AccFilePath
		log.Println("Creating file", filePath)
		if f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666); err != nil {
			log.Fatal(err)
		} else {
			if _, err := f.Write(b); err != nil {
				log.Fatal()
			}
		}
	}

	if b, err := markdown.ServicesCode(&metadata); err != nil {
		log.Fatal(err)
	} else {
		filePath := SvcFilePath
		log.Println("Creating file", filePath)
		if f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666); err != nil {
			log.Fatal(err)
		} else {
			if _, err := f.Write(b); err != nil {
				log.Fatal()
			}
		}
	}
}
