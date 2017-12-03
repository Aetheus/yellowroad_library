package user_routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/utils/api_response"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/database/entities"
	"yellowroad_library/containers"
)

type UserRouteHandlers struct {
	container containers.Container
}


type signUpForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}
func (this UserRouteHandlers) SignUp(c *gin.Context)  {
	//dependencies
	work := this.container.UnitOfWork()
	authService := this.container.AuthService(work)

	var user entities.User
	err := work.AutoCommit([]uow.WorkFragment{authService}, func() app_error.AppError {
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
func (this UserRouteHandlers) Login(c *gin.Context) {
	//dependencies
	work := this.container.UnitOfWork()
	authService := this.container.AuthService(work)

	var user entities.User
	var loginToken string
	err := work.AutoCommit([]uow.WorkFragment{authService}, func() app_error.AppError {
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

type verifyTokenForm struct {
	TokenString string `json:"auth_token"`
}
func (this UserRouteHandlers) VerifyToken(c *gin.Context) {
	//dependencies
	work := this.container.UnitOfWork()
	authService := this.container.AuthService(work)

	var user entities.User
	var form verifyTokenForm
	err := work.AutoCommit([]uow.WorkFragment{authService}, func () (err app_error.AppError){
		if formErr := c.BindJSON(&form); formErr != nil {
			return app_error.Wrap(formErr).SetHttpCode(http.StatusUnprocessableEntity)
		}

		user, err = authService.VerifyToken(form.TokenString)
		if (err != nil){
			return err
		}

		return nil
	})

	if(err != nil){
		c.JSON(api_response.ConvertErrWithCode(err))
	} else {
		c.JSON(api_response.SuccessWithCode(
			gin.H{
				"user" : user,
				"token" : form.TokenString,
			},
		))
	}
}