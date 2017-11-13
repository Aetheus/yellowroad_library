package user_repo

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
)

//go:generate moq -out Mock.go . UserRepository
type UserRepository interface {
	FindById(int) (entities.User, app_error.AppError)
	FindByUsername(string) (entities.User, app_error.AppError)
	Update(*entities.User) app_error.AppError
	Insert(*entities.User) app_error.AppError
	Delete(*entities.User) app_error.AppError
}
