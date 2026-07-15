package drive

import (
	"path/filepath"

	"github.com/sidekick-coder/atlas/internal/fs"
)

var requiredIgnores = []string{
	".atlas",
}

var defaultIgnores = []string{
	"**/node_modules",
	"**/package-lock.json",
	"**/vendor",
	"**/.git",
	"**/.DS_Store",
	"$RECYCLE.BIN",
	"System Volume Information",
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

func (d *Drive) Ignore(path string, ignores ...string) (bool, error) {
	patterns := CreateIgnorePatterns(ignores...)

	patterns = append(patterns, d.config.GetArrayString("scan.exclude")...)

	exclude, err := fs.MatchAny(path, patterns)

	if err != nil {
		return false, err
	}

	if exclude {
		return true, nil
	}

	pi := d.config.GetArrayString("scan.include")

	if len(pi) == 0 {
		return false, nil
	}

	include, err := fs.MatchAny(path, pi)

	if err != nil {
		return false, err
	}

	if !include {
		return true, nil
	}

	return false, nil

}

