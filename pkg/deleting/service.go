package deleting

import "errors"

// Event defines possible outcomes from the "adding actor"
type Event int

const (
	// Done means finished processing successfully
	Done Event = iota

	// Failed means processing did not finish successfully
	Failed

	// NotFound means that user was not found for deleting
	NotFound
)

func (e Event) GetMeaning() string {
	if e == Done {
		return "Done! user was deleted successfully."
	}

	if e == Failed {
		return "Failed! we couldn't delete the user."
	}

	return "Unknown result"
}

var ErrUserNotFound = errors.New("user not found")

// Service provides user deleting operations
type Service interface {
	DeleteUser(...int) <- chan Event
}

// Repository provides access to user repository
type Repository interface {
	// DeleteUser deletes a user with the specified ID
	DeleteUser(int) error
}

type service struct {
	r Repository
}

// NewService creates a deleting service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// DeleteUser deletes a user with the specified ID
func (s *service) DeleteUser(userIDs ...int) <- chan Event {
	results := make(chan Event)

	go func() {
		defer close(results)
		for _, id := range userIDs {
			err := s.r.DeleteUser(id)
			if err != nil {
				if err == ErrUserNotFound {
					results <- NotFound
					continue
				}
				results <- Failed
				continue
			}
			results <- Done
		}
	}()

	return results
}