package user

import (
	"fmt"
	"github.com/jpastorm/redis/model"
)

// User implement the UseCase interface
type User struct {
	storage Storage
}

// New returns a new User pointer
func New(s Storage) User {
	return User{storage: s}
}

// Create creates a new model.User
func (c User) Create(m *model.User) error {

	if err := c.storage.Create(m); err != nil {
		return fmt.Errorf("User: %w", err)
	}

	return nil
}