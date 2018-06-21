package game_domain

import (
	"yellowroad_library/utils/app_error"
	"net/http"
	"yellowroad_library/database/repo/chapter_repo"
	"yellowroad_library/database/repo/chapterpath_repo"
)

type NavigateToChapter struct {
	chapterRepo 	chapter_repo.ChapterRepository
	chapterPathRepo chapterpath_repo.ChapterPathRepository
}

func (this NavigateToChapter) Execute(
	request PathRequest,
) (response PathResponse, err app_error.AppError) {
	destinationChapter, err := this.chapterRepo.FindWithinBook(request.DestinationChapterId, request.BookId)
	if (err != nil) { return response, err }

	isDestinationChapterTheFirstChapter := destinationChapter.ID == destinationChapter.Book.FirstChapterId

	currentSaveAsJsonString := string(request.SaveData)
	newSaveAsJsonString := currentSaveAsJsonString
	if !isValidSave(currentSaveAsJsonString) {
		err = app_error.New(http.StatusBadRequest,"","Save format was invalid!")
		return response, err
	}

	if !isDestinationChapterTheFirstChapter {
		chapterPath, err :=  this.chapterPathRepo.FindById(request.ChapterPathId)
		if (err != nil){ return response, err }

		err = validateSaveRequirements(currentSaveAsJsonString, chapterPath.Requirements.ToString())
		if (err != nil) { return response, err }

		effect := chapterPath.Effects.ToString()
		newSaveAsJsonString, err = applyEffectOnSave(currentSaveAsJsonString, effect)
		if (err != nil) { return response, err }
	}

	response = NewPathResponse(destinationChapter,newSaveAsJsonString)
	return response, nil
}