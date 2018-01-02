package book_serv

import "testing"
import (
	. "github.com/smartystreets/goconvey/convey"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/utils/app_error"
)

func prepopulateMockData() uow.UnitOfWork {

	return &uow.UnitOfWorkMock{
		AutoCommitFunc: func(in1 []uow.WorkFragment, in2 func() app_error.AppError) app_error.AppError {
			return nil
		},
		CommitFunc: func() app_error.AppError {
			return nil
		},
		RollbackFunc: func() app_error.AppError {
			return nil
		},

	}
}

func TestDefaultBookService(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given a Unit of Work and prepopulated stored data", t, func(){
		work := prepopulateMockData()

		//prepopulated data




		Reset(func (){
			err := work.Rollback()
			So(err,ShouldBeNil)
		})

	})
}