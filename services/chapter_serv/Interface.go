package chapter_serv

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/repo/uow"
)

type ChapterService interface {
	CreateChapter(instigator entities.User, book_id int, chapter *entities.Chapter) app_error.AppError
	UpdateChapter(instigator entities.User, chapter *entities.Chapter) app_error.AppError
	DeleteChapter(instigator entities.User, chapter_id int) app_error.AppError
	SetUnitOfWork(work uow.UnitOfWork)

	DeletePathBetweenChapters(instigator entities.User, path_id int) app_error.AppError
	CreatePathBetweenChapters(instigator entities.User, path *entities.ChapterPath) app_error.AppError
	//UpdatePathBetweenChapters(instigator entities.User, path *entities.ChapterPath) app_error.AppError
}