package entities

type BookTagVote struct {
	Tag          	string			`json:"title"`

	BookId			int				`json:"book_id"`
	Book			*Book			`json:"book,omitempty" gorm:"foreignkey:BookId"`

	UserId      	int				`json:"user_id"`
	User		   	*User			`json:"user,omitempty" gorm:"foreignkey:UserId"`

	Direction 		int				`json:"direction"`

	//housekeeping attributes
	ID        		int				`json:"id"`
}

//for GORM; table name
func (BookTagVote) TableName() string {
	return "btags_votes"
}

//for GORM; associations
const ASSOC_BOOK_TAG_BOOK = "Book"
const ASSOC_BOOK_TAG_USER = "User"