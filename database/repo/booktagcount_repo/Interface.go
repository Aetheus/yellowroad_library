package booktagcount_repo

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
)

type BookTagCountRepository interface {
	Insert(*entities.BookTagCount) (app_error.AppError)
	Delete(*entities.BookTagCount) (app_error.AppError)
	SyncCount(tag string, book_id int) (count entities.BookTagCount, err app_error.AppError)
}