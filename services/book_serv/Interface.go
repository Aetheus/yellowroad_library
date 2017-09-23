package book_serv

import (
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
)

type BookService interface {
	CreateBook(creator entities.User, book *entities.Book) app_error.AppError
	DeleteBook(instigator entities.User, book *entities.Book) app_error.AppError
}