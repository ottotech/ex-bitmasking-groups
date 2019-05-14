package main

import (
	"github.com/ottotech/ex-bitmasking-groups/pkg/adding"
	"github.com/ottotech/ex-bitmasking-groups/pkg/deleting"
	"github.com/ottotech/ex-bitmasking-groups/pkg/groups"
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
	var deleter deleting.Service

	switch storageType {
	case Memory:
		s := new(memory.Storage)
		adder = adding.NewService(s)
		lister = listing.NewService(s)
		deleter = deleting.NewService(s)
	}

	app := new(rest.App)
	mux := http.NewServeMux()
	mux.Handle("/", app.UserList.Handler(lister))
	mux.Handle("/add", app.AddUser.Handler(adder))
	mux.Handle("/get/", app.GetUser.Handler(lister))
	mux.Handle("/delete/", app.DeleteUser.Handler(deleter))
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/dummy", groups.Wrapper(app.Dummy.Handler(), groups.GroupA))
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
