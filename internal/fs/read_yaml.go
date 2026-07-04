package fs

import (
	"github.com/goccy/go-yaml"
)

func ReadYaml(path string) (map[string]interface{}, error) {
	text, err := ReadText(path)

	if err != nil {
		return nil, err
	}

	var result map[string]interface{}

	err = yaml.Unmarshal([]byte(text), &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
