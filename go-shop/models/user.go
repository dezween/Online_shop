package models

// import "github.com/jinzhu/gorm"

type User struct {
	ID        uint   `gorm:"primary_key"`
	Login     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	FirstName string `gorm:"not null"`
	SecondName string `gorm:"not null"`
	Phone     string
	Email     string `gorm:"unique;not null"`
}
