package parser

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func parseJSCommand(commandData interface{}) (string, bool) {
	commandInt, ok := commandData.(int)
	if ok {
		return strconv.Itoa(commandInt), true
	}
	commandFloat32, ok := commandData.(float32)
	if ok {
		return strconv.FormatFloat(float64(commandFloat32), 'f', -1, 32), true
	}

	commandFloat64, ok := commandData.(float64)
	if ok {
		return strconv.FormatFloat(commandFloat64, 'f', -1, 64), true
	}

	commandString, ok := commandData.(string)
	if ok {
		return commandString, true
	}

	commandMap, ok := commandData.(map[string]interface{})
	if !ok {
		log.Fatalf("Failed to parse command %s\n", commandData)
	}

	value, ok := commandMap["type"]
	if !ok {
		log.Fatalf("No type found for command %s\n", commandMap)
	}

	commandType := value.(string)

	result := ""

	switch commandType {
	case "let":
		nameField := commandMap["name"].(string)
		valueField, _ := parseJSCommand(commandMap["value"])
		result += fmt.Sprintf("let %s = %s\n", nameField, valueField)

	case "const":
		nameField := commandMap["name"].(string)
		valueField, _ := parseJSCommand(commandMap["value"])
		result += fmt.Sprintf("const %s = %s\n", nameField, valueField)

	case "set":

	case "operator":
		operatorField := commandMap["operator"].(string)
		leftField, _ := parseJSCommand(commandMap["left"])
		rightField, _ := parseJSCommand(commandMap["right"])

		result += fmt.Sprintf("%s %s %s", leftField, operatorField, rightField)

	case "if":

	case "for":

	case "return":
		valueField, _ := parseJSCommand(commandMap["value"])
		result += fmt.Sprint(valueField)

	case "call":

	case "function":

	default:
		log.Fatalf("Unknown command type %s\n", commandType)
	}

	return result, false
}

func parseJS(jsData map[string]interface{}) {
	// Build css string
	fileName, exists := jsData["name"]
	if !exists {
		log.Fatal("No 'name' found in JS file definition.")
	}

	jsCommands, exists := jsData["commands"]
	if !exists {
		log.Fatal("No 'commands' found in JS file definition.")
	}

	commandsList, ok := jsCommands.([]interface{})
	if !ok {
		log.Fatal("Failed to parse commands map.")
	}

	result := ""
	for _, commandData := range commandsList {
		commandString, _ := parseJSCommand(commandData)
		result += commandString + "\n"
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
