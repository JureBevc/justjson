package parser

import (
	"log"
	"os"
	"strings"
)

var BuildPath string = "."

func ParseJsonData(jsonData *interface{}) {
	rootMap, ok := (*jsonData).(map[string]interface{})

	if !ok {
		log.Fatal("Failed to parse root JSON structure.")
	}

	if err := os.MkdirAll(BuildPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	definitions, ok := rootMap["definitions"]
	if !ok {
		log.Fatalln("No definitions found.")
	}

	definitionsList, ok := definitions.([]interface{})
	if !ok {
		log.Fatalln("Failed to parse root definitions.")
	}

	for _, value := range definitionsList {
		definitionMap, ok := value.(map[string]interface{})
		if !ok {
			log.Fatalf("Failed to create definition map %s\n", value)
		}
		definitionType, ok := definitionMap["type"]
		if !ok {
			log.Fatalf("No type found for definition %s\n", definitionMap)
		}
		definitionTypeString, ok := definitionType.(string)
		if !ok {
			log.Fatalf("Failed to load definition type string%s\n", definitionTypeString)
		}

		keyLower := strings.ToLower(definitionTypeString)

		if keyLower == "html" {
			parseHTML(definitionMap)
			continue
		}

		if keyLower == "css" {
			parseCSS(definitionMap)
			continue
		}

		if keyLower == "js" || keyLower == "javascript" {
			parseJS(definitionMap)
			continue
		}
	}
}
