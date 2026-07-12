package template

import (
	"bytes"
	"fmt"
	"maps"
	"text/template"

	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/internal/utils"
)

func Render(payload string, data map[string]any) (string, error) {
    t, err := template.New("").
		Option("missingkey=error").
		Parse(payload)

    if err != nil {
        return "", err
    }

    var buf bytes.Buffer

    err = t.Execute(&buf, data)

    if err != nil {
        return "", err
    }

    return buf.String(), nil
}

func Context(data ...map[string]any) map[string]any {
	ctx := map[string]any{}

	for _, d := range data {
		maps.Copy(ctx, d)
	}

	return ctx
}

func ContextConfig(config config.Config) map[string]any {
	ctx := map[string]any{}

	worksapce := map[string]any{}

	if wp, ok := config.Get("workspace.path"); ok {
		worksapce["path"] = wp
	}

	if wap, ok := config.Get("workspace.atlas_path"); ok {
		worksapce["atlas_path"] = wap
	}

	ctx["workspace"] = worksapce 

	return ctx
}

func ContextEntryInfo(info models.EntryInfo) map[string]any {
	ctx := map[string]any{}

	ei := map[string]string{
		"path": info.Path,
		"type": info.Type,
		"ext":  info.Ext,
		"basename": info.BaseName,
		"AbsolutePath": info.AbsolutePath,
	}

	ctx["entry_info"] = ei

	return ctx
}

func ParseArray(array []any, data map[string]any) ([]any, error) {
	flat, ok := utils.FlattenArray(array)

	if !ok {
		return nil, fmt.Errorf("failed to parse array")
	}

	sm := utils.StringifyMap(flat)

	if !ok {
		return nil, fmt.Errorf("failed to stringify map")
	}

	pm := map[string]string{}

	for k, v := range sm {
		rendered, err := Render(v, data)

		if err != nil {
			return nil, fmt.Errorf("failed to render template for key %s: %v", k, err)
		}

		pm[k] = rendered
	}

	pma := map[string]any{}

	for k, v := range pm {
		pma[k] = v
	}
	
	if !ok {
		return nil, fmt.Errorf("failed to convert map[string]string to map[string]any")
	}

	unflattened, ok := utils.UnflattenArray(pma)

	if !ok {
		return nil, fmt.Errorf("failed to unflatten array")
	}

	return unflattened, nil
}
