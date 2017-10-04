package user_registration_serv

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"unicode/utf8"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"yellowroad_library/database/repo/uow"
)

type DefaultUserRegistrationService struct {
	unitOfWork uow.UnitOfWork
	autocommit bool
}
var _ UserRegistrationService = DefaultUserRegistrationService{}
func Default(work uow.UnitOfWork, autocommit bool) DefaultUserRegistrationService{
	return DefaultUserRegistrationService{
		unitOfWork : work,
		autocommit : autocommit,
	}
}

func (this DefaultUserRegistrationService) Run (username string, password string, email string) (returnedUser *entities.User, returnedErr app_error.AppError) {
	if utf8.RuneCountInString(password) < 6 {
		encounteredError := app_error.New(http.StatusUnprocessableEntity, "","Password had an insufficient length (minimum 6 characters)")
		return nil, encounteredError
	}

	hashedPassword, encounteredError := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if encounteredError != nil {
		return nil, app_error.Wrap(encounteredError)
	}

	var user = entities.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	}

	userRepo := this.unitOfWork.UserRepo()
	if err := userRepo.Insert(&user); err != nil {
		return nil, app_error.Wrap(err)
	}

	if (this.autocommit){
		if err := this.unitOfWork.Commit(); err != nil {
			this.unitOfWork.Rollback();
			return nil, app_error.Wrap(err)
		}
	}

	return &user, nil
}
