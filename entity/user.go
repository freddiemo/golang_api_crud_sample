package entity

import "time"

type User struct {
	ID uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `json:"first_name" binding:"required,min=4,max=32" gorm:"type:varchar(32)"`
	LastName string `json:"last_name" binding:"required,min=4,max=32" gorm:"type:varchar(32)"`
	Email string `json:"email" binding:"required,email"`
	Phone    int `json:"phone" binding:"required"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}