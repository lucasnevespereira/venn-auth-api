package sms

import (
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type Service struct {
	client *twilio.RestClient
}

func NewService(accountSid, authToken string) *Service {
	return &Service{
		client: twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: accountSid,
			Password: authToken,
		}),
	}
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
