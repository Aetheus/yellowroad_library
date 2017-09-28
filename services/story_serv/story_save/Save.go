package story_save

import "yellowroad_library/utils/app_error"

type Save struct {
	jsonString string
}
func New() Save{
	return Save {}
}

//TODO: implement this
func (this Save) Encode() (encodedSaveString string, err app_error.AppError) {
	return "{\"place\":\"holder\"}", nil
}

//TODO: implement this
func DecodeSaveString(encodedSaveString string) (Save, app_error.AppError){
	return Save {}, nil
}