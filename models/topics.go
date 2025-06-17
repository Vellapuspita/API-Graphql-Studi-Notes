package models


type Topic struct {
	ID        uint      gorm:"column:id_topics;primary_key" json:"id"
	Topics    string    gorm:"column:topics;type:varchar(255);not null" json:"topics"
}

func (Topic) TableName() string {
	return "topics"
}