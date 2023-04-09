package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type smsRequest struct {
	Phone string `json:"phone"`
}

type smsResponse struct {
	Message string `json:"message"`
}

type verifyRequest struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type verifyResponse struct {
	Token string `json:"token"`
}

type createUserRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type createUserResponse struct {
	Message string `json:"message"`
}

type AuthServer struct {
	db *sql.DB
}

func (s *AuthServer) sendSMS(c *gin.Context) {
	var req smsRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: send SMS code to the phone number

	res := smsResponse{Message: "SMS sent"}
	c.JSON(http.StatusOK, res)
}

func (s *AuthServer) verifySMS(c *gin.Context) {
	var req verifyRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: verify SMS code for the phone number

	// Generate and return a JWT token
	token := "dummy-token"
	res := verifyResponse{Token: token}
	c.JSON(http.StatusOK, res)
}

func (s *AuthServer) createUser(c *gin.Context) {
	var req createUserRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: create user in the database

	res := createUserResponse{Message: "User created"}
	c.JSON(http.StatusOK, res)
}

func main() {
	db, err := sql.Open("postgres", "user=foo dbname=bar password=baz sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := &AuthServer{db: db}

	r := gin.Default()

	r.POST("/api/auth/sms", server.sendSMS)
	r.POST("/api/auth/verify", server.verifySMS)
	r.POST("/api/auth/create_user", server.createUser)

	log.Fatal(r.Run(":8080"))
}
