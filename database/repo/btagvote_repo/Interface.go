package btagvote_repo

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
)

type BookTagVoteRepository interface {
	FindByFields(searchFields entities.BookTagVote) ([]entities.BookTagVote, app_error.AppError)
	Insert(*entities.BookTagVote) (app_error.AppError)
	Upsert(*entities.BookTagVote) (app_error.AppError)
	Delete(*entities.BookTagVote) (app_error.AppError)
	DeleteByFields(searchFields entities.BookTagVote) (app_error.AppError)
}