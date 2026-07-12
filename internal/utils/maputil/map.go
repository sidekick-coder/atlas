package maputil

func Any[K comparable, V any](src map[K]V) map[any]any {
	out := map[any]any{}

	for k, v := range src {
		out[k] = v
	}

	return out
}

func Pick(m map[string]any, keys ...string) map[string]any {
	out := make(map[string]any, len(keys))

	for _, k := range keys {
		if v, ok := m[k]; ok {
			out[k] = v
		}
	}

	return out
}

func Except[K comparable, V any](src map[K]V, exclude ...K) map[K]V {
	skip := make(map[K]struct{}, len(exclude))
	for _, k := range exclude {
		skip[k] = struct{}{}
	}

	dst := make(map[K]V, len(src))
	for k, v := range src {
		if _, ok := skip[k]; ok {
			continue
		}
		dst[k] = v
	}

	return dst
}
