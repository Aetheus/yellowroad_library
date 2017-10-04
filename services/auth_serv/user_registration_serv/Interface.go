package user_registration_serv

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
)

type UserRegistrationService interface {
	Run (username string, password string, email string) (returnedUser *entities.User, returnedErr app_error.AppError)
}
