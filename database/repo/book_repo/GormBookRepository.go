package book_repo

import (
	"yellowroad_library/database/entities"
	"github.com/jinzhu/gorm"
	"yellowroad_library/utils/app_error"
	"net/http"
	"errors"
)

type GormBookRepository struct {
	dbConn *gorm.DB
}

var _ BookRepository = GormBookRepository{} //ensure interface implementation
func NewDefault(dbConn *gorm.DB) GormBookRepository {
	return GormBookRepository{
		dbConn: dbConn,
	}
}

func (repo GormBookRepository) FindById(id int) (entities.Book, app_error.AppError) {
	var book entities.Book
	dbConn := repo.dbConn

	queryResult := dbConn.
					Select("books.*, COUNT(chapters_c.id) as chapter_count").
					Joins("LEFT JOIN chapters as chapters_c on chapters_c.book_id = books.id").
					Preload(entities.ASSOC_BOOK_CREATOR,func(db *gorm.DB) *gorm.DB {
						return db.Select("username, id")
					}).
					Preload(entities.ASSOC_BOOK_FIRST_CHAPTER).
					Where("books.id = ?", id).
					Group("books.id").
					First(&book)

	if queryResult.RecordNotFound() {
		returnedErr := app_error.Wrap(queryResult.Error).
						SetEndpointMessage("No such book found").
						SetHttpCode(http.StatusNotFound)
		return book, returnedErr
	} else if queryResult.Error != nil {
		returnedErr := app_error.Wrap(queryResult.Error)
		return book, returnedErr
	}

	return book, nil
}

func (repo GormBookRepository) Paginate(options SearchOptions) ([]entities.Book, app_error.AppError) {
	results := []entities.Book{}
	startpage := options.StartPage
	perpage := options.PerPage

	if startpage < 1 {
		startpage = 1
	}

	//TODO: use SearchOptions to do stuff like filter, order, etc
	queryResult := repo.dbConn.
						Select("books.*, COUNT(chapters_c.id) as chapter_count").
						Joins("LEFT JOIN chapters as chapters_c on chapters_c.book_id = books.id").
						Preload(entities.ASSOC_BOOK_CREATOR,func(db *gorm.DB) *gorm.DB {
							return db.Select("username, id")
						}).
						Offset( (startpage - 1) * perpage).
						Limit(perpage).
						Group("books.id").
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
