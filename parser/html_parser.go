package parser

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func parseHTMLAttributes(attributesData []interface{}) string {

	result := ""
	for _, attributeData := range attributesData {
		attributeMap, ok := attributeData.(map[string]interface{})
		if !ok {
			log.Fatalf("Failed to parse attribute map %s\n", attributeData)
		}

		// Parse attribute name
		value, ok := attributeMap["name"]
		if !ok {
			log.Fatalf("No 'name' found for attribute %s\n", attributeMap)
		}

		attributeName, ok := value.(string)
		if !ok {
			log.Fatalf("Failed to create name string for attribute %s\n", attributeMap)
		}

		// Parse attribute value
		value, ok = attributeMap["value"]
		if !ok {
			log.Fatalf("No 'value' found for attribute %s\n", attributeMap)
		}

		attributeValue, ok := value.(string)
		if !ok {
			log.Fatalf("Failed to create value string for attribute %s\n", attributeMap)
		}

		result += fmt.Sprintf(" %s='%s'", attributeName, attributeValue)
	}
	return result
}

func parseHTMLElement(elementData interface{}) string {
	// Check for raw text
	stringData, ok := elementData.(string)
	if ok {
		return stringData
	}

	// Build HTML tag
	result := ""
	elementMap, ok := elementData.(map[string]interface{})
	if !ok {
		log.Fatalf("Failed to parse element data %s\n", elementData)
	}

	// Parse element tag name
	value, ok := elementMap["tag"]
	if !ok {
		log.Fatalf("Element has no 'tag':\n%s\n", elementData)
	}

	tagName, ok := value.(string)
	if !ok {
		log.Fatalf("Failed to parse tag name %s\n", value)
	}

	// Parse element attributes
	attributesResult := ""
	value, ok = elementMap["attributes"]
	if ok {
		attributesData, ok := value.([]interface{})
		if !ok {
			log.Fatalf("Failed to parse attributes %s\n", value)
		}
		attributesResult = parseHTMLAttributes(attributesData)
	}

	elementContent := ""
	value, ok = elementMap["elements"]
	if ok {
		childElements, ok := value.([]interface{})
		if !ok {
			log.Fatalf("Failed to parse elements %s\n", value)
		}

		for _, childElementData := range childElements {
			elementContent += parseHTMLElement(childElementData)
		}
	}

	result += fmt.Sprintf("<%s%s>\n", tagName, attributesResult)
	result += elementContent + "\n"
	result += fmt.Sprintf("</%s>", tagName)
	return result
}

func parseHTML(htmlData map[string]interface{}) {
	// Build html string
	fileName, exists := htmlData["name"]
	if !exists {
		log.Fatal("No 'name' found in HTML file definition.")
	}

	htmlElements, exists := htmlData["elements"]
	if !exists {
		log.Fatal("No 'elements' found in HTML file definition.")
	}

	htmlElementsList, ok := htmlElements.([]interface{})
	if !ok {
		log.Fatal("Failed to parse element list.")
	}

	result := "<html>\n"
	for _, elementData := range htmlElementsList {
		result += parseHTMLElement(elementData) + "\n"
	}
	result += "</html>"

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
