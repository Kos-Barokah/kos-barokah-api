package handler

import (
	"time"
)

type RegisterResponse struct {
	Name        string    `json:"name" form:"name"`
	Email       string    `json:"email" form:"email"`
	DateOfBirth time.Time `json:"date_of_birth"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role" form:"role"`
}

type RegisterCustomerResponse struct {
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Role        string `json:"role" form:"role"`
	CompanyName string `json:"company_name" form:"company_name"`
}

type LoginResponse struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Token any    `json:"token"`
}

type ProfileResponse struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"date_of_birth"`
	PhoneNumber string    `json:"phone_number"`
	Status      string    `json:"status"`
	// CustomerDetail *customer.Customer `json:"customer_detail"`
}

type UserInfo struct {
	Name   string `json:"name" form:"name"`
	Email  string `json:"email" form:"email"`
	Status string `json:"status" form:"status"`
}

type AdminResponse struct {
	Name        string    `json:"name" form:"name"`
	Email       string    `json:"email" form:"email"`
	PhoneNumber string    `json:"phone_number"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Status      string    `json:"status" form:"status"`
}

type AdminDetailResponse struct {
	ID          uint      `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Email       string    `json:"email" form:"email"`
	PhoneNumber string    `json:"phone_number"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Role        string    `json:"role" form:"role"`
	Status      string    `json:"status" form:"status"`
}
