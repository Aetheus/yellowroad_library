package story_route

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/services/auth_serv"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
	"yellowroad_library/services/book_serv"
	"yellowroad_library/utils/api_response"
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/utils/gin_tools"
	"net/http"
	"yellowroad_library/database/repo/uow"
)

func FetchSingleBook(c *gin.Context,work uow.UnitOfWork)  {

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

func FetchBooks(c *gin.Context, work uow.UnitOfWork) {

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

type createBookForm struct {
	Title string
	Description string
}
func CreateBook (
	c *gin.Context,
	work uow.UnitOfWork,
	authService auth_serv.AuthService,
	bookService book_serv.BookService,
) {
	var book entities.Book
	err := work.Auto([]uow.WorkFragment{authService,bookService}, func() app_error.AppError {
		var formData createBookForm

		//Get logged in user
		user, err := authService.GetLoggedInUser(c.Copy());
		if err != nil {
			return err
		}

		//Get form data to create book with
		if err := c.BindJSON(&formData); err != nil {
			var err app_error.AppError = app_error.Wrap(err)
			return err
		}

		//Create the book
		book = entities.Book {
			CreatorId: user.ID,
			Title: formData.Title,
			Description: formData.Description,
		}
		if err := bookService.CreateBook(user, &book); err != nil {
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



func DeleteBook(
	c *gin.Context,
	work uow.UnitOfWork,
	authService auth_serv.AuthService,
	bookService book_serv.BookService,
) {

	err := work.Auto([]uow.WorkFragment{authService, bookService}, func() app_error.AppError {
		book_id, err := gin_tools.GetIntParam("book_id",c)
		if err != nil {
			return err
		}

		user, err := authService.GetLoggedInUser(c.Copy())
		if err != nil {
			return err
		}

		err = bookService.DeleteBook(book_id,user)
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

//TODO : actually validate if you can update the book
func UpdateBook (
	c *gin.Context,
	work uow.UnitOfWork,
	authService auth_serv.AuthService,
	bookService book_serv.BookService,
) {

	var book entities.Book
	err := work.Auto([]uow.WorkFragment{authService, bookService}, func() app_error.AppError {
		bookRepo := work.BookRepo()

		book_id, err := gin_tools.GetIntParam("book_id",c)
		if err != nil {
			return err
		}

		book, err = bookRepo.FindById(book_id)
		if err != nil {
			return err
		}

		bookForm := entities.BookForm{}
		bindErr := c.BindJSON(&bookForm)
		if bindErr != nil {
			return app_error.Wrap(bindErr)
		}

		bookForm.Apply(&book)
		err = bookRepo.Update(&book)
		if err != nil {
			return err
		}
		return nil
	})

	if ( err != nil ){
		c.JSON(api_response.ConvertErrWithCode(err))
	}else {
		c.JSON(http.StatusOK, gin.H{"book": book})
	}
}