package gin_tools

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/utils/app_error"
	"net/http"
	"strconv"
	"encoding/json"
	"fmt"
)

func GetIntParam(key string,c *gin.Context) (int, app_error.AppError){
	if value, convErr := strconv.Atoi(c.Param(key)); convErr != nil {
		appErr := app_error.New(http.StatusUnprocessableEntity,
			"", key + " must be a valid integer value!")
		return 0, appErr
	} else {
		return value, nil
	}
}

func GetIntParamOrDefault(key string,defaultVal int, c *gin.Context) int{
	value, err := GetIntParam(key, c)
	if (err != nil){
		return defaultVal
	} else {
		return value
	}
}

func GetIntQuery(key string, c *gin.Context) (int, app_error.AppError) {
	if value, convErr := strconv.Atoi(c.Query(key)); convErr != nil {
		appErr := app_error.New(http.StatusUnprocessableEntity,
			"", key + " must be a valid integer value!")
		return 0, appErr
	} else {
		return value, nil
	}
}

func GetIntQueryOrDefault(key string, defaultVal int, c *gin.Context) int {
	convertedVal, err := GetIntQuery(key, c)
	if (err != nil) {
		return defaultVal
	}
	return convertedVal
}

func GetBoolQueryOrDefault(key string, defaultVal bool, c *gin.Context) bool{
	convertedVal, err := strconv.ParseBool(c.Query(key))
	if (err != nil) {
		return defaultVal
	}
	return convertedVal
}

func BindJSON(formPointer interface{}, c *gin.Context) (app_error.AppError){
	data, err := c.GetRawData()
	if err != nil {
		return app_error.Wrap(err)
	}

	err = json.Unmarshal(data, formPointer)
	errMessage := ""
	errorCode := http.StatusBadRequest;

	if jsonError, ok := err.(*json.UnmarshalTypeError); ok {
		errMessage = fmt.Sprintf("[%v]: Expected type '%v', got a '%v' instead",jsonError.Field, jsonError.Type.Name(), jsonError.Value,  )
		errorCode = http.StatusUnprocessableEntity
	} else {
		errMessage = "Cannot parse JSON due to an unexpected syntax error"
		errorCode = http.StatusBadRequest
	}

	if (err != nil) {
		return app_error.New(errorCode, err.Error(), errMessage)
	}

	return nil
}