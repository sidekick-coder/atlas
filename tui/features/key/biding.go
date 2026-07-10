package key

import (
	"strings"

	"github.com/sidekick-coder/atlas/internal/utils"
)

type Binding struct {
	id string
	keys []BindingKey
	tags []string
	help string
	desc string
}

type BindingKey struct {
	tokens []string
}

func normalizeToken(t string) string {
	t = strings.ToLower(t)

	switch t {
	case "cr":
		return "enter"
	case "bs":
		return "backspace"
	case "space":
		return "space"
	case "esc":
		return "esc"
	case "tab":
		return "tab"
	case "leader":
		return "<leader>"
	}

	// <C-s> -> ctrl+s
	if strings.HasPrefix(t, "c-") {
		return "ctrl+" + t[2:]
	}

	// <A-x> -> alt+x
	if strings.HasPrefix(t, "a-") || strings.HasPrefix(t, "m-") {
		return "alt+" + t[2:]
	}

	// <S-tab> -> shift+tab
	if strings.HasPrefix(t, "s-") {
		return "shift+" + t[2:]
	}

	return t
}

func parse(s string) []string {
	var out []string

	for len(s) > 0 {
		if strings.HasPrefix(s, "<leader>") {
			out = append(out, "<leader>")
			s = s[len("<leader>"):]
			continue
		}

		if s[0] == '<' {
			end := strings.IndexByte(s, '>')

			if end > 0 {
				token := normalizeToken(s[1:end])
				out = append(out, token)
				s = s[end+1:]
				continue
			}
		}

		out = append(out, string(s[0]))
		s = s[1:]
	}

	return out
}

func CreateBinding(keys ...string) Binding {
	bkeys := []BindingKey{}

	for _, k := range keys {
		tokens := parse(k)

		bk := BindingKey{
			tokens: tokens,
		}

		bkeys = append(bkeys, bk)

	}

	id, err := utils.CreateID()

	if err != nil {
		panic(err)
	}

	return Binding{
		id: id,
		keys: bkeys,
		tags: []string{},
	}
}

func (b Binding) SetHelp(help string) Binding {
	b.help = help
	return b
}

func (b Binding) SetDescription(desc string) Binding {
	b.desc = desc
	return b
}

func (b Binding) GetID() string {
	return b.id
}

func (b Binding) SetTags(tags ...string) Binding {
	b.tags = tags
	return b
}

func (b Binding) GetHelp() string {
	return b.help
}

func (b Binding) GetDescription() string {
	return b.desc
}

func (b Binding) GetTags() []string {
	return b.tags
}
