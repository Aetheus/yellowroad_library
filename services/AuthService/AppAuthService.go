package AuthService

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"errors"
	"unicode/utf8"

	"yellowroad_library/database/entities"
	"yellowroad_library/database/repositories/UserRepo"
	"yellowroad_library/services/TokenService"
)

type AppAuthService struct {
	userRepository UserRepo.UserRepository
	tokenService   TokenService.TokenService
}

func NewAppAuthService(userRepository UserRepo.UserRepository, tokenService TokenService.TokenService) AppAuthService {
	return AppAuthService{
		userRepository: userRepository,
		tokenService:   tokenService,
	}
}

func (service AppAuthService) RegisterUser(username string, password string, email string) (returnedUser *entities.User, returnedErr error) {
	var encounteredError error

	if utf8.RuneCountInString(password) < 6 {
		encounteredError = errors.New("Password had an insufficient length (minimum 6 characters)")
		return nil, encounteredError
	}

	hashedPassword, encounteredError := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if encounteredError != nil {
		return nil, encounteredError
	}

	var user = entities.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	}

	if error := service.userRepository.Insert(&user); error != nil {
		return nil, error
	}

	return &user, encounteredError
}

// return : user, login_token, err
func (service AppAuthService) LoginUser(username string, password string) (*entities.User, string, error) {
	var user *entities.User
	var err error

	//TODO : email as well
	user, err = service.userRepository.FindByUsername(username)
	if err != nil {
		return nil, "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", errors.New("Incorrect username or password")
	}

	token, err := service.tokenService.CreateTokenString(*user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (service AppAuthService) GetLoggedInUser(data interface{}) (*entities.User, error) {
	context, ok := data.(*gin.Context)

	if !ok {
		return nil, errors.New("Provided data was not a gin context struct")
	}

	if tokenClaim, err := getTokenClaim(context); err != nil {
		return nil, err
	} else {
		user, err := service.userRepository.FindById(tokenClaim.UserID)
		return user, err
	}
}

func getTokenClaim(c *gin.Context) (*TokenService.MyCustomClaims, error) {
	tokenClaim, exists := c.Get(TokenService.TOKEN_CLAIMS_CONTEXT_KEY)

	if !exists {
		return nil, errors.New("No token claim was provided")
	}

	claimsData := tokenClaim.(TokenService.MyCustomClaims)

	return &claimsData, nil
}
