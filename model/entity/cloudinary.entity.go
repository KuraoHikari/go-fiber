package entity

import (
	"time"

	"gorm.io/gorm"
)

type Cloudinary struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" form:"name"`
	Image     string         `json:"image" form:"image"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
