package template

import (
	"bytes"
	"fmt"
	"maps"
	"strings"
	"text/template"

	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/internal/utils"
)

type RenderOptions struct {
	AllowMissingKeys bool
}

func Eval(payload string, data map[string]any) (any, error) {
	input := payload

	// if starts with ={{var}}, return the value
	if len(input) > 0 && input[0] == '=' {
		key := strings.TrimPrefix(input, "={{")
		key = strings.TrimSuffix(key, "}}")
		key = strings.TrimSpace(key)
		key = strings.TrimPrefix(key, ".")

		return utils.Get(data, key), nil
	}

	t, err := template.New("").
		Option("missingkey=error").
		Parse(payload)

	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	err = t.Execute(&buf, data)

	if err != nil {
		return nil, err
	}

	return buf.String(), nil
}

func Render(payload string, data map[string]any, args ...RenderOptions) (string, error) {
	opts := RenderOptions{
		AllowMissingKeys: false,
	}

	if len(args) > 0 {
		opts = args[0]
	}

	t := template.New("render")

	if !opts.AllowMissingKeys {
		t = t.Option("missingkey=error")
	}

	t, err := t.Parse(payload)

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
		"path":         info.Path,
		"type":         info.Type,
		"ext":          info.Ext,
		"basename":     info.BaseName,
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

	pm := map[string]any{}

	for k, v := range sm {
		rendered, err := Eval(v, data)

		if err != nil {
			return nil, fmt.Errorf("failed to render template for key %s: %v", k, err)
		}

		pm[k] = rendered
	}

	if !ok {
		return nil, fmt.Errorf("failed to convert map[string]string to map[string]any")
	}

	unflattened, ok := utils.UnflattenArray(pm)

	if !ok {
		return nil, fmt.Errorf("failed to unflatten array")
	}

	return unflattened, nil
}
