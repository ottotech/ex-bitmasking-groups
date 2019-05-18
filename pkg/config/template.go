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
	// when testing we need to change the working directory to the root app
	if flag.Lookup("test.v") != nil {
		if err := os.Chdir("../../.."); err != nil {
			log.Fatal(err)
		}
	}

	// be aware of duplicated template names, this is done this way for simplicity
	TPL = template.Must(template.New("").Funcs(groups.Fm).ParseGlob("templates/*.gohtml"))
}
