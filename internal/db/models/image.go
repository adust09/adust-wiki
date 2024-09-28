package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Image struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	Filename    string    `gorm:"size:255;not null"`
	Size        int64     `gorm:"not null"`
	UploadURL   string    `gorm:"size:255;not null"`
	Description string    `gorm:"type:text"`
	Tags        string    `gorm:"type:text"`
	CreatedAt   time.Time
}

func (i *Image) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = uuid.New()
	return
}
