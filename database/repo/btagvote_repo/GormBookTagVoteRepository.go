package btagvote_repo

import (
	"github.com/jinzhu/gorm"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"errors"
)

type GormBookTagVoteRepository struct {
	dbConn *gorm.DB
}

var _ BookTagVoteRepository = GormBookTagVoteRepository{} //ensure interface implementation
func NewDefault(dbConn *gorm.DB) GormBookTagVoteRepository {
	return GormBookTagVoteRepository{
		dbConn : dbConn,
	}
}

func (this GormBookTagVoteRepository) FindByFields(searchFields entities.BookTagVote) (results []entities.BookTagVote, err app_error.AppError) {
	queryResult := this.dbConn.
						Where(searchFields).
						Find(&results)

	if queryResult.Error != nil {
		err = app_error.Wrap(queryResult.Error)
		return
	}

	return
}

func (this GormBookTagVoteRepository) Insert(booktag *entities.BookTagVote) (app_error.AppError){
	queryResult := this.dbConn.
		Set("gorm:save_associations", false).	//no magic! let the individual objects be saved on their own!
		Create(booktag)

	if queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}

func (this GormBookTagVoteRepository) Upsert(booktag *entities.BookTagVote) (app_error.AppError){
	queryResult := this.dbConn.
						Set("gorm:save_associations", false).	//no magic! let the individual objects be saved on their own!
						Where(entities.BookTagVote{Tag:booktag.Tag, BookId:booktag.BookId, UserId: booktag.UserId}).
						Assign(*booktag).
						FirstOrCreate(booktag)

	if queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}


func (this GormBookTagVoteRepository) Delete(booktag *entities.BookTagVote) (app_error.AppError) {
	if (booktag.ID == 0){
		err := errors.New("Invalid primary key value of 0 while attempting to delete")
		return app_error.Wrap(err)
	}

	if queryResult := this.dbConn.Delete(booktag); queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}
func (this GormBookTagVoteRepository) DeleteByFields(searchFields entities.BookTagVote) (app_error.AppError){
	queryResult := this.dbConn.
						Where(searchFields).
						Delete(entities.BookTagVote{})

	if queryResult.Error != nil {
		return app_error.Wrap(queryResult.Error)
	}

	return nil
}
