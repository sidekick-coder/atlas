package utils

import (
	"fmt"
	"strings"
)

func Flatten(input map[string]any, output map[string]any, prefix string) {
	for k, v := range input {
		key := k
		if prefix != "" {
			key = prefix + "." + k
		}

		switch val := v.(type) {
		case map[string]any:
			Flatten(val, output, key)

		case map[any]any:
			tmp := make(map[string]any)
			for k2, v2 := range val {
				if ks, ok := k2.(string); ok {
					tmp[ks] = v2
				}
			}
			Flatten(tmp, output, key)
		case []any:
			for i, v2 := range val {
				idxKey := fmt.Sprintf("%s[%d]", key, i)

				switch vv := v2.(type) {
				case map[string]any:
					Flatten(vv, output, idxKey)

				case map[any]any:
					tmp := make(map[string]any)
					for k2, v3 := range vv {
						if ks, ok := k2.(string); ok {
							tmp[ks] = v3
						}
					}
					Flatten(tmp, output, idxKey)

				default:
					output[idxKey] = fmt.Sprint(v2)
				}
			}
		default:
			output[key] = fmt.Sprint(val)
		}
	}
}

func FlattenMap(input map[string]any, prefix string) map[string]any {
	output := make(map[string]any)
	Flatten(input, output, prefix)
	return output
}

func FlattenArray(array []any, prefix ...string) (map[string]any, bool) {
	output := make(map[string]any)

	input := map[string]any{"array": array}

	Flatten(input, output, strings.Join(prefix, "."))

	result := make(map[string]any)

	for k, v := range output {
		nk := strings.TrimPrefix(k, "array")
		result[nk] = v
	}

	return result, true

}
