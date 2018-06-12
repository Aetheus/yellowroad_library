package btag_repo

import "testing"
import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/jinzhu/gorm"
	"yellowroad_library/test_utils"
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/database/repo/user_repo"
	"yellowroad_library/database/entities"
	"yellowroad_library/database/repo/btagvote_repo"
	"strconv"
)

func TestGormBookTagRepository(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given a GormBookRepository,UserRepository, GormBookTagVoteRepository and GormBookTagRepository", t, test_utils.WithRealGormDBConnection(func(gormDB *gorm.DB){
		var transaction = gormDB.Begin()
		var bookRepo = book_repo.NewDefault(transaction);
		var userRepo user_repo.UserRepository = user_repo.NewDefault(transaction)
		var bookTagRepo = btagvote_repo.NewDefault(transaction)
		var bookTagCountRepo = NewDefault(transaction)

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

				Convey("Inserting 30 'revenge' tag upvotes and 15 'murder' tag upvotes into the book should work", func () {
					for i := 0 ; i < 30; i++ {
						//create dummy user for this
						temp_user := entities.User {
							Username: "temp_reader_" + strconv.Itoa(i),
							Password : "dummypassword",
							Email : "temp_reader_" + strconv.Itoa(i) + "@read.com",
						}
						err = userRepo.Insert(&temp_user)
						So(err, ShouldBeNil)

						//"upvote" the "revenge" tag
						revengeTag := entities.BookTagVote{
							BookId	: newBook.ID,
							Tag 	: "revenge",
							UserId  : temp_user.ID,
							Direction : 1,
						}
						err = bookTagRepo.Upsert(&revengeTag)
						So(err, ShouldBeNil)

						//"upvote" the "murder" tag
						if (i < 15) {
							murderTag := entities.BookTagVote{
								BookId	: newBook.ID,
								Tag 	: "murder",
								UserId  : temp_user.ID,
								Direction : 1,
							}
							err = bookTagRepo.Upsert(&murderTag)
							So(err, ShouldBeNil)
						}

					}


					Convey("There should be 30 'revenge' tag votes and 15 'murder' tag votes", func (){
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

						Convey("Syncing the 'revenge' and 'murder' tag votes should bring the tag vote count to the correct values", func (){
							result, err := bookTagCountRepo.SyncCount("revenge",newBook.ID)
							So(err, ShouldBeNil)
							So(result.Count, ShouldEqual, 30)

							result, err = bookTagCountRepo.SyncCount("murder",newBook.ID)
							So(err, ShouldBeNil)
							So(result.Count, ShouldEqual, 15)
						})

						Convey("Downvoting the 'revenge' tag twice and syncing should bring the tag vote score to 28", func(){
							for i := 0 ; i < 2; i++ {
								//create dummy user for this
								temp_user := entities.User {
									Username: "temp_downvoter_" + strconv.Itoa(i),
									Password : "dummypassword",
									Email : "temp_downvoter_" + strconv.Itoa(i) + "@downvoter.com",
								}
								err = userRepo.Insert(&temp_user)
								So(err, ShouldBeNil)

								//"downvote" the "revenge" tag
								revengeTag := entities.BookTagVote{
									BookId	: newBook.ID,
									Tag 	: "revenge",
									UserId  : temp_user.ID,
									Direction : -1,
								}
								err = bookTagRepo.Upsert(&revengeTag)
								So(err, ShouldBeNil)
							}

							result, err := bookTagCountRepo.SyncCount("revenge",newBook.ID)
							So(err, ShouldBeNil)
							So(result.Count, ShouldEqual, 28)
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
