package rest

import (
	"github.com/ottotech/ex-bitmasking-groups/pkg/adding"
	_ "github.com/ottotech/ex-bitmasking-groups/pkg/config"
	"github.com/ottotech/ex-bitmasking-groups/pkg/deleting"
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

	u := storage.GetAllUsers()[0]
	if u.FirstName != "Fancy" || u.LastName != "Gopher" || u.Email != "gopher@hotmail.com" || u.GroupConfig != int(groups.GroupA|groups.GroupB) {
		t.Errorf("expected a user `Fancy Gopher gopher@hotmail.com %v`; got %v %v %v %v", int(groups.GroupA|groups.GroupB), u.FirstName, u.LastName, u.Email, u.GroupConfig)
	}
}

func TestGetUser_Handler(t *testing.T) {
	// request 1
	r1, err := http.NewRequest(http.MethodGet, "localhost:8080", nil)
	r1.URL.Path = "/get/1"
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// request 2
	r2, err := http.NewRequest(http.MethodGet, "localhost:8080", nil)
	r2.URL.Path = "/get/2"
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// request 3
	r3, err := http.NewRequest(http.MethodGet, "localhost:8080", nil)
	r3.URL.Path = "/get/2"
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
	rec3 := httptest.NewRecorder()

	// create user
	u := adding.User{FirstName: "Fancy", LastName: "Gopher", Email: "gopher@gmail.com", GroupConfig: 0}
	_ = storage.AddUser(u)

	// make requests
	app.GetUser.Handler(lister).ServeHTTP(rec1, r1)
	res1 := rec1.Result()
	app.GetUser.Handler(lister).ServeHTTP(rec2, r2)
	res2 := rec2.Result()
	app.GetUser.Handler(lister).ServeHTTP(rec3, r3)
	res3 := rec3.Result()

	// tests
	if res1.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res1.StatusCode)
	}

	if res2.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status 400; got %v", res2.StatusCode)
	}

	if res3.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status 400; got %v", res3.StatusCode)
	}
}

func TestDeleteUser_Handler(t *testing.T) {
	// request 1
	r1, err := http.NewRequest(http.MethodPost, "localhost:8080", nil)
	r1.URL.Path = "/delete/1"
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// request 2
	r2, err := http.NewRequest(http.MethodPost, "localhost:8080", nil)
	r2.URL.Path = "/delete/2"
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// request 3
	r3, err := http.NewRequest(http.MethodGet, "localhost:8080", nil)
	r3.URL.Path = "/delete/"
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// setup
	app := new(App)
	storage := new(memory.Storage)
	var deleter deleting.Service
	deleter = deleting.NewService(storage)
	rec1 := httptest.NewRecorder()
	rec2 := httptest.NewRecorder()
	rec3 := httptest.NewRecorder()

	// create user
	u := adding.User{FirstName: "Fancy", LastName: "Gopher", Email: "gopher@gmail.com", GroupConfig: 0}
	_ = storage.AddUser(u)

	// make requests
	app.DeleteUser.Handler(deleter).ServeHTTP(rec1, r1)
	res1 := rec1.Result()
	app.DeleteUser.Handler(deleter).ServeHTTP(rec2, r2)
	res2 := rec2.Result()
	app.DeleteUser.Handler(deleter).ServeHTTP(rec3, r3)
	res3 := rec3.Result()

	// tests
	if res1.StatusCode != http.StatusTemporaryRedirect {
		t.Errorf("expected status 307; got %v", res1.StatusCode)
	}

	if res2.StatusCode != http.StatusTemporaryRedirect {
		t.Errorf("expected status 307; got %v", res2.StatusCode)
	}

	if res3.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected status 405; got %v", res3.StatusCode)
	}

}
