package stat

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/sidekick-coder/atlas/internal/metadata/handler"
	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/internal/utils"
)

type Handler struct {
	id      string
	key     string
	options map[string]any
}

func Create(payload handler.Payload) Handler {
	key := "content"

	if k, ok := payload.Options["key"]; ok {
		key = k.(string)
	}

	return Handler{
		id:      payload.ID,
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
	return "json"
}

func (m Handler) Extract(info *models.EntryInfo) (map[string]string, error) {
	result := map[string]string{}

	result["basename"] = info.BaseName
	result["type"] = info.Type
	result["path"] = info.Path

	if info.Type == "file" {
		result["ext"] = strings.TrimPrefix(filepath.Ext(info.BaseName), ".")
	}

	parent := filepath.Dir(filepath.Clean(info.Path))

	if parent != "." {
		result["parent"] = parent
	}

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
