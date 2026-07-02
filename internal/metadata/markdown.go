package metadata 

import (
	"github.com/sidekick-coder/atlas/internal/drive/v2"
	"strings"
	"github.com/sidekick-coder/atlas/internal/utils"
	"github.com/adrg/frontmatter"
	"os"
	"path/filepath"
)

type MarkdownHandler struct {}

func (m MarkdownHandler) ID() string {
	return "markdown"
}

func (m MarkdownHandler) Extract(info *drive.EntryInfo) (map[string]string, error) {
	contents, err := os.ReadFile(filepath.Join(info.AbsolutePath))

	data := string(contents)

	if err != nil {
		return nil, err
	}

	result := map[string]any{}

	_, err = frontmatter.Parse(strings.NewReader(data), result)

	if err != nil {
		return nil, err
	}

	flat := utils.FlattenMap(result, "frontmatter")

	return utils.StringifyMap(flat), nil
}

func (m MarkdownHandler) Set(info *drive.EntryInfo, name string, value string) error {
	// Implement logic to set metadata in a Markdown file
	// For example, you could modify the front matter or other metadata formats
	return nil
}

