package shell

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/sidekick-coder/atlas/internal/config"
	"github.com/sidekick-coder/atlas/internal/metadata/handler"
	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/internal/template"
)

type Handler struct {
	id      string
	bin     string
	args    []string
	config  *config.Config
	options map[string]any
}

type ShellInputWorkspace struct {
	Path      string `json:"path"`
	AtlasPath string `json:"atlas_path"`
}

type ShellOutput struct {
	Metas map[string]string `json:"metas"`
}

func Create(p handler.Payload) handler.Handler {
	bin := ""
	args := []string{}

	if b, ok := p.Options["bin"]; ok {
		bin = b.(string)
	}

	if a, ok := p.Options["args"]; ok {
		args = a.([]string)
	}

	return Handler{
		id:      p.ID,
		options: p.Options,
		config:  p.Config,
		bin:     bin,
		args:    args,
	}
}

func (m Handler) GetID() string {
	return m.id
}

func (m Handler) GetTypeID() string {
	return "content"
}

func (m Handler) ID() string {
	return m.id
}

func (h Handler) Extract(info *models.EntryInfo) (map[string]string, error) {
	ctx := template.Context(
		template.ContextConfig(*h.config),
		template.ContextEntryInfo(*info),
	)

	bin, err := template.Render(h.bin, ctx)
	args := []string{}

	for _, arg := range h.args {
		renderedArg, err := template.Render(arg, ctx)

		if err != nil {
			return nil, err
		}

		args = append(args, renderedArg)
	}


	pb, err := json.Marshal(ctx)

	if err != nil {
		return nil, err
	}

	cmd := exec.Command(bin, args...)
	cmd.Stdin = strings.NewReader(string(pb))

	ob, err := cmd.CombinedOutput()

	if err != nil {
		return nil, fmt.Errorf("failed to execute command: %s, error: %v", err, string(ob))
	}

	os := strings.TrimSpace(string(ob))

	oj := ShellOutput{}

	err = json.Unmarshal([]byte(os), &oj)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal output: %s, error: %v", os, err)
	}

	return oj.Metas, nil

}

func (m Handler) Set(info *models.EntryInfo, name string, value string) (bool, error) {
	return false, nil
}

func (m Handler) Unset(info *models.EntryInfo, name string) error {
	return nil
}
