package book_route

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
)

func FetchSingleBook(bookRepo book_repo.BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {

		book_id, err := gin_tools.GetIntParam("book_id", c)
		if err != nil {
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		book, findErr := bookRepo.FindById(book_id)
		if findErr != nil {
			c.JSON(api_response.ConvertErrWithCode(findErr))
			return
		}

		c.JSON(api_response.SuccessWithCode(book))
		return
	}
}

func FetchBooks(repository book_repo.BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		page := gin_tools.GetIntQueryOrDefault("page",1,c)
		perpage := gin_tools.GetIntQueryOrDefault("perpage",15,c)

		//TODO: actually get some search options
		results, err := repository.Paginate(page,perpage, book_repo.SearchOptions{})
		if err != nil {
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		c.JSON(api_response.SuccessWithCode(results))
		return
	}
}

type createBookForm struct {
	Title string
	Description string
}
func CreateBook(authService auth_serv.AuthService, bookService book_serv.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entities.User
		var formData createBookForm

		//Get logged in user
		user, err := authService.GetLoggedInUser(c.Copy());
		if err != nil {
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		//Get form data to create book with
		if err := c.BindJSON(&formData); err != nil {
			var err app_error.AppError = app_error.Wrap(err)
			c.JSON( api_response.ConvertErrWithCode(err) )
			return
		}

		//Create the book
		book := entities.Book {
			CreatorId: user.ID,
			Title: formData.Title,
			Description: formData.Description,
		}
		if err := bookService.CreateBook(user, &book); err != nil {
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		c.JSON(api_response.SuccessWithCode(
			gin.H{"book": book},
		))
	}
}


func DeleteBook(authService auth_serv.AuthService, bookRepo book_repo.BookRepository, bookService book_serv.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		book_id, err := gin_tools.GetIntParam("book_id",c)
		if err != nil {
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		book, err := bookRepo.FindById(book_id)
		if  err != nil{
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		user, err := authService.GetLoggedInUser(c.Copy())
		if err != nil {
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		err = bookService.DeleteBook(user, &book)
		if err != nil {
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		c.JSON(api_response.SuccessWithCode(gin.H{}))
		return
	}
}

func UpdateBook (book_repo book_repo.BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		book_id, err := gin_tools.GetIntParam("book_id",c)
		if err != nil {
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		book, err := book_repo.FindById(book_id)
		if err != nil {
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		bookForm := entities.BookForm{}
		bindErr := c.BindJSON(&bookForm)
		if bindErr != nil {
			c.JSON(api_response.ConvertErrWithCode(app_error.Wrap(bindErr)))
			return
		}

		bookForm.Apply(&book)
		err = book_repo.Update(&book)
		if err != nil {
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		c.JSON(http.StatusOK, gin.H{ "book" : book })
		return
	}
}