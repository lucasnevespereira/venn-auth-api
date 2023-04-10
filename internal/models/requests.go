package models

type SendCodeRequest struct {
	PhoneNumber string `json:"phone_number"`
}

type VerifyCodeRequest struct {
	PhoneNumber string `json:"phone_number"`
	Code        string `json:"code"`
}
