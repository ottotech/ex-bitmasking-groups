package rest

import (
	"fmt"
	"github.com/ottotech/ex-bitmasking-groups/pkg/adding"
	"github.com/ottotech/ex-bitmasking-groups/pkg/groups"
	"github.com/ottotech/ex-bitmasking-groups/pkg/listing"
	"github.com/ottotech/ex-bitmasking-groups/pkg/utils"
	"net/http"
	"strconv"
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

		_ = r.ParseForm()
		messages := r.Form["message"]
		list := l.GetAllUsers()

		ctx := struct {
			Messages []string
			List     []listing.User
		}{
			messages,
			list,
		}

		utils.RenderTemplate(w, "list.gohtml", ctx)
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

		_ = r.ParseForm()
		firstName := r.PostFormValue("first_name")
		lastName := r.PostFormValue("last_name")
		email := r.PostFormValue("email")
		groupsIDs := r.PostForm["groups_ids"]
		groupConfig := 0

		for _, id := range groupsIDs {
			idInt, _ := strconv.Atoi(id) // for simplicity error is not handled
			groupConfig |= idInt
		}

		if firstName == "" || lastName == "" || email == "" {
			ctx := struct {
				Error string
			}{
				"Error: All fields are mandatory.",
			}
			utils.RenderTemplate(w, "add.gohtml", ctx)
			return
		}

		u := adding.User{
			FirstName:   firstName,
			LastName:    lastName,
			Email:       email,
			GroupConfig: groupConfig,
		}

		// get results of operation
		results := a.AddUser(u)

		// send message(s) through URL
		urlParameter := "?message=%v"
		url := "/"
		for r := range results {
			url += fmt.Sprintf(urlParameter, r.GetMeaning())
		}

		http.Redirect(w, r, url, http.StatusSeeOther)
		return
	})
}
