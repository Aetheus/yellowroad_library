package gorm_book_repo

import (
	"yellowroad_library/database/entities"

	"github.com/jinzhu/gorm"
	"yellowroad_library/utils/app_error"
	"net/http"
	"yellowroad_library/database/repo/book_repo"
)

type GormBookRepository struct {
	dbConn *gorm.DB
}
//ensure interface implementation
var _ book_repo.BookRepository = GormBookRepository{}

func New(dbConn *gorm.DB) GormBookRepository {
	return GormBookRepository{
		dbConn: dbConn,
	}
}

func (repo GormBookRepository) FindById(id int) (entities.Book, app_error.AppError) {
	var book entities.Book

	dbConn := repo.dbConn

	if queryResult := dbConn.Where("id = ?", id).First(&book); queryResult.Error != nil {
		var returnedErr app_error.AppError
		if queryResult.RecordNotFound() {
			returnedErr = app_error.Wrap(queryResult.Error).
							SetEndpointMessage("No such book found").
							SetHttpCode(http.StatusNotFound)
		} else {
			returnedErr = app_error.Wrap(queryResult.Error)
		}
		return book, returnedErr
	}

	return book, nil
}

func (repo GormBookRepository) Update(book *entities.Book) app_error.AppError {
	if queryResult := repo.dbConn.Save(book); queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}

func (repo GormBookRepository) Insert(book *entities.Book) app_error.AppError {
	if queryResult := repo.dbConn.Create(book); queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}
