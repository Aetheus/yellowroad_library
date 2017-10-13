package story_serv

import (
	"yellowroad_library/utils/app_error"
	//"yellowroad_library/database/repo/chapter_repo"
	"yellowroad_library/services/story_serv/story_save"
	"yellowroad_library/database/repo/uow"
)

type DefaultStoryService struct {
	work uow.UnitOfWork
}
var _ StoryService = DefaultStoryService{}

func Default(work uow.UnitOfWork) StoryService {
	return DefaultStoryService{
		work: work,
	}
}
var _ StoryServiceFactory = Default

func (this DefaultStoryService) SetUnitOfWork(work uow.UnitOfWork) {
	this.work = work
}

func (this DefaultStoryService) NavigateToChapter(request PathRequest, encodedSaveString string) (PathResponse,app_error.AppError) {

	destinationChapter, err := this.work.ChapterRepo().FindWithinBook(request.DestinationChapterId,request.BookId)
	if (err != nil) {
		return PathResponse{},err
	}

	saveIsEmpty := false
	if (len(encodedSaveString) == 0) {
		saveIsEmpty = true
	}

	destinationChapterIsFirstChapter := (destinationChapter.ID == destinationChapter.Book.FirstChapterId.Int)

	//if the save is empty and we're not on the first chapter, then consider ourselves to be in freemode
	if (request.IsFreeMode || saveIsEmpty && !destinationChapterIsFirstChapter) {
		response := NewPathResponse(destinationChapter,story_save.Save{})
		return response, nil
	}

	var currentSave story_save.Save
	if (saveIsEmpty) {
		currentSave = story_save.New()
	}else {
		currentSave, err = story_save.DecodeSaveString(encodedSaveString)
		if (err != nil){
			return PathResponse{}, err
		}
	}

	//TODO: actually do something here - get the chapter path, use it to validate/update the story save, etc
	response := NewPathResponse(destinationChapter,currentSave)
	return response, nil
}