package book_repo

import "yellowroad_library/database/entities"

type BookRepository interface {
	FindById(int) (*entities.Book, error)
	Update(*entities.Book) error
	Insert(*entities.Book) error
}
