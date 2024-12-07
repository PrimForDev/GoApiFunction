package models

type User struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"size:255;not null"`
	Email     string `gorm:"size:255;not null;unique"`
	FirstName string `gorm:"size:255;"`
	LastName  string `gorm:"size:255;"`
	Role      string `gorm:"size:255;not null"`
}
