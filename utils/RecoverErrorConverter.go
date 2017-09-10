package utils

import "errors"

func RecoverErrorConverter(r interface{}) error {
	var err error

	switch x := r.(type) {
	case string:
		err = errors.New(x)
	case error:
		err = x
	default:
		err = errors.New("Unknown error has occured")
	}

	return err
}
