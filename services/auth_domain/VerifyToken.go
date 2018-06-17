package auth_domain

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/repo/user_repo"
	"github.com/dgrijalva/jwt-go"
)

func NewVerifyToken(tokenStringValidator TokenStringValidator,userRepo user_repo.UserRepository) VerifyToken{
	return VerifyToken{tokenStringValidator ,userRepo}
}

type VerifyToken struct {
	tokenStringValidator TokenStringValidator
	userRepo user_repo.UserRepository
}

func (this VerifyToken) Execute(tokenString string) (user entities.User,err app_error.AppError) {
	claim, err := this.tokenStringValidator.ValidateTokenString(tokenString)
	if (err != nil){
		return
	}
	user, err = this.userRepo.FindById(claim.UserID)
	return
}

type LoginClaim struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

type TokenStringValidator interface{
	ValidateTokenString(string) (LoginClaim, app_error.AppError)
}