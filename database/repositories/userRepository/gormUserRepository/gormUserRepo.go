package gormUserRepository

import (
	"yellowroad_library/database/entities"

	"github.com/jinzhu/gorm"
	"yellowroad_library/utils/appError"
	"net/http"
)

type GormUserRepository struct {
	dbConn *gorm.DB
}

func New(dbConn *gorm.DB) GormUserRepository {
	return GormUserRepository{
		dbConn: dbConn,
	}
}

func (repo GormUserRepository) FindById(id int) (*entities.User, error) {
	var user entities.User

	dbConn := repo.dbConn

	if queryResult := dbConn.Where("id = ?", id).First(&user); queryResult.Error != nil {
		var returnedErr error
		if queryResult.RecordNotFound() {
			returnedErr = appError.Wrap(queryResult.Error).
									SetHttpCode(http.StatusNotFound).
									SetEndpointMessage("No user found")
		} else {
			returnedErr = appError.Wrap(queryResult.Error)
		}
		return nil, returnedErr
	}

	return &user, nil
}

func (repo GormUserRepository) FindByUsername(username string) (*entities.User, error) {
	var dbConn = repo.dbConn
	var user entities.User

	if queryResult := dbConn.Where("username = ?", username).First(&user); queryResult.Error != nil {
		var returnedErr error

		if queryResult.RecordNotFound() {
			returnedErr = appError.Wrap(queryResult.Error).
							SetHttpCode(http.StatusNotFound).
							SetEndpointMessage("Incorrect username or password")
		} else {
			returnedErr = appError.Wrap(queryResult.Error)
		}

		return nil, returnedErr
	}

	return &user, nil
}

func (repo GormUserRepository) Update(user *entities.User) error {
	if queryResult := repo.dbConn.Save(user); queryResult.Error != nil {
		return appError.Wrap(queryResult.Error)
	}

	return nil
}

func (repo GormUserRepository) Insert(user *entities.User) error {

	if queryResult := repo.dbConn.Create(user); queryResult.Error != nil {
		return appError.Wrap(queryResult.Error)
	}

	return nil
}

// Update(entities.User) error
// Create(entities.User) error
