package handler

import "time"

type RegisterInput struct {
	Name        string    `json:"name" form:"name" validate:"required"`
	Email       string    `json:"email" form:"email" validate:"required,email"`
	Password    string    `json:"password" form:"password" validate:"required"`
	DateOfBirth time.Time `json:"date_of_birth" form:"date_of_birth" validate:"required"`
	PhoneNumber string    `json:"phone_number" form:"phone_number" validate:"required"`
	Role        string    `json:"role" form:"role" validate:"required"`
}

type RegisterInputCustomer struct {
	Name        string    `json:"name" form:"name" validate:"required"`
	Email       string    `json:"email" form:"email" validate:"required,email"`
	Password    string    `json:"password" form:"password" validate:"required"`
	DateOfBirth time.Time `json:"date_of_birth" form:"date_of_birth" validate:"required"`
	PhoneNumber string    `json:"phone_number" form:"phone_number" validate:"required"`
	// Position       string `json:"position" form:"position" validate:"required"`
	// CompanyName    string `json:"company_name" form:"company_name" validate:"required"`
	// CompanyEmail   string `json:"company_email" form:"company_email" validate:"required"`
	// CompanyAddress string `json:"company_address" form:"company_address" validate:"required"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type ForgetPasswordInput struct {
	Email string `json:"email" form:"email"`
}

type ResetPasswordInput struct {
	Password        string `json:"password" form:"password" validate:"required"`
	PasswordConfirm string `json:"password_confirm" form:"password_confirm" validate:"required"`
}

type UpdateProfile struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RefreshInput struct {
	Token string `json:"access_token" form:"access_token"`
}
