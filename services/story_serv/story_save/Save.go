package story_save

import "yellowroad_library/utils/app_error"
import (
	"encoding/base64"
	"net/http"
	"bytes"
	"compress/zlib"
)

type Save struct {
	JsonString string
}
func New() Save{
	return Save {}
}

func (this Save) EncodedSaveString() (encodedSaveString string, err app_error.AppError) {
	compressedSave, compressErr := zlibToBytes(this.JsonString)
	if (compressErr != nil){
		err = app_error.Wrap(compressErr)
		return encodedSaveString, err
	}

	base64EncodedCompressedSave := base64.StdEncoding.EncodeToString(compressedSave)
	return base64EncodedCompressedSave, err
}

func DecodeSaveString(encodedSaveString string) (Save, app_error.AppError){
	//encodedSave is a base64 encoded, zlib string of a JSON object
	//1. A JSON object is stringified, then compressed to the zlib (not gzip!) format using zlib.
	//2. The zlib compressed bytes are then converted to a string.
	//3. The string is then base64 encoded to deal with illegal characters in the "string"

	//to reverse the process to get the raw JSON string, we have to
	//1. base64 decode the save string to a byte array.
	//2. unzlib the byte array.
	//3. convert the byte representation of the uncompresed string to an actual string object.
	var decodedSave Save
	invalidSaveError := app_error.New(http.StatusUnprocessableEntity,"","Invalid save string provided!")

	gzippedBytes, err := base64.StdEncoding.DecodeString(encodedSaveString)
	if (err != nil) {
		return decodedSave, invalidSaveError
	}

	saveString, err := unZlibToString(gzippedBytes)
	if (err != nil){
		return decodedSave, invalidSaveError
	}

	decodedSave.JsonString = saveString
	return decodedSave, nil
}

func zlibToBytes(inputValue string) (returnValue []byte, err error){
	var resultBuffer bytes.Buffer
	writer := zlib.NewWriter(&resultBuffer)

	_, err = writer.Write([]byte(inputValue))
	writer.Close()
	if (err != nil) {
		return returnValue, err
	}

	returnValue = resultBuffer.Bytes();

	return returnValue, nil
}


func unZlibToString(compressedData []byte) (string, error) {
	var result string
	byteReader := bytes.NewReader(compressedData)

	gzipReader, err := zlib.NewReader(byteReader)
	if err != nil {
		return result, err
	}

	var uncompressedBuffer bytes.Buffer
	_ , err = uncompressedBuffer.ReadFrom(gzipReader)
	if err != nil {
		return result, err
	}

	result = string(uncompressedBuffer.Bytes())
	return result, nil
}
