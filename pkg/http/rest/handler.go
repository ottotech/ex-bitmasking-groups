package rest

import (
	"fmt"
	"github.com/ottotech/ex-bitmasking-groups/pkg/adding"
	"github.com/ottotech/ex-bitmasking-groups/pkg/deleting"
	"github.com/ottotech/ex-bitmasking-groups/pkg/groups"
	"github.com/ottotech/ex-bitmasking-groups/pkg/listing"
	"github.com/ottotech/ex-bitmasking-groups/pkg/utils"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
)

type App struct {
	UserList   *UserList
	AddUser    *AddUser
	GetUser    *GetUser
	DeleteUser *DeleteUser
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
		groupsConfigurations := r.PostForm["groups_configurations"]
		groupConfig := 0

		for _, config := range groupsConfigurations {
			configInt, _ := strconv.Atoi(config) // for simplicity error is not handled
			groupConfig |= configInt
		}

		if firstName == "" || lastName == "" || email == "" {
			ctx := struct {
				Groups []groups.GroupData
				Error  string
			}{
				[]groups.GroupData{},
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
		customURL := "/"
		for r := range results {
			customURL += fmt.Sprintf(urlParameter, r.GetMeaning())
		}

		http.Redirect(w, r, customURL, http.StatusTemporaryRedirect)
		return
	})
}

type GetUser struct {
}

func (h *GetUser) Handler(l listing.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		// clean path to get the user ID
		p := path.Clean("/" + r.URL.Path)[1:]
		i := strings.Index(p, "/") + 1
		tail := p[i:]

		// check if id is valid
		id, err := strconv.Atoi(tail)
		if err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("This user ID is not valid: %v.", tail), http.StatusBadRequest)
			return
		}

		u, err := l.GetUser(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// one way to check if a user belongs to a particular group
		if groups.BelongsToGroup(groups.GroupA, u.GroupConfig) {
			msg := fmt.Sprintf("User: %v belongs to group A!", u.FirstName)
			fmt.Println(msg)
		}

		userGroups := groups.GetGroupsByConfiguration(u.GroupConfig)

		ctx := struct {
			User       listing.User
			UserGroups []groups.GroupData
		}{
			u,
			userGroups,
		}

		utils.RenderTemplate(w, "detail.gohtml", ctx)
		return
	})
}

type DeleteUser struct {
}

func (h *DeleteUser) Handler(d deleting.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		// clean path to get the user ID
		p := path.Clean("/" + r.URL.Path)[1:]
		i := strings.Index(p, "/") + 1
		tail := p[i:]

		// check if id is valid
		id, err := strconv.Atoi(tail)
		if err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("This user ID is not valid: %v.", tail), http.StatusBadRequest)
			return
		}

		// get results of operation
		results := d.DeleteUser(id)

		// send message(s) through URL
		urlParameter := "?message=%v"
		customURL := "/"
		for r := range results {
			customURL += fmt.Sprintf(urlParameter, r.GetMeaning())
		}

		http.Redirect(w, r, customURL, http.StatusTemporaryRedirect)
		return
	})
}
