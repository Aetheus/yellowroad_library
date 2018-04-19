package book_serv

import (
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
	"net/http"
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

func (this DefaultBookService) SetUnitOfWork(work uow.UnitOfWork) {
	this.work = work
}
func (this DefaultBookService) CreateBook(creator entities.User, form entities.Book_CreationForm) (entities.Book,app_error.AppError) {
	var BookRepo = this.work.BookRepo()
	var ChapterRepo = this.work.ChapterRepo()

	var book entities.Book
	form.Apply(&book)
	book.CreatorId = creator.ID
	if err := BookRepo.Insert(&book); err != nil {
		return book, app_error.Wrap(err)
	}

	var first_chapter =  entities.Chapter {
		Title: book.Title + ": Chapter 1",
		Body: "This is your first chapter. Fill it up!",
		CreatorId: creator.ID,
		BookId : book.ID,
	}
	if err := ChapterRepo.Insert(&first_chapter); err != nil {
		return book, app_error.Wrap(err)
	}

	book.FirstChapterId = first_chapter.ID
	if err := BookRepo.Update(&book); err != nil {
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

func (this DefaultBookService) VoteOnTag(currentUser entities.User, book_id int, tagName string, voteValue int) (newCount int, err app_error.AppError){
	_, err = this.work.BookRepo().FindById(book_id)
	if (err != nil){
		//quit early if we can't find the book or something else went wrong
		return
	}

	//prevent people from entering values larger than 1 or lower than -1
	if(voteValue > 0){
		voteValue = 1
	} else if (voteValue < 0) {
		voteValue = -1
	}

	vote := entities.BookTagVote{
		BookId : book_id,
		UserId : currentUser.ID,
		Tag : tagName,
		Direction : voteValue,
	}

	err = this.work.BookTagVoteRepo().Upsert(&vote)
	if (err != nil) {return newCount,err}

	result, err := this.work.BookTagVoteCountRepo().SyncCount(tagName,book_id)
	if (err != nil) {return newCount,err}
	newCount = result.Count

	return
}