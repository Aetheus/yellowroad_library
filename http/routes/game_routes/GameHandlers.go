package game_routes

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/utils/gin_tools"
	"yellowroad_library/utils/api_reply"
	"yellowroad_library/utils/app_error"
	"encoding/json"
	"yellowroad_library/domain/game_domain"
	"yellowroad_library/database/repo/uow"
)

type GameContainer interface{
	UnitOfWork() uow.UnitOfWork
	NavigateToChapter(uow.UnitOfWork) game_domain.NavigateToChapter
}
type GameHandlers struct {
	Container GameContainer
}

type NavigationForm struct {
	Save 			json.RawMessage		`json:"save"`
	IsFreeMode 		*bool 				`json:"mode"`//optional
	ChapterPathId	int  				`json:"chapter_path_id"`//chapter path can be ignored if we're on freemode or the first chapter
}
func (this GameHandlers) NavigateToSingleChapter(c *gin.Context) {
	/*Dependencies**************/
	work := this.Container.UnitOfWork()
	navigateToChapter := this.Container.NavigateToChapter(work)
	/***************************/

	var pathResponse game_domain.PathResponse
	var form NavigationForm
	err := work.AutoCommit(func() app_error.AppError {
		err := gin_tools.BindJSON(&form,c);
		if (err != nil){ return err }

		currentSave := form.Save
		chapterPathId := form.ChapterPathId
		isFreeMode := false
		if form.IsFreeMode != nil {
			isFreeMode = *form.IsFreeMode
		}

		bookId, err := gin_tools.GetIntParam("book_id",c)
		if (err != nil){ return err }

		chapterId, err := gin_tools.GetIntParam("chapter_id",c)
		if (err != nil){ return err }


		pathRequest := game_domain.NewPathRequest(bookId, chapterId, chapterPathId, currentSave, isFreeMode)
		pathResponse, err = navigateToChapter.Execute(pathRequest)
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