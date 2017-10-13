package book_serv

import (
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
	"yellowroad_library/database/repo/uow"
)

type BookService interface {
	CreateBook(creator entities.User, book *entities.Book) app_error.AppError
	DeleteBook(book_id int, instigator entities.User) app_error.AppError
	SetUnitOfWork(work uow.UnitOfWork)
}

type BookServiceFactory func(uow.UnitOfWork) BookService
