package book_serv

import (
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
	"yellowroad_library/database/repo/uow"
)

type BookService interface {
	CreateBook(currentUser entities.User, form entities.Book_CreationForm) (entities.Book,app_error.AppError)
	UpdateBook(currentUser entities.User, book_id int, form entities.Book_UpdateForm) (entities.Book,app_error.AppError)
	DeleteBook(currentUser entities.User, book_id int) app_error.AppError

	PlusTag(currentUser entities.User, book_id int, tagName string) (newCount int, err app_error.AppError)
	MinusTag(currentUser entities.User, book_id int, tagName string) (newCount int, err app_error.AppError)

	SetUnitOfWork(work uow.UnitOfWork)
}

