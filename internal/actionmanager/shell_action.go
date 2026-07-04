package actionmanager

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"text/template"
	"bytes"
)

type ShellAction struct {
	Options map[string]string
}

func (c ShellAction) Execute(params []string) error {
	command := c.Options["command"]

	if command == "" {
		return errors.New("command not specified in action options")
	}

	ctx := make(map[string]any)

	ctx["Params"] = params

	t := template.Must(template.New("command").Parse(command))

	var out bytes.Buffer 

	t.Execute(&out, ctx)

	command = out.String()

	cmd := exec.Command("sh", "-c", command)

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
