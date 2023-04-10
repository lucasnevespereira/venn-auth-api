package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"venn-auth-api/internal/models"
)

func (api *Api) sendSMS(c *gin.Context) {
	var req models.SmsRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: send SMS code to the phone number
	_, err = api.sms.Send(api.config.Twilio.PhoneNumber, req.Phone, "Hello from Venn Auth API")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	res := models.SmsResponse{Message: fmt.Sprintf("Sms sent to %s", req.Phone)}
	c.JSON(http.StatusOK, res)
}

func (api *Api) verifySMS(c *gin.Context) {
	var req models.VerifyRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: verify SMS code for the phone number

	// Generate and return a JWT token
	token := "dummy-token"
	res := models.VerifyResponse{Token: token}
	c.JSON(http.StatusOK, res)
}

func (api *Api) handleLogin(c *gin.Context) {
	var req models.LoginRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: verify SMS code for the phone number

	// Generate and return a JWT token
	token := "dummy-token"
	res := models.LoginResponse{Token: token}
	c.JSON(http.StatusOK, res)
}

func (api *Api) handleSignup(c *gin.Context) {
	var req models.SignupRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: verify SMS code for the phone number

	// TODO: create user in the database

	res := models.SignupResponse{Message: "Account created"}
	c.JSON(http.StatusOK, res)
}
