package gorm_booktag_repo

import (
	"yellowroad_library/database/repo/booktag_repo"
	"github.com/jinzhu/gorm"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"errors"
)

type GormBookTagRepository struct {
	dbConn *gorm.DB
}
var _ booktag_repo.BookTagRepository = GormBookTagRepository{}

func New(dbConn *gorm.DB) GormBookTagRepository{
	return GormBookTagRepository{
		dbConn : dbConn,
	}
}

func (this GormBookTagRepository) Insert(booktag *entities.BookTag) (app_error.AppError){
	queryResult := this.dbConn.
		Set("gorm:save_associations", false).	//no magic! let the individual objects be saved on their own!
		Create(booktag)

	if queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}
func (this GormBookTagRepository) Delete(booktag *entities.BookTag) (app_error.AppError) {
	if (booktag.ID == 0){
		err := errors.New("Invalid primary key value of 0 while attempting to delete")
		return app_error.Wrap(err)
	}

	if queryResult := this.dbConn.Delete(booktag); queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}
func (this GormBookTagRepository) DeleteByFields(tag string, user_id int, book_id int) (app_error.AppError){
	queryResult := this.dbConn.
						Where("tag = ? AND user_id = ? AND book_id = ?", tag,user_id,book_id).
						Delete(entities.BookTag{})

	if queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}
