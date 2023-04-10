package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"venn-auth-api/internal/models"
)

func (api *Api) sendCode(c *gin.Context) {
	var req models.SendCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Send verification code via SMS
	code, err := api.sms.SendVerificationCode(api.config.Twilio.PhoneNumber, req.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification code"})
		return
	}

	// TODO: Store the verification code in the database for the phone number
	err = api.db.SaveVerificationCode(req.PhoneNumber, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"code": code})
}

func (api *Api) verifyCode(c *gin.Context) {
	var req models.VerifyCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Retrieve the verification code from the database for the phone number
	user, err := api.db.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
	}

	if req.Code != user.VerificationCode {
		// TODO: Track the number of failed verification attempts and lock the account if necessary
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid verification code"})
		return
	}

	// Generate and return a JWT token
	token := "dummy-token"
	res := models.VerifyResponse{Token: token}
	c.JSON(http.StatusOK, res)
}
