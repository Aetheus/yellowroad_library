package story_route

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/services/chapter_serv"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/services/auth_serv"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/api_response"
	"yellowroad_library/utils/gin_tools"
)



func CreateChapter(
	c *gin.Context,
	work uow.UnitOfWork,
	authService auth_serv.AuthService,
	chapterService chapter_serv.ChapterService,
) {
	var form entities.ChapterCreationForm
	var newChapter entities.Chapter
	var chapterPath entities.ChapterPath

	err := work.Auto([]uow.WorkFragment{chapterService},func () app_error.AppError{
		book_id, err := gin_tools.GetIntParam("book_id",c)
		if (err != nil){
			return err
		}

		err = gin_tools.JSON(&form,c)
		if (err != nil) {
			return err
		}

		user, err := authService.GetLoggedInUser(c)
		if (err != nil){
			return err
		}

		//create the book
		newChapter, err = chapterService.CreateChapter(user, book_id, form)
		if (err != nil) {
			return err
		}

		//create the path, if necessary
		if (form.FromChapterPath != nil){
			form.FromChapterPath.ToChapterId = &newChapter.ID
			chapterPath,err = chapterService.CreatePathBetweenChapters(user,form.FromChapterPath)
			if (err != nil){
				return err
			}
		}

		return nil
	})

	if ( err != nil ){
		c.JSON(api_response.ConvertErrWithCode(err))
	}else {
		c.JSON(api_response.SuccessWithCode(gin.H{
			"chapter" : newChapter,
			"path_to_chapter" : chapterPath,
		}))
	}
}

func UpdateChapter (
	c *gin.Context,
	work uow.UnitOfWork,
	authService auth_serv.AuthService,
	chapterService chapter_serv.ChapterService,
) {
	var updatedChapter entities.Chapter
	var chapterForm entities.ChapterUpdateForm

	err := work.Auto([]uow.WorkFragment{authService, chapterService}, func () app_error.AppError {
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
		c.JSON(api_response.ConvertErrWithCode(err))
	}else {
		c.JSON(api_response.SuccessWithCode(gin.H{
			"chapter" : updatedChapter,
		}))
	}
}

func DeleteChapter(
	c *gin.Context,
	work uow.UnitOfWork,
	authService auth_serv.AuthService,
	chapterService chapter_serv.ChapterService,
){
	err := work.Auto([]uow.WorkFragment{authService, chapterService}, func () app_error.AppError {
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
		c.JSON(api_response.ConvertErrWithCode(err))
	}else {
		c.JSON(api_response.SuccessWithCode(gin.H{}))
	}
}