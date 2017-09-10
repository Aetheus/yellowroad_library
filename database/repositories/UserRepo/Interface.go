package UserRepo

import "yellowroad_library/database/entities"

type UserRepository interface {
	FindById(int) (*entities.User, error)
	FindByUsername(string) (*entities.User, error)
	Update(*entities.User) error
	Insert(*entities.User) error
}
