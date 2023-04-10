package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"venn-auth-api/configs"
	"venn-auth-api/internal/database"
	"venn-auth-api/internal/services/sms"
)

type Api struct {
	sms    *sms.Service
	db     *database.Client
	config configs.ApiConfig
}

func New() *Api {
	return &Api{config: configs.LoadApiConfig()}
}

func (api *Api) Run() error {

	db, err := database.NewClient(database.Config{
		User:     api.config.DB.User,
		Password: api.config.DB.Password,
		Host:     api.config.DB.Host,
		Port:     api.config.DB.Port,
		DBName:   api.config.DB.Name,
		SSLMode:  api.config.DB.SSLMode,
	})
	if err != nil {
		return err
	}
	api.db = db

	smsService := sms.NewService(api.config.Twilio.AccountSid, api.config.Twilio.AuthToken)
	api.sms = smsService

	router := gin.Default()
	auth := router.Group("/api/auth")

	auth.POST("/verify", api.verifyCode)
	auth.POST("/send", api.sendCode)

	err = router.Run(fmt.Sprintf(":%d", api.config.Port))
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
