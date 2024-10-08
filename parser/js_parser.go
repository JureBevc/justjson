package parser

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
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
		debug.PrintStack()
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
		result += fmt.Sprintf("let %s = %s", nameField, valueField)

	case "const":
		nameField := commandMap["name"].(string)
		valueField, _ := parseJSCommand(commandMap["value"])
		result += fmt.Sprintf("const %s = %s", nameField, valueField)

	case "set":
		variableField := commandMap["variable"].(string)
		valueField, _ := parseJSCommand(commandMap["value"])

		result += fmt.Sprintf("%s = %s", variableField, valueField)

	case "operator":
		operatorField := commandMap["operator"].(string)
		leftField, _ := parseJSCommand(commandMap["left"])
		rightField, _ := parseJSCommand(commandMap["right"])

		result += fmt.Sprintf("%s %s %s", leftField, operatorField, rightField)

	case "if":
		conditionField, _ := parseJSCommand(commandMap["condition"])
		thenCommands, isList := commandMap["then"].([]interface{})
		thenField := ""
		if isList {
			for _, thenCommand := range thenCommands {
				thenCommandString, _ := parseJSCommand(thenCommand)
				thenField += thenCommandString + "\n"
			}
		} else {
			thenField, _ = parseJSCommand(commandMap["then"])
		}
		result += fmt.Sprintf("if(%s){\n%s}", conditionField, thenField)

		_, elseExists := commandMap["else"]
		if elseExists {
			elseCommands, isList := commandMap["else"].([]interface{})
			elseField := ""
			if isList {
				for _, elseCommand := range elseCommands {
					elseCommandString, _ := parseJSCommand(elseCommand)
					elseField += elseCommandString + "\n"
				}
			} else {
				elseField, _ = parseJSCommand(commandMap["else"])
			}
			result += fmt.Sprintf("else{\n%s}", elseField)
		}

	case "for":
		initialField, _ := parseJSCommand(commandMap["initial"])
		conditionField, _ := parseJSCommand(commandMap["condition"])
		incrementField, _ := parseJSCommand(commandMap["increment"])

		commandsField := ""

		commandsList, isList := commandMap["commands"].([]interface{})
		if isList {
			for _, singleCommand := range commandsList {
				commandString, _ := parseJSCommand(singleCommand)
				commandsField += commandString + "\n"
			}
		} else {
			commandsField, _ = parseJSCommand(commandMap["commands"])
		}

		result += fmt.Sprintf("for(%s;%s;%s){\n%s}", initialField, conditionField, incrementField, commandsField)

	case "return":
		valueField, _ := parseJSCommand(commandMap["value"])
		result += fmt.Sprintf("return %s;", valueField)

	case "call":
		functionField, _ := parseJSCommand(commandMap["function"])

		parametersField := ""
		parametersList, isList := commandMap["parameters"].([]interface{})
		if isList {
			for i, singleParam := range parametersList {
				paramString, _ := parseJSCommand(singleParam)
				parametersField += paramString
				if i+1 < len(parametersList) {
					parametersField += ","
				}
			}
		} else {
			parametersField, _ = parseJSCommand(commandMap["parameters"])
		}
		result += fmt.Sprintf("%s(%s);", functionField, parametersField)

	case "function":
		nameField, _ := parseJSCommand(commandMap["name"])

		parametersField := ""
		parametersList, isList := commandMap["parameters"].([]interface{})
		if isList {
			for i, singleParam := range parametersList {
				paramString, _ := singleParam.(string)
				parametersField += paramString
				if i+1 < len(parametersList) {
					parametersField += ","
				}
			}
		} else {
			parametersField, _ = parseJSCommand(commandMap["parameters"])
		}

		commandsField := ""
		commandsList, isList := commandMap["commands"].([]interface{})
		if isList {
			for _, singleCommand := range commandsList {
				commandString, _ := parseJSCommand(singleCommand)
				commandsField += commandString + "\n"
			}
		} else {
			commandsField, _ = parseJSCommand(commandMap["commands"])
		}
		result += fmt.Sprintf("function %s(%s){\n%s}", nameField, parametersField, commandsField)

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
