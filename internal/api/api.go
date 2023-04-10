package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"venn-auth-api/configs"
	"venn-auth-api/internal/database"
)

type Api struct {
	db     *sql.DB
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
	defer db.Close()

	api.db = db

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
