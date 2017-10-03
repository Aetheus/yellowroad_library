package gorm_book_repo

import "testing"
import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/jinzhu/gorm"
	"yellowroad_library/database/repo"
	"yellowroad_library/database/entities"
	"yellowroad_library/config"
)

//this is relative to the test script location, not to the main package
const pathToConfigFile = "../../config_for_mocks.json"
var configFile = config.Load(pathToConfigFile)

func TestGormBookRepository(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given a book repository", t, repo.WithGormDBConnection(configFile,func(gormDB *gorm.DB){
		var book_repo GormBookRepository = New(gormDB)




		Convey("Inserting a book entity should work", func (){
			newBook := entities.Book {
				CreatorId: 1,
				Title: "test title",
				Description: "test description",
			}
			err := book_repo.Insert(&newBook)
			So(err, ShouldBeNil)

			Convey("Finding the inserted book by ID should work", func () {
				found_book, err := book_repo.FindById(newBook.ID)
				So(err, ShouldBeNil)
				So(found_book.ID, ShouldEqual, newBook.ID)


				Convey("Updating the inserted book should work", func (){
					newTitle := "this is a new title"
					newDescription := "this is a new description"

					found_book.Description = newDescription
					found_book.Title = newTitle

					currentUpdatedAt := found_book.UpdatedAt

					err := book_repo.Update(&found_book)
					So(err, ShouldBeNil)
					So(found_book.UpdatedAt, ShouldNotEqual, currentUpdatedAt)


					Convey("Deleting the inserted book should work", func (){
						err := book_repo.Delete(&found_book)
						So(err, ShouldBeNil)
					})

				})

			})
		})


		Reset(func (){
			//
		})

	}))

}

//func() {
//
//
//	Convey("When the Book Repository is created", func() {
//
//
//
//
//	})
//
//}