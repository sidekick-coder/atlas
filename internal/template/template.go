package template

import (
	"text/template"
	"bytes"
)

func Render(s string, data map[string]any) (string, error) {
    t, err := template.New("").Parse(s)

    if err != nil {
        return "", err
    }

    var buf bytes.Buffer

    err = t.Execute(&buf, data)

    if err != nil {
        return "", err
    }

    return buf.String(), nil
}
