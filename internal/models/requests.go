package models

type SmsRequest struct {
	Phone string `json:"phone"`
}

type VerifyRequest struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type SignupRequest struct {
	PhoneNumber string `json:"phone_number"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Code        string `json:"code"`
}
