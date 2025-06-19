package models

import "time"

type StudyNotes struct {
	ID        uint      `gorm:"column:id;primaryKey" json:"id"`
	IDUsers   uint      `gorm:"column:id_users" json:"id_users"`
	IDTopics  uint      `gorm:"column:id_topics" json:"id_topics"`
	Judul     string    `gorm:"column:judul" json:"judul"`
	CreatedBy uint      `gorm:"column:created_by" json:"created_by"`
	IsGrup    bool      `gorm:"column:is_grup" json:"is_grup"`
	Deskripsi string    `gorm:"column:deskripsi;type:text" json:"deskripsi"`
	Content   string    `gorm:"column:content;type:text" json:"content"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`

	// Relasi many-to-many dengan Users melalui tabel "collab"
	Users []Users `gorm:"many2many:collab;joinForeignKey:ID;joinReferences:ID"`
}

func (StudyNotes) TableName() string {
	return "studynotes"
}
