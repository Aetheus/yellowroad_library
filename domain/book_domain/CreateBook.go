package book_domain

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/database/repo/chapter_repo"
)

type CreateBook struct {
	bookRepo 	book_repo.BookRepository
	chapterRepo chapter_repo.ChapterRepository
}

func NewCreateBook(bookRepo book_repo.BookRepository, chapterRepo chapter_repo.ChapterRepository) CreateBook{
	return CreateBook {bookRepo, chapterRepo,}
}

func (this CreateBook) Execute(creator entities.User, form entities.Book_CreationForm) (entities.Book,app_error.AppError) {
	var book entities.Book

	form.Apply(&book)
	book.CreatorId = creator.ID

	if err := this.bookRepo.Insert(&book); err != nil {
		return book, app_error.Wrap(err)
	}

	var first_chapter =  entities.Chapter {
		Title: book.Title + ": Chapter 1",
		Body: "This is your first chapter. Fill it up!",
		CreatorId: creator.ID,
		BookId : book.ID,
	}
	if err := this.chapterRepo.Insert(&first_chapter); err != nil {
		return book, app_error.Wrap(err)
	}

	book.FirstChapterId = first_chapter.ID
	if err := this.bookRepo.Update(&book); err != nil {
		return book, app_error.Wrap(err)
	}

	return book, nil
}