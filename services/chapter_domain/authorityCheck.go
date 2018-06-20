package chapter_domain

import (
	"yellowroad_library/utils/app_error"
	"net/http"
	"yellowroad_library/database/repo/book_repo"
)

type authorityCheck struct {
	bookRepo book_repo.BookRepository
}

const auth_create_chapter = "CREATE_CHAPTER"
const auth_create_path = "CREATE_PATH"
const auth_update_chapter = "UPDATE_CHAPTER"
const auth_update_path = "UPDATE_PATH"
const auth_delete_chapter = "DELETE_CHAPTER"
const auth_delete_path = "DELETE_PATH"

const unauthorizedErrMessage = "You are not authorized to perform this action!"

func isValidAction(action string) (isValid bool) {
	switch(action){
		case auth_create_chapter,
				auth_create_path,
			auth_update_chapter,
				auth_update_path,
			auth_delete_chapter,
				auth_delete_path:
			isValid = true
		default:
			isValid = false
	}
	return isValid
}

func newAuthorityCheck(repository book_repo.BookRepository) authorityCheck {
	return authorityCheck{repository}
}

func (this authorityCheck) Execute(action string, userId int, bookId int) app_error.AppError {
	//TODO: more advanced permission system; and make use of "action"
	if (!isValidAction(action)){
		invalidActionMessage := "'" + action + "' is not a valid action!"
		return app_error.New(http.StatusBadRequest,invalidActionMessage,unauthorizedErrMessage)
	}

	book, err := this.bookRepo.FindById(bookId)
	if(err != nil){
		return err
	}
	if (book.CreatorId != userId){
		return app_error.New(http.StatusUnauthorized,"",unauthorizedErrMessage)
	}
	return nil
}