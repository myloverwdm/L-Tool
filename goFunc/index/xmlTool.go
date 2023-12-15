package index

import (
	"LTool/goFunc"
	"encoding/json"
	"fmt"
	"strings"
)

const className = "lTool"

func jsonToXml(jsonData []byte) ([]byte, string, error) {
	var data map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, className, err
	}

	xmlData, err := json2xml(data, "")
	if err != nil {
		return nil, className, err
	}
	classType, ok := data["@type"]
	if ok {
		fullClassName := fmt.Sprintf("%v", classType)
		parts := strings.Split(fullClassName, ".")
		lastPart := parts[len(parts)-1]
		return xmlData, strings.ToLower(string(lastPart[0])) + lastPart[1:], nil
	}
	return xmlData, className, nil
}

func json2xml(data map[string]interface{}, indent string) ([]byte, error) {
	var xmlData []byte
	for key, value := range data {
		xmlData = append(xmlData, []byte(fmt.Sprintf("%s<%s>", indent, key))...)
		switch value := value.(type) {
		case map[string]interface{}:
			childData, err := json2xml(value, indent+"  ")
			if err != nil {
				return nil, err
			}
			xmlData = append(xmlData, childData...)
		default:
			xmlData = append(xmlData, []byte(fmt.Sprintf("%v", value))...)
		}
		xmlData = append(xmlData, []byte(fmt.Sprintf("</%s>\n", key))...)
	}

	return xmlData, nil
}

func add2Space(input string) string {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if line != "" {
			lines[i] = "  " + line
		} else {
			lines[i] = line
		}
	}
	return strings.Join(lines, "\n")
}

// JsonToXml xml和json互转
func JsonToXml(value string) goFunc.LTResponse {
	xmlData, bookstore, err := jsonToXml([]byte(value))
	if err != nil {
		return goFunc.LTResponse{Success: false, ErrorMsg: err.Error()}
	}
	return goFunc.LTResponse{Success: true, Data: "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<" + bookstore + ">\n" + add2Space(string(xmlData)) + "</" + bookstore + ">"}
}
