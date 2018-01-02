package gorm_btagvote_repo

import "testing"
import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/jinzhu/gorm"
	"yellowroad_library/test_utils"
	"yellowroad_library/database/entities"
	"yellowroad_library/database/repo/user_repo"
	"yellowroad_library/database/repo/user_repo/gorm_user_repo"
	"yellowroad_library/database/repo/book_repo/gorm_book_repo"
)

func TestGormBookTagRepository(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given a GormBookRepository,UserRepository and GormBookTagRepository", t, test_utils.WithRealGormDBConnection(func(gormDB *gorm.DB){
		var transaction = gormDB.Begin()
		var bookRepo = gorm_book_repo.New(transaction);
		var userRepo user_repo.UserRepository = gorm_user_repo.New(transaction)
		var bookTagRepo = New(transaction)

		Convey("Given several valid users", func (){
			authorUser := entities.User{
				Username :"absolutelyrandomtestuserguyhere",
				Password : "ah well, whatever",
				Email : "thisisjust@dummytextanyway.com",
			}
			err := userRepo.Insert(&authorUser)
			So(err, ShouldBeNil)

			reader_1_user := entities.User {
				Username: "reader_1",
				Password : "dummypassword",
				Email : "reader_1@read.com",
			}
			err = userRepo.Insert(&reader_1_user)
			So(err, ShouldBeNil)

			Convey("Inserting a book should work", func (){
				newBook := entities.Book {
					CreatorId:   authorUser.ID,
					Title:       "test title",
					Description: "test description",
				}
				err := bookRepo.Insert(&newBook)
				So(err, ShouldBeNil)

				Convey("Inserting a book tag should work", func (){
					firstTag := entities.BookTagVote{
						BookId	: newBook.ID,
						Tag 	: "plot_twist",
						UserId  : authorUser.ID,
					}
					err = bookTagRepo.Insert(&firstTag)
					So(err, ShouldBeNil)

					Convey("Finding the newly inserted tag should work", func (){
						results, err := bookTagRepo.FindByFields(entities.BookTagVote{
							BookId:newBook.ID,
							Tag: "plot_twist",
						})
						So(err, ShouldBeNil)
						So(len(results), ShouldEqual, 1)

						foundTag := results[0]
						So(foundTag.BookId, ShouldEqual, firstTag.BookId)
						So(foundTag.Tag, ShouldEqual, firstTag.Tag)
						So(foundTag.UserId, ShouldEqual, firstTag.UserId)
					})

					Convey("Inserting the same tag but with a different user should work", func (){
						duplicateTag := entities.BookTagVote{
							BookId	: newBook.ID,
							Tag 	: "plot_twist",
							UserId  : reader_1_user.ID,
						}
						err = bookTagRepo.Insert(&duplicateTag)
						So(err, ShouldBeNil)

						Convey("Finding the tag should give 2 results", func (){
							results, err := bookTagRepo.FindByFields(entities.BookTagVote{
								BookId:newBook.ID,
								Tag: "plot_twist",
							})
							So(err, ShouldBeNil)
							So(len(results), ShouldEqual, 2)

							firstFoundTag := results[0]
							So(firstFoundTag.BookId, ShouldEqual, firstTag.BookId)
							So(firstFoundTag.Tag, ShouldEqual, firstTag.Tag)
							So(firstFoundTag.UserId, ShouldEqual, firstTag.UserId)

							secondFoundTag := results[1]
							So(secondFoundTag.Tag, ShouldEqual, duplicateTag.Tag)
							So(secondFoundTag.UserId, ShouldEqual, duplicateTag.UserId)
							So(secondFoundTag.BookId, ShouldEqual, duplicateTag.BookId)


							Convey("Deleting the tags using the tag text and book id as search options should delete both tags", func (){
								err := bookTagRepo.DeleteByFields(entities.BookTagVote{
									Tag: "plot_twist",
									BookId: newBook.ID,
								})
								So(err,ShouldBeNil)

								results, err := bookTagRepo.FindByFields(entities.BookTagVote{
									BookId:newBook.ID,
									Tag: "plot_twist",
								})
								So(err, ShouldBeNil)
								So(len(results), ShouldEqual, 0)
							})
						})


					})

					Convey("Inserting the same tag again with the same user should return an error", func (){
						err = bookTagRepo.Insert(&entities.BookTagVote{
							BookId	: newBook.ID,
							Tag 	: "plot_twist",
							UserId  : authorUser.ID,
						})
						So(err, ShouldNotBeNil)
					})

					Convey("Deleting the tag should work", func (){
						err = bookTagRepo.Delete(&firstTag)
						So(err,ShouldBeNil)
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
