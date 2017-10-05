package book_crud

import (
	"yellowroad_library/utils/app_error"
	"yellowroad_library/utils/api_response"
	"yellowroad_library/database/entities"
	"github.com/gin-gonic/gin"
	"yellowroad_library/services/auth_serv"
	"yellowroad_library/containers"
	"yellowroad_library/services/book_serv/book_create"
)

type createBookForm struct {
	Title string
	Description string
}

func CreateBookHandler(container containers.Container) gin.HandlerFunc{
	return func(c *gin.Context) {
		bookCreateServ := container.BookCreateService(nil,true)
		authServ := container.GetAuthService()

		CreateBook(c,authServ,bookCreateServ)
	}
}

func CreateBook(c *gin.Context, authService auth_serv.AuthService, createBookService book_create.BookCreateService) {
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
	if err := createBookService.CreateBook(user, &book); err != nil {
		c.JSON(api_response.ConvertErrWithCode(err))
		return
	}

	c.JSON(api_response.SuccessWithCode(
		gin.H{"book": book},
	))
}