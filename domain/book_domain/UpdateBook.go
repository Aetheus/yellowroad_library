package book_domain

import (
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"net/http"
)

type UpdateBook struct {
	bookRepo book_repo.BookRepository
}

func NewUpdateBook(repository book_repo.BookRepository) UpdateBook {
	return UpdateBook{repository}
}

func (this UpdateBook) Execute(currentUser entities.User,book_id int,form entities.Book_UpdateForm) (entities.Book,app_error.AppError){
	var book entities.Book

	book, err := this.bookRepo.FindById(book_id)
	if (err != nil){
		return book, err
	}

	//TODO: if we implement a "contributors" system, then this should do more checking later on
	if (book.CreatorId != currentUser.ID){
		return book, app_error.New(http.StatusUnauthorized, "","You are not authorized to delete this book!")
	}

	form.Apply(&book)

	err = this.bookRepo.Update(&book)
	if (err != nil){
		return book, err
	}

	return book, nil
}