package story_route

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/services/story_serv"
	"yellowroad_library/utils/gin_tools"
	"yellowroad_library/utils/api_response"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/utils/app_error"
)

func NavigateToSingleChapter(
	c *gin.Context,
	work uow.UnitOfWork,
	storyService story_serv.StoryService,
) {
	var newSaveString string
	var pathResponse story_serv.PathResponse
	err := work.Auto([]uow.WorkFragment{storyService}, func() app_error.AppError {
		bookId, err := gin_tools.GetIntParam("book_id",c)
		if (err != nil){
			return err
		}

		chapterId, err := gin_tools.GetIntParam("chapter_id",c)
		if (err != nil){
			return err
		}

		saveString := c.Query("save")

		chapterPathId := gin_tools.GetIntQueryOrDefault("chapter_path_id", 0,c)	//chapter path can be ignored if we're on freemode
		isFreeMode := gin_tools.GetBoolQueryOrDefault("freemode",false,c)			//free mode is off by default

		pathRequest := story_serv.NewPathRequest(isFreeMode, bookId, chapterId, chapterPathId)
		pathResponse, err = storyService.NavigateToChapter(pathRequest,saveString)
		if (err != nil) {
			return err
		}

		newSaveString , err = pathResponse.NewSave.EncodedSaveString()
		if (err != nil) {
			return err
		}
		return nil
	})


	if(err != nil){
		c.JSON(api_response.ConvertErrWithCode(err))
	} else {
		c.JSON(api_response.SuccessWithCode(gin.H{
			"chapter" : pathResponse.DestinationChapter,
			"save" : newSaveString,
		}))
	}
}