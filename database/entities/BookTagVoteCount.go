package entities

type BookTagVoteCount struct {
	Tag          	string			`json:"title"`
	Count			int				`json:"count"`

	BookId			int				`json:"book_id"`
	Book			*Book			`json:"book,omitempty" gorm:"foreignkey:BookId"`

	//housekeeping attributes
	ID        		int				`json:"id"`
}

//for GORM
func (BookTagVoteCount) TableName() string {
	return "btags_vote_count"
}