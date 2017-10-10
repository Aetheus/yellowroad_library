package story_serv

import (
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
	"yellowroad_library/services/story_serv/story_save"
	"yellowroad_library/database/repo/uow"
)

type StoryService interface {
	NavigateToChapter(request PathRequest, encodedSaveString string) (response PathResponse, err app_error.AppError)
}

type StoryServiceFactory func(uow.UnitOfWork) StoryService

//all the necessary parameters to find out if you can navigate to a chapter in a book
type PathRequest struct {
	//if IsFreeMode == true and the DestinationChapterId + BookId combo are valid,
	//skip all further checking
	IsFreeMode bool
	BookId		int
	DestinationChapterId int

	ChapterPathId int
}
func NewPathRequest(IsFreeMode bool,BookId int,DestinationChapterId int,ChapterPathId int) PathRequest {
	return PathRequest{
		IsFreeMode ,
		BookId ,
		DestinationChapterId ,
		ChapterPathId,
	}
}

type PathResponse struct {
	DestinationChapter entities.Chapter
	NewSave            story_save.Save
}
func NewPathResponse(destinationChapter entities.Chapter,newSave story_save.Save) PathResponse {
	return PathResponse{
		destinationChapter,newSave,
	}
}
