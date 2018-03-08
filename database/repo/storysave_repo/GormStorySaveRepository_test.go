package storysave_repo

import "testing"
import (
	. "github.com/smartystreets/goconvey/convey"
	"yellowroad_library/test_utils"
	"github.com/jinzhu/gorm"
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/database/entities"
	"yellowroad_library/database"
	"encoding/json"
	"yellowroad_library/database/repo/chapter_repo"
	"yellowroad_library/database/repo/user_repo"
)

func TestGormStorySaveRepository(t *testing.T){
	// Only pass t into top-level Convey calls
	Convey("Given a GormBookRepository, UserRepository and GormStorySaveRepository", t, test_utils.WithRealGormDBConnection(func(gormDB *gorm.DB){
		transaction := gormDB.Begin()

		storySaveRepo := NewDefault(transaction)
		userRepo := user_repo.NewDefault(transaction)
		bookRepo := book_repo.NewDefault(transaction)
		chapterRepo := chapter_repo.NewDefault(transaction)

		Convey("Creating test user, book and chapter", func (){
			user := entities.User{
				Username:  "Jill Bates",
				Password:  "pass123132",
				Email:     "jill@bates.com",
			}
			err := userRepo.Insert(&user)
			So(err, ShouldBeNil)


			book  := entities.Book{
				Title:          "Silly Corn Valley",
				Description:    "The place where people write Corn for Operating Cereals",
				Permissions:    "",
				CreatorId:      user.ID,
			}
			err = bookRepo.Insert(&book)
			So(err, ShouldBeNil)


			chapter := entities.Chapter{
				Title:     "Have you ...",
				Body:      "EVERRR THE HEARD THE WOLFF CRYYYY",
				BookId:    book.ID,
				CreatorId: user.ID,
			}
			err = chapterRepo.Insert(&chapter)
			So(err, ShouldBeNil)


			Convey("Inserting some Story Saves should work", func (){
				rawBytes, err := json.Marshal(map[string]interface{} {
					"have_heard_the_wolf_cry_to_the_blue_corn_moon" : true,
				})
				So(err, ShouldBeNil)

				storySave := entities.StorySave{
					Token:     "HaveYouEverHeardTheWolfCryToTheBlueCornMoon",
					Save:      database.Jsonb{json.RawMessage(rawBytes)},
					CreatedBy: user.ID,
					BookId:    book.ID,
					ChapterId: chapter.ID,
				}
				err = storySaveRepo.Insert(&storySave)
				So(err, ShouldBeNil)

				secondStorySave := entities.StorySave{
					Token:     "OrAskedTheGrinningBobcatWhyHeGrinned",
					Save:      database.Jsonb{json.RawMessage(rawBytes)},
					CreatedBy: user.ID,
					BookId:    book.ID,
					ChapterId: chapter.ID,
				}
				err = storySaveRepo.Insert(&secondStorySave)
				So(err, ShouldBeNil)


				Convey("Finding the newly created Save using the token should work", func (){
					searchResult, err := storySaveRepo.FindByToken(storySave.Token)
					So(err, ShouldBeNil)
					So(searchResult.Save.ToString(), ShouldEqual, storySave.Save.ToString())
				})

				Convey("Updating the created Save should work", func (){
					rawBytes, err = json.Marshal(map[string]interface{} {
						"have_heard_the_wolf_cry_to_the_blue_corn_moon" : false,
					})
					So(err, ShouldBeNil)

					storySave.Save = database.Jsonb{ json.RawMessage(rawBytes) }
					err = storySaveRepo.Update(&storySave)
					So(err, ShouldBeNil)

					Convey("Finding the updated Save should yield a result with updated fields", func (){
						searchResult, err := storySaveRepo.FindByToken(storySave.Token)
						So(err,ShouldBeNil)
						So(searchResult.Save.ToString(), ShouldEqual, storySave.Save.ToString())
					})
				})


				Convey("Deleting the created Save should work", func (){
					err = storySaveRepo.Delete(&storySave)
					So(err, ShouldBeNil)

					Convey("Finding the deleted Save should yield no results", func (){
						_, err := storySaveRepo.FindByToken(storySave.Token)
						So(err,ShouldNotBeNil)
						So(err.EndpointMessage(), ShouldEqual, "No save found")
					})

					Convey("Finding the Save that wasn't deleted should still yield a result", func(){
						_, err := storySaveRepo.FindByToken(secondStorySave.Token)
						So(err, ShouldBeNil)
					})
				})

			})
		})


		Reset(func (){
			transaction.Rollback()
			So(transaction.Error,ShouldBeNil)
		})
	}))

}
