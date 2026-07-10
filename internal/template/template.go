package template

import (
	"bytes"
	"maps"
	"text/template"

	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/models"
)

func Render(s string, data map[string]any) (string, error) {
    t, err := template.New("").Parse(s)

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
