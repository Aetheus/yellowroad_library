package gorm_booktagcount_repo

import (
	"yellowroad_library/database/repo/booktagcount_repo"
	"github.com/jinzhu/gorm"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"errors"
)

type GormBookTagCountRepository struct {
	dbConn *gorm.DB
}

var _ booktagcount_repo.BookTagCountRepository = GormBookTagCountRepository{}

func New(dbConn *gorm.DB) GormBookTagCountRepository {
	return GormBookTagCountRepository{
		dbConn : dbConn,
	}
}

func (this GormBookTagCountRepository) Insert(booktag_count *entities.BookTagCount) (app_error.AppError){
	queryResult := this.dbConn.
		Set("gorm:save_associations", false).	//no magic! let the individual objects be saved on their own!
		Create(booktag_count)

	if queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}
func (this GormBookTagCountRepository) Delete(booktag_count *entities.BookTagCount) (app_error.AppError) {
	if (booktag_count.ID == 0){
		err := errors.New("Invalid primary key value of 0 while attempting to delete")
		return app_error.Wrap(err)
	}

	if queryResult := this.dbConn.Delete(booktag_count); queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}
func (this GormBookTagCountRepository) SyncCount(tag string, book_id int) (count entities.BookTagCount, err app_error.AppError) {
	//count the rows in BookTags that match the given tag and book_id
	var total_tags_in_book int
	queryResult := this.dbConn.
						Model(&entities.BookTag{}).
						Where("tag = ? AND book_id = ?", tag, book_id).
						Count(&total_tags_in_book)
	if queryResult.Error != nil {
		err = app_error.Wrap(queryResult.Error)
		return count, err
	}

	//upsert the value into the BookTagCount table
	queryResult = this.dbConn.
						Where(entities.BookTagCount{Tag:tag, BookId:book_id}).
						Assign(entities.BookTagCount{Count: total_tags_in_book}).
						FirstOrCreate(&count)

	if queryResult.Error != nil {
		err = app_error.Wrap(queryResult.Error)
		return count, err
	}

	return count, err
}
