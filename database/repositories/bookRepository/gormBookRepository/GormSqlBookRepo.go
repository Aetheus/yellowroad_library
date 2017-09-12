package gormBookRepository

import (
	"errors"
	"yellowroad_library/database/entities"

	"github.com/jinzhu/gorm"
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

	if results := dbConn.Where("id = ?", id).First(&book); results.Error != nil {
		var returnedErr error
		if results.RecordNotFound() {
			returnedErr = errors.New("No such book")
		} else {
			returnedErr = errors.New("Unknown error occured")
		}
		return nil, returnedErr
	}

	return &book, nil
}

func (repo GormBookRepository) Update(book *entities.Book) error {
	if queryResult := repo.dbConn.Save(book); queryResult.Error != nil {
		return queryResult.Error
	}

	return nil
}

func (repo GormBookRepository) Insert(book *entities.Book) error {

	if result := repo.dbConn.Create(book); result.Error != nil {
		return result.Error
	}

	return nil
}
