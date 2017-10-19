package book_serv

import (
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
	"yellowroad_library/database/repo/uow"
)

type BookService interface {
	CreateBook(currentUser entities.User, form entities.BookCreationForm) (entities.Book,app_error.AppError)
	UpdateBook(currentUser entities.User, book_id int, form entities.BookUpdateForm) (entities.Book,app_error.AppError)
	DeleteBook(currentUser entities.User, book_id int) app_error.AppError
	SetUnitOfWork(work uow.UnitOfWork)
}

type BookServiceFactory func(uow.UnitOfWork) BookService
