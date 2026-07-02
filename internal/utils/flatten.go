package utils 

import (
	"fmt"
)

func Flatten( input map[string]any, output map[string]any, prefix string) {
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
