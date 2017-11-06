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
)


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
		work := uow.NewAppUnitOfWork(db)

		Convey("Given a valid dataset", func (){
			dataset, err := createAuthorAndBookAndChaptersAndChapterPaths(work)
			So(err,ShouldBeNil)

			Convey("Given a DefaultStoryService", func(){
				storyServ := Default(work)

				Convey("Navigating to the first chapter should produce no error", func (){
					pathRequest := PathRequest{
						IsFreeMode : false,
						BookId : dataset.book.ID,
						DestinationChapterId : dataset.firstChapter.ID,
					}
					encodedSaveString := ""

					firstChapterResponse, err := storyServ.NavigateToChapter(pathRequest, encodedSaveString)
					So(err, ShouldBeNil)

					Convey("Navigating to the second chapter from the first one should produce no error", func (){
						pathRequest := PathRequest{
							IsFreeMode : false,
							BookId : dataset.book.ID,
							DestinationChapterId : dataset.secondChapter.ID,
							ChapterPathId: dataset.pathBetweenFirstAndSecondChapter.ID,
						}
						encodedSaveString, err := firstChapterResponse.NewSave.EncodedSaveString()
						So(err, ShouldBeNil)

						_, err = storyServ.NavigateToChapter(pathRequest, encodedSaveString)
						So(err, ShouldBeNil)
					})
				})



			});
		})

		Reset(func(){
			err := work.Rollback()
			So(err,ShouldBeNil)
		})
	}))

}
