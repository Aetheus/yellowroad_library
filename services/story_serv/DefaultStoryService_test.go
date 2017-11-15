package story_serv

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"yellowroad_library/test_utils"
	//"yellowroad_library/database/entities"
	//"yellowroad_library/database/repo/user_repo"
	//"yellowroad_library/database/repo/user_repo/gorm_user_repo"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/database/entities"
	"encoding/json"
	"yellowroad_library/database"
	"github.com/jinzhu/gorm"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/repo/chapterpath_repo"
	"yellowroad_library/database/repo/chapter_repo"
)

func createMockUnitOfWork(
	mock_user_id int, mock_book_id int, mock_first_chapter_id int, mock_second_chapter_id int,
	first_to_second_chapter_path_id int,
) uow.UnitOfWork {
	return &uow.UnitOfWorkMock{
		AutoFunc: func(in1 []uow.WorkFragment, in2 func() app_error.AppError) app_error.AppError {
			return nil
		},
		CommitFunc: func() app_error.AppError {
			return nil
		},
		RollbackFunc: func() app_error.AppError {
			return nil
		},
		ChapterPathRepoFunc: func() chapterpath_repo.ChapterPathRepository {
			panic("TODO: mock out the ChapterPathRepo method")
			return &chapterpath_repo.ChapterPathRepositoryMock{
				FindByIdFunc: func(chapterId int) (entities.ChapterPath, app_error.AppError) {
					switch(chapterId) {
					case first_to_second_chapter_path_id : return entities.ChapterPath{
						FromChapterId: mock_first_chapter_id,
						ToChapterId:   mock_second_chapter_id,
						Effects : database.Jsonb {json.RawMessage(`{
							"/morale": {"op":"SET","arg":50 },
							"/health": {"op":"SET","arg":-5  }
						}`),},
						Requirements : database.Jsonb {json.RawMessage(`{}`),},
					}, nil
					default :
						panic("unexpected behaviour")
					}
				},
			}
		},
		ChapterRepoFunc: func() chapter_repo.ChapterRepository {
			return &chapter_repo.ChapterRepositoryMock {
				FindWithinBookFunc :func(chapter_id int, book_id int) (entities.Chapter, app_error.AppError){
					if chapter_id == mock_first_chapter_id && book_id == mock_book_id {
						return entities.Chapter{
							Title :     "First things first ... ",
							Body :      "You wake up. You're hungry. What's for breakfast? Bacon and eggs? Or Salad?",
							BookId :    mock_book_id,
							Book : &entities.Book {
								Title :       "Game of Thrones, Season XIIVIIXII",
								Description : "From the best selling author of Game of Thrones, Season XIIVIIXI",
								CreatorId :   mock_user_id,
							},
							CreatorId : mock_user_id,
							Creator   : &entities.User {
								ID:        mock_user_id,
								Username : "Robert Baratheon",
								Password : "onanopenfield",
								Email:     "bobbyb@rightfulkingofwesteros.com",
							},
						}, nil
					} else if chapter_id == mock_second_chapter_id && book_id == mock_book_id {
						return entities.Chapter{
							Title :     "Bacon & Eggs",
							Body :      "It tastes good. You can feel the cholesterol starting to gum up your arteries, but eh, whatever. What next?",
							BookId :    mock_book_id,
							Book : &entities.Book {
								Title :       "Game of Thrones, Season XIIVIIXII",
								Description : "From the best selling author of Game of Thrones, Season XIIVIIXI",
								CreatorId :   mock_user_id,
							},
							CreatorId : mock_user_id,
							Creator   : &entities.User {
								ID:        mock_user_id,
								Username : "Robert Baratheon",
								Password : "onanopenfield",
								Email:     "bobbyb@rightfulkingofwesteros.com",
							},
						}, nil
					} else {
						panic("TODO: unexpected arguments")
					}
				},
			}
		},
	}
}


