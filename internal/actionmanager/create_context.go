package actionmanager

import (
	"maps"
)

func (am *ActionManager) CreateContext(ctx ...map[string]any) map[string]any {
	result := make(map[string]any)

	wp, ok := am.config.Get("workspace.path")

	if ok {
		result["WorkspacePath"] = wp
	}

	ap, ok := am.config.Get("workspace.atlas_path")

	if ok {
		result["AtlasPath"] = ap
	}
	
	for _, c := range ctx {
		maps.Copy(result, c)
	}

	return result
}
