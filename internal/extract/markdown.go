package extract

import (
	"strings"
	"github.com/sidekick-coder/atlas/internal/utils"

	"github.com/adrg/frontmatter"
)

func Markdown(content string, meta *map[string]any) error {
	mdFrontmatter := map[string]any{}

	_, err := frontmatter.Parse(strings.NewReader(content), mdFrontmatter)

	if err != nil {
		return err
	}

	frontmatterFlat := utils.FlattenMap(mdFrontmatter, "frontmatter")

	for k, v := range frontmatterFlat {
		(*meta)[k] = v
	}

	return nil
}
