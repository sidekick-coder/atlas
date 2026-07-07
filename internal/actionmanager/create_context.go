package actionmanager

import (
	"maps"
)

func (am *ActionManager) CreateContext(ctx ...map[string]any) map[string]any {
	result := make(map[string]any)

	result["WorkspacePath"] = am.config.Get("workspace.path")
	result["AtlasPath"] = am.config.Get("workspace.atlas_path")
	
	for _, c := range ctx {
		maps.Copy(result, c)
	}

	return result
}
