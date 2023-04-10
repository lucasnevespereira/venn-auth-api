package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	PhoneNumber      string `gorm:"phone_number,unique"`
	VerificationCode string `gorm:"verification_code"`
}
