package chapterpath_repo

import (
	"yellowroad_library/test_utils"
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/database/repo/user_repo"
	"yellowroad_library/database/entities"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/jinzhu/gorm"
	"yellowroad_library/database/repo/chapter_repo"
	"yellowroad_library/database"
	"encoding/json"
	"fmt"
)

func TestGormChapterPathRepository (t *testing.T) {

	Convey("Given a GormChapterPathRepository, ChapterRepository, BookRepository and UserRepository", t, test_utils.WithRealGormDBConnection(func(gormDB *gorm.DB){
		transaction := gormDB.Begin()

		var chapterPathRepo = NewDefault(transaction)
		var chapterRepo chapter_repo.ChapterRepository = chapter_repo.NewDefault(transaction)
		var bookRepo book_repo.BookRepository = book_repo.NewDefault(transaction)
		var userRepo user_repo.UserRepository = user_repo.NewDefault(transaction)

		Convey("Given a valid user, book and chapters", func (){
			//create a user and book to test this
			newUser := entities.User{
				Username : "Robert Baratheon",
				Password : "onanopenfield",
				Email: "bobbyb@rightfulkingofwesteros.com",
			}
			userRepoErr := userRepo.Insert(&newUser)
			So(userRepoErr, ShouldBeNil)

			newBook := entities.Book{
				Title : "Game of Thrones, Season XIIVIIXII",
				Description : "From the best selling author of Game of Thrones, Season XIIVIIXI",
				CreatorId : newUser.ID,
			}
			bookRepoErr := bookRepo.Insert(&newBook)
			So(bookRepoErr, ShouldBeNil)

			firstChapter := entities.Chapter{
				Title : "Fish N Chips",
				Body : "I'm running out of ideas for flavour text",
				BookId : newBook.ID,
				CreatorId : newUser.ID,
			}
			err := chapterRepo.Insert(&firstChapter)
			So(err, ShouldBeNil)

			secondChapter := entities.Chapter{
				Title : "Bacon and Eggs",
				Body : "Food is an original idea, right? Right?",
				BookId : newBook.ID,
				CreatorId : newUser.ID,
			}
			err = chapterRepo.Insert(&secondChapter)
			So(err, ShouldBeNil)


			Convey("Inserting a new Chapter Path with all the right fields should work", func (){
				newChapterPath := entities.ChapterPath{
					FromChapterId:firstChapter.ID,
					ToChapterId:secondChapter.ID,
					Effects : database.Jsonb {json.RawMessage(`{"morale": "+1"}`),},
					Requirements : database.Jsonb {json.RawMessage(`{}`),},
				}

				err = chapterPathRepo.Insert(&newChapterPath)
				So(err, ShouldBeNil)


				Convey("Finding the new Chapter Path by its own ID should work", func (){
					foundChapterPath, err := chapterPathRepo.FindById(newChapterPath.ID)

					fmt.Println("Found chapter was:")
					res, _ := json.Marshal(foundChapterPath)
					fmt.Println(string(res))

					So(err, ShouldBeNil)
					So(foundChapterPath.ID, ShouldEqual, newChapterPath.ID)
				})

				Convey("Updating the new Chapter Path with valid parameters should work", func (){
					newEffect := json.RawMessage(`{"morale": "+5"}`)
					newRequirements := json.RawMessage(`{"hunger": "3"}`)
					oldUpdatedAt := newChapterPath.UpdatedAt

					newChapterPath.Effects = database.Jsonb {newEffect}
					newChapterPath.Requirements = database.Jsonb{newRequirements}

					chapterPathRepo.Update(&newChapterPath)

					So(newChapterPath.UpdatedAt, ShouldNotEqual, oldUpdatedAt)
				})

				Convey("Deleting the new Chapter Path should work", func (){
					oldDeletedAt := newChapterPath.DeletedAt
					err = chapterPathRepo.Delete(&newChapterPath)

					So(err, ShouldBeNil)
					So(oldDeletedAt, ShouldNotEqual, newChapterPath.DeletedAt)
				})



			})


		})


		Reset(func (){
			transaction.Rollback()
		})

	}))

}