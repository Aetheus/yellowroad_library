package commands

import (
	"log"
	"yellowroad_library/utils/app_error"
)

func LogErrorAndExit(appErr app_error.AppError){
	log.Fatalf("Error:\n\t%s\nStacktrace:\n%s",appErr.Error(),appErr.Stacktrace())
}