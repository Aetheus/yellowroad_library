package book_domain

import (
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"net/http"
)

type DeleteBook struct {
	bookRepo book_repo.BookRepository
}

func NewDeleteBook(repository book_repo.BookRepository) DeleteBook {
	return DeleteBook{repository}
}

func (this DeleteBook) Execute(currentUser entities.User, book_id int) app_error.AppError{
	book, err := this.bookRepo.FindById(book_id)
	if  err != nil{
		return app_error.Wrap(err)
	}

	//TODO: if we implement a "contributors" system, then this should do more checking later on
	if (book.CreatorId != currentUser.ID){
		return app_error.New(http.StatusUnauthorized,
			"","You are not authorized to delete this book!")
	}

	if err := this.bookRepo.Delete(&book); err != nil {
		return app_error.Wrap(err)
	}

	return nil
}
