package book_repo

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
)

//go:generate moq -out Mock.go . BookRepository
type BookRepository interface {
	FindById(int) (entities.Book, app_error.AppError)
	Update(*entities.Book) app_error.AppError
	Insert(*entities.Book) app_error.AppError
	Delete(*entities.Book) app_error.AppError
    Paginate(startpage int, perpage int, options SearchOptions) ([]entities.Book, app_error.AppError)
}

//TODO : define this
type SearchOptions struct {

}