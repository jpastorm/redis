package user

import "github.com/jpastorm/redis/model"

// UseCase interfaces of use cases of User
type UseCase interface {
	Create(m *model.User) error
	//GetWhere(filter model.Fields) (model.User, error)
}

// Storage interface for db engines
type Storage interface {
	Create(m *model.User) error
	//GetWhere(filter model.Fields, sort model.SortFields) (model.User, error)
}
