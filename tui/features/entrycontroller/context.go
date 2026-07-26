package entrycontroller

import (
	"path/filepath"

	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/internal/utils/maputil"
	"github.com/sidekick-coder/atlas/tui/app/program"
)

func CreateContext(e models.Entry) map[string]any {
	config := program.GetConfig()

	em := e.ToMap()

	em["absolute_path"] = filepath.Join(config.GetWorkspaceDir(), e.Path)

	em["update"] = func(payload map[string]any) {
		msg := UpdateMsg{
			Path:   e.Path,
			Values: maputil.String(payload),
		}

		program.Send(msg)
	}

	return em
}
