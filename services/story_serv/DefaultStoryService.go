package story_serv

import (
	"yellowroad_library/utils/app_error"
	//"yellowroad_library/database/repo/chapter_repo"
	//"yellowroad_library/services/story_serv/story_save"
	"yellowroad_library/database/repo/uow"
)

type DefaultStoryService struct {
	work uow.UnitOfWork
}
var _ StoryService = DefaultStoryService{}

func Default(work uow.UnitOfWork) StoryService {
	return DefaultStoryService{ work: work }
}

func (this DefaultStoryService) SetUnitOfWork(work uow.UnitOfWork) {
	this.work = work
}

func (this DefaultStoryService) NavigateToChapter(request PathRequest) (PathResponse,app_error.AppError) {

	destinationChapter, err := this.work.ChapterRepo().FindWithinBook(request.DestinationChapterId,request.BookId)
	if (err != nil) {
		return PathResponse{},err
	}

	destinationChapterIsFirstChapter := ( int64(destinationChapter.ID) == destinationChapter.Book.FirstChapterId.Int64 )

	currentSaveAsJsonString := string(request.SaveData)
	newSaveAsJsonString := currentSaveAsJsonString
	saveIsEmpty := currentSaveAsJsonString == ""

	//if the save is empty and we're not on the first chapter, then consider ourselves to be in freemode
	if (request.IsFreeMode || saveIsEmpty && !destinationChapterIsFirstChapter) {
		response := NewPathResponse(destinationChapter,currentSaveAsJsonString)
		return response, nil
	}


	if request.ChapterPathId != 0 {
		chapterPath, err :=  this.work.ChapterPathRepo().FindById(request.ChapterPathId)
		if (err != nil){
			return PathResponse{}, err
		}

		err = ValidateSaveRequirements(currentSaveAsJsonString, chapterPath.Requirements.ToString())
		if (err != nil) {
			return PathResponse{}, err
		}

		effect := chapterPath.Effects.ToString()
		newSaveAsJsonString, err = ApplyEffectOnSave(currentSaveAsJsonString, effect)
		if (err != nil) {
			return PathResponse{}, err
		}
	}

	response := NewPathResponse(destinationChapter,newSaveAsJsonString)
	return response, nil
}