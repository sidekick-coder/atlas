package utils

import "strconv"

func ParseInt(value any) (int, bool) {
	if v, ok := value.(int); ok {
		return v, true
	}

	if v, ok := value.(float64); ok {
		return int(v), true
	}

	if s, ok := value.(string); ok {
		v, ok := strconv.Atoi(s)

		if ok == nil {
			return v, true
		}
	}

	return 0, false
}
