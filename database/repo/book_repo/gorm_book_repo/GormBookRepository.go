package gorm_book_repo

import (
	"yellowroad_library/database/entities"

	"github.com/jinzhu/gorm"
	"yellowroad_library/utils/app_error"
	"net/http"
	"yellowroad_library/database/repo/book_repo"
	"errors"
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

func preloadAssociations(dbConn *gorm.DB) *gorm.DB{
	for i := 0; i < len(entities.BookAssociations); i++ {
		dbConn = dbConn.Preload(entities.BookAssociations[i])
	}

	return dbConn
}

func (repo GormBookRepository) FindById(id int) (entities.Book, app_error.AppError) {
	var book entities.Book

	dbConn := repo.dbConn

	queryResult := preloadAssociations(dbConn).
						Where("id = ?", id).
						First(&book)

	if queryResult.Error != nil {
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

func (repo GormBookRepository) Paginate(startpage int, perpage int, options book_repo.SearchOptions) ([]entities.Book, app_error.AppError) {
	results := []entities.Book{}

	if startpage < 1 {
		startpage = 1
	}

	//TODO: use SearchOptions to do stuff like filter, order, etc
	queryResult := repo.dbConn.
						Offset( (startpage - 1) * perpage).
						Limit(perpage).
						Find(&results)

	if queryResult.Error != nil {
		var returnedErr app_error.AppError
		if queryResult.RecordNotFound() {
			returnedErr = app_error.Wrap(queryResult.Error).
				SetEndpointMessage("No books found").
				SetHttpCode(http.StatusNotFound)
		} else {
			returnedErr = app_error.Wrap(queryResult.Error)
		}
		return results, returnedErr
	}

	return results, nil
}

func (repo GormBookRepository) Update(book *entities.Book) app_error.AppError {
	queryResult := repo.dbConn.
					Set("gorm:save_associations", false).	//no magic! let the individual objects be saved on their own!
					Save(book)

	if queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}

func (repo GormBookRepository) Insert(book *entities.Book) app_error.AppError {
	queryResult := repo.dbConn.
					Set("gorm:save_associations", false).	//no magic! let the individual objects be saved on their own!
					Create(book)

	if queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}

//soft delete
func (repo GormBookRepository) Delete(book *entities.Book) app_error.AppError {
	if (book.ID == 0){
		err := errors.New("Invalid primary key value of 0 while attempting to delete")
		return app_error.Wrap(err)
	}

	if queryResult := repo.dbConn.Delete(book); queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}
