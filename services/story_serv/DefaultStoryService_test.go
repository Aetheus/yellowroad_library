package story_serv

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	//"yellowroad_library/database/entities"
	//"yellowroad_library/database/repo/user_repo"
	//"yellowroad_library/database/repo/user_repo/gorm_user_repo"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/database/entities"
	"encoding/json"
	"yellowroad_library/database"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/repo/chapterpath_repo"
	"yellowroad_library/database/repo/chapter_repo"
	"fmt"
)

func createMockUnitOfWork(
	mock_user_id int, mock_book_id int, mock_first_chapter_id int, mock_second_chapter_id int,
	mock_third_chapter_a_id int, mock_third_chapter_b_id int,
	first_to_second_chapter_path_id int, second_to_third_a_chapter_path_id int, second_to_third_b_chapter_path_id int,
) uow.UnitOfWork {
	book := entities.Book{
		Title:       "Game of Thrones, Season XIIVIIXII",
		Description: "From the best selling author of Game of Thrones, Season XIIVIIXI",
		CreatorId:   mock_user_id,
	}
	creator := entities.User{
		ID:       mock_user_id,
		Username: "Robert Baratheon",
		Password: "onanopenfield",
		Email:    "bobbyb@rightfulkingofwesteros.com",
	}

	return &uow.UnitOfWorkMock{
		AutoCommitFunc: func(in1 []uow.WorkFragment, in2 func() app_error.AppError) app_error.AppError {
			return nil
		},
		CommitFunc: func() app_error.AppError {
			return nil
		},
		RollbackFunc: func() app_error.AppError {
			return nil
		},
		ChapterPathRepoFunc: func() chapterpath_repo.ChapterPathRepository {
			return &chapterpath_repo.ChapterPathRepositoryMock{
				FindByIdFunc: func(chapter_path_id int) (entities.ChapterPath, app_error.AppError) {
					switch(chapter_path_id) {
						case first_to_second_chapter_path_id : return entities.ChapterPath{
							ID : first_to_second_chapter_path_id,
							FromChapterId: mock_first_chapter_id,
							ToChapterId:   mock_second_chapter_id,
							Effects : database.Jsonb {json.RawMessage(`{
								"/morale": {"op":"SET","arg":50 },
								"/health": {"op":"SET","arg":-5  }
							}`),},
							Requirements : database.Jsonb {json.RawMessage(`{
								"type" : "object"
							}`)},
						}, nil
						case second_to_third_a_chapter_path_id : return entities.ChapterPath{
							ID : second_to_third_a_chapter_path_id,
							FromChapterId: mock_second_chapter_id,
							ToChapterId:   mock_third_chapter_a_id,
							Effects: database.Jsonb{json.RawMessage(`{
												"/morale": {"op":"INCR", "arg":5 },
												"/status": "RELAXED"
											}`),},
							Requirements: 	database.Jsonb{
												json.RawMessage(
													`{
														"type" : "object",
														"properties" : {
															"health" : {
																"type": "integer",
																"title": "health",
																"minimum":5
															}
														}
													}`,
												),
											},
						}, nil
						case second_to_third_b_chapter_path_id : return entities.ChapterPath{
							ID : second_to_third_b_chapter_path_id,
							FromChapterId: mock_second_chapter_id,
							ToChapterId:   mock_third_chapter_b_id,
							Effects : database.Jsonb {json.RawMessage(`{
								"/morale": {"op":"INCR", "arg":5 },
								"/status": "RELAXED"
							}`),},
							Requirements : 	database.Jsonb {
												json.RawMessage(
													`{
														"type" : "object",
														"properties" : {
															"morale" : {
																"type": "integer",
																"title": "morale",
																"minimum":50
															}
														}
													}`,
												),
											},
						}, nil
						default : panic(fmt.Sprintf("unexpected behaviour. Arguments were: chapter_path_id(%s) ", chapter_path_id) )
					}
				},
			}
		},
		ChapterRepoFunc: func() chapter_repo.ChapterRepository {
			return &chapter_repo.ChapterRepositoryMock {
				FindWithinBookFunc :func(chapter_id int, book_id int) (entities.Chapter, app_error.AppError) {

					if (book_id == mock_book_id ){
						switch(chapter_id){
							case(mock_first_chapter_id):{
								return entities.Chapter{
									ID : mock_first_chapter_id,
									Title:  "First things first ... ",
									Body:   "You wake up. You're hungry. What's for breakfast? Bacon and eggs? Or Salad?",
									BookId: mock_book_id,
									Book: &book,
									CreatorId: mock_user_id,
									Creator: &creator,
								}, nil
							}
							case(mock_second_chapter_id) : {
								return entities.Chapter{
									ID : mock_second_chapter_id,
									Title:  "Bacon & Eggs",
									Body:   "It tastes good. You can feel the cholesterol starting to gum up your arteries, but eh, whatever. What next?",
									BookId: mock_book_id,
									Book: &book,
									CreatorId: mock_user_id,
									Creator: &creator,
								}, nil
							}
							case(mock_third_chapter_a_id) : {
								return entities.Chapter{
									ID : mock_third_chapter_a_id,
									Title :     "Lounge around at home",
									Body :      "You decide to be a potato. You sit on a couch. You do nothing. You begin to feel relaxed. That is, until something draws your attention ...",
									BookId :    mock_book_id,
									Book: &book,
									CreatorId : mock_user_id,
									Creator: &creator,
								}, nil
							}
							case(mock_third_chapter_b_id) : {
								return entities.Chapter{
									ID: mock_third_chapter_b_id,
									Title :     "Go job hunting",
									Body :      `You put on your smartest shirt and start wandering the commercial district, applying to anywhere with a 'Help Wanted' sign out in the front. It's rather stressful. Eventually, you hear a voice calling out to you. This voice was ...`,
									BookId :    mock_book_id,
									Book: &book,
									CreatorId : mock_user_id,
									Creator: &creator,
								}, nil
							}
							default : panic(fmt.Sprintf("unexpected behaviour. Arguments were: chapter_id(%s), book_id(%s)  ", chapter_id, book_id) )
						}
					}

					panic(fmt.Sprintf("unexpected behaviour. Arguments were: chapter_id(%s), book_id(%s)  ", chapter_id, book_id) )
				},
			}
		},
	}
}


