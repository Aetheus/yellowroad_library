package chapter_domain

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/repo/chapter_repo"
	"yellowroad_library/database/repo/book_repo"
)

type CreateChapter struct {
	chapterRepo chapter_repo.ChapterRepository
	bookRepo book_repo.BookRepository
}

func NewCreateChapter(
	chapterRepo chapter_repo.ChapterRepository,
	bookRepo book_repo.BookRepository,
) CreateChapter {
	return CreateChapter{chapterRepo, bookRepo}
}

func (this CreateChapter) Execute(
	user entities.User, form entities.Chapter_CreationForm,
) (entities.Chapter,app_error.AppError) {
	var newChapter entities.Chapter
	book_id := *form.BookId

	//TODO: find a way to move the authcheck creation out of here
	err := newAuthorityCheck(this.bookRepo).Execute(auth_create_chapter, user.ID, book_id)
	if (err != nil){
		return newChapter, err
	}

	form.Apply(&newChapter)
	newChapter.BookId = book_id
	newChapter.CreatorId = user.ID

	err = this.chapterRepo.Insert(&newChapter)
	if err != nil {
		return newChapter, err
	}

	return newChapter, nil
}