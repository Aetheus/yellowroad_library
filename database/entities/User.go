package entities

import (
	"time"
	"yellowroad_library/utils"
)

type User struct {
	Username string
	Password string `json:"-"` //this excludes it from being accidentally serialized into a JSON and shown to users
	Email    string

	//housekeeping attributes
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt utils.NullTime
}


//fields that we allow to edit in our handlers (e.g: for the "update" routes of CRUD)
type UserForm struct {
	Username *string
	Password *string //this shouldn't be applied
	Email    *string
}
func (this UserForm) apply(user *User){
	if(this.Username != nil) { user.Username = *this.Username }
	//no need to apply password since it needs to be hashed+salted first anyway
	if(this.Email != nil) { user.Email = *this.Email }
}
func (this UserForm) isPasswordChange() bool{
	if(this.Password != nil){
		return true
	} else {
		return false
	}
}