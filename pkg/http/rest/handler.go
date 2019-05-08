package rest

import (
	"github.com/ottotech/ex-bitmasking-groups/pkg/adding"
	"github.com/ottotech/ex-bitmasking-groups/pkg/groups"
	"github.com/ottotech/ex-bitmasking-groups/pkg/listing"
	"github.com/ottotech/ex-bitmasking-groups/pkg/utils"
	"log"
	"net/http"
)

type App struct {
	UserList *UserList
	AddUser  *AddUser
}

type UserList struct {
}

func (h *UserList) Handler(l listing.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		list := l.GetAllUsers()
		utils.RenderTemplate(w, "list.gohtml", list)
		return
	})
}

type AddUser struct {
}

func (h *AddUser) Handler(a adding.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			groupList := groups.GetAllGroups()

			ctx := struct {
				Groups []groups.GroupData
				Error  string
			}{
				groupList,
				"",
			}

			utils.RenderTemplate(w, "add.gohtml", ctx)
			return
		}

		firstName := r.PostFormValue("first_name")
		lastName := r.PostFormValue("last_name")
		email := r.PostFormValue("email")

		if firstName == "" || lastName == "" || email == "" {
			ctx := struct {
				Error string
			}{
				"Error: All fields are mandatory.",
			}
			utils.RenderTemplate(w, "add.gohtml", ctx)
			return
		}

		u := adding.User{FirstName: firstName, LastName: lastName, Email: email}
		results := a.AddUser(u)

		for r := range results {
			log.Println(r.GetMeaning())
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	})
}
