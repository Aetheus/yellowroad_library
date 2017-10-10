package uow

import (
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/database/repo/chapter_repo"
	"yellowroad_library/database/repo/chapterpath_repo"
	"yellowroad_library/database/repo/user_repo"
	"yellowroad_library/utils/app_error"
	"github.com/jinzhu/gorm"
	"yellowroad_library/database/repo/book_repo/gorm_book_repo"
	"yellowroad_library/database/repo/chapter_repo/gorm_chapter_repo"
	"yellowroad_library/database/repo/chapterpath_repo/gorm_chapterpath_repo"
	"yellowroad_library/database/repo/user_repo/gorm_user_repo"
)

type UnitOfWork interface {
	BookRepo() (book_repo.BookRepository)
	ChapterRepo() (chapter_repo.ChapterRepository)
	ChapterPathRepo() (chapterpath_repo.ChapterPathRepository)
	UserRepo() (user_repo.UserRepository)

	Finish() (app_error.AppError) //for use in a defer; calls rollback if the transaction had a commit error or hasn't been commited yet
	Commit() (app_error.AppError)
	Rollback() (app_error.AppError)
}

type SimpleUnitOfWorkFactory func() UnitOfWork

type AppUnitOfWork struct {
	bookRepo *book_repo.BookRepository
	chapterRepo *chapter_repo.ChapterRepository
	chapterPathRepo *chapterpath_repo.ChapterPathRepository
	userRepo *user_repo.UserRepository

	hasCommitted bool
	hasCommitErrors bool
	hasRolledback bool

	transaction *gorm.DB
}
var _ UnitOfWork = AppUnitOfWork{}
func NewAppUnitOfWork(db *gorm.DB) UnitOfWork{
	transaction := db.Begin();
	return AppUnitOfWork{
		transaction: transaction,
		hasCommitted: false,
		hasCommitErrors : false,
		hasRolledback : false,
	}
}

func (this AppUnitOfWork) Finish() (app_error.AppError) {
	if(!this.hasCommitted || this.hasCommitErrors && !this.hasRolledback ){
		this.Rollback()
	}
	return nil
}

func (this AppUnitOfWork) Commit() (app_error.AppError) {
	this.transaction.Commit()
	this.hasCommitted = true

	if(this.transaction.Error != nil){
		return app_error.Wrap(this.transaction.Error)
		this.hasCommitErrors = true
	}

	return nil
}

func (this AppUnitOfWork) Rollback() (app_error.AppError) {
	currentNumErrors := len(this.transaction.GetErrors())

	this.transaction.Rollback()
	this.hasRolledback = true

	newNumErrors := len(this.transaction.GetErrors())

	if (currentNumErrors != newNumErrors) {
		//transaction.Error should hold the latest error
		return app_error.Wrap(this.transaction.Error)
	}

	return nil
}

func (this AppUnitOfWork) BookRepo() (book_repo.BookRepository) {
	if this.bookRepo == nil {
		var bookRepo book_repo.BookRepository = gorm_book_repo.New(this.transaction)
		this.bookRepo = &bookRepo
	}
	return *this.bookRepo
}

func (this AppUnitOfWork) ChapterRepo() (chapter_repo.ChapterRepository) {
	if this.chapterRepo == nil {
		var chapterRepo chapter_repo.ChapterRepository = gorm_chapter_repo.New(this.transaction)
		this.chapterRepo = &chapterRepo
	}
	return *this.chapterRepo
}

func (this AppUnitOfWork) ChapterPathRepo() (chapterpath_repo.ChapterPathRepository) {
	if this.chapterPathRepo == nil {
		var chapterPathRepo chapterpath_repo.ChapterPathRepository =
			gorm_chapterpath_repo.New(this.transaction)
		this.chapterPathRepo = &chapterPathRepo
	}
	return *this.chapterPathRepo
}

func (this AppUnitOfWork) UserRepo() (user_repo.UserRepository) {
	if this.userRepo == nil {
		var userRepo user_repo.UserRepository =
			gorm_user_repo.New(this.transaction)
		this.userRepo = &userRepo
	}
	return *this.userRepo
}






