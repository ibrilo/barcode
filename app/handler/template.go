package handler

import (
	"html/template"
	"regexp"

	"github.com/pkg/errors"
)

func loadTemplate(filename string) (*template.Template, error) {
	re := regexp.MustCompile(`(\w+).html`)
	var templateName string
	subs := re.FindAllStringSubmatch(filename, -1)[0]
	if len(re.FindStringIndex(filename)) > 0 {
		templateName = subs[1]
	}

	t, err := template.New(templateName).ParseFiles(filename)
	if err != nil {
		return nil, errors.Wrap(err, "parse template file template error")
	}

	return t, nil
}
