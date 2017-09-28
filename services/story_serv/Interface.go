package story_serv

import (
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
)

type StoryService interface {
	NavigateToChapter(request PathRequest, encodedSaveString string) (response PathResponse, err app_error.AppError)
	EncodeSave(save Save) (encodedSaveString string, err app_error.AppError)
	DecodeSave(encodedSaveString string) (save Save, err app_error.AppError)
}

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
	destinationChapter entities.Chapter
	newSave            Save
}
func NewPathResponse(destinationChapter entities.Chapter,newSave Save) PathResponse {
	return PathResponse{
		destinationChapter,newSave,
	}
}

//TODO populate this struct
//save data ...
type Save struct {}
func NewSave() Save{
	return Save {}
}