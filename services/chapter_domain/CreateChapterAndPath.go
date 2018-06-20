package chapter_domain

import (
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/repo/chapter_repo"
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/database/repo/chapterpath_repo"
)

type CreateChapterAndPath struct {
	createChapter CreateChapter
	createPathBetweenChapters CreatePathBetweenChapters
}

func NewCreateChapterAndPath(
	chapterRepo chapter_repo.ChapterRepository,
	chapterPathRepo chapterpath_repo.ChapterPathRepository,
	bookRepo book_repo.BookRepository,
) CreateChapterAndPath {
	return CreateChapterAndPath{
		NewCreateChapter(chapterRepo,bookRepo),
		NewCreatePathBetweenChapters(chapterRepo,chapterPathRepo,bookRepo),
	}
}

func (this CreateChapterAndPath) CreateChapterAndPath(
	user entities.User,
	form entities.Chapter_And_Path_CreationForm,
) (chapter entities.Chapter,chapter_path entities.ChapterPath, err app_error.AppError){
	chapter_form := *form.ChapterForm
	path_form := *form.ChapterPathForm

	chapter, err = this.createChapter.Execute(user, chapter_form)
	if (err != nil){
		return
	}
	path_form.ToChapterId = &chapter.ID

	chapter_path, err = this.createPathBetweenChapters.Execute(user,path_form)
	return
}
