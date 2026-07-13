package content

import (
	"fmt"
	"strings"
)

type Transformer struct {
	Type    string                       `json:"type"`
	Options map[string]any               `json:"options"`
	Apply   func(string) (string, error) `json:"-"`
}

func ApplyRemoveTransformer(input string, options map[string]any) (string, error) {
	from, ok := options["from"].(string)

	if !ok {
		return "", fmt.Errorf("invalid start option for remove transformer")
	}

	to, ok := options["to"].(string)

	if !ok {
		return "", fmt.Errorf("invalid end option for remove transformer")
	}

	start := strings.Index(input, from)

	if start == -1 {
		return input, nil
	}

	search:= input[start+len(from):]
	end := strings.Index(search, to)

	if end == -1 {
		return input, nil
	}

	end += start + len(from)

	output := input[:start] + input[end+len(to):]

	return output, nil
}

func TrimTransformer(input string, options map[string]any) (string, error) {
	return strings.TrimSpace(input), nil
}

func CreateTransformer(payload map[string]any) (Transformer, error) {
	t := ""
	rest := map[string]any{}

	for k, v := range payload {
		if k == "type" {
			t, _ = v.(string)
		} else {
			rest[k] = v
		}
	}

	if t == "" {
		return Transformer{}, fmt.Errorf("transformer type is required")
	}

	if t == "remove" {
		tr := Transformer{
			Type:    t,
			Options: rest,
			Apply: func(input string) (string, error) {
				return ApplyRemoveTransformer(input, rest)
			},
		}

		return tr, nil
	}

	if t == "trim" {
		tr := Transformer{
			Type:    t,
			Options: rest,
			Apply: func(input string) (string, error) {
				return TrimTransformer(input, rest)
			},
		}

		return tr, nil
	}

	return Transformer{}, fmt.Errorf("unknown transformer type: %s", t)
}

func CreateTransformers(payload []any) ([]Transformer, error) {
	transformers := []Transformer{}

	for _, tr := range payload {
		if trMap, ok := tr.(map[string]any); ok {
			t, err := CreateTransformer(trMap)

			if err != nil {
				return nil, err
			}

			transformers = append(transformers, t)
		}
	}

	return transformers, nil
}
