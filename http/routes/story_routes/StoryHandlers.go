package story_routes

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/services/story_serv"
	"yellowroad_library/utils/gin_tools"
	"yellowroad_library/utils/api_reply"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/utils/app_error"
	"encoding/json"
	"yellowroad_library/containers"
)

type StoryHandlers struct {
	Container containers.Container
}

func (this StoryHandlers) NavigateToSingleChapter(c *gin.Context) {
	/*Dependencies**************/
	work := this.Container.UnitOfWork()
	storyService := this.Container.StoryService(work)
	/***************************/

	var newSaveString string
	var pathResponse story_serv.PathResponse
	var saveData struct {
		Raw  json.RawMessage
		Code string
	}

	err := work.AutoCommit([]uow.WorkFragment{storyService}, func() app_error.AppError {
		bookId, err := gin_tools.GetIntParam("book_id",c)
		if (err != nil){ return err }

		chapterId, err := gin_tools.GetIntParam("chapter_id",c)
		if (err != nil){ return err }

		saveString := c.Query("save")

		chapterPathId := gin_tools.GetIntQueryOrDefault("chapter_path_id", 0,c)	//chapter path can be ignored if we're on freemode
		isFreeMode := gin_tools.GetBoolQueryOrDefault("freemode",false,c)			//free mode is off by default

		pathRequest := story_serv.NewPathRequest(isFreeMode, bookId, chapterId, chapterPathId)
		pathResponse, err = storyService.NavigateToChapter(pathRequest,saveString)
		if (err != nil) { return err }

		newSaveString , err = pathResponse.NewSave.EncodedSaveString()
		if (err != nil) { return err }

		saveData.Code = newSaveString
		saveData.Raw = []byte(pathResponse.NewSave.JsonString)

		return nil
	})


	if(err != nil){
		api_reply.Failure(c,err)
	} else {
		api_reply.Success(c,gin.H{
			"save" : saveData, "chapter" : pathResponse.DestinationChapter,
		})
	}
}