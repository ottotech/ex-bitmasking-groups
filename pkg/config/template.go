package config

import (
	"flag"
	"github.com/ottotech/ex-bitmasking-groups/pkg/groups"
	"html/template"
	"log"
	"os"
)

var TPL *template.Template

func init() {
	// be aware of duplicated template names, this is done this way for simplicity
	templateList := []string{
		"templates/add.gohtml",
		"templates/list.gohtml",
		"templates/detail.gohtml",
	}

	// when testing we need to change the working directory to the root app
	if flag.Lookup("test.v") != nil {
		if err := os.Chdir("../../.."); err != nil {
			log.Fatal(err)
		}
	}
	TPL = template.Must(template.New("").Funcs(groups.Fm).ParseFiles(templateList...))
}
