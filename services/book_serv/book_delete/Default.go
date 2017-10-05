package book_delete

import (
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"net/http"
)

type DefaultBookDeleteService struct {
	isAutocommit bool
	unitOfWork   uow.UnitOfWork
}
func Default(work uow.UnitOfWork, isAutocommit bool) BookDeleteService {
	return DefaultBookDeleteService{
		isAutocommit,
		work,
	}
}

func (this DefaultBookDeleteService) DeleteBook(book_id int, instigator entities.User) app_error.AppError {
	unitOfWork := this.unitOfWork
	book, err := unitOfWork.BookRepo().FindById(book_id)
	if  err != nil{
		return app_error.Wrap(err)
	}

	//TODO: if we implement a "contributors" system, then this should do more checking later on
	if (book.CreatorId != instigator.ID){
		return app_error.New(http.StatusUnauthorized,
			"","You are not authorized to delete this book!")
	}

	if err := unitOfWork.BookRepo().Delete(&book); err != nil {
		if (this.isAutocommit){
			unitOfWork.Rollback()
		}
		return app_error.Wrap(err)
	}


	if (this.isAutocommit){
		unitOfWork.Commit()
	}
	return nil
}
