package userRoute

import (
	"net/http"

	"yellowroad_library/services/authService"

	"github.com/gin-gonic/gin"
)

type signUpForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}
func SignUp(authService authService.AuthService) gin.HandlerFunc {
	return func (c *gin.Context) {
		form := signUpForm{}

		if err := c.BindJSON(&form) ; err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H { "error" : err.Error() })
			return
		}

		user, err := authService.RegisterUser(form.Username,form.Password,form.Email)

		if (err != nil) {
			c.JSON(http.StatusUnprocessableEntity, gin.H { "error" : err.Error() })
			return
		}

		c.JSON(http.StatusOK, gin.H {
			"user" : user,
		})
		return
	}
}

type loginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func Login(authService authService.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		form := loginForm{}
		if err := c.BindJSON(&form); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H { "error" : err.Error() })
			return
		}

		user, loginToken, err := authService.LoginUser(form.Username, form.Password)

		if (err != nil){
			//TODO : return an appropriate error code / standardize the API return format and errors with custom structs
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}


		c.JSON(http.StatusOK, gin.H{
			"user" : user,
			"token" : loginToken,
		})
		return
	}
}