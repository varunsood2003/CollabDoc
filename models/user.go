package models

type User struct {
	ID        string     `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Email     string     `json:"email" gorm:"unique"`
	Password  string     `json:"password"`
	Documents []Document `json:"documents" gorm:"foreignKey:OwnerID"`
}
