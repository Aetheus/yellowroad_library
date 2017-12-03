package book_crud_routes

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/api_response"
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/utils/gin_tools"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/containers"
)

type BookCrudHandlers struct {
	container containers.Container
}

func (this BookCrudHandlers) FetchSingleBook(c *gin.Context)  {
	//dependencies
	work := this.container.UnitOfWork()

	var book entities.Book
	err := work.Auto([]uow.WorkFragment{}, func() app_error.AppError {
		bookRepo := work.BookRepo()

		book_id, err := gin_tools.GetIntParam("book_id", c)
		if err != nil {
			c.JSON(api_response.ConvertErrWithCode(err))
			return err
		}

		var findErr app_error.AppError
		book, findErr = bookRepo.FindById(book_id)
		if findErr != nil {
			return err
		}

		return nil
	})

	if ( err != nil ){
		c.JSON(api_response.ConvertErrWithCode(err))
	}else {
		c.JSON(api_response.SuccessWithCode(book))
	}
}

func (this BookCrudHandlers) FetchBooks(c *gin.Context) {
	//dependencies
	work := this.container.UnitOfWork()

	var results []entities.Book
	err := work.Auto([]uow.WorkFragment{}, func() app_error.AppError {
		repository := work.BookRepo()

		page := gin_tools.GetIntQueryOrDefault("page",1,c)
		perpage := gin_tools.GetIntQueryOrDefault("perpage",15,c)

		//TODO: actually get some search options
		var paginateErr app_error.AppError
		results, paginateErr = repository.Paginate(page,perpage, book_repo.SearchOptions{})
		if paginateErr != nil {
			return paginateErr
		}
		return nil
	})

	if ( err != nil ){
		c.JSON(api_response.ConvertErrWithCode(err))
	}else {
		c.JSON(api_response.SuccessWithCode(results))
	}
}


func (this BookCrudHandlers) CreateBook (c *gin.Context) {
	//dependencies
	work := this.container.UnitOfWork()
	authService := this.container.AuthService(work)
	bookService := this.container.BookService(work)

	var book entities.Book
	err := work.Auto([]uow.WorkFragment{authService,bookService}, func() app_error.AppError {
		var form entities.Book_CreationForm

		//Get logged in user
		user, err := authService.GetLoggedInUser(c);
		if err != nil {
			return err
		}

		//Get form data to create book with
		if err := c.BindJSON(&form); err != nil {
			var err app_error.AppError = app_error.Wrap(err)
			return err
		}

		book, err = bookService.CreateBook(user, form)
		if err != nil {
			return err
		}
		return nil
	})


	if ( err != nil ){
		c.JSON(api_response.ConvertErrWithCode(err))
	}else {
		c.JSON(api_response.SuccessWithCode(
			gin.H{"book": book},
		))
	}
}



func (this BookCrudHandlers) DeleteBook (c *gin.Context) {
	//dependencies
	work := this.container.UnitOfWork()
	authService := this.container.AuthService(work)
	bookService := this.container.BookService(work)

	err := work.Auto([]uow.WorkFragment{authService, bookService}, func() app_error.AppError {
		book_id, err := gin_tools.GetIntParam("book_id",c)
		if err != nil {
			return err
		}

		user, err := authService.GetLoggedInUser(c)
		if err != nil {
			return err
		}

		err = bookService.DeleteBook(user, book_id)
		if err != nil {
			return err
		}
		return nil
	})


	if ( err != nil ){
		c.JSON(api_response.ConvertErrWithCode(err))
	}else {
		c.JSON(api_response.SuccessWithCode(gin.H{}))
	}
}

func (this BookCrudHandlers) UpdateBook (c *gin.Context) {
	//dependencies
	work := this.container.UnitOfWork()
	authService := this.container.AuthService(work)
	bookService := this.container.BookService(work)

	var book entities.Book
	err := work.Auto([]uow.WorkFragment{authService, bookService}, func() app_error.AppError {
		var form entities.Book_UpdateForm

		book_id, err := gin_tools.GetIntParam("book_id",c)
		if err != nil {
			return err
		}

		err = gin_tools.JSON(&form,c)
		if err != nil {
			return err
		}

		user, err := authService.GetLoggedInUser(c)
		if err != nil {
			return err
		}

		book, err = bookService.UpdateBook(user,book_id,form)
		if err != nil {
			return err
		}
		return nil
	})

	if ( err != nil ){
		c.JSON(api_response.ConvertErrWithCode(err))
	}else {
		c.JSON(api_response.SuccessWithCode(gin.H{"book": book}))
	}
}