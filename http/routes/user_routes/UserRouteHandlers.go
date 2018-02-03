package user_routes

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/utils/api_reply"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/database/entities"
	"yellowroad_library/containers"
	"yellowroad_library/utils/gin_tools"
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

		if err := gin_tools.BindJSON(&form,c) ; err != nil {
			return err
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
		if formErr := gin_tools.BindJSON(&form,c); formErr != nil {
			return formErr
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
		if formErr := gin_tools.BindJSON(&form,c); formErr != nil {
			return formErr
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