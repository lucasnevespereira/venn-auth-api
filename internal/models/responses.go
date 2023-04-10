package models

type SmsResponse struct {
	Message string `json:"message"`
}

type VerifyResponse struct {
	Token string `json:"token"`
}

type CreateUserResponse struct {
	Message string `json:"message"`
}

type SignupResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
