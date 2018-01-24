package storysave_repo

import (
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
)

type StorySaveRepository interface {
	FindByToken(token string) (entities.StorySave, app_error.AppError)
	Insert(*entities.StorySave) (app_error.AppError)
	Update(*entities.StorySave) (app_error.AppError)
	Delete(*entities.StorySave) (app_error.AppError)
}