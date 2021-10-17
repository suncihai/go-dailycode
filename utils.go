package main

import (
	"bytes"
	"encoding/json"
)

const (
	emptyString = ""
	tabLine ="\t"
)

func prettyJSON(data interface{}) (string) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(emptyString, tabLine)

	err := encoder.Encode(data)
	if err != nil {
		return emptyString
	}
	return buffer.String()
}