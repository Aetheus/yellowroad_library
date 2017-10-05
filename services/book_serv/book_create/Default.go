package book_create

import (
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database"
)

type DefaultBookCreateService struct {
	isAutocommit bool
	unitOfWork   uow.UnitOfWork
}
func Default(work uow.UnitOfWork, isAutocommit bool) BookCreateService {
	return DefaultBookCreateService{
		isAutocommit,
		work,
	}
}

func (this DefaultBookCreateService) CreateBook(creator entities.User, book *entities.Book) app_error.AppError {
	//TODO: do some extra checking here (eg: check if the creator is banned or not, etc)
	book.CreatorId = creator.ID
	book.FirstChapterId = database.NullInt{Int:0}

	unitOfWork := this.unitOfWork
	if err := unitOfWork.BookRepo().Insert(book); err != nil {
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