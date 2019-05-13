package rest

import (
	_ "github.com/ottotech/ex-bitmasking-groups/pkg/config"
	"github.com/ottotech/ex-bitmasking-groups/pkg/listing"
	"github.com/ottotech/ex-bitmasking-groups/pkg/storage/memory"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserList(t *testing.T) {
	// request 1
	r1, err := http.NewRequest("GET", "localhost:8080", nil)
	r1.URL.Path = "/"
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	// request 2
	r2, err := http.NewRequest("POST", "localhost:8080", nil)
	r2.URL.Path = "/url-not-defined"
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// setup
	app := new(App)
	storage := new(memory.Storage)
	var lister listing.Service
	lister = listing.NewService(storage)

	// make requests
	rec1 := httptest.NewRecorder()
	rec2 := httptest.NewRecorder()
	app.UserList.Handler(lister).ServeHTTP(rec1, r1)
	app.UserList.Handler(lister).ServeHTTP(rec2, r2)
	res1 := rec1.Result()
	res2 := rec2.Result()

	// tests
	if res1.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res1.StatusCode)
	}

	if res2.StatusCode != http.StatusNotFound {
		t.Errorf("expected status 404; got %v", res2.StatusCode)
	}

}
