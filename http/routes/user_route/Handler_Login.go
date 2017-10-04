package user_route

import (
	"net/http"

	"yellowroad_library/services/auth_serv"

	"github.com/gin-gonic/gin"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/utils/api_response"
)

type loginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func Login(authService auth_serv.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		form := loginForm{}
		if err := c.BindJSON(&form); err != nil {
			var err  = app_error.Wrap(err).SetHttpCode(http.StatusUnprocessableEntity)
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		user, loginToken, err := authService.LoginUser(form.Username, form.Password)
		if (err != nil){
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		c.JSON(api_response.SuccessWithCode(
			gin.H{
				"user" : user,
				"token" : loginToken,
			},
		))
		return
	}
}