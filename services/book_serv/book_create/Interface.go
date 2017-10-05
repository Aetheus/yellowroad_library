package book_create

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
)

type BookCreateService interface {
	CreateBook(creator entities.User, book *entities.Book) app_error.AppError
}
