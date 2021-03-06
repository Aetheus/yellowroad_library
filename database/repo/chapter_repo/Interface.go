package chapter_repo

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
)


//go:generate moq -out Mock.go . ChapterRepository
type ChapterRepository interface {
	FindById(int) (entities.Chapter, app_error.AppError)
	FindWithinBook(chapter_id int, book_id int) (entities.Chapter, app_error.AppError)
	ChaptersIndex(book_id int) ([]entities.Chapter, app_error.AppError)
	Update(*entities.Chapter) app_error.AppError
	Insert(*entities.Chapter) app_error.AppError
	Delete(*entities.Chapter) app_error.AppError

	//TODO: we need a "FindInId" or something for fetching multiple chapters
	//TODO: we need to add an "associations" param for most of these functions to fetch the associations
}