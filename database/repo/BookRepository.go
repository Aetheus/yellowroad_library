package repo

import (
	"time"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"database/sql"
)

type BookRepository struct {
	DB *sql.Tx
}

//func (this *BookRepository) FindById(id int) (entities.Book, app_error.AppError) {
//
//}
//
//func (this *BookRepository) Update(book *entities.Book) app_error.AppError {
//
//}

type BookInsertParams struct {
	Title string			`json:"title"`
	Description string		`json:"description"`
}
func (this *BookRepository) Insert(params BookInsertParams) (entities.Book, app_error.AppError) {
	var book entities.Book
	var appErr app_error.AppError

	nowTime := time.Now()

	err := this.DB.
		QueryRow(`
			INSERT INTO books(title, description, created_at, updated_at)
			VALUES($1,$2,$3,$4)
			RETURNING id, title, description, created_at
		`, params.Title, params.Description, nowTime, nowTime).
		Scan(&book.ID, &book.Title, &book.Description, &book.CreatedAt)

	if err != nil {
		appErr = app_error.Wrap(err)
	}

	return book, appErr
}

//func (this *BookRepository) Delete(*entities.Book) app_error.AppError {
//
//}
//
//type BookSearchParams struct {
//	StartPage int
//	PerPage int
//}
//func (this *BookRepository) Search(options BookSearchParams) ([]entities.Book, app_error.AppError) {
//
//}