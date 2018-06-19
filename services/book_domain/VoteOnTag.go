package book_domain

import (
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/database/repo/btag_repo"
	"yellowroad_library/database/repo/btagvote_repo"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
)

type VoteOnTag struct {
	bookRepo book_repo.BookRepository
	bookTagRepo btag_repo.BookTagRepository
	bookTagVoteRepo btagvote_repo.BookTagVoteRepository
}

func NewVoteOnTag(
	bookRepo 		book_repo.BookRepository,
	bookTagRepo 	btag_repo.BookTagRepository,
	bookTagVoteRepo btagvote_repo.BookTagVoteRepository,
) VoteOnTag {
	return VoteOnTag {
		bookRepo,bookTagRepo,bookTagVoteRepo,
	}
}

func (this VoteOnTag) Execute(currentUser entities.User, book_id int, tagName string, voteValue int) (newCount int, err app_error.AppError){
	_, err = this.bookRepo.FindById(book_id)
	if (err != nil){
		//quit early if we can't find the book or something else went wrong
		return
	}

	//prevent people from entering values larger than 1 or lower than -1
	if(voteValue > 0){
		voteValue = 1
	} else if (voteValue < 0) {
		voteValue = -1
	}

	vote := entities.BookTagVote{
		BookId : book_id,
		UserId : currentUser.ID,
		Tag : tagName,
		Direction : voteValue,
	}

	err = this.bookTagVoteRepo.Upsert(&vote)
	if (err != nil) {return newCount,err}

	//TODO: move this to a queue since it isn't really critical;newCount could simply be currentCount++
	result, err := this.bookTagRepo.SyncCount(tagName,book_id)
	if (err != nil) {return newCount,err}
	newCount = result.Count

	return
}