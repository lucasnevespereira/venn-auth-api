package database

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	db *gorm.DB
}

type Config struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
	SSLMode  string
}

func NewClient(config Config) (*Client, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=UTC",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Client{db: db}, nil
}

// SaveVerificationCode saves the verification code for the given phone number.
// If a user with the given phone number already exists, it updates the verification code.
// Otherwise, it creates a new user with the phone number and verification code.
func (c *Client) SaveVerificationCode(phoneNumber, code string) error {
	var user User
	result := c.db.Where("phone_number = ?", phoneNumber).First(&user)

	if result.Error != nil {
		// If the error is not "record not found", return it
		if result.Error != gorm.ErrRecordNotFound {
			return result.Error
		}

		// If the user doesn't exist, create a new one
		user = User{
			PhoneNumber:      phoneNumber,
			VerificationCode: code,
		}
		result = c.db.Create(&user)
	} else {
		// If the user already exists, update the verification code
		result = c.db.Model(&user).Update("verification_code", code)
	}

	return result.Error
}

func (c *Client) GetUserByPhoneNumber(phoneNumber string) (*User, error) {
	var user User
	result := c.db.Where(&User{PhoneNumber: phoneNumber}).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
