package chapter_domain

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/repo/chapterpath_repo"
	"yellowroad_library/database/repo/chapter_repo"
	"yellowroad_library/database/repo/book_repo"
)

type UpdatePathBetweenChapters struct {
	chapterPathRepo chapterpath_repo.ChapterPathRepository
	chapterRepo chapter_repo.ChapterRepository
	authCheck authorityCheck
}

func NewUpdatePathBetweenChapters(
	chapterRepo chapter_repo.ChapterRepository,
	chapterPathRepo chapterpath_repo.ChapterPathRepository,
	bookRepo book_repo.BookRepository,
) UpdatePathBetweenChapters {
	return UpdatePathBetweenChapters{
		chapterPathRepo,
		chapterRepo,
		newAuthorityCheck(bookRepo), //TODO: find a way to move the authcheck creation out of here
	}
}

func (this UpdatePathBetweenChapters) Execute(
	user entities.User,
	path_id int,
	form entities.ChapterPath_UpdateForm,
) (entities.ChapterPath, app_error.AppError) {

	path, err := this.chapterPathRepo.FindById(path_id)
	if (err != nil) {
		return path, err
	}

	//TODO: if chapterPathRepo.FindById had an "associations" param, we wouldn't need to do this
	chapter, err := this.chapterRepo.FindById(path.FromChapterId)
	if (err != nil) {return path, err}

	err = this.authCheck.Execute(auth_update_path, user.ID, chapter.BookId)
	if (err != nil){
		return path, err
	}

	form.Apply(&path)
	err = this.chapterPathRepo.Update(&path)
	if (err != nil){
		return path, err
	}

	return path, nil
}