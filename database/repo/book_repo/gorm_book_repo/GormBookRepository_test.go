package gorm_book_repo

import "testing"
import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/jinzhu/gorm"
	"yellowroad_library/test_utils"
	"yellowroad_library/database/entities"
	book_repo "yellowroad_library/database/repo/book_repo"
)

func TestGormBookRepository(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given a book repository", t, test_utils.WithGormDBConnection(func(gormDB *gorm.DB){
		var transaction = gormDB.Begin()
		var bookRepo book_repo.BookRepository = New(transaction);

		Convey("Inserting a book entity should work", func (){
			newBook := entities.Book {
				CreatorId: 1,
				Title: "test title",
				Description: "test description",
			}
			err := bookRepo.Insert(&newBook)
			So(err, ShouldBeNil)

			Convey("Finding the inserted book by ID should work", func () {
				found_book, err := bookRepo.FindById(newBook.ID)
				So(err, ShouldBeNil)
				So(found_book.ID, ShouldEqual, newBook.ID)


				Convey("Updating the inserted book should work", func (){
					newTitle := "this is a new title"
					newDescription := "this is a new description"

					found_book.Description = newDescription
					found_book.Title = newTitle

					currentUpdatedAt := found_book.UpdatedAt

					err := bookRepo.Update(&found_book)
					So(err, ShouldBeNil)
					So(found_book.UpdatedAt, ShouldNotEqual, currentUpdatedAt)


					Convey("Deleting the inserted book should work", func (){
						err := bookRepo.Delete(&found_book)
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
