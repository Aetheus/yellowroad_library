package database

import (
	"database/sql/driver"
	"fmt"
	"encoding/json"
	"errors"
)

type Jsonb struct {
	json.RawMessage
}

func (j Jsonb) Value() (driver.Value, error) {
	val, err:= j.MarshalJSON()
	stringVal := string(val)

	return stringVal, err
	//return j.MarshalJSON()
}

func (j Jsonb) ToString() string {
	return string(j.RawMessage)
}

func (j *Jsonb) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	return json.Unmarshal(bytes, j)
}