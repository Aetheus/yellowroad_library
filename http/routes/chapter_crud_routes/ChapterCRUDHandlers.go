package chapter_crud_routes

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/api_reply"
	"yellowroad_library/utils/gin_tools"
	"yellowroad_library/http/middleware/auth_middleware"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/services/chapter_domain"
)


type ChapterCrudContainer interface {
	UnitOfWork() uow.UnitOfWork
	CreateChapterAndPath(uow.UnitOfWork) chapter_domain.CreateChapterAndPath
	UpdateChapter(uow.UnitOfWork) chapter_domain.UpdateChapter
	DeleteChapter(uow.UnitOfWork) chapter_domain.DeleteChapter
	CreatePathBetweenChapters(uow.UnitOfWork) chapter_domain.CreatePathBetweenChapters
	UpdatePathBetweenChapters(uow.UnitOfWork) chapter_domain.UpdatePathBetweenChapters
}
type ChapterCrudHandlers struct {
	Container ChapterCrudContainer
}

func (this ChapterCrudHandlers) FetchSingleChapter(c *gin.Context) {
	/*Dependencies**************/
	work := this.Container.UnitOfWork()
	chapterRepo := work.ChapterRepo()
	/***************************/

	var chapter entities.Chapter

	err := work.AutoCommit(func() app_error.AppError {
		book_id, err := gin_tools.GetIntParam("book_id",c)
		if (err != nil) { return err }

		chapter_id, err := gin_tools.GetIntParam("chapter_id",c)
		if (err != nil) { return err }

		chapter, err = chapterRepo.FindWithinBook(chapter_id, book_id)
		if (err != nil) { return err }

		return nil
	})

	if ( err != nil ){
		api_reply.Failure(c,err)
	}else {
		api_reply.Success(c, gin.H{
			"chapter" : chapter,
		})
	}
}

func (this ChapterCrudHandlers) FetchChaptersIndex(c *gin.Context) {
	/*Dependencies**************/
	work := this.Container.UnitOfWork()
	chapterRepo := work.ChapterRepo()
	/***************************/

	var chapters []entities.Chapter
	err := work.AutoCommit(func() app_error.AppError {
		book_id ,err := gin_tools.GetIntParam("book_id",c)
		if (err != nil){
			return err
		}
		chapters, err = chapterRepo.ChaptersIndex(book_id)
		if (err != nil){
			return err
		}

		return nil
	})

	if ( err != nil ){
		api_reply.Failure(c,err)
	}else {
		api_reply.Success(c, gin.H{
			"chapters" : chapters,
		})
	}
}

func (this ChapterCrudHandlers) CreateChapter(c *gin.Context) {
	/*Dependencies**************/
	work := this.Container.UnitOfWork()
	createChapterAndPath := this.Container.CreateChapterAndPath(work)
	/***************************/

	var form entities.Chapter_And_Path_CreationForm
	var newChapter entities.Chapter
	var newPath entities.ChapterPath

	err := work.AutoCommit(func () app_error.AppError{
		book_id, err := gin_tools.GetIntParam("book_id",c)
		if (err != nil){
			return err
		}

		err = gin_tools.BindJSON(&form,c)
		if (err != nil) {
			return err
		}
		form.ChapterForm.BookId = &book_id //ensures that the bookId is the book id of this route

		user, err := auth_middleware.GetUser(c)
		if (err != nil){
			return err
		}

		newChapter, newPath, err = createChapterAndPath.Execute(user,form)
		if (err != nil) {
			return err
		}

		return nil
	})

	if ( err != nil ){
		api_reply.Failure(c,err)
	}else {
		api_reply.Success(c, gin.H{
			"chapter" : newChapter,
		})
	}
}

func (this ChapterCrudHandlers) UpdateChapter (c *gin.Context) {
	/*Dependencies**************/
	work := this.Container.UnitOfWork()
	updateChapter := this.Container.UpdateChapter(work)
	/***************************/

	var updatedChapter entities.Chapter
	var chapterForm entities.Chapter_UpdateForm

	err := work.AutoCommit(func () app_error.AppError {
		chapterId, err := gin_tools.GetIntParam("chapter_id",c)
		if (err != nil){
			return err
		}

		err = gin_tools.BindJSON(&chapterForm,c)
		if (err != nil) {
			return err
		}

		currentUser, err := auth_middleware.GetUser(c)
		if (err != nil){
			return err
		}

		updatedChapter, err = updateChapter.Execute(currentUser,chapterId, chapterForm)
		return nil
	})


	if ( err != nil ){
		api_reply.Failure(c, err)
	}else {
		api_reply.Success(c, gin.H{ "chapter" : updatedChapter })
	}
}

func (this ChapterCrudHandlers) DeleteChapter(c *gin.Context){
	/*Dependencies**************/
	work := this.Container.UnitOfWork()
	deleteChapter := this.Container.DeleteChapter(work)
	/***************************/

	err := work.AutoCommit(func () app_error.AppError {
		chapterId, err := gin_tools.GetIntParam("chapter_id",c)
		if (err != nil){
			return err
		}

		user, err := auth_middleware.GetUser(c)
		if (err != nil){
			return err
		}

		err = deleteChapter.Execute(user,chapterId)
		if(err != nil){
			return err
		}

		return nil
	})

	if ( err != nil ){
		api_reply.Failure(c,err)
	}else {
		api_reply.Success(c,gin.H{})
	}
}

func (this ChapterCrudHandlers) CreateChapterPath(c *gin.Context) {
	/*Dependencies**************/
	work := this.Container.UnitOfWork()
	createPathBetweenChapters := this.Container.CreatePathBetweenChapters(work)
	/***************************/

	var newPath entities.ChapterPath
	var form entities.ChapterPath_CreationForm

	err := work.AutoCommit(func () app_error.AppError{
		err := gin_tools.BindJSON(&form,c)
		if (err != nil){
			return err
		}

		currentUser, err := auth_middleware.GetUser(c)
		if (err != nil){
			return err
		}

		newPath, err = createPathBetweenChapters.Execute(currentUser, form)
		if (err != nil) {
			return err
		}

		return nil
	})

	if (err != nil){
		api_reply.Failure(c,err)
	}else{
		api_reply.Success(c,newPath)
	}
}

func (this ChapterCrudHandlers) UpdateChapterPath(c *gin.Context){
	/*Dependencies**************/
	work := this.Container.UnitOfWork()
	updatePathBetweenChapters := this.Container.UpdatePathBetweenChapters(work)
	/***************************/

	var chapter_path entities.ChapterPath

	err := work.AutoCommit(func () app_error.AppError{
		var update_form entities.ChapterPath_UpdateForm

		chapter_path_id, err := gin_tools.GetIntParam("chapter_path_id", c)
		if err != nil {
			return err
		}

		err = gin_tools.BindJSON(&update_form,c)
		if err != nil {
			return err
		}

		currentUser, err := auth_middleware.GetUser(c)
		if err != nil {
			return err
		}

		chapter_path, err = updatePathBetweenChapters.Execute(currentUser,chapter_path_id, update_form)
		if err != nil {
			return err
		}

		return nil
	})

	if (err != nil){
		api_reply.Failure(c,err)
	}else{
		api_reply.Success(c,chapter_path)
	}
}