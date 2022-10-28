package entity

import "time"

type User struct {
	ID uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `json:"first_name" binding:"required" gorm:"type:varchar(32)"`
	LastName string `json:"last_name" binding:"required" gorm:"type:varchar(32)"`
	Phone    int `json:"phone" binding:"required"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}