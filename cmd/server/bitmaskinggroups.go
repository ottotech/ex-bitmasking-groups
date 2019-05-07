package main

import (
	"github.com/ottotech/ex-bitmasking-groups/pkg/adding"
	"github.com/ottotech/ex-bitmasking-groups/pkg/http/rest"
	"github.com/ottotech/ex-bitmasking-groups/pkg/listing"
	"github.com/ottotech/ex-bitmasking-groups/pkg/storage/memory"
	"log"
	"net/http"
)

const (
	Memory int = 1
)

func main() {
	// setup storage
	storageType := Memory

	var adder adding.Service
	var lister listing.Service

	switch storageType {
	case Memory:
		s := new(memory.Storage)
		adder = adding.NewService(s)
		lister = listing.NewService(s)
	}

	app := new(rest.App)
	mux := http.NewServeMux()
	mux.Handle("/", app.UserList.Handler(lister))
	mux.Handle("/add", app.AddUser.Handler(adder))
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
