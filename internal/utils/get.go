package utils

import "strings"

func Get(m map[string]any, path string) any {
	parts := strings.Split(path, ".")

	var cur any = m

	for _, p := range parts {
		obj, ok := cur.(map[string]any)

		if !ok {
			return nil
		}

		cur, ok = obj[p]

		if !ok {
			return nil
		}
	}

	return cur
}
