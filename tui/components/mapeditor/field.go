package mapeditor

import "fmt"

type Field struct {
	Label    string
	FielName string
}

func CreateFieldFromMap(payload any) (Field, error) {
	f := Field{}

	m, ok := payload.(map[string]any)

	if !ok {
		return f, fmt.Errorf("invalid field type: %T", payload)
	}

	if label, ok := m["label"].(string); ok {
		f.Label = label
	}

	if name, ok := m["name"].(string); ok {
		f.FielName = name
	}

	return f, nil
}

func CreateFieldsFromArray(payload any) ([]Field, error) {
	fields := []Field{}

	array, ok := payload.([]any)

	if !ok {
		return fields, fmt.Errorf("invalid fields type: %T", payload)
	}

	for _, item := range array {
		field, err := CreateFieldFromMap(item)

		if err != nil {
			return fields, err
		}

		fields = append(fields, field)
	}

	return fields, nil
}
