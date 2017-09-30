package story_route

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/services/story_serv"
	"yellowroad_library/utils/gin_tools"
	"yellowroad_library/utils/api_response"
)

func NavigateToSingleChapter(service story_serv.StoryService) gin.HandlerFunc {
	return func (c *gin.Context){
		bookId, err := gin_tools.GetIntParam("book_id",c)
		if (err != nil){
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		chapterId, err := gin_tools.GetIntParam("chapter_id",c)
		if (err != nil){
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		//chapter path can be ignored if we're on freemode
		chapterPathId := gin_tools.GetIntQueryOrDefault("chapter_path_id", 0,c)

		//free mode is off by default
		isFreeMode := gin_tools.GetBoolQueryOrDefault("freemode",false,c)

		saveString := c.Query("save")

		pathRequest := story_serv.NewPathRequest(isFreeMode, bookId, chapterId, chapterPathId)
		pathResponse, err := service.NavigateToChapter(pathRequest,saveString)
		if (err != nil) {
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		saveString , _ = pathResponse.NewSave.Encode()
		c.JSON(api_response.SuccessWithCode(gin.H{
			"chapter" : pathResponse.DestinationChapter,
			"save" : saveString,
		}))
		return
	}
}