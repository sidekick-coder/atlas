package maputil

import (
	"fmt"
	"slices"
	"strings"
)

func Any[K comparable, V any](src map[K]V) map[K]any {
	out := map[K]any{}

	for k, v := range src {
		out[k] = v
	}

	return out
}

func KeyString[K comparable, V any](src map[K]V) map[string]any {
	out := map[string]any{}

	for k, v := range src {
		out[fmt.Sprintf("%v", k)] = v
	}

	return out
}

func String(src map[string]any) map[string]string {
	out := map[string]string{}

	for k, v := range src {
		out[k] = fmt.Sprintf("%v", v)
	}

	return out
}

func FromString(source string) map[string]any {
	out := map[string]any{}

	// Split the source string into key-value pairs
	pairs := slices.Collect(strings.SplitSeq(source, ";"))

	for _, pair := range pairs {
		// Split each pair into key and value
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) == 2 {
			key := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			out[key] = value
		}
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
