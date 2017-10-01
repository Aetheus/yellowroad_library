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
	return j.MarshalJSON()
}

func (j *Jsonb) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	return json.Unmarshal(bytes, j)
}