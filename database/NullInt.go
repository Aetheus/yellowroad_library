package database

import "database/sql/driver"

//almost identical to SQL's NullInt64, but using Int instead of Int64
type NullInt struct {
	Int int
	Valid bool
}

func (n *NullInt) Scan(value interface{}) error {
	if value == nil {
		n.Int = 0
		n.Valid = false
		return nil
	}

	n.Valid = true
	return nil
}

func (n NullInt) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int, nil
}