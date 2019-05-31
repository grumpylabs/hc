// +build ignore

// Creates swift code for all HomeKit service and characteristic types and logs
// that to the console.

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/grumpylabs/hcf/gen"
	"github.com/grumpylabs/hcf/gen/swift"
)

// LibPath is the location of the top directory
var LibPath = os.ExpandEnv("$GOPATH/src/github.com/grumpylabs/hcf")

// GenPath is relative to the LibPath and has the gen files
var GenPath = filepath.Join(LibPath, "gen")

// MetadataPath is the relative to LibPath JSON file that contains the output from plutil.
var MetadataPath = filepath.Join(GenPath, "metadata.json")

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

	if b, err := swift.CharacteristicEnumDecl(metadata); err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(b))
	}

	if b, err := swift.ServiceEnumDecl(metadata); err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(b))
	}
}
