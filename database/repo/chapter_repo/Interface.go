package chapter_repo

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
)

type ChapterRepository interface {
	FindById(int) (entities.Chapter, app_error.AppError)
	FindWithinBook(chapter_id int, book_id int) (entities.Chapter, app_error.AppError)
	Update(*entities.Chapter) app_error.AppError
	Insert(*entities.Chapter) app_error.AppError
	Delete(*entities.Chapter) app_error.AppError
}