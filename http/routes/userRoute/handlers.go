package userRoute

import (
	"net/http"

	"yellowroad_library/services/authService"

	"github.com/gin-gonic/gin"
)

func Login(authService authService.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var username = c.GetString("username")
		var password = c.GetString("password")

		user, loginToken, err := authService.LoginUser(username, password)

		if (err != nil){
			//TODO : return an appropriate error code / standardize the API return format and errors with custom structs
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
			return
		}


		c.JSON(http.StatusOK, gin.H{
			"user" : user,
			"token" : loginToken,
		})
		return
	}
}