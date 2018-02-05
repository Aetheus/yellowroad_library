package story_routes

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/services/story_serv"
	"yellowroad_library/utils/gin_tools"
	"yellowroad_library/utils/api_reply"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/containers"
	"encoding/json"
)

type StoryHandlers struct {
	Container containers.Container
}


type NavigationForm struct {
	Save 		json.RawMessage
	IsFreeMode 	*bool //optional
}
func (this StoryHandlers) NavigateToSingleChapter(c *gin.Context) {
	/*Dependencies**************/
	work := this.Container.UnitOfWork()
	storyService := this.Container.StoryService(work)
	/***************************/

	var pathResponse story_serv.PathResponse
	var form NavigationForm
	err := work.AutoCommit([]uow.WorkFragment{storyService}, func() app_error.AppError {
		err := gin_tools.BindJSON(&form,c);
		if (err != nil){ return err }

		currentSave := form.Save
		isFreeMode := false
		if form.IsFreeMode != nil {
			isFreeMode = *form.IsFreeMode
		}

		bookId, err := gin_tools.GetIntParam("book_id",c)
		if (err != nil){ return err }

		chapterId, err := gin_tools.GetIntParam("chapter_id",c)
		if (err != nil){ return err }

		chapterPathId := gin_tools.GetIntQueryOrDefault("chapter_path_id", 0,c)	//chapter path can be ignored if we're on freemode

		pathRequest := story_serv.NewPathRequest(bookId, chapterId, chapterPathId, currentSave, isFreeMode)
		pathResponse, err = storyService.NavigateToChapter(pathRequest)
		if (err != nil) { return err }

		return nil
	})


	if(err != nil){
		api_reply.Failure(c,err)
	} else {
		api_reply.Success(c,gin.H{
			"save" : pathResponse.NewSaveData, "chapter" : pathResponse.DestinationChapter,
		})
	}
}