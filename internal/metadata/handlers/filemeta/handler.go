package filemeta

import (
	"fmt"
	"maps"

	"github.com/goccy/go-yaml"
	"github.com/sidekick-coder/atlas/internal/fs"
	"github.com/sidekick-coder/atlas/internal/metadata/handler"
	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/internal/template"
	"github.com/sidekick-coder/atlas/internal/utils"
	"github.com/sidekick-coder/atlas/internal/utils/maputil"
)


type Handler struct {
	id      string
	options map[string]any
	filename string 
	format string
}

func Create(payload handler.Payload) handler.Handler {
	h := Handler{
		id:      payload.ID,
		options: payload.Options,
		filename: "{{if eq .type \"directory\"}}{{.absolute_dirname}}/{{.basename}}.metas.yml{{else}}{{.absolute_path}}.metas.yml{{end}}",
		format: "yml",
	}

	if f, ok := payload.Options["filename"]; ok  {
		h.filename = f.(string)
	}

	return h
}

func (m Handler) GetID() string {
	return m.id
}

func (m Handler) GetTypeID() string {
	return "filemeta"
}

func (m Handler) ID() string {
	return m.id
}

func UnmarshalYml(yml string) (map[string]any, error) {
	result := map[string]any{}

	err := yaml.Unmarshal([]byte(yml), &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func MarshalYml(metas map[string]any) (string, error) {
	result := ""

	yamlBytes, err := yaml.Marshal(metas)

	if err != nil {
		return "", err
	}

	result = string(yamlBytes)

	return result, nil
}

func (m Handler) Marshal(metas map[string]any) (string, error) {
	if m.format == "yml" {
		return MarshalYml(metas)
	}

	return "", fmt.Errorf("setter handler %s: unsupported format %s", m.id, m.format)
}

func (m Handler) Unmarshal(content string) (map[string]any, error) {
	if m.format == "yml" {
		return UnmarshalYml(content)
	}

	return nil, fmt.Errorf("setter handler %s: unsupported format %s", m.id, m.format)
}

func (m Handler) Extract(payload handler.ExtractPayload) (map[string]string, error) {
	ctx := maputil.Any(payload.Metas)
	result := map[string]string{}


	filename, err := template.Render(m.filename, ctx)

	if err != nil {
		return nil, err
	}

	result["filemeta.filename"] = filename

	if !fs.Exists(filename) {
		return result, nil
	}

	content, err := fs.ReadText(filename)

	if err != nil {
		return result, err
	}

	metas, err := m.Unmarshal(content)

	if err != nil {
		return result, err
	}

	flat := utils.FlattenMap(metas, "")

	maps.Copy(result, maputil.String(flat))

	return result, nil
}

func (m Handler) Set(info *models.EntryInfo, name string, value string) (bool, error) {
	return false, nil
}

func (m Handler) Unset(info *models.EntryInfo, name string) error {
	return nil
}
