package appAuthService

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"errors"
	"unicode/utf8"

	"yellowroad_library/database/entities"
	"yellowroad_library/database/repositories/userRepository"
	"yellowroad_library/services/tokenService"
	"yellowroad_library/utils/appError"
	"net/http"
)

type AppAuthService struct {
	userRepository userRepository.UserRepository
	tokenService   tokenService.TokenService
}

func New(userRepository userRepository.UserRepository, tokenService tokenService.TokenService) AppAuthService {
	return AppAuthService{
		userRepository: userRepository,
		tokenService:   tokenService,
	}
}

func (service AppAuthService) RegisterUser(username string, password string, email string) (returnedUser *entities.User, returnedErr error) {

	if utf8.RuneCountInString(password) < 6 {
		encounteredError := appError.New(http.StatusUnprocessableEntity, "","Password had an insufficient length (minimum 6 characters)")
		return nil, encounteredError
	}

	hashedPassword, encounteredError := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if encounteredError != nil {
		return nil, appError.Wrap(encounteredError)
	}

	var user = entities.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	}

	if err := service.userRepository.Insert(&user); err != nil {
		return nil, appError.Wrap(err)
	}

	return &user, nil
}

// return : user, login_token, err
func (service AppAuthService) LoginUser(username string, password string) (*entities.User, string, error) {
	var user *entities.User
	var err error

	//TODO : email as well
	user, err = service.userRepository.FindByUsername(username)
	if err != nil {
		return nil, "", appError.Wrap(err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", appError.New(http.StatusUnauthorized, "","Incorrect username or password")
	}

	token, err := service.tokenService.CreateTokenString(*user)
	if err != nil {
		return nil, "", appError.Wrap(err)
	}

	return user, token, nil
}

func (service AppAuthService) GetLoggedInUser(data interface{}) (*entities.User, error) {
	context, ok := data.(*gin.Context)

	if !ok {
		err := errors.New("Provided data was not a gin context struct");
		return nil, appError.Wrap(err)
	}

	if tokenClaim, err := getTokenClaim(context); err != nil {
		return nil, appError.Wrap(err)
	} else {
		user, err := service.userRepository.FindById(tokenClaim.UserID)
		return user, appError.Wrap(err)
	}
}

func getTokenClaim(c *gin.Context) (*tokenService.MyCustomClaims, error) {
	tokenClaim, exists := c.Get(tokenService.TOKEN_CLAIMS_CONTEXT_KEY)

	if !exists {
		err := appError.Wrap(errors.New("No token claim was provided")).
							SetHttpCode(http.StatusUnauthorized).
							SetEndpointMessage("No login token provided");
		return nil, err
	}

	claimsData := tokenClaim.(tokenService.MyCustomClaims)

	return &claimsData, nil
}
