package metadata 

import (
	"os"
	"path/filepath"
	"strings"
	"fmt"
	"github.com/goccy/go-yaml"
	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/internal/utils"
	"github.com/adrg/frontmatter"
)

type MarkdownHandler struct {}

func ExtractFromContent(content string) (string, map[string]any, error) {
	result := map[string]any{}

	bodyRaw, err := frontmatter.Parse(strings.NewReader(content), result)

	body := string(bodyRaw)

	if err != nil {
		return "", nil, err
	}

	flat := utils.FlattenMap(result, "frontmatter")

	return body, flat, nil
}

func Marshal(body string, frontmatter map[string]any) (string, error) {
	result := ""

	yamlBytes, err := yaml.Marshal(frontmatter)

	if err != nil {
		return "", err
	}

	result = fmt.Sprintf("---\n%s---\n%s", string(yamlBytes), body)

	return result, nil
}

func (m MarkdownHandler) ID() string {
	return "markdown"
}

func (m MarkdownHandler) Extract(info *models.EntryInfo) (map[string]string, error) {
	contents, err := os.ReadFile(filepath.Join(info.AbsolutePath))

	if err != nil {
		return nil, err
	}

	body, flat, err := ExtractFromContent(string(contents))

	if err != nil {
		return nil, err
	}

	result := utils.StringifyMap(flat)
	
	result["body"] = body

	return result, nil
}

func (m MarkdownHandler) Set(info *models.EntryInfo, name string, value string) (bool, error) {
	isFrontmatterField := strings.HasPrefix(name, "frontmatter.")

	if !isFrontmatterField {
		return false, nil
	}

	contents, err := os.ReadFile(filepath.Join(info.AbsolutePath))

	data := string(contents)

	if err != nil {
		return false, err
	}

	body, metas, err := ExtractFromContent(data)

	if err != nil {
		return false, err
	}

	metas[name] = value

	unflattened := utils.Unflatten(metas)

	newContents, err := Marshal(body, unflattened["frontmatter"].(map[string]any))

	if err != nil {
		return false, err
	}

	err = os.WriteFile(filepath.Join(info.AbsolutePath), []byte(newContents), 0644)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (m MarkdownHandler) Unset(info *models.EntryInfo, name string) error {
	isFrontmatterField := strings.HasPrefix(name, "frontmatter.")

	if !isFrontmatterField {
		return nil
	}

	contents, err := os.ReadFile(filepath.Join(info.AbsolutePath))

	data := string(contents)

	if err != nil {
		return err
	}

	body, metas, err := ExtractFromContent(data)

	if err != nil {
		return err
	}

	delete(metas, name)

	unflattened := utils.Unflatten(metas)

	newContents, err := Marshal(body, unflattened["frontmatter"].(map[string]any))

	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(info.AbsolutePath), []byte(newContents), 0644)

	if err != nil {
		return err
	}

	return nil
}
