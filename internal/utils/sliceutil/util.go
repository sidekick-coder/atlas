package sliceutil

func MapString(arr []any) []string {
	out := make([]string, 0, len(arr))

	for _, x := range arr {
		if s, ok := x.(string); ok {
			out = append(out, s)
		}
	}

	return out
}
