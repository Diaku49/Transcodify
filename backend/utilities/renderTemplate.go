package utilities

import (
	"bytes"
	"html/template"
)

func RenderTemplate(template *template.Template, data interface{}) (string, error) {
	buf := bytes.NewBufferString("")
	if err := template.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
