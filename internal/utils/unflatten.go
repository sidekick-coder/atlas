package utils

import (
	"strconv"
	"strings"
)

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// supports: a.b[0].c
func splitKey(key string) []string {
	key = strings.ReplaceAll(key, "]", "")
	key = strings.ReplaceAll(key, "[", ".")
	return strings.Split(key, ".")
}

func Unflatten(input map[string]any) map[string]any {
	output := make(map[string]any)

	for key, value := range input {
		parts := splitKey(key)
		current := output

		for i := 0; i < len(parts); i++ {
			part := parts[i]
			if part == "" {
				continue
			}

			last := i == len(parts)-1

			// ARRAY CASE
			if i+1 < len(parts) && isInt(parts[i+1]) {
				idx := atoi(parts[i+1])
				i++              // consume index
				last = i == len(parts)-1 // recompute after consuming index

				arr, ok := current[part]
				var slice []any

				if ok {
					if v, ok := arr.([]any); ok {
						slice = v
					}
				}

				for len(slice) <= idx {
					slice = append(slice, nil)
				}

				// FINAL VALUE (IMPORTANT FIX)
				if last {
					slice[idx] = value
					current[part] = slice
					break
				}

				// ensure container exists
				if slice[idx] == nil {
					slice[idx] = map[string]any{}
				}

				next, ok := slice[idx].(map[string]any)
				if !ok {
					next = map[string]any{}
					slice[idx] = next
				}

				current[part] = slice
				current = next
				continue
			}

			// FINAL KEY (NOT ARRAY)
			if last {
				current[part] = value
				continue
			}

			// MAP CASE
			nextRaw, ok := current[part]
			if !ok {
				nextRaw = map[string]any{}
				current[part] = nextRaw
			}

			next, ok := nextRaw.(map[string]any)
			if !ok {
				next = map[string]any{}
				current[part] = next
			}

			current = next
		}
	}

	return output
}
