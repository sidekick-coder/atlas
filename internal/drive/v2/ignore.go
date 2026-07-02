package drive

import (
	"path/filepath"
)

var requiredIgnores = []string{
	".atlas",
}

func CreateIgnorePatterns(ignores ...string) []string {
	patterns := append(requiredIgnores, ignores...)

	return patterns
}

func ShouldIgnore(path string, patterns []string) bool {
	for _, pattern := range patterns {
		ok, err := filepath.Match(pattern, path)

		if err == nil && ok {
			return true
		}
	}
	return false
}
