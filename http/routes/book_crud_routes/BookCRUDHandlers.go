package book_crud_routes

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/api_reply"
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/utils/gin_tools"
	"yellowroad_library/http/middleware/auth_middleware"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/services/book_domain"
)

type BookCrudContainer interface {
	UnitOfWork() uow.UnitOfWork
	CreateBook(uow.UnitOfWork) book_domain.CreateBook
	DeleteBook(uow.UnitOfWork) book_domain.DeleteBook
	UpdateBook(uow.UnitOfWork) book_domain.UpdateBook
}
type BookCrudHandlers struct {
	Container BookCrudContainer
}

func (this BookCrudHandlers) FetchSingleBook(c *gin.Context)  {
	/*Dependencies**************/
	work := this.Container.UnitOfWork()
	/***************************/

	var book entities.Book
	err := work.AutoCommit(func() app_error.AppError {
		bookRepo := work.BookRepo()

		book_id, err := gin_tools.GetIntParam("book_id", c)
		if err != nil {
			return err
		}

		book, err = bookRepo.FindById(book_id)
		return err
	})

	if ( err != nil ){
		api_reply.Failure(c,err)
	}else {
		api_reply.Success(c,book)
	}
}

func (this BookCrudHandlers) FetchBooks(c *gin.Context) {
	/*Dependencies**************/
	work := this.Container.UnitOfWork()
	/***************************/

	var results []entities.Book
	err := work.AutoCommit(func() app_error.AppError {
		repository := work.BookRepo()

		page := gin_tools.GetIntQueryOrDefault("page",1,c)
		perpage := gin_tools.GetIntQueryOrDefault("perpage",15,c)

		var paginateErr app_error.AppError
		results, paginateErr = repository.Paginate(book_repo.SearchOptions{
			StartPage: page,
			PerPage: perpage,
		})
		if paginateErr != nil {
			return paginateErr
		}
		return nil
	})

	if ( err != nil ){
		api_reply.Failure(c,err)
	}else {
		api_reply.Success(c,results)
	}
}

func (this BookCrudHandlers) CreateBook (c *gin.Context) {
	/*Dependencies**************/
	work := this.Container.UnitOfWork()
	createBook := this.Container.CreateBook(work)
	/***************************/

	var book entities.Book
	err := work.AutoCommit(func() app_error.AppError {
		var form entities.Book_CreationForm

		//Get logged in user
		user, err := auth_middleware.GetUser(c)
		if err != nil {
			return err
		}

		//Get form data to create book with
		if err := gin_tools.BindJSON(&form,c); err != nil {
			return err
		}

		book, err = createBook.Execute(user, form)
		if err != nil {
			return err
		}
		return nil
	})


	if ( err != nil ){
		api_reply.Failure(c,err)
	}else {
		api_reply.Success(c,gin.H{"book": book})
	}
}

func (this BookCrudHandlers) DeleteBook (c *gin.Context) {
	/*Dependencies**************/
	work := this.Container.UnitOfWork()
	deleteBook := this.Container.DeleteBook(work)
	/***************************/

	err := work.AutoCommit(func() app_error.AppError {
		book_id, err := gin_tools.GetIntParam("book_id",c)
		if err != nil {
			return err
		}

		user, err := auth_middleware.GetUser(c)
		if err != nil {
			return err
		}

		err = deleteBook.Execute(user, book_id)
		if err != nil {
			return err
		}
		return nil
	})


	if ( err != nil ){
		api_reply.Failure(c,err)
	}else {
		api_reply.Success(c,gin.H{})
	}
}

func (this BookCrudHandlers) UpdateBook (c *gin.Context) {
	/*Dependencies**************/
	work := this.Container.UnitOfWork()
	updateBook := this.Container.UpdateBook(work)
	/***************************/

	var book entities.Book
	err := work.AutoCommit(func() app_error.AppError {
		var form entities.Book_UpdateForm

		book_id, err := gin_tools.GetIntParam("book_id",c)
		if err != nil {
			return err
		}

		err = gin_tools.BindJSON(&form,c)
		if err != nil {
			return err
		}

		user, err := auth_middleware.GetUser(c)
		if err != nil {
			return err
		}

		book, err = updateBook.Execute(user,book_id,form)
		return err
	})

	if ( err != nil ){
		api_reply.Failure(c,err)
	}else {
		api_reply.Success(c,gin.H{"book": book})
	}
}