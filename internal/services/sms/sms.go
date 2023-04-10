package sms

import (
	"fmt"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"math"
	"math/rand"
)

type Service struct {
	client *twilio.RestClient
}

const codeLength = 6

func NewService(accountSid, authToken string) *Service {
	return &Service{
		client: twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: accountSid,
			Password: authToken,
		}),
	}
}

func (s *Service) SendVerificationCode(from, to string) (string, error) {
	// Generate a random 6-digit code
	code := rand.Intn(int(math.Pow10(codeLength)-1)+1) + int(math.Pow10(codeLength-1))
	body := fmt.Sprintf("Venn verification code %d", code)

	// Send SMS with code to the phone number
	_, err := s.Send(from, to, body)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", code), nil
}

func (s *Service) Send(from, to, body string) (*twilioApi.ApiV2010Message, error) {
	params := &twilioApi.CreateMessageParams{
		To:   &to,
		From: &from,
		Body: &body,
	}
	message, err := s.client.Api.CreateMessage(params)
	if err != nil {
		return nil, err
	}
	return message, nil
}
