package btagvotecount_repo

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
)

type BookTagVoteCountRepository interface {
	Insert(*entities.BookTagVoteCount) (app_error.AppError)
	Delete(*entities.BookTagVoteCount) (app_error.AppError)
	SyncCount(tag string, book_id int) (count entities.BookTagVoteCount, err app_error.AppError)
}