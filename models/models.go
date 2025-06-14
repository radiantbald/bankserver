package models

type User struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"column:name;size:255"`
	Surname     string `gorm:"column:surname;size:255"`
	PhoneNumber string `gorm:"column:phone_number;size:255"`
	Verified    bool   `gorm:"column:verified"`
}
