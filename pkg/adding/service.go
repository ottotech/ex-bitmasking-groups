package adding

import (
	"errors"
)

// Event defines possible outcomes from the "adding actor"
type Event int

const (
	// UserCreatedSuccessFully means finished processing successfully
	UserCreatedSuccessFully Event = iota

	// UserAlreadyExists means the given user is a duplicate of an existing one
	UserAlreadyExists

	// UserCouldNotBeCreated means processing did not finish successfully
	UserCouldNotBeCreated
)

func (e Event) GetMeaning() string {
	if e == UserCreatedSuccessFully {
		return "User was created successfully"
	}

	if e == UserAlreadyExists {
		return "Duplicate user"
	}

	if e == UserCouldNotBeCreated {
		return "User couldn't be created"
	}

	return "Unknown result"
}

var ErrDuplicatedUser = errors.New("user already exists")

// Service provides user adding operations
type Service interface {
	AddUser(...User) <-chan Event
}

// Repository provides access to user repository
type Repository interface {
	// AddUser saves a given user to the repository
	AddUser(User) error
}

type service struct {
	r Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddUser adds the given user(s) to the repository
func (s *service) AddUser(users ...User) <-chan Event {
	results := make(chan Event)

	go func() {
		defer close(results)
		for _, u := range users {
			err := s.r.AddUser(u)
			if err != nil {
				if err == ErrDuplicatedUser {
					results <- UserAlreadyExists
					continue
				}
				results <- UserCouldNotBeCreated
				continue
			}
			results <- UserCreatedSuccessFully
		}
	}()

	return results
}
