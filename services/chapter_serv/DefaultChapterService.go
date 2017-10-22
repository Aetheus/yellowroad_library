package chapter_serv

import (
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"net/http"
)

type DefaultChapterService struct {
	work uow.UnitOfWork
}

func Default(work uow.UnitOfWork) ChapterService{
	return DefaultChapterService{
		work : work,
	}
}

func (this DefaultChapterService) crudAuthorityCheck(instigatorId int,bookId int) app_error.AppError {
	//TODO: more advanced permission system
	book, err := this.work.BookRepo().FindById(bookId)
	if(err != nil){
		return err
	}
	if (book.CreatorId != instigatorId){
		return app_error.New(http.StatusUnauthorized,"","You are not authorized to perform this action!")
	}
	return nil
}

func (this DefaultChapterService) CreateChapter(
	instigator entities.User, book_id int , form entities.Chapter_CreationForm,
) (entities.Chapter,app_error.AppError) {
	var newChapter entities.Chapter

	err := this.crudAuthorityCheck(instigator.ID, book_id)
	if (err != nil){
		return newChapter, err
	}

	form.Apply(&newChapter)
	newChapter.BookId = book_id
	newChapter.CreatorId = instigator.ID

	err = this.work.ChapterRepo().Insert(&newChapter)
	if err != nil {
		return newChapter, err
	}

	return newChapter, nil
}

func (this DefaultChapterService) DeleteChapter(instigator entities.User, chapter_id int) app_error.AppError {
	chapter, err := this.work.ChapterRepo().FindById(chapter_id)
	if (err != nil) {
		return err
	}

	err = this.crudAuthorityCheck(instigator.ID, chapter.BookId)
	if(err != nil){
		return err
	}

	err = this.work.ChapterRepo().Delete(&chapter)
	if(err != nil){
		return err
	}

	return nil
}

func (this DefaultChapterService) UpdateChapter(
	instigator entities.User, chapter_id int, form entities.Chapter_UpdateForm,
) (entities.Chapter,app_error.AppError) {
	var chapter entities.Chapter

	//TODO : do a permission check
	//err := this.crudAuthorityCheck(instigator.ID, chapter.BookId)

	chapter, err := this.work.ChapterRepo().FindById(chapter_id)
	if(err != nil){
		return chapter, err
	}
	form.Apply(&chapter)

	err = this.work.ChapterRepo().Update(&chapter)
	if(err != nil){
		return chapter, err
	}

	return chapter, nil
}


//TODO refactor this to use goroutines so the various checks can take advantage of concurrency instead of being performed in serial
func (this DefaultChapterService) CreatePathBetweenChapters(
	instigator entities.User,form entities.ChapterPath_CreationForm,
) (entities.ChapterPath, app_error.AppError) {
	var chapterRepo = this.work.ChapterRepo()
	var chapterPathRepo = this.work.ChapterPathRepo()

	var path entities.ChapterPath
	form.Apply(&path)

	//check if these two chapters exist
	fromChapter, err := chapterRepo.FindById(path.FromChapterId)
	if (err != nil) {
		return path, err
	}
	toChapter, err := chapterRepo.FindById(path.ToChapterId)
	if (err != nil) {
		return path, err
	}
	if(fromChapter.BookId != toChapter.BookId){
		errMessage := "These two chapters aren't a part of the same book!"
		return path, app_error.New(http.StatusUnprocessableEntity,"",errMessage)
	}

	//TODO implement check if a path between the two chapters already exists
	//check if a chapter path between these two chapters already exists
	//chapterPathRepo

	//standard CRUD check
	err = this.crudAuthorityCheck(instigator.ID,fromChapter.BookId)
	if (err != nil){
		return path, err
	}

	err = chapterPathRepo.Insert(&path)
	if (err != nil){
		 return path, err
	}

	return path, nil
}

func (this DefaultChapterService) UpdatePathBetweenChapters(
	instigator entities.User, path_id int, form entities.ChapterPath_UpdateForm,
) (entities.ChapterPath, app_error.AppError) {
	var chapterRepo = this.work.ChapterPathRepo()
	var path entities.ChapterPath
	path, err := chapterRepo.FindById(path_id)
	if (err != nil) {
		return path, err
	}
	form.Apply(&path)

	err = chapterRepo.Update(&path)
	if (err != nil){
		return path, err
	}

	return path, nil
}

func (this DefaultChapterService)  DeletePathBetweenChapters(instigator entities.User,path_id int) app_error.AppError {
	var chapterRepo = this.work.ChapterPathRepo()
	path, err := chapterRepo.FindById(path_id)
	if (err != nil ){
		return err
	}

	//standard CRUD check
	err = this.crudAuthorityCheck(instigator.ID,path.FromChapter.BookId)
	if (err != nil){
		return err
	}

	err = chapterRepo.Delete(&path)
	if (err != nil){
		return err
	}

	return nil
}



func (this DefaultChapterService) SetUnitOfWork(work uow.UnitOfWork) {
	this.work = work
}



