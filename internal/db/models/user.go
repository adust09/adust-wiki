package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username            string    `gorm:"size:100;not null;unique"`
	Email               string    `gorm:"size:100;not null;unique"`
	PasswordHash        string    `gorm:"not null"`
	PasswordResetToken  string    `gorm:"size:255"` // Field for password reset token
	PasswordResetExpiry time.Time `gorm:""`         // Field for token expiry time
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (er error) {
	u.ID = uuid.New()
	return
}
