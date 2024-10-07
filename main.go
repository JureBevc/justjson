package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/JureBevc/justjson/parser"
)

func main() {

	// Parse command line arguments
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		log.Fatalln("Missing path to JSON")
	}

	jsonPath := args[0]

	// Read JSON
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	// Parse JSON
	var jsonData interface{}

	// Create a new JSON decoder and decode the file content into result
	decoder := json.NewDecoder(jsonFile)
	if err := decoder.Decode(&jsonData); err != nil {
		log.Fatal(err)
	}

	absoluteJsonPath, err := filepath.Abs(jsonPath)
	if err != nil {
		log.Fatal(err)
	}

	absoluteJsonDirectory := filepath.Dir(absoluteJsonPath)

	parser.BuildPath = filepath.Join(absoluteJsonDirectory, "build")
	parser.ParseJsonData(&jsonData)

}
