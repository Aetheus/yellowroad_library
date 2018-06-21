package game_domain

import (
	"yellowroad_library/utils/app_error"
	"strings"
	"fmt"
	"net/http"
	"github.com/xeipuuv/gojsonschema"
	jmutate "github.com/Aetheus/jmutate_go"
	"encoding/json"
)

//TODO : move
func isValidSave(saveJsonString string) bool {
	var jsonMessage json.RawMessage
	return json.Unmarshal([]byte(saveJsonString), &jsonMessage) == nil
}

func validateSaveRequirements(saveDataAsJsonString string,requirementsAsJsonString string) (err app_error.AppError) {
	if strings.Trim(requirementsAsJsonString, " ") == ""{
		return nil
	}

	schemaLoader := gojsonschema.NewStringLoader(requirementsAsJsonString)
	documentLoader := gojsonschema.NewStringLoader(saveDataAsJsonString)

	result, validationErr := gojsonschema.Validate(schemaLoader,documentLoader)
	if (validationErr != nil ){
		err = app_error.Wrap(validationErr)
		return
	}

	if !result.Valid(){
		reason := ""
		for _, desc := range result.Errors(){
			reason = reason + fmt.Sprintf("- %s\n", desc)
		}
		reason = strings.TrimRight(reason,"\n")
		err = app_error.New(http.StatusUnprocessableEntity,"",reason)
	}

	return
}

func applyEffectOnSave(saveDataAsJsonString string, effectAsJsonString string)(
	newSaveDataAsJsonString string, err app_error.AppError,
) {
	mutation, jMutateErr := jmutate.New([]byte(effectAsJsonString))
	if jMutateErr != nil {
		err = app_error.Wrap(err)
		return
	}

	newSaveJson, jMutateErr := mutation.Apply([]byte(saveDataAsJsonString))
	if jMutateErr != nil {
		err = app_error.Wrap(err)
		return
	}

	newSaveDataAsJsonString = string(newSaveJson)
	return
}