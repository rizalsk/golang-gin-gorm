package models

import (
	"gorm.io/gorm"
)

type User struct {
	// ID            uint           `gorm:"primaryKey" json:"id"`
	gorm.Model
	Name          string `gorm:"size:100;not null" json:"name"`
	Email         string `gorm:"size:100;not null;uniqueIndex;" json:"email"`
	Password      string `gorm:"size:100;not null" json:"password"`
	RememberToken string `gorm:"size:100;default:null" json:"remember_token"`
	// CreatedAt     time.Time      `sql:"DEFAULT:'current_timestamp'"`
	// UpdatedAt     time.Time      `sql:"DEFAULT:ON UPDATE current_timestamp"`
	// DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (u *User) TableName() string {
	return "users"
}
