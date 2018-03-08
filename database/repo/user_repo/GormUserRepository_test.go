package user_repo

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"yellowroad_library/test_utils"
	"github.com/jinzhu/gorm"
	"yellowroad_library/database/entities"
)

func TestGormUserRepository(t *testing.T){
	Convey("Given a User Repository", t, test_utils.WithRealGormDBConnection(func(gormDB *gorm.DB){
		transaction := gormDB.Begin()
		userRepo := NewDefault(transaction)

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
				foundUser, err := userRepo.FindById(seneca.ID)
				So(err, ShouldBeNil)
				So(foundUser.ID, ShouldEqual, seneca.ID)
			})

			Convey("Updating the new user with valid fields should work", func (){
				currentUpdatedAt := seneca.UpdatedAt
				newEmail := "seneca_the_younger@deadromanphilosophers.com"
				seneca.Email = newEmail
				err := userRepo.Update(&seneca)
				So(err, ShouldBeNil)

				found_user_again, _ := userRepo.FindById(seneca.ID)
				So(found_user_again.Email,ShouldEqual, newEmail)
				So(found_user_again.UpdatedAt, ShouldNotEqual, currentUpdatedAt)


			})

			Convey ("Deleting the new user should work", func (){
				currentDeletedAt := seneca.DeletedAt
				err := userRepo.Delete(&seneca)
				So(err,ShouldBeNil)
				So(seneca.DeletedAt, ShouldNotEqual, currentDeletedAt)
			})

		})

		Reset(func (){
			transaction.Rollback()
			So(transaction.Error,ShouldBeNil)
		})

	}))


}