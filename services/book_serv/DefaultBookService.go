package book_serv

import (
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
	"net/http"
	"yellowroad_library/database/repo/uow"
	"gopkg.in/guregu/null.v3"
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

func (this DefaultBookService) SetUnitOfWork(work uow.UnitOfWork) {
	this.work = work
}
func (this DefaultBookService) CreateBook(creator entities.User, form entities.Book_CreationForm) (entities.Book,app_error.AppError) {
	var book entities.Book

	form.Apply(&book)
	book.CreatorId = creator.ID
	book.FirstChapterId = null.IntFrom(0)

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

func (this DefaultBookService) PlusTag(currentUser entities.User, book_id int, tagName string) (newCount int, err app_error.AppError) {
	searchFields := entities.BookTag{
		BookId : book_id,
		UserId : currentUser.ID,
		Tag : tagName,
	}
	existingVote, err := this.work.BookTagRepo().FindByFields(searchFields)

	if (len(existingVote) == 0 ){
		err = this.work.BookTagRepo().Insert(&searchFields)
		if (err != nil) {return}

		result, err := this.work.BookTagCountRepo().SyncCount(tagName,book_id)
		if (err != nil) {return}
		newCount = result.Count

		return
	} else {
		//if the user already "plused" this tag for this book, just return at this point
		return
	}
}

func (this DefaultBookService) MinusTag(currentUser entities.User, book_id int, tagName string) (newCount int, err app_error.AppError) {
	searchFields := entities.BookTag{
		BookId : book_id,
		UserId : currentUser.ID,
		Tag : tagName,
	}
	existingVote, err := this.work.BookTagRepo().FindByFields(searchFields)

	if (len(existingVote) == 0 ){
		//if the user has no existing vote for this tag, just return at this point
		return
	} else {
		err = this.work.BookTagRepo().DeleteByFields(searchFields)
		if (err != nil) {return}

		result, err := this.work.BookTagCountRepo().SyncCount(tagName,book_id)
		if (err != nil) {return}
		newCount = result.Count

		return
	}
}