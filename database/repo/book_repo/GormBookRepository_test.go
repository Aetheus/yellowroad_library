package book_repo

import "testing"
import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/jinzhu/gorm"
	"yellowroad_library/test_utils"
	"yellowroad_library/database/entities"
	"yellowroad_library/database/repo/user_repo"
	"strconv"
)

func TestGormBookRepository(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given a GormBookRepository and UserRepository", t, test_utils.WithRealGormDBConnection(func(gormDB *gorm.DB){
		var transaction = gormDB.Begin()
		var bookRepo = NewDefault(transaction);
		var userRepo user_repo.UserRepository = user_repo.NewDefault(transaction)


		Convey("Given a valid user", func (){
			newUser := entities.User{
				Username :"absolutelyrandomtestuserguyhere",
				Password : "ah well, whatever",
				Email : "thisisjust@dummytextanyway.com",
			}
			err := userRepo.Insert(&newUser)
			So(err, ShouldBeNil)

			Convey("Inserting a book should work", func (){
				newBook := entities.Book {
					CreatorId: newUser.ID,
					Title: "test title",
					Description: "test description",
				}
				err := bookRepo.Insert(&newBook)
				So(err, ShouldBeNil)

				Convey("Finding the inserted book by ID should work", func () {
					found_book, err := bookRepo.FindById(newBook.ID)
					So(err, ShouldBeNil)
					So(found_book.ID, ShouldEqual, newBook.ID)
				})

				Convey("Updating the inserted book with valid fields should work", func (){
					newTitle := "this is a new title"
					newDescription := "this is a new description"

					newBook.Description = newDescription
					newBook.Title = newTitle

					currentUpdatedAt := newBook.UpdatedAt

					err := bookRepo.Update(&newBook)
					So(err, ShouldBeNil)
					So(newBook.UpdatedAt, ShouldNotEqual, currentUpdatedAt)
				})

				Convey("Deleting the inserted book should work", func (){
					currentDeletedAt := newBook.DeletedAt
					err := bookRepo.Delete(&newBook)
					So(err, ShouldBeNil)
					So(newBook.DeletedAt, ShouldNotEqual, currentDeletedAt)
				})
			})

			Convey("Given several (30) books", func(){
				for i := 0 ; i < 30; i++{
					newBook := entities.Book {
						CreatorId: newUser.ID,
						Title: "test title "+strconv.Itoa(i),
						Description: "test description "+strconv.Itoa(i),
					}
					err := bookRepo.Insert(&newBook)
					So(err, ShouldBeNil)
				}

				Convey("If perpage = 5, only 5 Books should be returned", func (){
					results,err := bookRepo.Paginate(SearchOptions{
						StartPage: 1,
						PerPage: 5,
					})
					So(err,ShouldBeNil)
					So(len(results), ShouldEqual, 5)
				})

				Convey("There should be at least 3 pages of results (when perpage = 10)", func (){
					for i := 0 ; i < 3; i++ {
						results,err := bookRepo.Paginate(SearchOptions{
							StartPage: 1,
							PerPage: 10,
						})
						So(err,ShouldBeNil)
						So(len(results), ShouldEqual, 10)
					}
				})
			})
		})


		Reset(func (){
			transaction.Rollback()
			So(transaction.Error,ShouldBeNil)
		})

	}))

}
