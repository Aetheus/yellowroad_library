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
	bookRepo := work.BookRepo()

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

	work.Commit()
	c.JSON(api_response.SuccessWithCode(book))
	return
}

func FetchBooks(c *gin.Context, work uow.UnitOfWork) {
	repository := work.BookRepo()

	page := gin_tools.GetIntQueryOrDefault("page",1,c)
	perpage := gin_tools.GetIntQueryOrDefault("perpage",15,c)

	//TODO: actually get some search options
	results, err := repository.Paginate(page,perpage, book_repo.SearchOptions{})
	if err != nil {
		c.JSON(api_response.ConvertErrWithCode(err))
		return
	}

	work.Commit()
	c.JSON(api_response.SuccessWithCode(results))
	return
}

type createBookForm struct {
	Title string
	Description string
}
func CreateBook (
	c *gin.Context,
	work uow.UnitOfWork,
	authService auth_serv.AuthService, bookService book_serv.BookService,
) {
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

	work.Commit()
	c.JSON(api_response.SuccessWithCode(
		gin.H{"book": book},
	))
}



func DeleteBook(
	c *gin.Context,
	work uow.UnitOfWork,
	authService auth_serv.AuthService,
	bookService book_serv.BookService,
) {
	book_id, err := gin_tools.GetIntParam("book_id",c)
	if err != nil {
		c.JSON(api_response.ConvertErrWithCode(err))
		return
	}

	user, err := authService.GetLoggedInUser(c.Copy())
	if err != nil {
		c.JSON(api_response.ConvertErrWithCode(err))
		return
	}

	err = bookService.DeleteBook(book_id,user)
	if err != nil {
		c.JSON(api_response.ConvertErrWithCode(err))
		return
	}

	work.Commit()
	c.JSON(api_response.SuccessWithCode(gin.H{}))
	return
}

//TODO : actually validate if you can update the book
func UpdateBook (
	c *gin.Context,
	work uow.UnitOfWork,
	authService auth_serv.AuthService,
	bookService book_serv.BookService,
) {
	bookRepo := work.BookRepo()

	book_id, err := gin_tools.GetIntParam("book_id",c)
	if err != nil {
		c.JSON(api_response.ConvertErrWithCode(err))
		return
	}

	book, err := bookRepo.FindById(book_id)
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
	err = bookRepo.Update(&book)
	if err != nil {
		c.JSON(api_response.ConvertErrWithCode(err))
		return
	}

	work.Commit()
	c.JSON(http.StatusOK, gin.H{ "book" : book })
	return
}