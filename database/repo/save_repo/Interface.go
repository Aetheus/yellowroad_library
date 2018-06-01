package save_repo

import (
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
)

type SaveRepository interface {
	FindById(id int) (entities.SaveState, app_error.AppError)
	Insert(*entities.SaveState) (app_error.AppError)
	Update(*entities.SaveState) (app_error.AppError)
	Delete(*entities.SaveState) (app_error.AppError)
}