package uow

import (
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/database/repo/chapter_repo"
	"yellowroad_library/database/repo/chapterpath_repo"
	"yellowroad_library/database/repo/user_repo"
	"yellowroad_library/utils/app_error"
	"github.com/jinzhu/gorm"
	"yellowroad_library/database/repo/btagvote_repo"
	"yellowroad_library/database/repo/btagvotecount_repo"
)

type WorkFragment interface {
	SetUnitOfWork(work UnitOfWork)
}

//go:generate moq -out Mock.go . UnitOfWork
type UnitOfWork interface {
	BookRepo() (book_repo.BookRepository)
	ChapterRepo() (chapter_repo.ChapterRepository)
	ChapterPathRepo() (chapterpath_repo.ChapterPathRepository)
	UserRepo() (user_repo.UserRepository)
	BookTagVoteRepo() (btagvote_repo.BookTagVoteRepository)
	BookTagVoteCountRepo() (btagvotecount_repo.BookTagVoteCountRepository)

	AutoCommit([]WorkFragment, func() app_error.AppError) app_error.AppError
	Commit() (app_error.AppError)
	Rollback() (app_error.AppError)
}

type AppUnitOfWork struct {
	bookRepo *book_repo.BookRepository
	chapterRepo *chapter_repo.ChapterRepository
	chapterPathRepo *chapterpath_repo.ChapterPathRepository
	userRepo *user_repo.UserRepository
	bookTagRepo *btagvote_repo.BookTagVoteRepository
	bookTagCountRepo *btagvotecount_repo.BookTagVoteCountRepository

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

//TODO: remove the workfragment concept entirely
func (this AppUnitOfWork) AutoCommit(fragments []WorkFragment,callback func() app_error.AppError) app_error.AppError {
	for _, fragment := range fragments {
		fragment.SetUnitOfWork(this)
	}

	err := callback()
	if (err != nil) {
		rollbackErr := this.Rollback()

		if(rollbackErr != nil){
			return rollbackErr
		}else {
			return err
		}
	}

	commitErr := this.Commit()
	return commitErr
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
		var bookRepo book_repo.BookRepository = book_repo.NewDefault(this.transaction)
		this.bookRepo = &bookRepo
	}
	return *this.bookRepo
}

func (this AppUnitOfWork) ChapterRepo() (chapter_repo.ChapterRepository) {
	if this.chapterRepo == nil {
		var chapterRepo chapter_repo.ChapterRepository = chapter_repo.NewDefault(this.transaction)
		this.chapterRepo = &chapterRepo
	}
	return *this.chapterRepo
}

func (this AppUnitOfWork) ChapterPathRepo() (chapterpath_repo.ChapterPathRepository) {
	if this.chapterPathRepo == nil {
		var chapterPathRepo chapterpath_repo.ChapterPathRepository =
			chapterpath_repo.NewDefault(this.transaction)
		this.chapterPathRepo = &chapterPathRepo
	}
	return *this.chapterPathRepo
}

func (this AppUnitOfWork) UserRepo() (user_repo.UserRepository) {
	if this.userRepo == nil {
		var userRepo user_repo.UserRepository = user_repo.NewDefault(this.transaction)
		this.userRepo = &userRepo
	}
	return *this.userRepo
}


func (this AppUnitOfWork) BookTagVoteRepo() (btagvote_repo.BookTagVoteRepository){
	if this.bookTagRepo == nil {
		var bookTagRepo btagvote_repo.BookTagVoteRepository = btagvote_repo.NewDefault(this.transaction)
		this.bookTagRepo = &bookTagRepo
	}
	return *this.bookTagRepo
}
func (this AppUnitOfWork) BookTagVoteCountRepo() (btagvotecount_repo.BookTagVoteCountRepository){
	if this.bookTagCountRepo == nil {
		var bookTagRepo btagvotecount_repo.BookTagVoteCountRepository = btagvotecount_repo.NewDefault(this.transaction)
		this.bookTagCountRepo = &bookTagRepo
	}
	return *this.bookTagCountRepo
}





