package models

type Document struct {
	ID      string  `json:"id" gorm:"primaryKey"`
	Title   string  `json:"title"`
	Content *string `json:"content,omitempty"`
	OwnerID string  `json:"ownerId"`
}
