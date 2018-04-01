package game_serv

import (
	"yellowroad_library/utils/app_error"
	//"yellowroad_library/database/repo/chapter_repo"
	//"yellowroad_library/services/story_serv/story_save"
	"yellowroad_library/database/repo/uow"
	"net/http"
)

type DefaultGameService struct {
	work uow.UnitOfWork
}
var _ GameService = DefaultGameService{}

func Default(work uow.UnitOfWork) GameService {
	return DefaultGameService{ work: work }
}

func (this DefaultGameService) SetUnitOfWork(work uow.UnitOfWork) {
	this.work = work
}

func (this DefaultGameService) NavigateToChapter(
	request PathRequest,
) (response PathResponse, err app_error.AppError) {

	destinationChapter, err := this.work.ChapterRepo().FindWithinBook(request.DestinationChapterId,request.BookId)
	if (err != nil) { return response, err }

	destinationChapterIsFirstChapter := ( int64(destinationChapter.ID) == destinationChapter.Book.FirstChapterId.Int64 )

	currentSaveAsJsonString := string(request.SaveData)
	newSaveAsJsonString := currentSaveAsJsonString
	if( !isValidSave(currentSaveAsJsonString) ){
		err = app_error.New(http.StatusBadRequest,"","Save format was invalid!")
		return response, err
	}


	if request.ChapterPathId != 0 && !destinationChapterIsFirstChapter {
		chapterPath, err :=  this.work.ChapterPathRepo().FindById(request.ChapterPathId)
		if (err != nil){ return response, err }

		err = ValidateSaveRequirements(currentSaveAsJsonString, chapterPath.Requirements.ToString())
		if (err != nil) { return response, err }

		effect := chapterPath.Effects.ToString()
		newSaveAsJsonString, err = ApplyEffectOnSave(currentSaveAsJsonString, effect)
		if (err != nil) { return response, err }
	}

	response = NewPathResponse(destinationChapter,newSaveAsJsonString)
	return response, nil
}