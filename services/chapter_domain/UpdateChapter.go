package chapter_domain

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/database/repo/chapter_repo"
)

type UpdateChapter struct {
	bookRepo book_repo.BookRepository
	chapterRepo chapter_repo.ChapterRepository
}

func (this UpdateChapter) Execute(
	user entities.User, chapter_id int, form entities.Chapter_UpdateForm,
)  (entities.Chapter,app_error.AppError) {
	var chapter entities.Chapter

	//TODO: find a way to move the authcheck creation out of here
	err := newAuthorityCheck(this.bookRepo).Execute(auth_update_chapter, user.ID, chapter.BookId)
	if(err != nil){
		return chapter, err
	}

	chapter, err = this.chapterRepo.FindById(chapter_id)
	if(err != nil){
		return chapter, err
	}
	form.Apply(&chapter)

	err = this.chapterRepo.Update(&chapter)
	if(err != nil){
		return chapter, err
	}

	return chapter, nil
}

func NewUpdateChapter(
	bookRepo book_repo.BookRepository,
	chapterRepo chapter_repo.ChapterRepository,
) UpdateChapter {
	return UpdateChapter{bookRepo,chapterRepo}
}