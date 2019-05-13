package rest

import (
	"github.com/ottotech/ex-bitmasking-groups/pkg/adding"
	_ "github.com/ottotech/ex-bitmasking-groups/pkg/config"
	"github.com/ottotech/ex-bitmasking-groups/pkg/groups"
	"github.com/ottotech/ex-bitmasking-groups/pkg/listing"
	"github.com/ottotech/ex-bitmasking-groups/pkg/storage/memory"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestUserList_Handler(t *testing.T) {
	// request 1
	r1, err := http.NewRequest("GET", "localhost:8080", nil)
	r1.URL.Path = "/"
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	// request 2
	r2, err := http.NewRequest("GET", "localhost:8080", nil)
	r2.URL.Path = "/url-not-defined"
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// setup
	app := new(App)
	storage := new(memory.Storage)
	var lister listing.Service
	lister = listing.NewService(storage)
	rec1 := httptest.NewRecorder()
	rec2 := httptest.NewRecorder()

	// make requests
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

func TestAddUser_Handler(t *testing.T) {
	// request 1
	r1, err := http.NewRequest(http.MethodGet, "localhost:8080", nil)
	r1.URL.Path = "/add"
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	// request 2
	gA := strconv.Itoa(int(groups.GroupA))
	gB := strconv.Itoa(int(groups.GroupB))
	params := url.Values{"first_name": {"Fancy"}, "last_name": {"Gopher"}, "email": {"gopher@hotmail.com"}, "groups_configurations": {gA, gB}}
	r2, err := http.NewRequest(http.MethodPost, "localhost:8080", strings.NewReader(params.Encode()))
	r2.URL.Path = "/add"
	r2.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// setup
	app := new(App)
	storage := new(memory.Storage)
	var adder adding.Service
	adder = adding.NewService(storage)
	rec1 := httptest.NewRecorder()
	rec2 := httptest.NewRecorder()

	// make requests
	app.AddUser.Handler(adder).ServeHTTP(rec1, r1)
	res1 := rec1.Result()
	app.AddUser.Handler(adder).ServeHTTP(rec2, r2)
	res2 := rec2.Result()

	// tests
	if res1.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res1.StatusCode)
	}

	if res2.StatusCode != http.StatusTemporaryRedirect {
		t.Errorf("expected status 307; got %v", res2.StatusCode)
	}
}
