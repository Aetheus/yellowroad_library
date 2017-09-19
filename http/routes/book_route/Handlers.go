package book_route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"yellowroad_library/services/auth_serv"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
	"yellowroad_library/services/book_serv"
	"yellowroad_library/utils/api_response"
)

func FetchSingleBook() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Book": "nah jk"})
		return
	}
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
type createBookForm struct {
	Title string
	Description string
}