package maputil

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

func GetString(m map[string]any, path string) (string, bool) {
	v := Get(m, path)

	if v == nil {
		return "", false
	}

	s, ok := v.(string)

	if !ok {
		return "", false
	}

	return s, true
}

func GetMap(m map[string]any, path string) (map[string]any, bool) {
	v := Get(m, path)

	if v == nil {
		return nil, false
	}

	mm, ok := v.(map[string]any)

	if !ok {
		return nil, false
	}

	return mm, true
}
