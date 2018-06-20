package chapter_domain

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/repo/chapter_repo"
	"yellowroad_library/database/repo/book_repo"
)

type DeleteChapter struct {
	chapterRepo chapter_repo.ChapterRepository
	bookRepo book_repo.BookRepository
}

func NewDeleteChapter(
	chapterRepo chapter_repo.ChapterRepository,
	bookRepo book_repo.BookRepository,
) DeleteChapter {
	return DeleteChapter{chapterRepo,bookRepo}
}

func (this DeleteChapter) Execute(user entities.User, chapter_id int) app_error.AppError {
	chapter, err := this.chapterRepo.FindById(chapter_id)
	if (err != nil) {
		return err
	}

	//TODO: find a way to move the authcheck creation out of here
	err = newAuthorityCheck(this.bookRepo).Execute(auth_delete_chapter, user.ID, chapter.BookId)
	if(err != nil){
		return err
	}

	err = this.chapterRepo.Delete(&chapter)
	if(err != nil){
		return err
	}

	return nil
}