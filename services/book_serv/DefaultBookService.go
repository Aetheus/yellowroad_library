package book_serv

import (
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
	"net/http"
	"yellowroad_library/database"
	"yellowroad_library/database/repo/uow"
)

type DefaultBookService struct {
	work  uow.UnitOfWork
}

//ensure interface implementation
var _ BookService = DefaultBookService{}

func Default(work uow.UnitOfWork) BookService {

	return DefaultBookService{
		work: work,
	}
}
var _ BookServiceFactory = Default


func (this DefaultBookService) SetUnitOfWork(work uow.UnitOfWork) {
	this.work = work
}
func (this DefaultBookService) CreateBook(creator entities.User, form entities.Book_CreationForm) (entities.Book,app_error.AppError) {
	var book entities.Book

	form.Apply(&book)
	book.CreatorId = creator.ID
	book.FirstChapterId = database.NullInt{Int:0}

	if err := this.work.BookRepo().Insert(&book); err != nil {
		return book, app_error.Wrap(err)
	}

	return book, nil
}

func (this DefaultBookService) UpdateBook(currentUser entities.User,book_id int,form entities.Book_UpdateForm) (entities.Book,app_error.AppError){
	var book entities.Book

	book, err := this.work.BookRepo().FindById(book_id)
	if (err != nil){
		return book, err
	}

	//TODO: if we implement a "contributors" system, then this should do more checking later on
	if (book.CreatorId != currentUser.ID){
		return book, app_error.New(http.StatusUnauthorized, "","You are not authorized to delete this book!")
	}


	form.Apply(&book)
	err = this.work.BookRepo().Update(&book)
	if (err != nil){
		return book, err
	}

	return book, nil
}


func (this DefaultBookService) DeleteBook(currentUser entities.User, book_id int) app_error.AppError{
	book, err := this.work.BookRepo().FindById(book_id)
	if  err != nil{
		return app_error.Wrap(err)
	}

	//TODO: if we implement a "contributors" system, then this should do more checking later on
	if (book.CreatorId != currentUser.ID){
		return app_error.New(http.StatusUnauthorized,
			"","You are not authorized to delete this book!")
	}

	if err := this.work.BookRepo().Delete(&book); err != nil {
		return app_error.Wrap(err)
	}

	return nil
}