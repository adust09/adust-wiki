package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	Username     string    `gorm:"size:255;not null"`
	Email        string    `gorm:"size:255;not null;unique"`
	PasswordHash string    `gorm:"size:255;not null"`
	CreatedAt    time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (er error) {
	u.ID = uuid.New()
	return
}
