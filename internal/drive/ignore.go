package drive

import (
	"path/filepath"
)

var requiredIgnores = []string{
	".atlas",
}

var defaultIgnores = []string{
	"node_modules",
	"vendor",
	".git",
	".DS_Store",
}

func CreateIgnorePatterns(ignores ...string) []string {
	patterns := []string{}

	if len(ignores) > 0 {
		patterns = append(patterns, ignores...)
	}

	if len(patterns) == 0 {
		patterns = append(patterns, defaultIgnores...)
	}

	patterns = append(patterns, requiredIgnores...)

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
