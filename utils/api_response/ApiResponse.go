package api_response

import "yellowroad_library/utils/app_error"

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

func Success(data interface{}) ApiResponse{
	return ApiResponse {
		Status : SUCCESS,
		Data : data,
	}
}

//converts an AppError into either a Fail or Error response, depending on its HTTP code
func ConvertErr(err app_error.AppError) ApiResponse{
	isError := isErrorCode(err.HttpCode())

	if (isError) {
		return Error(err.EndpointMessage())
	} else {
		var dummyData struct{} //TODO add "data" field to AppError and use it here instead
		return Fail(err.EndpointMessage(), dummyData)
	}
}

func Fail(message string, data interface{}) ApiResponse{
	return ApiResponse {
		Status : FAIL,
		Data : data,
		Message : message,
	}
}

func Error(message string) ApiResponse{
	return ApiResponse {
		Status : ERROR,
		Message : message,
	}
}

//if it's an error code (HTTP 5xx range), returns true. Otherwise, false
func isErrorCode(code int) bool{
	if (code >= 500 && code < 600 ){
		return true
	} else {
		return false
	}
}