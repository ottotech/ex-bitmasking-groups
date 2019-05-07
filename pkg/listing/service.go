package listing

import "errors"

var ErrUserNotFound = errors.New("user not found when updating")

// Service provides user listing operations
type Service interface {
	GetUser(int) (User, error)
	GetAllUsers() []User
}

// Repository provides access to user repository
type Repository interface {
	// GetUser returns a user with the given ID
	GetUser(int) (User, error)
	// GetAllUsers returns all users in the repository
	GetAllUsers() []User
}

type service struct {
	r Repository
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetUser returns a user with the given ID
func (s *service) GetUser(id int) (User, error) {
	return s.r.GetUser(id)
}

// GetAllUsers returns all users in the repository
func (s *service) GetAllUsers() []User {
	return s.r.GetAllUsers()
}
