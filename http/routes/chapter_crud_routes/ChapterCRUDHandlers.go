package chapter_crud_routes

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/api_reply"
	"yellowroad_library/utils/gin_tools"
	"yellowroad_library/containers"
)

type ChapterCrudHandlers struct {
	Container containers.Container
}

func (this ChapterCrudHandlers) FetchSingleChapter(c *gin.Context) {
	/*Dependencies**************/
	work := this.Container.UnitOfWork()
	chapterRepo := work.ChapterRepo()
	/***************************/

	var chapter entities.Chapter

	err := work.AutoCommit([]uow.WorkFragment{}, func() app_error.AppError {
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
	err := work.AutoCommit([]uow.WorkFragment{}, func() app_error.AppError {
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
	authService := this.Container.AuthService(work)
	chapterService  := this.Container.ChapterService(work)
	/***************************/

	var form entities.Chapter_And_Path_CreationForm
	var newChapter entities.Chapter
	var newPath entities.ChapterPath

	err := work.AutoCommit([]uow.WorkFragment{chapterService},func () app_error.AppError{
		book_id, err := gin_tools.GetIntParam("book_id",c)
		if (err != nil){
			return err
		}

		err = gin_tools.BindJSON(&form,c)
		if (err != nil) {
			return err
		}
		form.ChapterForm.BookId = &book_id //ensures that the bookId is the book id of this route

		user, err := authService.GetLoggedInUser(c)
		if (err != nil){
			return err
		}

		newChapter, newPath, err = chapterService.CreateChapterAndPath(user,form)
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
	authService := this.Container.AuthService(work)
	chapterService  := this.Container.ChapterService(work)
	/***************************/

	var updatedChapter entities.Chapter
	var chapterForm entities.Chapter_UpdateForm

	err := work.AutoCommit([]uow.WorkFragment{authService, chapterService}, func () app_error.AppError {
		chapterId, err := gin_tools.GetIntParam("chapter_id",c)
		if (err != nil){
			return err
		}

		err = gin_tools.BindJSON(&chapterForm,c)
		if (err != nil) {
			return err
		}

		currentUser, err := authService.GetLoggedInUser(c)
		if (err != nil){
			return err
		}

		updatedChapter, err = chapterService.UpdateChapter(currentUser,chapterId, chapterForm)
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
	authService := this.Container.AuthService(work)
	chapterService  := this.Container.ChapterService(work)
	/***************************/

	err := work.AutoCommit([]uow.WorkFragment{authService, chapterService}, func () app_error.AppError {
		chapterId, err := gin_tools.GetIntParam("chapter_id",c)
		if (err != nil){
			return err
		}

		user, err := authService.GetLoggedInUser(c)
		if (err != nil){
			return err
		}

		err = chapterService.DeleteChapter(user,chapterId)
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
	authService := this.Container.AuthService(work)
	chapterService  := this.Container.ChapterService(work)
	/***************************/

	var newPath entities.ChapterPath
	var form entities.ChapterPath_CreationForm

	err := work.AutoCommit([]uow.WorkFragment{ authService, chapterService }, func () app_error.AppError{
		err := gin_tools.BindJSON(&form,c)
		if (err != nil){
			return err
		}

		currentUser, err := authService.GetLoggedInUser(c)
		if (err != nil){
			return err
		}

		newPath, err = chapterService.CreatePathBetweenChapters(currentUser, form)
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
	authService := this.Container.AuthService(work)
	chapterService  := this.Container.ChapterService(work)
	/***************************/

	var chapter_path entities.ChapterPath

	err := work.AutoCommit([]uow.WorkFragment{}, func () app_error.AppError{
		var update_form entities.ChapterPath_UpdateForm

		chapter_path_id, err := gin_tools.GetIntParam("chapter_path_id", c)
		if err != nil {
			return err
		}

		err = gin_tools.BindJSON(&update_form,c)
		if err != nil {
			return err
		}

		currentUser, err := authService.GetLoggedInUser(c)
		if err != nil {
			return err
		}

		chapter_path, err = chapterService.UpdatePathBetweenChapters(currentUser,chapter_path_id, update_form)
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