package setter

import (
	"github.com/sidekick-coder/atlas/internal/metadata/handler"
	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/internal/template"
	"github.com/sidekick-coder/atlas/internal/utils/maputil"
)

type Key struct {
	Name  string
	Value string
}

type Handler struct {
	id      string
	options map[string]any
	keys    []Key
}

func ParseKeys(payload any) []Key {
	arr, ok := payload.([]any)

	if !ok {
		return []Key{}
	}

	keys := []Key{}

	for _, item := range arr {
		obj, ok := item.(map[string]any)

		if !ok {
			continue
		}

		name, ok := obj["name"].(string)

		if !ok {
			continue
		}

		value, ok := obj["value"].(string)

		if !ok {
			continue
		}

		keys = append(keys, Key{
			Name:  name,
			Value: value,
		})
	}

	return keys
}

func Create(payload handler.Payload) handler.Handler {
	keys := ParseKeys(payload.Options["keys"])

	return Handler{
		id:      payload.ID,
		options: payload.Options,
		keys:    keys,
	}
}

func (m Handler) GetID() string {
	return m.id
}

func (m Handler) GetTypeID() string {
	return "setter"
}

func (m Handler) ID() string {
	return m.id
}

func (m Handler) Extract(payload handler.ExtractPayload) (map[string]string, error) {
	result := map[string]string{}
	ctx := maputil.Any(payload.Metas)

	for _, key := range m.keys {
		v, err := template.Render(key.Value, ctx, template.RenderOptions{
			AllowMissingKeys: true,
		})

		if err != nil {
			return nil, err
		}

		result[key.Name] = v
	}

	return result, nil
}

func (m Handler) Set(info *models.EntryInfo, name string, value string) (bool, error) {
	return true, nil
}

func (m Handler) Unset(info *models.EntryInfo, name string) error {
	return nil
}
