package entity

import "time"

type Patient struct {
	ID uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `json:"first_name" gorm:"type:varchar(32)"`
	LastName string `json:"last_name" gorm:"type:varchar(32)"`
	Email     string `json:"email" gorm:"type:varchar(30);UNIQUE" binding:"required,email"`
	DocumentId int `json:"document_id"`
	Phone    int `json:"phone"`
	Age int `json:"age" binding:"gte=1, lte=130"`
	Address string `json:"address" gorm:"type:varchar(256)"`
	OwnerID  uint64	`json:"owner_id" binding:"required"`
	Owner      User    `json:"owner" gorm:"foreignkey:OwnerID"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}