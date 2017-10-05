package book_delete

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
)

type BookDeleteService interface{
	DeleteBook(book_id int, instigator entities.User) app_error.AppError
	//DeleteBook(instigator entities.User, book *entities.Book) app_error.AppError
}