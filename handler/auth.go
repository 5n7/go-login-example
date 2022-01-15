package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) GetHello(c *gin.Context) {
	email, exists := c.Get("email")
	if !exists {
		log.Printf("failed to get email from context\n")
		c.Status(http.StatusUnauthorized)
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("Hello, %v", email))
}
