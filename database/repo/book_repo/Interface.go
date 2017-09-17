package book_repo

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
)

type BookRepository interface {
	FindById(int) (*entities.Book, app_error.AppError)
	Update(*entities.Book) app_error.AppError
	Insert(*entities.Book) app_error.AppError
}
