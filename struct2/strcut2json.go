package struct2

import (
	"bytes"
	"encoding/json"
	"strings"
)

// have all fields
func ToString(obj interface{}) (string, error) {
	beforeValue, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	midResultMap := make(map[string]interface{}, 0)

	jen := json.NewDecoder(bytes.NewBuffer(beforeValue))
	jen.UseNumber()
	err = jen.Decode(&midResultMap)
	if err != nil {
		return "", err
	}

	eValue, err := JsonMarshal(midResultMap)
	if err != nil {
		return "", err
	}

	return string(eValue), nil
}

// The default behavior is to escape &, <, and > to \u0026, \u003c, and \u003e
// to avoid certain safety problems that can arise when embedding JSON in HTML.
func JsonMarshal(value map[string]interface{}) (string, error) {
	buffer := &bytes.Buffer{}

	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(value)
	if err != nil {
		return "", err
	}

	result := strings.Trim(buffer.String(), "\n")

	return result, nil
}

func JsonUnMarshal(value string) (map[string]interface{}, error) {
	midResultMap := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(value), &midResultMap)
	if err != nil {
		return nil, err
	}

	return midResultMap, nil
}

func Set2List(s []interface{}) []string {
	l := make([]string, 0)
	for i := range s {
		if s[i] != "" {
			l = append(l, s[i].(string))
		}
	}
	return l
}
