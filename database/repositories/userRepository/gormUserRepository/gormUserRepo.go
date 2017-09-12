package gormUserRepository

import (
	"errors"
	"yellowroad_library/database/entities"

	"github.com/jinzhu/gorm"
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

	if results := dbConn.Where("id = ?", id).First(&user); results.Error != nil {
		var returnedErr error
		if results.RecordNotFound() {
			returnedErr = errors.New("No such user")
		} else {
			returnedErr = errors.New("Unknown error occured")
		}
		return nil, returnedErr
	}

	return &user, nil
}

func (repo GormUserRepository) FindByUsername(username string) (*entities.User, error) {
	var dbConn = repo.dbConn
	var user entities.User

	//TODO : email as well
	if queryResult := dbConn.Where("username = ?", username).First(&user); queryResult.Error != nil {
		var returnedErr error

		if queryResult.RecordNotFound() {
			returnedErr = errors.New("Incorrect username")
		} else {
			returnedErr = errors.New("Unknown error occured")
		}

		return nil, returnedErr
	}

	return &user, nil
}

func (repo GormUserRepository) Update(user *entities.User) error {
	if queryResult := repo.dbConn.Save(user); queryResult.Error != nil {
		return queryResult.Error
	}

	return nil
}

func (repo GormUserRepository) Insert(user *entities.User) error {

	if result := repo.dbConn.Create(user); result.Error != nil {
		return result.Error
	}

	return nil
}

// Update(entities.User) error
// Create(entities.User) error
