package gorm_user_repo

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"yellowroad_library/test_utils"
	"github.com/jinzhu/gorm"
	"yellowroad_library/database/entities"
)


func TestGormUserRepository(t *testing.T){
	Convey("Given a User Repository", t, test_utils.WithGormDBConnection(func(gormDB *gorm.DB){
		transaction := gormDB.Begin()
		userRepo := New(transaction)

		Convey("Inserting a new user should work", func (){
			seneca := entities.User {
				Username : "Seneca",
				Password : "defaultValue",
				Email : "seneca_the_younger@deadromans.com",
			}
			err := userRepo.Insert(&seneca)
			So(err, ShouldBeNil)

			Convey("Finding the new user by Username should work", func (){
				_, err  := userRepo.FindByUsername(seneca.Username)
				So(err, ShouldBeNil)
			})

			Convey("Finding the new user by ID should work", func (){
				found_user, err := userRepo.FindById(seneca.ID)
				So(err, ShouldBeNil)

				Convey("Updating the new user should work", func (){
					newEmail := "seneca_the_younger@deadromanphilosophers.com"
					found_user.Email = newEmail
					err := userRepo.Update(&found_user)
					So(err, ShouldBeNil)

					found_user_again, _ := userRepo.FindById(seneca.ID)
					So(found_user_again.Email,ShouldEqual, newEmail)

					Convey ("Deleting the new user should work", func (){
						err := userRepo.Delete(&seneca)
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