package chapter_serv

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/repo/uow"
)

type ChapterService interface {
	CreateChapter(currentUser entities.User, bookID int, form entities.ChapterCreationForm) (entities.Chapter,app_error.AppError)
	UpdateChapter(currentUser entities.User, chapterID int, form entities.ChapterUpdateForm) (entities.Chapter,app_error.AppError)
	DeleteChapter(currentUser entities.User, chapterID int) app_error.AppError
	SetUnitOfWork(work uow.UnitOfWork)

	DeletePathBetweenChapters(instigator entities.User, path_id int) app_error.AppError
	CreatePathBetweenChapters(instigator entities.User, path *entities.ChapterPath) app_error.AppError
	//UpdatePathBetweenChapters(instigator entities.User, path *entities.ChapterPath) app_error.AppError
}