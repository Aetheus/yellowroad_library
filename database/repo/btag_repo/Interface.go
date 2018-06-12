package btag_repo

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
)

type BookTagRepository interface {
	Insert(*entities.BookTag) (app_error.AppError)
	Delete(*entities.BookTag) (app_error.AppError)
	SyncCount(tag string, book_id int) (count entities.BookTag, err app_error.AppError)
}