package book_crud

import (
	"yellowroad_library/services/auth_serv"
	"github.com/gin-gonic/gin"
	"yellowroad_library/utils/gin_tools"
	"yellowroad_library/utils/api_response"
	"yellowroad_library/services/book_serv/book_delete"
	"yellowroad_library/containers"
)

func DeleteBookHandler(container containers.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		authServ := container.GetAuthService()
		bookDeleteServ := container.BookDeleteService(nil,true)

		DeleteBook(c,authServ,bookDeleteServ)
	}
}
func DeleteBook(c *gin.Context, authService auth_serv.AuthService, bookDeleteServ book_delete.BookDeleteService) {
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

	err = bookDeleteServ.DeleteBook(book_id, user)
	if err != nil {
		c.JSON(api_response.ConvertErrWithCode(err))
		return
	}

	c.JSON(api_response.SuccessWithCode(gin.H{}))
	return
}