package chapterpath_repo

import (
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
)

//go:generate moq -out Mock.go . ChapterPathRepository
type ChapterPathRepository interface {
	FindById(chapterId int) (entities.ChapterPath, app_error.AppError)
	FindByChapterIds(fromChapterId int, toChapterId int) (entities.ChapterPath, app_error.AppError)
	Update(*entities.ChapterPath) app_error.AppError
	Insert(*entities.ChapterPath) app_error.AppError
	Delete(*entities.ChapterPath) app_error.AppError
}

