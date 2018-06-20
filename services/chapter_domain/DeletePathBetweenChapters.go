package chapter_domain

import (
	"yellowroad_library/database/repo/chapterpath_repo"
	"yellowroad_library/database/repo/chapter_repo"
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
)

type DeletePathBetweenChapters struct {
	chapterPathRepo chapterpath_repo.ChapterPathRepository
	chapterRepo chapter_repo.ChapterRepository
	authCheck authorityCheck
}

func NewDeletePathBetweenChapters(
	chapterPathRepo chapterpath_repo.ChapterPathRepository,
	chapterRepo chapter_repo.ChapterRepository,
	bookRepo book_repo.BookRepository,
) DeletePathBetweenChapters {
	return DeletePathBetweenChapters{
		chapterPathRepo,
		chapterRepo,
		newAuthorityCheck(bookRepo), //TODO: find a way to move the authcheck creation out of here
	}
}

func (this DeletePathBetweenChapters) Execute(
	user entities.User,
	path_id int,
) app_error.AppError {
	path, err := this.chapterPathRepo.FindById(path_id)
	if (err != nil ){
		return err
	}

	//TODO: if chapterPathRepo.FindById had an "associations" param, we wouldn't need to do this
	chapter, err := this.chapterRepo.FindById(path.FromChapterId)
	if (err != nil) {return err}

	err = this.authCheck.Execute(auth_delete_path, user.ID, chapter.BookId)
	if (err != nil) {return err}

	err = this.chapterPathRepo.Delete(&path)
	if (err != nil){
		return err
	}

	return nil
}