package config

import (
	"flag"
	"github.com/ottotech/ex-bitmasking-groups/pkg/groups"
	"html/template"
	"log"
)

var TPL *template.Template

func init() {
	// be aware of duplicated template names, this is done this way for simplicity
	templateList := []string{
		"templates/add.gohtml",
		"templates/list.gohtml",
		"templates/detail.gohtml",
	}
	var err error
	TPL, err = template.New("").Funcs(groups.Fm).ParseFiles(templateList...)
	if err != nil {
		// TODO: research if this is a robust approach to check if the test pkg has been called
		// TODO: research why templates don't work when testing
		if flag.Lookup("test.v") != nil {
			log.Println(err)
		} else {
			log.Fatal(err)
		}
	}
}
