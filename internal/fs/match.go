package fs 

import (
	"github.com/bmatcuk/doublestar/v4"
)

func Match(path string, pattern string) (bool, error) {
	matched, err := doublestar.Match(pattern, path)

	if err != nil {
		return false, err
	}

	return matched, nil
}

func MatchAny(path string, patterns []string) (bool, error) {
	for _, pattern := range patterns {
		matched, err := Match(path, pattern)

		if err != nil {
			return false, err
		}

		if matched {
			return true, nil
		}
	}

	return false, nil
}
