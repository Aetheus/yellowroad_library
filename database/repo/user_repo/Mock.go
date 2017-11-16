// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package user_repo

import (
	"sync"
	"yellowroad_library/database/entities"
	"yellowroad_library/utils/app_error"
)

var (
	lockUserRepositoryMockDelete         sync.RWMutex
	lockUserRepositoryMockFindById       sync.RWMutex
	lockUserRepositoryMockFindByUsername sync.RWMutex
	lockUserRepositoryMockInsert         sync.RWMutex
	lockUserRepositoryMockUpdate         sync.RWMutex
)

// UserRepositoryMock is a mock implementation of UserRepository.
//
//     func TestSomethingThatUsesUserRepository(t *testing.T) {
//
//         // make and configure a mocked UserRepository
//         mockedUserRepository := &UserRepositoryMock{
//             DeleteFunc: func(in1 *entities.User) app_error.AppError {
// 	               panic("TODO: mock out the Delete method")
//             },
//             FindByIdFunc: func(in1 int) (entities.User, app_error.AppError) {
// 	               panic("TODO: mock out the FindById method")
//             },
//             FindByUsernameFunc: func(in1 string) (entities.User, app_error.AppError) {
// 	               panic("TODO: mock out the FindByUsername method")
//             },
//             InsertFunc: func(in1 *entities.User) app_error.AppError {
// 	               panic("TODO: mock out the Insert method")
//             },
//             UpdateFunc: func(in1 *entities.User) app_error.AppError {
// 	               panic("TODO: mock out the Update method")
//             },
//         }
//
//         // TODO: use mockedUserRepository in code that requires UserRepository
//         //       and then make assertions.
//
//     }
type UserRepositoryMock struct {
	// DeleteFunc mocks the Delete method.
	DeleteFunc func(in1 *entities.User) app_error.AppError

	// FindByIdFunc mocks the FindById method.
	FindByIdFunc func(in1 int) (entities.User, app_error.AppError)

	// FindByUsernameFunc mocks the FindByUsername method.
	FindByUsernameFunc func(in1 string) (entities.User, app_error.AppError)

	// InsertFunc mocks the Insert method.
	InsertFunc func(in1 *entities.User) app_error.AppError

	// UpdateFunc mocks the Update method.
	UpdateFunc func(in1 *entities.User) app_error.AppError

	// calls tracks calls to the methods.
	calls struct {
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// In1 is the in1 argument value.
			In1 *entities.User
		}
		// FindById holds details about calls to the FindById method.
		FindById []struct {
			// In1 is the in1 argument value.
			In1 int
		}
		// FindByUsername holds details about calls to the FindByUsername method.
		FindByUsername []struct {
			// In1 is the in1 argument value.
			In1 string
		}
		// Insert holds details about calls to the Insert method.
		Insert []struct {
			// In1 is the in1 argument value.
			In1 *entities.User
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// In1 is the in1 argument value.
			In1 *entities.User
		}
	}
}

// Delete calls DeleteFunc.
func (mock *UserRepositoryMock) Delete(in1 *entities.User) app_error.AppError {
	if mock.DeleteFunc == nil {
		panic("moq: UserRepositoryMock.DeleteFunc is nil but UserRepository.Delete was just called")
	}
	callInfo := struct {
		In1 *entities.User
	}{
		In1: in1,
	}
	lockUserRepositoryMockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	lockUserRepositoryMockDelete.Unlock()
	return mock.DeleteFunc(in1)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedUserRepository.DeleteCalls())
func (mock *UserRepositoryMock) DeleteCalls() []struct {
	In1 *entities.User
} {
	var calls []struct {
		In1 *entities.User
	}
	lockUserRepositoryMockDelete.RLock()
	calls = mock.calls.Delete
	lockUserRepositoryMockDelete.RUnlock()
	return calls
}

