package models

import "time"

type Collab struct {
	ID        	uint      	gorm:"column:id_collab;primary_key" json:"id"
	IDUser    	int      	gorm:"column:id_user" json:"id_user"         // relasi ke user
	IDNotes   	int      	gorm:"column:id_notes" json:"id_notes"       // relasi ke study notes
	CreatedAt 	time.Time 	gorm:"column:created_at" json:"created_at"
	UpdatedAt 	time.Time 	gorm:"column:updated_at" json:"updated_at"
}

func (Collab) TableName() string {
	return "collab"
}