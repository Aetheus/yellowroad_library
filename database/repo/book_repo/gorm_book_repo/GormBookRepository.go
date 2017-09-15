package gorm_book_repo

import (
	"yellowroad_library/database/entities"

	"github.com/jinzhu/gorm"
	"yellowroad_library/utils/app_error"
	"net/http"
)

type GormBookRepository struct {
	dbConn *gorm.DB
}

func New(dbConn *gorm.DB) GormBookRepository {
	return GormBookRepository{
		dbConn: dbConn,
	}
}

func (repo GormBookRepository) FindById(id int) (*entities.Book, error) {
	var book entities.Book

	dbConn := repo.dbConn

	if queryResult := dbConn.Where("id = ?", id).First(&book); queryResult.Error != nil {
		var returnedErr error
		if queryResult.RecordNotFound() {
			returnedErr = app_error.Wrap(queryResult.Error).
							SetEndpointMessage("No such book found").
							SetHttpCode(http.StatusNotFound)
		} else {
			returnedErr = app_error.Wrap(queryResult.Error)
		}
		return nil, returnedErr
	}

	return &book, nil
}

func (repo GormBookRepository) Update(book *entities.Book) error {
	if queryResult := repo.dbConn.Save(book); queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}

func (repo GormBookRepository) Insert(book *entities.Book) error {
	if queryResult := repo.dbConn.Create(book); queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}
