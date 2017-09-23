package gorm_user_repo

import (
	"yellowroad_library/database/entities"

	"github.com/jinzhu/gorm"
	"yellowroad_library/utils/app_error"
	"net/http"
	"yellowroad_library/database/repo/user_repo"
	"errors"
)

type GormUserRepository struct {
	dbConn *gorm.DB
}
//ensure interface implementation
var _ user_repo.UserRepository = GormUserRepository{}

func New(dbConn *gorm.DB) GormUserRepository {
	return GormUserRepository{
		dbConn: dbConn,
	}
}

func (repo GormUserRepository) FindById(id int) (entities.User, app_error.AppError) {
	var user entities.User

	dbConn := repo.dbConn

	if queryResult := dbConn.Where("id = ?", id).First(&user); queryResult.Error != nil {
		var returnedErr app_error.AppError
		if queryResult.RecordNotFound() {
			returnedErr = app_error.Wrap(queryResult.Error).
									SetHttpCode(http.StatusNotFound).
									SetEndpointMessage("No user found")
		} else {
			returnedErr = app_error.Wrap(queryResult.Error)
		}
		return user, returnedErr
	}

	return user, nil
}

func (repo GormUserRepository) FindByUsername(username string) (entities.User, app_error.AppError) {
	var dbConn = repo.dbConn
	var user entities.User

	if queryResult := dbConn.Where("username = ?", username).First(&user); queryResult.Error != nil {
		var returnedErr app_error.AppError

		if queryResult.RecordNotFound() {
			returnedErr = app_error.Wrap(queryResult.Error).
							SetHttpCode(http.StatusNotFound).
							SetEndpointMessage("Incorrect username or password")
		} else {
			returnedErr = app_error.Wrap(queryResult.Error)
		}

		return user, returnedErr
	}

	return user, nil
}

func (repo GormUserRepository) Update(user *entities.User) app_error.AppError {
	if queryResult := repo.dbConn.Save(user); queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}

func (repo GormUserRepository) Insert(user *entities.User) app_error.AppError {

	if queryResult := repo.dbConn.Create(user); queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}

func (repo GormUserRepository) Delete(user *entities.User) app_error.AppError {
	if user.ID == 0 {
		err := errors.New("Invalid primary key value of 0 while attempting to delete")
		return app_error.Wrap(err)
	}

	if queryResult := repo.dbConn.Delete(user); queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}
// Update(entities.User) error
// Create(entities.User) error
