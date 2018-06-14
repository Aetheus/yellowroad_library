package auth_serv

import (
	"golang.org/x/crypto/bcrypt"

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
	user, findErr := service.work.UserRepo().FindByUsername(username)
	if findErr != nil {
		if (findErr.HttpCode() == http.StatusNotFound){
			findErr = app_error.New(http.StatusUnauthorized, "","Incorrect username or password")
		}
		return user, "", findErr
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

func (service DefaultAuthService) VerifyToken(tokenString string) (user entities.User,err app_error.AppError) {
	claim, err := service.tokenService.ValidateTokenString(tokenString)
	if (err != nil){
		return
	}
	user, err = service.work.UserRepo().FindById(claim.UserID)
	return
}

func (service DefaultAuthService) GetLoggedInUser(loginClaimAdapter LoginClaimExtractor) (entities.User, app_error.AppError) {
	var user entities.User
	var err app_error.AppError

	tokenClaim, err := loginClaimAdapter.GetLoginClaim()
	if err != nil {
		return user, app_error.Wrap(err)
	}

	user, err = service.work.UserRepo().FindById(tokenClaim.UserID)
	return user, err
}