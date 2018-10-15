package repo

import (
	"fmt"
	"net/http"
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
			RETURNING id, title, description, created_at, updated_at
		`, params.Title, params.Description, nowTime, nowTime).
		Scan(&book.ID, &book.Title, &book.Description, &book.CreatedAt, &book.UpdatedAt)

	if err != nil {
		appErr = app_error.Wrap(err)
	}

	return book, appErr
}


func (this *BookRepository) Delete(id int) app_error.AppError {
	res, err := this.DB.Exec(`
		DELETE FROM books WHERE id = $1;
	`, id)

	if err != nil {
		return app_error.Wrap(err)
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		errMessage := fmt.Sprintf("Book with id of %d was not found", id)
		return app_error.ClientError(http.StatusNotFound, errMessage)
	}

	return nil
}


//type BookSearchParams struct {
//	StartPage int
//	PerPage int
//}
//func (this *BookRepository) Search(options BookSearchParams) ([]entities.Book, app_error.AppError) {
//
//}