func TestDefaultStoryService_NavigateToChapter(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given a UnitOfWork", t,  func(){
		mock_user_id := 1
		mock_book_id := 20
		mock_first_chapter_id := 15
		mock_second_chapter_id := 17
		mock_third_chapter_a_id := 18
		mock_third_chapter_b_id := 19
		first_to_second_chapter_path_id := 21
		second_to_third_a_chapter_path_id := 22
		second_to_third_b_chapter_path_id := 23

		work := createMockUnitOfWork(
			mock_user_id,
			mock_book_id,
			mock_first_chapter_id,
			mock_second_chapter_id,
			mock_third_chapter_a_id,
			mock_third_chapter_b_id,
			first_to_second_chapter_path_id,
			second_to_third_a_chapter_path_id,
			second_to_third_b_chapter_path_id,
		)

			Convey("Given a DefaultStoryService", func(){
				storyServ := Default(work)

				Convey("Navigating to the first chapter should produce no error", func (){
					pathRequest := PathRequest{
						IsFreeMode : false,
						BookId : mock_book_id,
						DestinationChapterId : mock_first_chapter_id,
						SaveData: json.RawMessage("{}"),
					}

					firstChapterResponse, err := storyServ.NavigateToChapter(pathRequest)
					So(err, ShouldBeNil)

					Convey("Navigating to the second chapter from the first one should produce no error", func (){
						pathRequest := PathRequest{
							IsFreeMode : false,
							BookId : mock_book_id,
							DestinationChapterId : mock_second_chapter_id,
							ChapterPathId: first_to_second_chapter_path_id,
							SaveData: firstChapterResponse.NewSaveData,
						}

						So(err, ShouldBeNil)

						secondChapterResponse, err := storyServ.NavigateToChapter(pathRequest)
						So(err, ShouldBeNil)

						Convey("Save from navigating to the second chapter should reflect second chapter's effects", func (){
							newSaveDocument := map[string]interface{}{}
							json.Unmarshal(secondChapterResponse.NewSaveData, &newSaveDocument)
							So(newSaveDocument["health"], ShouldEqual, -5)
							So(newSaveDocument["morale"], ShouldEqual, 50)
						})

						Convey("Navigating to third chapter A from the second chapter should produce an error since it doesn't fulfill the requirements", func (){
							pathRequest := PathRequest {
								IsFreeMode : false,
								BookId : mock_book_id,
								DestinationChapterId:mock_third_chapter_a_id,
								ChapterPathId : second_to_third_a_chapter_path_id,
								SaveData : secondChapterResponse.NewSaveData,
							}

							So(err, ShouldBeNil)

							_, err = storyServ.NavigateToChapter(pathRequest)
							So(err, ShouldNotBeNil)
							So(err.Error(), ShouldEqual, "- health: Must be greater than or equal to 5")
						})

						//Convey("Navigating to third chapter A from the second chapter with Free Mode set to true should produce no error", func (){
						//	pathRequest := PathRequest {
						//		IsFreeMode : true,
						//		BookId : mock_book_id,
						//		DestinationChapterId:mock_third_chapter_a_id,
						//		SaveData : secondChapterResponse.NewSaveData,
						//	}
						//
						//	So(err, ShouldBeNil)
						//
						//	thirdChapterAResponse, err := storyServ.NavigateToChapter(pathRequest)
						//	So(err, ShouldBeNil)
						//	So(thirdChapterAResponse.DestinationChapter.ID, ShouldEqual, mock_third_chapter_a_id)
						//	So(thirdChapterAResponse.DestinationChapter.Title, ShouldEqual,"Lounge around at home" )
						//	So(thirdChapterAResponse.DestinationChapter.Body, ShouldEqual,"You decide to be a potato. You sit on a couch. You do nothing. You begin to feel relaxed. That is, until something draws your attention ..." )
						//})
					})
				})

			});


	})

}
