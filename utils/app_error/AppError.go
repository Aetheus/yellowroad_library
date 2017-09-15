package app_error

import (
	"github.com/ansel1/merry"
	"net/http"
)

func New (httpErrorCode int, contextMessage string, endpointMessage string) AppError {
	enhancedError := merry.New(contextMessage).WithUserMessage(endpointMessage).WithHTTPCode(httpErrorCode)
	return &appError {
		enhancedError : enhancedError,
	}
}

func unhandled(err error) AppError {
	enhancedError := merry.Wrap(err).
		WithHTTPCode(http.StatusInternalServerError).
		WithUserMessage("An unhandled error occurred")

	return &appError {
		enhancedError : enhancedError,
	}
}

//converts regular errors to appError if they aren't already an instance of it
//if it is, then it just returns
// by default, all non merry.Error/appError errors will be converted to appErrors with 500 HTTP code
func Wrap (err error) AppError {
	if appErr, ok := err.(AppError); ok {
		return appErr
	} else if merryErr, ok := err.(merry.Error); ok {
		//if this is already a merry.Error instance, just wrap it and return
		return &appError {
			enhancedError : merryErr,
		}
	} else {
		return unhandled(err)
	}
}


/* appError:
	A "facade" class - it wraps around any "enhanced" implementation of errors in order to provide us nice
	things like stack traces and the like. Why wrap those classes? Because we can then swap them out easily
	if we need to.
*/
type AppError interface{
	Error() string
	Stacktrace() string

	EndpointMessage() string
	SetEndpointMessage(message string) AppError

	HttpCode() int
	SetHttpCode(code int) AppError
}



type appError struct{
	enhancedError merry.Error
}

//This makes it fulfill the Error interface; this message should NOT be returned to the endpoint users!
func (this appError) Error() string {
	return this.enhancedError.Error()
}
func (this appError) Stacktrace () string {
	return merry.Stacktrace(this.enhancedError)
}

func (this appError) EndpointMessage() string {
	return merry.UserMessage(this.enhancedError)
}
func (this *appError) SetEndpointMessage(message string) AppError {
	this.enhancedError = this.enhancedError.WithUserMessage(message)
	return this
}

func (this *appError) SetHttpCode(code int) AppError {
	this.enhancedError = this.enhancedError.WithHTTPCode(code)
	return this
}
func (this appError) HttpCode() int {
	return merry.HTTPCode(this.enhancedError)
}


