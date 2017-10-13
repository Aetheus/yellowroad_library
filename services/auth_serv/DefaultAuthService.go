package auth_serv

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"errors"
	"unicode/utf8"

	"yellowroad_library/database/entities"
	"yellowroad_library/services/token_serv"
	"yellowroad_library/utils/app_error"
	"net/http"
	"yellowroad_library/database/repo/uow"
)

type DefaultAuthService struct {
	work uow.UnitOfWork
	tokenService token_serv.TokenService
}
//ensure interface implementation
var _ AuthService = DefaultAuthService{}

func Default(work uow.UnitOfWork, tokenService token_serv.TokenService) AuthService {
	return DefaultAuthService{
		work,
		tokenService,
	}
}
var _ AuthServiceFactory = Default

func (this DefaultAuthService) SetUnitOfWork(work uow.UnitOfWork) {
	this.work = work
}

func (service DefaultAuthService) RegisterUser(username string, password string, email string) (returnedUser entities.User, returnedErr app_error.AppError) {

	var user entities.User

	if utf8.RuneCountInString(password) < 6 {
		encounteredError := app_error.New(http.StatusUnprocessableEntity, "","Password had an insufficient length (minimum 6 characters)")
		return user, encounteredError
	}

	hashedPassword, encounteredError := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if encounteredError != nil {
		return user, app_error.Wrap(encounteredError)
	}

	user = entities.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	}

	if err := service.work.UserRepo().Insert(&user); err != nil {
		return user, app_error.Wrap(err)
	}

	return user, nil
}

// return : user, login_token, err
func (service DefaultAuthService) LoginUser(username string, password string) (entities.User, string, app_error.AppError) {
	var user entities.User
	var err error

	//TODO : email as well
	user, err = service.work.UserRepo().FindByUsername(username)
	if err != nil {
		return user, "", app_error.Wrap(err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return user, "", app_error.New(http.StatusUnauthorized, "","Incorrect username or password")
	}

	token, err := service.tokenService.CreateTokenString(user)
	if err != nil {
		return user, "", app_error.Wrap(err)
	}

	return user, token, nil
}

func (service DefaultAuthService) GetLoggedInUser(data interface{}) (entities.User, app_error.AppError) {
	var user entities.User
	var err app_error.AppError

	context, ok := data.(*gin.Context)
	if !ok {
		err := errors.New("Provided data was not a gin context struct");
		return user, app_error.Wrap(err)
	}

	tokenClaim, err := getTokenClaim(context)
	if err != nil {
		return user, app_error.Wrap(err)
	}

	user, err = service.work.UserRepo().FindById(tokenClaim.UserID)
	return user, err
}

func getTokenClaim(c *gin.Context) (token_serv.LoginClaim, app_error.AppError) {
	var tokenClaim token_serv.LoginClaim
	potentialClaim, exists := c.Get(token_serv.TOKEN_CLAIMS_CONTEXT_KEY)

	if !exists {
		err := app_error.Wrap(errors.New("No token claim was provided")).
							SetHttpCode(http.StatusUnauthorized).
							SetEndpointMessage("No login token provided");
		return tokenClaim, err
	}

	tokenClaim = potentialClaim.(token_serv.LoginClaim)

	return tokenClaim, nil
}
