package utils 

import (
	"fmt"
)

func StringifyMap(m map[string]any) map[string]string {
	out := make(map[string]string, len(m))

	for k, v := range m {
		out[k] = fmt.Sprint(v)
	}

	return out
}

