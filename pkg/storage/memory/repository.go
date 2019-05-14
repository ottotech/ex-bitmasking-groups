package memory

import (
	"github.com/ottotech/ex-bitmasking-groups/pkg/adding"
	"github.com/ottotech/ex-bitmasking-groups/pkg/deleting"
	"github.com/ottotech/ex-bitmasking-groups/pkg/listing"
)

// Memory storage keeps data in memory
type Storage struct {
	users []User
}

// AddUser saves the given user in repository
func (m *Storage) AddUser(u adding.User) error {

	// validate user uniqueness by email
	for i := range m.users {
		if m.users[i].Email == u.Email {
			return adding.ErrDuplicatedUser
		}
	}

	newUser := User{
		Id:          len(m.users) + 1,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		GroupConfig: u.GroupConfig,
	}

	m.users = append(m.users, newUser)
	return nil
}

// DeleteUser deletes a user with the specified ID
func (m *Storage) DeleteUser(userID int) error {
	for i := range m.users {
		if m.users[i].Id == userID {
			m.users[i] = m.users[len(m.users)-1]
			m.users = m.users[:len(m.users)-1]
			return nil
		}
	}
	return deleting.ErrUserNotFound
}

// GetUser returns a user with the given ID
func (m *Storage) GetUser(userID int) (listing.User, error) {
	for i := range m.users {
		if m.users[i].Id == userID {
			u := listing.User{
				Id:          m.users[i].Id,
				FirstName:   m.users[i].FirstName,
				LastName:    m.users[i].LastName,
				Email:       m.users[i].Email,
				GroupConfig: m.users[i].GroupConfig,
			}
			return u, nil
		}
	}
	return listing.User{}, listing.ErrUserNotFound
}

// GetAllUsers returns all users in the repository
func (m *Storage) GetAllUsers() []listing.User {
	var list []listing.User

	for i := range m.users {
		u := listing.User{
			Id:          m.users[i].Id,
			FirstName:   m.users[i].FirstName,
			LastName:    m.users[i].LastName,
			Email:       m.users[i].Email,
			GroupConfig: m.users[i].GroupConfig,
		}
		list = append(list, u)
	}

	return list
}
