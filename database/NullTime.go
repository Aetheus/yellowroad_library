package database


import "github.com/lib/pq"

//just use lib/pq's NullTime since most of the code is lifted from there anyway
type NullTime pq.NullTime


//import "time"
//import "database/sql/driver"
//
//type NullTime struct {
//	Time  time.Time
//	Valid bool // Valid is true if Time is not NULL
//}
//
//func (nt *NullTime) Scan(value interface{}) error {
//
//	nt.Time, nt.Valid = value.(time.Time)
//	return nil
//}
//
//// Value implements the driver Valuer interface.
//func (nt NullTime) Value() (driver.Value, error) {
//	if !nt.Valid {
//		return nil, nil
//	}
//	return nt.Time, nil
//}