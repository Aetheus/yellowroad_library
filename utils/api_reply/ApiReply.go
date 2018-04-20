package api_reply

import (
	"yellowroad_library/utils/app_error"
	"net/http"
	"github.com/gin-gonic/gin"
)

var SUCCESS = "success"
var FAIL = "fail"
var ERROR = "error"

/* ApiResponse:
*	Follows the JSend specification for the most part, except we also allow "fail" responses to have messages
*/
type ApiResponse struct  {
	Data interface{} `json:"data,omitempty"`
	Status string `json:"status"`
	Message string `json:"message,omitempty"`
}

func SuccessResponse(data interface{}) ApiResponse{
	return ApiResponse {
		Status : SUCCESS,
		Data : data,
	}
}

//converts an AppError into either a FailResponse or ErrorResponse response, depending on its HTTP code
func ConvertErr(err app_error.AppError) ApiResponse{
	isServerError := isServerErrorCode(err.HttpCode())

	if (isServerError) {
		return ErrorResponse(err.EndpointMessage())
	} else {
		var dummyData struct{} //TODO add "data" field to AppError and use it here instead
		return FailResponse(err.EndpointMessage(), dummyData)
	}
}

func FailResponse(message string, data interface{}) ApiResponse{
	return ApiResponse {
		Status : FAIL,
		Data : data,
		Message : message,
	}
}

func ErrorResponse(message string) ApiResponse{
	return ApiResponse {
		Status : ERROR,
		Message : message,
	}
}

//if it's a server error code (HTTP 5xx range), returns true. Otherwise, false for others like client error codes (4xx)
func isServerErrorCode(code int) bool{
	if (code >= 500 && code < 600 ){
		return true
	} else {
		return false
	}
}

// Convenience methods for sending Gin responses
func Success(c *gin.Context, payload interface{}){
	c.JSON(http.StatusOK, SuccessResponse(payload))
}
func Failure(c *gin.Context, err app_error.AppError) {
	c.JSON(err.HttpCode(),ConvertErr(err))
}