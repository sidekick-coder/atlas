package content

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/sidekick-coder/atlas/internal/metadata/handler"
	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/internal/utils"
)


type Handler struct {
	id           string
	key          string
	transformers []Transformer
	options      map[string]any
}


func Create(payload handler.Payload) handler.Handler {
	key := "content"
	transformers := []Transformer{}

	if k, ok := payload.Options["key"]; ok {
		key = k.(string)
	}

	if t, ok := payload.Options["transformers"].([]any); ok {
		trs, err := CreateTransformers(t)

		if err == nil {
			transformers = trs
		}
	}

	return Handler{
		id:      payload.ID,
		transformers: transformers,
		options: payload.Options,
		key:     key,
	}
}

func (m Handler) GetID() string {
	return m.id
}

func (m Handler) GetTypeID() string {
	return "content"
}

func UnmarshalFromBytes(content []byte) (map[string]any, error) {
	result := map[string]any{}

	err := json.Unmarshal(content, &result)

	if err != nil {
		return nil, err
	}

	flat := utils.FlattenMap(result, "")

	return flat, nil
}

func MarshalToBytes(metas map[string]any) ([]byte, error) {
	data, err := json.MarshalIndent(metas, "", "  ")

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m Handler) ID() string {
	return "content"
}

func (m Handler) Extract(info *models.EntryInfo) (map[string]string, error) {
	contents, err := os.ReadFile(filepath.Join(info.AbsolutePath))

	if err != nil {
		return nil, err
	}

	content := string(contents)

	for _, t := range m.transformers {
		out, err := t.Apply(content)

		if err != nil {
			return nil, err
		}

		content = out
	}

	result := map[string]string{}
	result[m.key] = content

	return result, nil
}

func (m Handler) Set(info *models.EntryInfo, name string, value string) (bool, error) {
	if name != m.key {
		return false, nil
	}

	err := os.WriteFile(filepath.Join(info.AbsolutePath), []byte(value), 0644)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (m Handler) Unset(info *models.EntryInfo, name string) error {
	return nil
}
