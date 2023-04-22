package teststore

import (
	"github.com/denisandreenko/fire/internal/app/model"
	"github.com/denisandreenko/fire/internal/app/store"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	return &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}
}
