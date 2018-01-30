package gin_tools

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/utils/app_error"
	"net/http"
	"strconv"
	"encoding/json"
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

func GetJsonParam(key string,c *gin.Context) json.RawMessage {
	value := c.Param(key)
	return json.RawMessage(value)
}

func GetJsonParamOrDefault(key string, defaultValAsString string, c *gin.Context) json.RawMessage {
	value := c.Param(key)
	if (value == "") {
		value = defaultValAsString
	}
	return json.RawMessage(value)
}


func JSON(formPointer interface{}, c *gin.Context) (app_error.AppError){
	bindErr := c.BindJSON(formPointer)
	if (bindErr != nil) {
		return app_error.Wrap(bindErr).SetHttpCode(http.StatusUnprocessableEntity)
	}

	return nil
}