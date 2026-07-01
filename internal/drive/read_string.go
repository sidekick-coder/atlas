package drive

import (
	"os"
	"path/filepath"
)

func ReadAsString(root string, path ...string) (error, string) {
	absolutePath := filepath.Join(append([]string{root}, path...)...)

	data, err := os.ReadFile(absolutePath)

	if err != nil {
		return err, ""
	}

	return nil, string(data)
}
