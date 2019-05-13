package rest

import (
	"fmt"
	_ "github.com/ottotech/ex-bitmasking-groups/pkg/config"
	"github.com/ottotech/ex-bitmasking-groups/pkg/listing"
	"github.com/ottotech/ex-bitmasking-groups/pkg/storage/memory"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserList(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080", nil)
	req.URL.Path = "/"
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	app := new(App)
	rec := httptest.NewRecorder()
	storage := new(memory.Storage)
	var lister listing.Service
	lister = listing.NewService(storage)
	app.UserList.Handler(lister).ServeHTTP(rec, req)

	res := rec.Result()
	fmt.Println(res.StatusCode)

}
