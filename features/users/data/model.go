package data

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Name           string    `gorm:"column:name;type:varchar(255)"`
	Email          string    `gorm:"column:email;unique;type:varchar(255)"`
	Password       string    `gorm:"column:password;type:varchar(255)"`
	Role           string    `gorm:"column:role;type:enum('admin', 'customer')"`
	DateOfBirth    time.Time `gorm:"column:date_of_birth"`
	PhoneNumber    string    `gorm:"column:phone_number;type:varchar(255)"`
	TokenResetPass string    `gorm:"column:token_reset_pass;type:varchar(255)"`
	Points         uint      `gorm:"column:points"`
	Status         string    `gorm:"column:status;type:enum('active', 'not_active')"`
}

type UserResetPass struct {
	*gorm.Model
	Email     string    `json:"email" gorm:"email"`
	Code      string    `json:"code" gorm:"code"`
	ExpiresAt time.Time `json:"expiresAt" gorm:"expires_at"`
}
