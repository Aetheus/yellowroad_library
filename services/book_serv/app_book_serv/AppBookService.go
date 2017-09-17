package app_book_serv

import (
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/database/repo/user_repo"
	"yellowroad_library/database/entities"
)

type BookService struct {
	bookRepo book_repo.BookRepository
	userRepo user_repo.UserRepository
}

func New(bookRepo book_repo.BookRepository, userRepo user_repo.UserRepository ) BookService {
	return BookService{
		bookRepo : bookRepo,
		userRepo : userRepo,
	}
}

func (this BookService) CreateBook(creator entities.User, book *entities.Book) app_error.AppError {
	//do some extra checking here (eg: check if the creator is banned or not, etc)
	book.CreatorId = creator.ID
	book.FirstChapterId = 0

	if err := this.bookRepo.Insert(book); err != nil {
		return app_error.Wrap(err)
	}

	return nil
}