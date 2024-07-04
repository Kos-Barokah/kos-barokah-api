package users

import (
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Role           string    `json:"role"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	PhoneNumber    string    `json:"phone_number"`
	TokenResetPass string    `json:"token_reset_pass"`
	Points         uint      `json:"points" gorm:"column:points"`
	Status         string    `json:"status"`
}

type Admin struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Role           string `json:"role"`
	PhoneNumber    string `json:"phone_number"`
	TokenResetPass string `json:"token_reset_pass"`
	Status         string `json:"status"`
}

type UserInfo struct {
	Name   string         `json:"name"`
	Email  string         `json:"email"`
	Access map[string]any `json:"token"`
}

type UserCredential struct {
	Name   string         `json:"name"`
	Email  string         `json:"email"`
	Access map[string]any `json:"token"`
}

type UpdateAdmin struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Role        string    `json:"role"`
	PhoneNumber string    `json:"phone_number"`
	Status      string    `json:"status"`
}

type UserResetPass struct {
	Email     string
	Code      string
	ExpiresAt time.Time
}

type UpdateProfile struct {
	Name     string
	Email    string
	Password string
}

type UserHandlerInterface interface {
	Register() echo.HandlerFunc
	RegisterCustomer() echo.HandlerFunc
	Login() echo.HandlerFunc
	LoginCustomer() echo.HandlerFunc
	ForgetPasswordWeb() echo.HandlerFunc
	ResetPassword() echo.HandlerFunc
	ForgetPasswordVerify() echo.HandlerFunc
	UpdateProfile() echo.HandlerFunc
	RefreshToken() echo.HandlerFunc
	GetProfile() echo.HandlerFunc
}

type UserServiceInterface interface {
	Register(newData User) (*User, error)
	RegisterCustomer(newData User) (*User, error)
	Login(email string, password string) (*UserCredential, error)
	LoginCustomer(email string, password string) (*UserCredential, error)
	GenerateJwt(email string) (*UserCredential, error)
	ForgetPasswordWeb(email string) error
	TokenResetVerify(code string) (*UserResetPass, error)
	ResetPassword(code, email string, password string) error
	UpdateProfile(id int, newData UpdateProfile) (bool, error)
	AddPoints(id int, value int) (bool, error)
	DeductPoints(id int, value int) (bool, error)
	GetProfile(id int) (*User, error)
}

type UserDataInterface interface {
	Register(newData User) (*User, error)
	Login(email string, password string) (*User, error)
	LoginCustomer(email string, password string) (*User, error)
	GetByID(id int) (User, error)
	GetByEmail(email string) (*User, error)
	InsertCode(email string, code string) error
	DeleteCode(email string) error
	GetByCode(code string) (*UserResetPass, error)
	ResetPassword(code, email string, password string) error
	UpdateProfile(id int, newData UpdateProfile) (bool, error)
	AddPoints(userID int, value int) (bool, error)
	DeductPoints(userID int, value int) (bool, error)
}
