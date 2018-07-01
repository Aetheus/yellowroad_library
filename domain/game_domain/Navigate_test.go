package game_domain

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"yellowroad_library/database/repo/chapterpath_repo"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/repo/chapter_repo"
	"encoding/json"
	"yellowroad_library/database"
	"fmt"
)

// mock return values
var (
	creator = entities.User{
		ID: 50, Username: "Robert Baratheon", Password: "onanopenfield", Email: "bobbyb@rightfulkingofwesteros.com",
	}
	firstChapterId = 101
	firstBook = entities.Book{
		ID:10, Title: "Game of Thrones, Season XIIVIIXII", CreatorId: creator.ID,
		Description: "From the best selling author of Game of Thrones, Season XIIVIIXI",
		FirstChapterId: firstChapterId,
	}
	firstChapter = entities.Chapter{
		ID : firstChapterId, Title:  "First things first..." ,
		Body:   "You wake up. You're hungry. What's for breakfast? Bacon and eggs? Or Salad?",
		BookId: firstBook.ID, Book: &firstBook,
		CreatorId: creator.ID, Creator: &creator,
	}
	secondChapter = entities.Chapter{
		ID : 102, Title:  "Bacon & Eggs",
		Body:   "It tastes good. You can feel the cholesterol starting to gum up your arteries, but eh, whatever. What next?",
		BookId: firstBook.ID, Book: &firstBook,
		CreatorId: creator.ID, Creator: &creator,
	}
	thirdChapterA = entities.Chapter{
		ID : 103, Title: "Lounge around at home",
		Body: "You decide to be a potato. You sit on a couch. You do nothing. You begin to feel relaxed. That is, until something draws your attention ...",
		BookId: firstBook.ID, Book: &firstBook,
		CreatorId: creator.ID, Creator: &creator,
	}
	thirdChapterB = entities.Chapter{
		ID: 104, Title: "Go job hunting",
		Body : `You put on your smartest shirt and start wandering the commercial district, applying to anywhere with a 'Help Wanted' sign out in the front. It's rather stressful. Eventually, you hear a voice calling out to you. This voice was ...`,
		BookId: firstBook.ID, Book: &firstBook,
		CreatorId: creator.ID, Creator: &creator,
	}
	pathFromFirstChapterToSecondChapter = entities.ChapterPath{
		ID : 1001, FromChapterId: firstChapter.ID, ToChapterId:   secondChapter.ID,
		Effects 	: 	database.Jsonb {json.RawMessage(`
							{ "/morale": {"op":"SET","arg":50 }, "/health": {"op":"SET","arg":-5} }
						`)},
		Requirements : 	database.Jsonb {json.RawMessage(`
							{"type" : "object"}
						`)},
	}
	pathFromSecondChapterToThirdChapterA = entities.ChapterPath{
		ID : 1002,
		FromChapterId: secondChapter.ID,
		ToChapterId:   thirdChapterA.ID,
		Effects: database.Jsonb{json.RawMessage(`
					{"/morale": {"op":"INCR", "arg":5 }, "/status": "RELAXED"}
				`)},
		Requirements: 	database.Jsonb{
			json.RawMessage(
				`{
					"type": "object",
					"properties" : {
						"health" : { "type": "integer", "title": "health", "minimum":5 }
					}
				}`,
			),
		},
	}
	pathFromSecondChapterToThirdChapterB = entities.ChapterPath{
		ID : 1003,
		FromChapterId: secondChapter.ID,
		ToChapterId:   thirdChapterB.ID,
		Effects : 	database.Jsonb {json.RawMessage(`
						{"/morale": {"op":"INCR", "arg":5 },"/status": "RELAXED"}
					`)},
		Requirements : 	database.Jsonb {
			json.RawMessage(
				`{
					"type" : "object",
					"properties" : {
						"morale" : { "type": "integer", "title": "morale", "minimum":50 }
					}
				}`,
			),
		},
	}
)

// mocks
var (
	chapterRepo = &chapter_repo.ChapterRepositoryMock{
		FindWithinBookFunc: func(chapter_id int, book_id int) (entities.Chapter, app_error.AppError) {
			if (book_id == firstChapter.BookId && chapter_id == firstChapter.ID) {
				return firstChapter, nil
			} else if (book_id == secondChapter.BookId && chapter_id == secondChapter.ID) {
				return secondChapter, nil
			} else if (book_id == thirdChapterA.BookId && chapter_id == thirdChapterA.ID) {
				return thirdChapterA, nil
			} else if (book_id == thirdChapterB.BookId && chapter_id == thirdChapterB.ID) {
				return thirdChapterB, nil
			}
			panic(fmt.Sprint("Unexpected inputs:", chapter_id, book_id))
		},
	}
	chapterPathRepo = &chapterpath_repo.ChapterPathRepositoryMock{
		FindByIdFunc: func(chapterPathId int) (entities.ChapterPath, app_error.AppError) {
			switch(chapterPathId) {
				case pathFromFirstChapterToSecondChapter.ID: return pathFromFirstChapterToSecondChapter, nil
				case pathFromSecondChapterToThirdChapterA.ID: return pathFromSecondChapterToThirdChapterA,nil
				case pathFromSecondChapterToThirdChapterB.ID: return pathFromSecondChapterToThirdChapterB,nil
			}

			panic(fmt.Sprint("Unexpected inputs:", chapterPathId))
		},
	}
)

func TestDefaultStoryService_NavigateToChapter(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given an instance of NavigateToChapter", t,  func(){
		navigateToChapter := NewNavigateToChapter(chapterRepo,chapterPathRepo)

		Convey("Navigating to the first chapter with an empty save should work", func (){
			request := PathRequest{
				IsFreeMode : false,
				BookId : firstChapter.BookId,
				DestinationChapterId : firstChapter.ID,
				SaveData: json.RawMessage("{}"),
			}

			firstChapterResponse, err := navigateToChapter.Execute(request)

			So(err,ShouldBeNil)

			Convey("Navigating to the second chapter from the first one should work", func () {
				request = PathRequest{
					IsFreeMode:           false,
					BookId:               secondChapter.BookId,
					DestinationChapterId: secondChapter.ID,
					ChapterPathId:        pathFromFirstChapterToSecondChapter.ID,
					SaveData:             firstChapterResponse.NewSaveData,
				}

				secondChapterResponse, err := navigateToChapter.Execute(request)
				So(err, ShouldBeNil)

				Convey("Save from navigating to the second chapter should reflect second chapter's effects", func (){
					newSaveDocument := map[string]interface{}{}
					json.Unmarshal(secondChapterResponse.NewSaveData, &newSaveDocument)
					So(newSaveDocument["health"], ShouldEqual, -5)
					So(newSaveDocument["morale"], ShouldEqual, 50)
				})

				Convey("Navigating to third chapter A from the second chapter should return an error since it doesn't fulfill the requirements", func (){
					request = PathRequest {
						IsFreeMode : false,
						BookId : thirdChapterA.BookId,
						DestinationChapterId:thirdChapterA.ID,
						ChapterPathId : pathFromSecondChapterToThirdChapterA.ID,
						SaveData : secondChapterResponse.NewSaveData,
					}

					_, err = navigateToChapter.Execute(request)
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, "- health: Must be greater than or equal to 5")
				})
			})
		})
	})
}