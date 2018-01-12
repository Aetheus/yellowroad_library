package user_routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/utils/api_reply"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/database/entities"
	"yellowroad_library/containers"
)

type UserRouteHandlers struct {
	Container containers.Container
}

type signUpForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}
func (this UserRouteHandlers) SignUp(c *gin.Context)  {
	/*Dependencies**************/
	work := this.Container.UnitOfWork()
	authService := this.Container.AuthService(work)
	/***************************/


	var user entities.User
	err := work.AutoCommit([]uow.WorkFragment{authService}, func() (errOrNil app_error.AppError) {
		form := signUpForm{}

		if err := c.BindJSON(&form) ; err != nil {
			return app_error.Wrap(err).SetHttpCode(http.StatusUnprocessableEntity)
		}

		user, errOrNil = authService.RegisterUser(form.Username,form.Password,form.Email)
		return errOrNil
	});


	if(err != nil){
		api_reply.Failure(c, err)
	} else {
		api_reply.Success(c, gin.H{"user" : user})
	}
}

type loginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func (this UserRouteHandlers) Login(c *gin.Context) {
	/*Dependencies**************/
	work := this.Container.UnitOfWork()
	authService := this.Container.AuthService(work)
	/***************************/


	var user entities.User
	var loginToken string
	err := work.AutoCommit([]uow.WorkFragment{authService}, func() (errOrNil app_error.AppError) {
		form := loginForm{}
		if formErr := c.BindJSON(&form); formErr != nil {
			return app_error.Wrap(formErr).SetHttpCode(http.StatusUnprocessableEntity)
		}

		user, loginToken, errOrNil = authService.LoginUser(form.Username, form.Password)
		return errOrNil
	});


	if(err != nil){
		api_reply.Failure(c, err)
	} else {
		api_reply.Success(c, gin.H{ "user" : user, "token" : loginToken})
	}
}

type verifyTokenForm struct {
	TokenString string `json:"auth_token"`
}
func (this UserRouteHandlers) VerifyToken(c *gin.Context) {
	/*Dependencies**************/
	work := this.Container.UnitOfWork()
	authService := this.Container.AuthService(work)
	/***************************/


	var user entities.User
	var form verifyTokenForm
	err := work.AutoCommit([]uow.WorkFragment{authService}, func () (errOrNil app_error.AppError){
		if formErr := c.BindJSON(&form); formErr != nil {
			return app_error.Wrap(formErr).SetHttpCode(http.StatusUnprocessableEntity)
		}

		user, errOrNil = authService.VerifyToken(form.TokenString)
		return errOrNil
	})


	if(err != nil){
		api_reply.Failure(c, err)
	} else {
		api_reply.Success(c,gin.H{ "user" : user, "token" : form.TokenString})
	}
}