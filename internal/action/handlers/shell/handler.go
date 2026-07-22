package shell

import (
	"errors"
	"github.com/sidekick-coder/atlas/internal/template"
	"os/exec"
	"strings"
)

type Handler struct {}

func Create() Handler {
	return Handler{}
}

func (c Handler) Execute(ctx map[string]any) (map[string]any, error) {
	result := make(map[string]any)

	command, ok := ctx["command"].(string)

	if !ok || command == "" {
		return result, errors.New("command not specified in action options")
	}

	rendered, err := template.Render(command, ctx)

	if err != nil {
		return result, errors.New("failed to render command template: " + err.Error())
	}

	cmd := exec.Command("sh", "-c", rendered)

	output, err := cmd.CombinedOutput()

	if err != nil {
		return result, err
	}

	out := strings.TrimSpace(string(output))

	if len(out) > 0 {
		result["output"] = out
	}

	return result, nil
}
