package book_serv

import (
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/database/repo/user_repo"
	"yellowroad_library/database/entities"
	"net/http"
	"yellowroad_library/database"
)

type DefaultBookService struct {
	bookRepo book_repo.BookRepository
	userRepo user_repo.UserRepository
}
//ensure interface implementation
var _ BookService = DefaultBookService{}

func Default(bookRepo book_repo.BookRepository, userRepo user_repo.UserRepository ) DefaultBookService {
	return DefaultBookService{
		bookRepo : bookRepo,
		userRepo : userRepo,
	}
}

func (this DefaultBookService) CreateBook(creator entities.User, book *entities.Book) app_error.AppError {
	//TODO: do some extra checking here (eg: check if the creator is banned or not, etc)
	book.CreatorId = creator.ID
	book.FirstChapterId = database.NullInt{Int:0}

	if err := this.bookRepo.Insert(book); err != nil {
		return app_error.Wrap(err)
	}

	return nil
}

func (this DefaultBookService) DeleteBook(instigator entities.User, book *entities.Book) app_error.AppError{
	//TODO: if we implement a "contributors" system, then this should do more checking later on
	if (book.CreatorId != instigator.ID){
		return app_error.New(http.StatusUnauthorized,
				"","You are not authorized to delete this book!")
	}

	if err := this.bookRepo.Delete(book); err != nil {
		return app_error.Wrap(err)
	}

	return nil
}