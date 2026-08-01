package field

import (
	"fmt"

	"github.com/sidekick-coder/atlas/internal/utils/maputil"
)

type Data struct {
	Type  string
	Name  string
	Options map[string]any
}

func CreateData(payload any) (Data, error) {
	f := Data{}

	m, ok := payload.(map[string]any)

	if !ok {
		return f, fmt.Errorf("invalid field type: %T", payload)
	}

	if name, ok := m["name"].(string); ok {
		f.Name = name
	}

	if t, ok := m["type"].(string); ok {
		f.Type = t
	}

	f.Options = maputil.Except(m, "type", "name")

	return f, nil
}

func CreateDataList(payload any) ([]Data, error) {
	fields := []Data{}

	list, ok := payload.([]any)

	if !ok {
		return fields, fmt.Errorf("invalid field list type: %T", payload)
	}

	for _, item := range list {
		field, err := CreateData(item)

		if err != nil {
			return fields, err
		}

		fields = append(fields, field)
	}

	return fields, nil
}