// FindById calls FindByIdFunc.
func (mock *UserRepositoryMock) FindById(in1 int) (entities.User, app_error.AppError) {
	if mock.FindByIdFunc == nil {
		panic("moq: UserRepositoryMock.FindByIdFunc is nil but UserRepository.FindById was just called")
	}
	callInfo := struct {
		In1 int
	}{
		In1: in1,
	}
	lockUserRepositoryMockFindById.Lock()
	mock.calls.FindById = append(mock.calls.FindById, callInfo)
	lockUserRepositoryMockFindById.Unlock()
	return mock.FindByIdFunc(in1)
}

// FindByIdCalls gets all the calls that were made to FindById.
// Check the length with:
//     len(mockedUserRepository.FindByIdCalls())
func (mock *UserRepositoryMock) FindByIdCalls() []struct {
	In1 int
} {
	var calls []struct {
		In1 int
	}
	lockUserRepositoryMockFindById.RLock()
	calls = mock.calls.FindById
	lockUserRepositoryMockFindById.RUnlock()
	return calls
}

// FindByUsername calls FindByUsernameFunc.
func (mock *UserRepositoryMock) FindByUsername(in1 string) (entities.User, app_error.AppError) {
	if mock.FindByUsernameFunc == nil {
		panic("moq: UserRepositoryMock.FindByUsernameFunc is nil but UserRepository.FindByUsername was just called")
	}
	callInfo := struct {
		In1 string
	}{
		In1: in1,
	}
	lockUserRepositoryMockFindByUsername.Lock()
	mock.calls.FindByUsername = append(mock.calls.FindByUsername, callInfo)
	lockUserRepositoryMockFindByUsername.Unlock()
	return mock.FindByUsernameFunc(in1)
}

// FindByUsernameCalls gets all the calls that were made to FindByUsername.
// Check the length with:
//     len(mockedUserRepository.FindByUsernameCalls())
func (mock *UserRepositoryMock) FindByUsernameCalls() []struct {
	In1 string
} {
	var calls []struct {
		In1 string
	}
	lockUserRepositoryMockFindByUsername.RLock()
	calls = mock.calls.FindByUsername
	lockUserRepositoryMockFindByUsername.RUnlock()
	return calls
}

// Insert calls InsertFunc.
func (mock *UserRepositoryMock) Insert(in1 *entities.User) app_error.AppError {
	if mock.InsertFunc == nil {
		panic("moq: UserRepositoryMock.InsertFunc is nil but UserRepository.Insert was just called")
	}
	callInfo := struct {
		In1 *entities.User
	}{
		In1: in1,
	}
	lockUserRepositoryMockInsert.Lock()
	mock.calls.Insert = append(mock.calls.Insert, callInfo)
	lockUserRepositoryMockInsert.Unlock()
	return mock.InsertFunc(in1)
}

// InsertCalls gets all the calls that were made to Insert.
// Check the length with:
//     len(mockedUserRepository.InsertCalls())
func (mock *UserRepositoryMock) InsertCalls() []struct {
	In1 *entities.User
} {
	var calls []struct {
		In1 *entities.User
	}
	lockUserRepositoryMockInsert.RLock()
	calls = mock.calls.Insert
	lockUserRepositoryMockInsert.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *UserRepositoryMock) Update(in1 *entities.User) app_error.AppError {
	if mock.UpdateFunc == nil {
		panic("moq: UserRepositoryMock.UpdateFunc is nil but UserRepository.Update was just called")
	}
	callInfo := struct {
		In1 *entities.User
	}{
		In1: in1,
	}
	lockUserRepositoryMockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	lockUserRepositoryMockUpdate.Unlock()
	return mock.UpdateFunc(in1)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//     len(mockedUserRepository.UpdateCalls())
func (mock *UserRepositoryMock) UpdateCalls() []struct {
	In1 *entities.User
} {
	var calls []struct {
		In1 *entities.User
	}
	lockUserRepositoryMockUpdate.RLock()
	calls = mock.calls.Update
	lockUserRepositoryMockUpdate.RUnlock()
	return calls
}