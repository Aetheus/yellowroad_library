package user_routes

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/utils/api_reply"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/gin_tools"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/domain/auth_domain"
)


type UserContainer interface {
	UnitOfWork() uow.UnitOfWork
	RegisterUser(uow.UnitOfWork) auth_domain.RegisterUser
	LoginUser(uow.UnitOfWork) auth_domain.LoginUser
	VerifyToken() auth_domain.VerifyToken
}
type UserRouteHandlers struct {
	Container UserContainer
}

type signUpForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}
func (this UserRouteHandlers) SignUp(c *gin.Context)  {
	/*Dependencies**************/
	work := this.Container.UnitOfWork()
	registerUser := this.Container.RegisterUser(work)
	/***************************/

	var user entities.User
	var loginToken string
	err := work.AutoCommit(func() (errOrNil app_error.AppError) {
		form := signUpForm{}

		if err := gin_tools.BindJSON(&form,c) ; err != nil {
			return err
		}

		user, loginToken, errOrNil = registerUser.Execute(form.Username,form.Password,form.Email)
		return errOrNil
	});


	if(err != nil){
		api_reply.Failure(c, err)
	} else {
		api_reply.Success(c, gin.H{"user" : user, "token": loginToken})
	}
}

type loginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func (this UserRouteHandlers) Login(c *gin.Context) {
	/*Dependencies**************/
	work := this.Container.UnitOfWork()
	loginUser := this.Container.LoginUser(work)
	/***************************/

	var user entities.User
	var loginToken string
	err := work.AutoCommit(func() (errOrNil app_error.AppError) {
		form := loginForm{}
		if formErr := gin_tools.BindJSON(&form,c); formErr != nil {
			return formErr
		}

		user, loginToken, errOrNil = loginUser.Execute(form.Username, form.Password)
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
	verifyToken := this.Container.VerifyToken()
	/***************************/


	var user entities.User
	var form verifyTokenForm
	err := work.AutoCommit(func () (errOrNil app_error.AppError){
		if formErr := gin_tools.BindJSON(&form,c); formErr != nil {
			return formErr
		}

		user, errOrNil = verifyToken.Execute(form.TokenString)
		return errOrNil
	})


	if(err != nil){
		api_reply.Failure(c, err)
	} else {
		api_reply.Success(c,gin.H{ "user" : user, "token" : form.TokenString})
	}
}