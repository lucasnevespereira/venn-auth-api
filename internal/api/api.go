package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"venn-auth-api/configs"
	"venn-auth-api/internal/database"
	"venn-auth-api/internal/services/sms"
)

type Api struct {
	sms    *sms.Service
	db     *gorm.DB
	config configs.ApiConfig
}

func New() *Api {
	return &Api{config: configs.LoadApiConfig()}
}

func (api *Api) Run() error {

	db, err := database.Open(database.Config{
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
	auth.POST("/sms", api.sendSMS)
	auth.POST("/verify", api.verifySMS)
	auth.POST("/signup", api.handleSignup)
	auth.POST("/login", api.handleLogin)

	err = router.Run(fmt.Sprintf(":%d", api.config.Port))
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
