package app_story_serv

import (
	"yellowroad_library/services/story_serv"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/repo/chapter_repo"
	"yellowroad_library/services/story_serv/story_save"
)

type AppStoryService struct {
	chapterRepo chapter_repo.ChapterRepository
}
var _ story_serv.StoryService = AppStoryService{}

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

	saveIsEmpty := false
	if (len(encodedSaveString) == 0) {
		saveIsEmpty = true
	}

	destinationChapterIsFirstChapter := (destinationChapter.ID == destinationChapter.Book.FirstChapterId.Int)

	//if the save is empty and we're not on the first chapter, then consider ourselves to be in freemode
	if (request.IsFreeMode || saveIsEmpty && !destinationChapterIsFirstChapter) {
		response := story_serv.NewPathResponse(destinationChapter,story_save.Save{})
		return response, nil
	}

	var currentSave story_save.Save
	if (saveIsEmpty) {
		currentSave = story_save.New()
	}else {
		currentSave, err = story_save.DecodeSaveString(encodedSaveString)
		if (err != nil){
			return story_serv.PathResponse{}, err
		}
	}

	//TODO: actually do something here - get the chapter path, use it to validate/update the story save, etc
	response := story_serv.NewPathResponse(destinationChapter,currentSave)
	return response, nil
}