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
	var chapterForm entities.ChapterForm
	var newChapter entities.Chapter
	var chapterPath entities.ChapterPath

	err := work.Auto([]uow.WorkFragment{chapterService},func () app_error.AppError{
		book_id, err := gin_tools.GetIntParam("book_id",c)
		if (err != nil){
			return err
		}
		bindErr := c.BindJSON(&chapterForm)
		if (bindErr != nil){
			return app_error.Wrap(bindErr)
		}
		user, err := authService.GetLoggedInUser(c)
		if (err != nil){
			return err
		}

		//create the book
		chapterForm.Apply(&newChapter)
		err = chapterService.CreateChapter(user, book_id, &newChapter)
		if (err != nil) {
			return err
		}

		//create the path, if necessary
		if (chapterForm.FromChapterPath != nil){
			chapterForm.FromChapterPath.Apply(&chapterPath)
			chapterPath.ToChapterId = newChapter.ID
			err = chapterService.CreatePathBetweenChapters(user,&chapterPath)
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