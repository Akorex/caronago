package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct{
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	FirstName string `gorm:"not null" json:"first_name"`
	LastName string `gorm:"not null" json:"last_name"`
	Email string `gorm:"not null; uniqueIndex" json:"email"`
	Password string `gorm:"not null" json:"-"`
	IsVerified bool `gorm:"not null; default:false" json:"is_verified"`
	PasswordResetToken *string `gorm:"default:null" json:"-"`
	PasswordResetExpires *time.Time `gorm:"default:null" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

}