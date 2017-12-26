package entities

type BookTag struct {
	Tag          	string			`json:"title"`

	BookId			int				`json:"book_id"`
	Book			*Book			`json:"book,omitempty" gorm:"ForeignKey:BookId"`

	UserId      	int				`json:"user_id"`
	User		   	*User			`json:"user,omitempty" gorm:"ForeignKey:UserId"`

	//housekeeping attributes
	ID        		int				`json:"id"`
}

//for GORM; table name
func (BookTag) TableName() string {
	return "book_tags"
}

//for GORM; associations
const BOOK_TAG_BOOK = "Book"
const BOOK_TAG_USER = "User"