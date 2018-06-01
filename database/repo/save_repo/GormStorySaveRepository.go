package save_repo

import (
	"github.com/jinzhu/gorm"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"net/http"
)

type GormStorySaveRepository struct {
	db *gorm.DB
}

var _ SaveRepository = GormStorySaveRepository{} //ensure interface implementation
func NewDefault(db *gorm.DB) GormStorySaveRepository{
	return GormStorySaveRepository{db}
}

func (this GormStorySaveRepository) FindById(id int) (entities.SaveState, app_error.AppError) {
	var save entities.SaveState

	queryResult := this.db.
						Where("id = ?", id).
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

func (this GormStorySaveRepository) Insert(save *entities.SaveState) (app_error.AppError) {
	queryResult := this.db.
						Set("gorm:save_associations", false).
						Create(save)

	if queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}

func (this GormStorySaveRepository) Update(save *entities.SaveState) (app_error.AppError) {
	queryResult := this.db.
						Set("gorm:save_associations", false).
						Save(save)

	if queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}

func (this GormStorySaveRepository) Delete(save *entities.SaveState) (app_error.AppError) {
	queryResult := this.db.
						Set("gorm:save_associations", false).
						Delete(save)

	if queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}



