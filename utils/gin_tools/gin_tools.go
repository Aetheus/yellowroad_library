package gin_tools

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/utils/app_error"
	"net/http"
	"strconv"
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