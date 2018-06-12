package btag_repo

import (
	"github.com/jinzhu/gorm"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"errors"
)

type GormBookTagRepository struct {
	dbConn *gorm.DB
}

var _ BookTagRepository = GormBookTagRepository{} //ensure interface implementation
func NewDefault(dbConn *gorm.DB) GormBookTagRepository {
	return GormBookTagRepository{
		dbConn : dbConn,
	}
}

func (this GormBookTagRepository) Insert(booktag_count *entities.BookTag) (app_error.AppError){
	queryResult := this.dbConn.
		Set("gorm:save_associations", false).	//no magic! let the individual objects be saved on their own!
		Create(booktag_count)

	if queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}
func (this GormBookTagRepository) Delete(booktag_count *entities.BookTag) (app_error.AppError) {
	if (booktag_count.ID == 0){
		err := errors.New("Invalid primary key value of 0 while attempting to delete")
		return app_error.Wrap(err)
	}

	if queryResult := this.dbConn.Delete(booktag_count); queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}
func (this GormBookTagRepository) SyncCount(tag string, book_id int) (count entities.BookTag, err app_error.AppError) {
	//count the rows in BookTags that match the given tag and book_id
	var result struct {
		Total int
	}
	var total_tags_in_book int
	queryResult := this.dbConn.
						Model(&entities.BookTagVote{}).
						Select("SUM(direction) as total").
						Where("tag = ? AND book_id = ?", tag, book_id).
						Scan(&result)
	if queryResult.Error != nil {
		err = app_error.Wrap(queryResult.Error)
		return count, err
	}
	total_tags_in_book = result.Total


	//upsert the value into the BookTag table
	queryResult = this.dbConn.
						Where(entities.BookTag{Tag:tag, BookId:book_id}).
						Assign(entities.BookTag{Count: total_tags_in_book}).
						FirstOrCreate(&count)

	if queryResult.Error != nil {
		err = app_error.Wrap(queryResult.Error)
		return count, err
	}

	return count, err
}
