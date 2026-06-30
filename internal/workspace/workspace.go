package workspace 

import (
	"os"
	"path/filepath"
)

var Path string

func Get() (string, error) {
    if Path != "" {
        return filepath.Abs(Path)
    }

    return os.Getwd()
}
