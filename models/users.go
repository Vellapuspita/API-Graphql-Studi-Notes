package models

import "time"

type User struct {
	ID             uint      `gorm:"column:id_users;primaryKey" json:"id"`
	Name           string    `gorm:"column:username" json:"name"`
	Email          string    `gorm:"column:email" json:"email"`
	Password       string    `gorm:"column:password" json:"password"`
	ConfirPassword string    `gorm:"column:confirm_password" json:"confirm_password"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`

	// Relasi many-to-many dengan StudyNotes melalui tabel "collab"
	StudyNotes []StudyNote `gorm:"many2many:collab;joinForeignKey:ID;joinReferences:ID"`
}

func (User) TableName() string {
	return "users"
}
