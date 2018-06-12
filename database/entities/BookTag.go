package entities

type BookTag struct {
	Tag          	string			`json:"title"`
	Count			int				`json:"count"`

	BookId			int				`json:"book_id"`
	Book			*Book			`json:"book,omitempty" gorm:"foreignkey:BookId"`

	//housekeeping attributes
	ID        		int				`json:"id"`
}

//for GORM
func (BookTag) TableName() string {
	return "book_tags"
}