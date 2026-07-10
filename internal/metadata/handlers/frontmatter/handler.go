package frontmatter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/goccy/go-yaml"
	"github.com/sidekick-coder/atlas/internal/metadata/handler"
	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/internal/utils"
)

type Handler struct {
	id string
	prefix string
	options map[string]any
}

func Create(payload handler.Payload) handler.Handler {
	prefix := "frontmatter."

	if p, ok := payload.Options["prefix"]; ok { 
		prefix = p.(string)
	}

	return Handler{
		id: payload.ID,
		options: payload.Options,
		prefix: prefix,
	}
}

func(m Handler) GetID() string {
	return m.id
}

func(m Handler) GetTypeID() string {
	return "markdown"
}

func (h Handler) ExtractFromContent(content string) (string, map[string]any, error) {
	result := map[string]any{}

	bodyRaw, err := frontmatter.Parse(strings.NewReader(content), result)

	body := string(bodyRaw)

	if err != nil {
		return "", nil, err
	}

	flat := utils.FlattenMap(result, h.prefix)

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

func (h Handler) Extract(info *models.EntryInfo) (map[string]string, error) {
	contents, err := os.ReadFile(filepath.Join(info.AbsolutePath))

	if err != nil {
		return nil, err
	}

	_, flat, err := h.ExtractFromContent(string(contents))

	if err != nil {
		return nil, err
	}

	result := utils.StringifyMap(flat)

	return result, nil
}

func (h Handler) Set(info *models.EntryInfo, name string, value string) (bool, error) {
	isFrontmatterField := strings.HasPrefix(name, h.prefix)

	if !isFrontmatterField {
		return false, nil
	}

	contents, err := os.ReadFile(filepath.Join(info.AbsolutePath))

	data := string(contents)

	if err != nil {
		return false, err
	}

	body, metas, err := h.ExtractFromContent(data)

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

func (h Handler) Unset(info *models.EntryInfo, name string) error {
	isFrontmatterField := strings.HasPrefix(name, h.prefix)

	if !isFrontmatterField {
		return nil
	}

	contents, err := os.ReadFile(filepath.Join(info.AbsolutePath))

	data := string(contents)

	if err != nil {
		return err
	}

	body, metas, err := h.ExtractFromContent(data)

	if err != nil {
		return err
	}

	delete(metas, name)

	unflattened := utils.Unflatten(metas)

	newContents, err := Marshal(body, unflattened[h.prefix].(map[string]any))

	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(info.AbsolutePath), []byte(newContents), 0644)

	if err != nil {
		return err
	}

	return nil
}

