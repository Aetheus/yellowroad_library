package entities

import (
	"time"
	"github.com/lib/pq"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"-"` //this excludes it from being accidentally serialized into a JSON and shown to users
	Email    string `json:"email"`

	//housekeeping attributes
	ID        int `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt pq.NullTime `json:"deleted_at"`
}


//fields that we allow to edit in our handlers (e.g: for the "update" routes of CRUD)
type UserForm struct {
	Username *string	`json:"username"`
	Password *string 	`json:"password"` //this shouldn't be applied
	Email    *string	`json:"email"`
}
func (this UserForm) Apply(user *User){
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