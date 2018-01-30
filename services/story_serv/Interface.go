package story_serv

import (
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
	"yellowroad_library/database/repo/uow"
	"encoding/json"
)

type StoryService interface {
	NavigateToChapter(request PathRequest) (response PathResponse, err app_error.AppError)
	SetUnitOfWork(work uow.UnitOfWork)
}


//all the necessary parameters to find out if you can navigate to a chapter in a book
type PathRequest struct {
	BookId		int
	DestinationChapterId int
	ChapterPathId int

	SaveData	json.RawMessage

	//if IsFreeMode == true and the DestinationChapterId + BookId combo are valid,
	//skip all further checking
	IsFreeMode bool
}

func NewPathRequest(BookId int, DestinationChapterId int,ChapterPathId int,SaveData	json.RawMessage,IsFreeMode bool) PathRequest {
	return PathRequest{
		BookId:BookId,
		DestinationChapterId:DestinationChapterId,
		ChapterPathId:ChapterPathId,
		SaveData:SaveData,
		IsFreeMode:IsFreeMode,
	}
}

type PathResponse struct {
	DestinationChapter entities.Chapter
	NewSaveData        json.RawMessage
}
func NewPathResponse(destinationChapter entities.Chapter,newSaveAsJsonString string) PathResponse {
	return PathResponse{
		destinationChapter,json.RawMessage(newSaveAsJsonString),
	}
}
