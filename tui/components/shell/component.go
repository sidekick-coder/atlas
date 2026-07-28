package shell

import (
	"fmt"
	"os"
	"os/exec"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/internal/template"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/components/viewport"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/context"
)

type Component struct {
	props    map[string]any
	ctx      *context.Feature
	viewport *viewport.Component
}

func Create() *Component {
	ctx := context.Create()

	ctx.SetLabel("shell")

	return &Component{
		viewport: viewport.Create(),
		ctx:      ctx,
	}
}

func (c *Component) Init() tea.Cmd {
	return chain.Init(
		c.Load,
		c.ctx.Init,
	)
}

func (c *Component) Dispose() tea.Cmd {
	return chain.Dispose(
		c.UnloadBindings,
		c.ctx.Dispose,
	)
}

func (c *Component) Activate() tea.Cmd {
	return chain.Cmd(
		c.LoadBindings,
		c.ctx.Activate,
	)
}

func (c *Component) Deactivate() tea.Cmd {
	return chain.Cmd(
		c.UnloadBindings,
		c.ctx.Deactivate,
	)
}

func (c *Component) Context() *context.Feature {
	return c.ctx
}

func (c *Component) SetProps(props map[string]any) {
	c.props = props

	c.ctx.SetAll(props)

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
	ctx := c.ctx.GetEntriesMap()

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
