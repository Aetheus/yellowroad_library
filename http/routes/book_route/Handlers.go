package book_route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"yellowroad_library/services/auth_serv"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
	"yellowroad_library/services/book_serv"
	"fmt"
)

func FetchSingleBook() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Book": "nah jk"})
		return
	}
}

func CreateBook(authService auth_serv.AuthService, bookService book_serv.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *entities.User
		var formData createBookForm

		//Get logged in user
		user, err := authService.GetLoggedInUser(c.Copy());
		if err != nil {
			c.JSON(err.HttpCode(), gin.H { "message": err.EndpointMessage(), })
			return
		}

		//Get form data to create book with
		if err := c.BindJSON(&formData); err != nil {
			var err app_error.AppError = app_error.Wrap(err)
			fmt.Println(err.Stacktrace())
			c.JSON(err.HttpCode(),gin.H { "message": err.EndpointMessage()} )
			return
		}

		//Create the book
		book := entities.Book {
			CreatorId: user.ID, Title: formData.Title, Description: formData.Description,
		}
		if err := bookService.CreateBook(*user, &book); err != nil {
			c.JSON(err.HttpCode(), gin.H { "message": err.EndpointMessage()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"book": book})
	}
}
type createBookForm struct {
	Title string
	Description string
}