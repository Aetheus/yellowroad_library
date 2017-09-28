package app_story_serv

import (
	"yellowroad_library/services/story_serv"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/repo/chapter_repo"
)

type AppStoryService struct {
	chapterRepo chapter_repo.ChapterRepository
}

func New(chapterRepo chapter_repo.ChapterRepository) AppStoryService{
	return AppStoryService{
		chapterRepo : chapterRepo,
	}
}

func (this AppStoryService) NavigateToChapter(request story_serv.PathRequest, encodedSaveString string) (story_serv.PathResponse,app_error.AppError) {

	destinationChapter, err := this.chapterRepo.FindWithinBook(request.DestinationChapterId,request.BookId)
	if (err != nil) {
		return story_serv.PathResponse{},err
	}

	if (request.IsFreeMode) {
		response := story_serv.NewPathResponse(destinationChapter,story_serv.Save{})
		return response, nil
	}

	var currentSave story_serv.Save
	if (len(encodedSaveString) == 0) {
		currentSave = story_serv.NewSave()
	}else {
		currentSave, err = this.DecodeSave(encodedSaveString)
		if (err != nil){
			return story_serv.PathResponse{}, err
		}
	}

	//TODO: actually do something here - get the chapter path, use it to validate/update the story save, etc
	response := story_serv.NewPathResponse(destinationChapter,currentSave)
	return response, nil
}

func (this AppStoryService) EncodeSave(save story_serv.Save) (encodedSaveString string, err app_error.AppError) {
	//TODO: implement this
	panic("implement me")
}

func (this AppStoryService) DecodeSave(encodedSaveString string) (save story_serv.Save, err app_error.AppError) {
	//TODO: implement this
	panic("implement me")
}
