package booktag_repo

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
)

type BookTagRepository interface {
	FindByFields(searchFields entities.BookTag) ([]entities.BookTag, app_error.AppError)
	Insert(*entities.BookTag) (app_error.AppError)
	Delete(*entities.BookTag) (app_error.AppError)
	DeleteByFields(searchFields entities.BookTag) (app_error.AppError)
}