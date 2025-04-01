package models

import (
	"github.com/google/uuid"
)

type Document struct {
	ID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title   string
	Content string
	OwnerID uuid.UUID `gorm:"type:uuid;not null"`
	Owner   User      `gorm:"foreignKey:OwnerID"`
}
