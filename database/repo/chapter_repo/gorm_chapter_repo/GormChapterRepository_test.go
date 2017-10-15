package gorm_chapter_repo

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"yellowroad_library/test_utils"
	"github.com/jinzhu/gorm"
	"yellowroad_library/database/entities"
	"yellowroad_library/database/repo/book_repo/gorm_book_repo"
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/database/repo/user_repo"
	"yellowroad_library/database/repo/user_repo/gorm_user_repo"
)


func TestGormChapterRepository (t *testing.T) {

	Convey("Given a GormChapterRepository, BookRepository and UserRepository", t, test_utils.WithGormDBConnection(func(gormDB *gorm.DB){
		transaction := gormDB.Begin()

		var chapterRepo = New(transaction)
		var bookRepo book_repo.BookRepository = gorm_book_repo.New(transaction)
		var userRepo user_repo.UserRepository = gorm_user_repo.New(transaction)

		Convey("Given a valid book", func (){
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


			Convey("Inserting a new chapter with all the right fields should work", func (){
				newChapter := entities.Chapter{
					Title : "Fish N Chips",
					Body : "I'm running out of ideas for flavour text",
					BookId : newBook.ID,
					CreatorId : newUser.ID,
				}

				err := chapterRepo.Insert(&newChapter)
				So(err, ShouldBeNil)


				Convey("Finding the new chapter by ID should work", func () {
					foundChapter, err := chapterRepo.FindById(newChapter.ID)
					So(err, ShouldBeNil)
					So(foundChapter.ID, ShouldEqual, newChapter.ID)
				})

				Convey("Finding the new chapter inside of the book by ID should work", func (){
					foundChapter ,err := chapterRepo.FindWithinBook(newChapter.ID, newBook.ID)
					So(err, ShouldBeNil)
					So(foundChapter.ID, ShouldEqual, newChapter.ID)
				})

				Convey("Updating the new chapter with valid fields should work", func () {
					oldUpdatedAt := newChapter.UpdatedAt
					newTitle := "New Beginnings"
					newChapter.Title = newTitle
					err := chapterRepo.Update(&newChapter)
					So(err, ShouldBeNil)
					So(newChapter.UpdatedAt, ShouldNotEqual, oldUpdatedAt)

				})

				Convey("Updating the new chapter with invalid fields should return an error", func () {
					newChapter.BookId = -1
					newChapter.CreatorId = -1

					err := chapterRepo.Update(&newChapter)
					So(err, ShouldNotBeNil)
				})

				Convey("Deleting the new chapter should work", func (){
					oldDeletedAt := newChapter.DeletedAt
					chapterId := newChapter.ID
					err := chapterRepo.Delete(&newChapter)
					So(err, ShouldBeNil)
					So(newChapter.DeletedAt, ShouldNotEqual, oldDeletedAt)

					newChapter, err = chapterRepo.FindById(chapterId)
					So(err, ShouldNotBeNil)
				})
			})

			Convey("Inserting a new chapter with an incorrect book id should cause an error", func (){
				invalidChapter := entities.Chapter {
					Title : "this title",
					Body : "this body",
					BookId : -1,
					CreatorId : newUser.ID,
				}
				err := chapterRepo.Insert(&invalidChapter)
				So(err,ShouldNotBeNil)
			})

			Convey("Inserting a new chapter with an incorrect creator id should cause an error", func (){
				invalidChapter := entities.Chapter {
					Title : "this title",
					Body : "this body",
					BookId : newBook.ID,
					CreatorId : -1,
				}
				err := chapterRepo.Insert(&invalidChapter)
				So(err,ShouldNotBeNil)
			})

		})


		Reset(func (){
			transaction.Rollback()
		})

	}))

}