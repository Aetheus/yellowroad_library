package chapter_domain

import (
	"yellowroad_library/utils/app_error"
	"net/http"
	"yellowroad_library/database/entities"
	"yellowroad_library/database/repo/chapter_repo"
	"yellowroad_library/database/repo/chapterpath_repo"
	"yellowroad_library/database/repo/book_repo"
)

type CreatePathBetweenChapters struct {
	chapterRepo chapter_repo.ChapterRepository
	chapterPathRepo chapterpath_repo.ChapterPathRepository
	bookRepo book_repo.BookRepository
}

func NewCreatePathBetweenChapters(
	chapterRepo chapter_repo.ChapterRepository,
	chapterPathRepo chapterpath_repo.ChapterPathRepository,
	bookRepo book_repo.BookRepository,
) CreatePathBetweenChapters {
	return CreatePathBetweenChapters{
		chapterRepo,
		chapterPathRepo,
		bookRepo,
	}
}

//TODO refactor this to use goroutines so the various checks can take advantage of concurrency instead of being performed in serial
func (this CreatePathBetweenChapters) Execute(
	user entities.User,
	form entities.ChapterPath_CreationForm,
) (entities.ChapterPath, app_error.AppError) {
	var path entities.ChapterPath

	err := form.Apply(&path)
	if (err != nil) {
		return path, err
	}

	fromChapter, err := this.chapterRepo.FindById(path.FromChapterId)
	if (err != nil) {
		return path, err
	}

	toChapter, err := this.chapterRepo.FindById(path.ToChapterId)
	if (err != nil) {
		return path, err
	}

	if(fromChapter.BookId != toChapter.BookId){
		errMessage := "These two chapters aren't a part of the same book!"
		return path, app_error.New(http.StatusUnprocessableEntity,"",errMessage)
	}

	_, err = this.chapterPathRepo.FindByChapterIds(path.FromChapterId, path.ToChapterId)
	if (err == nil){
		endPointMessage := "A path already exists between these two chapters!"
		return path, app_error.New(http.StatusConflict, "",endPointMessage)
	} else if (err != nil && err.HttpCode() != http.StatusNotFound){
		return path, err
	}

	//TODO: find a way to move the authcheck creation out of here
	err = newAuthorityCheck(this.bookRepo).Execute(auth_create_path, user.ID, fromChapter.BookId)
	if (err != nil){
		return path, err
	}

	err = this.chapterPathRepo.Insert(&path)
	if (err != nil){
		return path, err
	}

	return path, nil
}