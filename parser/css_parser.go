package parser

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func parseCSS(cssData map[string]interface{}) {
	// Build css string
	fileName, exists := cssData["name"]
	if !exists {
		log.Fatal("No 'name' found in CSS file definition.")
	}

	cssContent, exists := cssData["content"]
	if !exists {
		log.Fatal("No 'content' found in CSS file definition.")
	}

	contentMap, ok := cssContent.(map[string]interface{})
	if !ok {
		log.Fatal("Failed to parse content map.")
	}

	result := ""
	for selectorKey, selectorValue := range contentMap {
		result += fmt.Sprintf("%s {\n", selectorKey)

		propertyMap, ok := selectorValue.(map[string]interface{})
		if !ok {
			log.Fatalf("Failed to parse property map %s\n", selectorValue)
		}
		for propertyKey, propertyvalue := range propertyMap {
			result += fmt.Sprintf("\t%s: %s;\n", propertyKey, propertyvalue)
		}

		result += "}\n"
	}

	// Write to file
	filePath := filepath.Join(BuildPath, fileName.(string))

	htmlFile, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer htmlFile.Close()

	_, err = htmlFile.WriteString(result)
	if err != nil {
		log.Fatal(err)
	}
}
