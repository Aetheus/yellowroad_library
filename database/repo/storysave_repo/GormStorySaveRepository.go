package storysave_repo

import (
	"github.com/jinzhu/gorm"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"net/http"
)

type GormStorySaveRepository struct {
	db *gorm.DB
}

var _ StorySaveRepository = GormStorySaveRepository{} //ensure interface implementation
func NewDefault(db *gorm.DB) GormStorySaveRepository{
	return GormStorySaveRepository{db}
}

func (this GormStorySaveRepository) FindByToken(token string) (entities.StorySave, app_error.AppError) {
	var save entities.StorySave

	queryResult := this.db.
						Where("token = ?", token).
						First(&save)

	if(queryResult.Error != nil) {
		var returnedErr app_error.AppError
		if queryResult.RecordNotFound() {
			returnedErr = app_error.Wrap(queryResult.Error).
				SetHttpCode(http.StatusNotFound).
				SetEndpointMessage("No save found")
		} else {
			returnedErr = app_error.Wrap(queryResult.Error)
		}
		return save, returnedErr
	}

	return save, nil
}

func (this GormStorySaveRepository) Insert(save *entities.StorySave) (app_error.AppError) {
	queryResult := this.db.
						Set("gorm:save_associations", false).
						Create(save)

	if queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}

func (this GormStorySaveRepository) Update(save *entities.StorySave) (app_error.AppError) {
	queryResult := this.db.
						Set("gorm:save_associations", false).
						Save(save)

	if queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}

func (this GormStorySaveRepository) Delete(save *entities.StorySave) (app_error.AppError) {
	queryResult := this.db.
						Set("gorm:save_associations", false).
						Delete(save)

	if queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}



