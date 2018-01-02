package gorm_btagvotecount_repo

import "testing"
import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/jinzhu/gorm"
	"yellowroad_library/test_utils"
	"yellowroad_library/database/repo/book_repo/gorm_book_repo"
	"yellowroad_library/database/repo/user_repo"
	"yellowroad_library/database/repo/user_repo/gorm_user_repo"
	"yellowroad_library/database/entities"
	"yellowroad_library/database/repo/btagvote_repo/gorm_btagvote_repo"
	"strconv"
)

func TestGormBookTagCountRepository(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given a GormBookRepository,UserRepository, GormBookTagVoteRepository and GormBookTagVoteCountRepository", t, test_utils.WithRealGormDBConnection(func(gormDB *gorm.DB){
		var transaction = gormDB.Begin()
		var bookRepo = gorm_book_repo.New(transaction);
		var userRepo user_repo.UserRepository = gorm_user_repo.New(transaction)
		var bookTagRepo = gorm_btagvote_repo.New(transaction)
		var bookTagCountRepo = New(transaction)

		Convey("Given a valid user", func (){
			authorUser := entities.User{
				Username :"absolutelyrandomtestuserguyhere",
				Password : "ah well, whatever",
				Email : "thisisjust@dummytextanyway.com",
			}
			err := userRepo.Insert(&authorUser)
			So(err, ShouldBeNil)

			Convey("Inserting a book should work", func (){
				newBook := entities.Book {
					CreatorId:   authorUser.ID,
					Title:       "test title",
					Description: "test description",
				}
				err := bookRepo.Insert(&newBook)
				So(err, ShouldBeNil)

				Convey("Inserting 30 'revenge' tags and 15 'murder' tags into the book should work", func () {
					for i := 0 ; i < 30; i++ {
						//create dummy user for this
						temp_user := entities.User {
							Username: "temp_reader_" + strconv.Itoa(i),
							Password : "dummypassword",
							Email : "temp_reader_" + strconv.Itoa(i) + "@read.com",
						}
						err = userRepo.Insert(&temp_user)
						So(err, ShouldBeNil)

						//insert "revenge" tag
						revengeTag := entities.BookTagVote{
							BookId	: newBook.ID,
							Tag 	: "revenge",
							UserId  : temp_user.ID,
						}
						err = bookTagRepo.Insert(&revengeTag)
						So(err, ShouldBeNil)

						//insert "murder" tag
						if (i < 15) {
							murderTag := entities.BookTagVote{
								BookId	: newBook.ID,
								Tag 	: "murder",
								UserId  : temp_user.ID,
							}
							err = bookTagRepo.Insert(&murderTag)
							So(err, ShouldBeNil)
						}

					}


					Convey("There should be 30 'revenge' tags and 15 'murder' tags", func (){
						revengeResults, err := bookTagRepo.FindByFields(entities.BookTagVote{
							BookId:newBook.ID,
							Tag: "revenge",
						})
						So(err, ShouldBeNil)
						So(len(revengeResults), ShouldEqual, 30)

						murderResults, err := bookTagRepo.FindByFields(entities.BookTagVote{
							BookId:newBook.ID,
							Tag: "murder",
						})
						So(err, ShouldBeNil)
						So(len(murderResults), ShouldEqual, 15)

						Convey("Syncing the 'revenge' tags and 'murder' tags should bring the count to the correct values", func (){
							result, err := bookTagCountRepo.SyncCount("revenge",newBook.ID)
							So(err, ShouldBeNil)
							So(result.Count, ShouldEqual, 30)

							result, err = bookTagCountRepo.SyncCount("murder",newBook.ID)
							So(err, ShouldBeNil)
							So(result.Count, ShouldEqual, 15)
						})
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
