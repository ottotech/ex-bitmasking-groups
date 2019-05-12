package config

import (
	"github.com/ottotech/ex-bitmasking-groups/pkg/groups"
	"html/template"
)

var TPL *template.Template

func init() {
	// be aware of duplicated template names, this is done this way for simplicity
	templateList := []string{
		"templates/add.gohtml",
		"templates/list.gohtml",
		"templates/detail.gohtml",
	}
	TPL = template.Must(template.New("").Funcs(groups.Fm).ParseFiles(templateList...))
}
