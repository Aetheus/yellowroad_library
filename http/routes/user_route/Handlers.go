package user_route

import (
	"net/http"

	"yellowroad_library/services/auth_serv"

	"github.com/gin-gonic/gin"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/utils/api_response"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/database/entities"
)

type signUpForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}
func SignUp(
	c *gin.Context,
	work uow.UnitOfWork,
	authService auth_serv.AuthService,
)  {

	var user entities.User
	err := work.Auto([]uow.WorkFragment{authService}, func() app_error.AppError {
		form := signUpForm{}

		if err := c.BindJSON(&form) ; err != nil {
			var err = app_error.Wrap(err).SetHttpCode(http.StatusUnprocessableEntity)
			return err
		}

		var registrationErr app_error.AppError
		user, registrationErr = authService.RegisterUser(form.Username,form.Password,form.Email)
		if (registrationErr != nil) {
			return registrationErr
		}

		return nil
	});


	if(err != nil){
		c.JSON(api_response.ConvertErrWithCode(err))
	} else {
		c.JSON(api_response.SuccessWithCode(
			gin.H {
				"user" : user,
			},
		))
	}

}

type loginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func Login(
	c *gin.Context,
	work uow.UnitOfWork,
	authService auth_serv.AuthService,
) {

	var user entities.User
	var loginToken string
	err := work.Auto([]uow.WorkFragment{authService}, func() app_error.AppError {
		form := loginForm{}
		if err := c.BindJSON(&form); err != nil {
			var err  = app_error.Wrap(err).SetHttpCode(http.StatusUnprocessableEntity)
			return err
		}

		var loginErr app_error.AppError
		user, loginToken, loginErr = authService.LoginUser(form.Username, form.Password)
		if (loginErr != nil){
			return loginErr
		}

		return nil
	});


	if(err != nil){
		c.JSON(api_response.ConvertErrWithCode(err))
	} else {
		c.JSON(api_response.SuccessWithCode(
			gin.H{
				"user" : user,
				"token" : loginToken,
			},
		))
	}

}