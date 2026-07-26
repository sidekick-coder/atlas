package shell

import (
	"fmt"
	"maps"
	"os"
	"os/exec"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/template"
	"github.com/sidekick-coder/atlas/tui/app/program"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/components/viewport"
)

type Component struct {
	props    map[string]any
	viewport *viewport.Component
}

func Create() *Component {
	return &Component{
		viewport: viewport.Create(),
	}
}

func (c *Component) Init() tea.Cmd {
	return c.Load()
}

func (c *Component) Dispose() tea.Cmd {
	return nil
}

func (c *Component) Activate() tea.Cmd {
	return c.LoadBindings()
}

func (c *Component) Deactivate() tea.Cmd {
	return c.UnloadBindings()
}

func (c *Component) SetProps(props map[string]any) {
	c.props = props

	c.Load()
}

func Execute(options map[string]any) (string, error) {
	bin := ""
	args := []string{}
	command := ""

	if b, ok := options["bin"]; ok {
		bin = b.(string)
	}

	if a, ok := options["args"]; ok {
		args = a.([]string)
	}

	if c, ok := options["command"]; ok {
		command = c.(string)
	}

	if bin != "" {
		cmd := exec.Command(bin, args...)
		output, err := cmd.CombinedOutput()

		if err != nil {
			return "", err
		}

		return string(output), nil
	}

	if command != "" {
		cmd := exec.Command("sh", "-c", command)
		cmd.Env = append(os.Environ(), "COLOR=always")
		output, err := cmd.CombinedOutput()

		outputStr := string(output)

		if err != nil {
			return "", fmt.Errorf("command execution failed: %w, output: %s", err, outputStr)
		}

		return string(output), nil
	}

	return "", nil
}

func (c *Component) Load() tea.Cmd {
	app := program.GetApp()
	config := app.Config()

	ctx :=  map[string]any{}

	ctx["workspace"] = config.GetMap("workspace")

	maps.Copy(ctx, c.props)

	computed, err := template.EvaluateMap(c.props, ctx)

	if err != nil {
		return toast.Error(err.Error())
	}

	output, err := Execute(computed)

	if err != nil {
		return toast.Error(err.Error())
	}

	content := string(output)

	c.viewport.SetContent(content)

	return nil
}