//TODO : rather than actually create and insert all these elements because we use real repos,
//maybe we should just use some mock repos?
func createAuthorAndBookAndChaptersAndChapterPaths(work uow.UnitOfWork) (
	testDataset struct {
		author entities.User
		book entities.Book
		firstChapter entities.Chapter
		secondChapter entities.Chapter
		thirdChapterA entities.Chapter
		thirdChapterB entities.Chapter
		pathBetweenFirstAndSecondChapter entities.ChapterPath
		pathBetweenSecondAndThirdChapterA entities.ChapterPath
		pathBetweenSecondAndThirdChapterB entities.ChapterPath
	},
	err error,
){
	testDataset.author = entities.User {
		Username : "Robert Baratheon",
		Password : "onanopenfield",
		Email: "bobbyb@rightfulkingofwesteros.com",
	}
	if err = work.UserRepo().Insert(&testDataset.author);err != nil{
		return
	}

	testDataset.book = entities.Book{
		Title :       "Game of Thrones, Season XIIVIIXII",
		Description : "From the best selling author of Game of Thrones, Season XIIVIIXI",
		CreatorId :   testDataset.author.ID,
	}
	if 	err = work.BookRepo().Insert(&testDataset.book);err != nil{
		return
	}

	testDataset.firstChapter = entities.Chapter{
		Title :     "First things first ... ",
		Body :      "You wake up. You're hungry. What's for breakfast? Bacon and eggs? Or Salad?",
		BookId :    testDataset.book.ID,
		CreatorId : testDataset.author.ID,
	}
	if 	err = work.ChapterRepo().Insert(&testDataset.firstChapter);err != nil{
		return
	}

	testDataset.secondChapter = entities.Chapter{
		Title :     "Bacon & Eggs",
		Body :      "It tastes good. You can feel the cholesterol starting to gum up your arteries, but eh, whatever. What next?",
		BookId :    testDataset.book.ID,
		CreatorId : testDataset.author.ID,
	}
	if err = work.ChapterRepo().Insert(&testDataset.secondChapter); err != nil{
		return
	}

	testDataset.thirdChapterA = entities.Chapter {
		Title :     "Lounge around at home",
		Body :      "You decide to be a potato. You sit on a couch. You do nothing. You begin to feel relaxed. That is, until something draws your attention ...",
		BookId :    testDataset.book.ID,
		CreatorId : testDataset.author.ID,
	}
	if err = work.ChapterRepo().Insert(&testDataset.thirdChapterA); err != nil{
		return
	}

	testDataset.thirdChapterB = entities.Chapter {
		Title :     "Go job hunting",
		Body :      `You put on your smartest shirt and start wandering the commercial district, applying to anywhere with a 'Help Wanted' sign out in the front. It's rather stressful. Eventually, you hear a voice calling out to you. This voice was ...`,
		BookId :    testDataset.book.ID,
		CreatorId : testDataset.author.ID,
	}
	if err = work.ChapterRepo().Insert(&testDataset.thirdChapterB); err != nil{
		return
	}

	testDataset.pathBetweenFirstAndSecondChapter = entities.ChapterPath{
		FromChapterId: testDataset.firstChapter.ID,
		ToChapterId:   testDataset.secondChapter.ID,
		Effects : database.Jsonb {json.RawMessage(`{
			"/morale": {"op":"SET","arg":50 },
			"/health": {"op":"SET","arg":-5  }
		}`),},
		Requirements : database.Jsonb {json.RawMessage(`{}`),},
	}
	if err = work.ChapterPathRepo().Insert(&testDataset.pathBetweenFirstAndSecondChapter) ; err != nil{
		return
	}

	testDataset.pathBetweenSecondAndThirdChapterA = entities.ChapterPath{
		FromChapterId: testDataset.secondChapter.ID,
		ToChapterId:   testDataset.thirdChapterA.ID,
		Effects : database.Jsonb {json.RawMessage(`{
			"/morale": {"op":"INCR", "arg":5 },
			"/status": "RELAXED"
		}`),},
		Requirements : database.Jsonb {json.RawMessage(`{
			"type" : "object",
			"properties" : {
				"health" : {
			  		"type": "integer",
      				"title": "health",
					"minimum":5
				}
			}
		}`),},
	}
	if err = work.ChapterPathRepo().Insert(&testDataset.pathBetweenSecondAndThirdChapterA); err != nil{
		return
	}

	testDataset.pathBetweenSecondAndThirdChapterB = entities.ChapterPath{
		FromChapterId: testDataset.secondChapter.ID,
		ToChapterId:   testDataset.thirdChapterB.ID,
		Effects : database.Jsonb {json.RawMessage(`{
			"/morale": {"op":"INCR", "arg":5 },
			"/status": "RELAXED"
		}`),},
		Requirements : database.Jsonb {json.RawMessage(`{
			"type" : "object",
			"properties" : {
				"morale" : {
			  		"type": "integer",
      				"title": "morale",
					"minimum":50
				}
			}
		}`),},
	}
	if err = work.ChapterPathRepo().Insert(&testDataset.pathBetweenSecondAndThirdChapterB); err != nil{
		return
	}

	return
}

func TestGormBookRepository(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given a UnitOfWork", t, test_utils.WithRealGormDBConnection(func(db *gorm.DB){
		mock_user_id := 1
		mock_book_id := 20
		mock_first_chapter_id := 15
		mock_second_chapter_id := 17
		first_to_second_chapter_path_id := 21
		work := createMockUnitOfWork(mock_user_id, mock_book_id, mock_first_chapter_id, mock_second_chapter_id, first_to_second_chapter_path_id)

			Convey("Given a DefaultStoryService", func(){
				storyServ := Default(work)

				Convey("Navigating to the first chapter should produce no error", func (){
					pathRequest := PathRequest{
						IsFreeMode : false,
						BookId : mock_book_id,
						DestinationChapterId : mock_first_chapter_id,
					}
					encodedSaveString := ""

					firstChapterResponse, err := storyServ.NavigateToChapter(pathRequest, encodedSaveString)
					So(err, ShouldBeNil)

					Convey("Navigating to the second chapter from the first one should produce no error", func (){
						pathRequest := PathRequest{
							IsFreeMode : false,
							BookId : mock_book_id,
							DestinationChapterId : mock_second_chapter_id,
							ChapterPathId: first_to_second_chapter_path_id,
						}
						encodedSaveString, err := firstChapterResponse.NewSave.EncodedSaveString()
						So(err, ShouldBeNil)

						response, err := storyServ.NavigateToChapter(pathRequest, encodedSaveString)
						So(err, ShouldBeNil)
						So(response.NewSave.JsonString, ShouldNotEqual, "")
						So(response.NewSave.JsonString, ShouldNotEqual, "{}")
					})
				})

			});

		Reset(func(){
			err := work.Rollback()
			So(err,ShouldBeNil)
		})
	}))

}
