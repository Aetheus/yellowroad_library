package chapter_serv

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/repo/uow"
)

type ChapterService interface {
	CreateChapter(currentUser entities.User, form entities.Chapter_CreationForm) (entities.Chapter,app_error.AppError)
	UpdateChapter(currentUser entities.User, chapterID int, form entities.Chapter_UpdateForm) (entities.Chapter,app_error.AppError)
	DeleteChapter(currentUser entities.User, chapterID int) app_error.AppError


	CreatePathBetweenChapters(instigator entities.User, form entities.ChapterPath_CreationForm) (entities.ChapterPath, app_error.AppError)
	UpdatePathBetweenChapters(instigator entities.User, path_id int, form entities.ChapterPath_UpdateForm) (entities.ChapterPath, app_error.AppError)
	DeletePathBetweenChapters(instigator entities.User, path_id int) app_error.AppError

	CreateChapterAndPath(
		instigator entities.User,
		chapter_form entities.Chapter_CreationForm,
		path_form entities.ChapterPath_CreationForm,
	) (entities.Chapter, entities.ChapterPath, app_error.AppError)

	SetUnitOfWork(work uow.UnitOfWork)
}