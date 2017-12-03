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
	container containers.Container
}

func (this ChapterCrudHandlers) CreateChapter(c *gin.Context) {
	/*Dependencies**************/
	work := this.container.UnitOfWork()
	authService := this.container.AuthService(work)
	chapterService  := this.container.ChapterService(work)
	/***************************/

	var form entities.Chapter_And_Path_CreationForm
	var newChapter entities.Chapter
	var chapterPath entities.ChapterPath

	err := work.AutoCommit([]uow.WorkFragment{chapterService},func () app_error.AppError{
		book_id, err := gin_tools.GetIntParam("book_id",c)
		if (err != nil){
			return err
		}

		err = gin_tools.JSON(&form,c)
		if (err != nil) {
			return err
		}
		form.ChapterForm.BookId = &book_id //ensures that the bookId is the book id of this route

		user, err := authService.GetLoggedInUser(c)
		if (err != nil){
			return err
		}

		newChapter, chapterPath, err = chapterService.CreateChapterAndPath(user,form)
		if (err != nil) {
			return err
		}

		return nil
	})

	if ( err != nil ){
		api_reply.Failure(c,err)
	}else {
		api_reply.Success(c, gin.H{
			"chapter" : newChapter, "path_to_chapter" : chapterPath,
		})
	}
}

func (this ChapterCrudHandlers) UpdateChapter (c *gin.Context) {
	/*Dependencies**************/
	work := this.container.UnitOfWork()
	authService := this.container.AuthService(work)
	chapterService  := this.container.ChapterService(work)
	/***************************/

	var updatedChapter entities.Chapter
	var chapterForm entities.Chapter_UpdateForm

	err := work.AutoCommit([]uow.WorkFragment{authService, chapterService}, func () app_error.AppError {
		chapterId, err := gin_tools.GetIntParam("chapter_id",c)
		if (err != nil){
			return err
		}

		err = gin_tools.JSON(&chapterForm,c)
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
	work := this.container.UnitOfWork()
	authService := this.container.AuthService(work)
	chapterService  := this.container.ChapterService(work)
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

func (this ChapterCrudHandlers) CreatePathAwayFromThisChapter(c *gin.Context) {
	/*Dependencies**************/
	work := this.container.UnitOfWork()
	authService := this.container.AuthService(work)
	chapterService  := this.container.ChapterService(work)
	/***************************/

	var newPath entities.ChapterPath
	var form entities.ChapterPath_CreationForm

	err := work.AutoCommit([]uow.WorkFragment{ authService, chapterService }, func () app_error.AppError{
		chapterId, err := gin_tools.GetIntParam("chapter_id",c)
		if (err != nil){
			return err
		}

		err = gin_tools.JSON(&form,c)
		if (err != nil){
			return err
		}
		form.FromChapterId = &chapterId

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