package booktag_repo

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
)

type BookTagRepository interface {
	Insert(*entities.BookTag) (app_error.AppError)
	Delete(*entities.BookTag) (app_error.AppError)
	DeleteByFields(tag string, user_id int, book_id int) (app_error.AppError)
}