package actionmanager

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"github.com/sidekick-coder/atlas/internal/template"
)

type ShellAction struct {
	Options map[string]string
}

func (c ShellAction) Execute(ctx ActionContext) error {
	command := c.Options["command"]

	if command == "" {
		return errors.New("command not specified in action options")
	}

	rendered, err := template.Render(command, ctx)

	if err != nil {
		return errors.New("failed to render command template: " + err.Error())
	}

	cmd := exec.Command("sh", "-c", rendered)

	output, err := cmd.CombinedOutput()

	if err != nil {
		return err
	}

	outputStr := strings.TrimSpace(string(output))

	if len(outputStr) > 0 {
		fmt.Println(outputStr)
	}

	return nil
}
