package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/km2/go-login-example/model"
	"github.com/km2/go-login-example/repository"
)

type LoginHandler struct {
	authRepository repository.AuthRepository
}

func NewLoginHandler(authRepository repository.AuthRepository) *LoginHandler {
	return &LoginHandler{authRepository: authRepository}
}

func (h *LoginHandler) GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (h *LoginHandler) PostLogin(c *gin.Context) {
	var a model.Auth
	if err := c.Bind(&a); err != nil {
		log.Printf("failed to bind JSON: %v\n", err)
		c.Status(http.StatusBadRequest)
		return
	}

	log.Printf("email: %s, password: %s\n", a.Email, a.Password)

	auth, err := h.authRepository.GetAuthByEmail(a.Email)
	if err != nil {
		log.Printf("failed to get auth by email: %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	if a.Password != auth.Password {
		log.Printf("password is incorrect (%s vs %s)\n", a.Password, auth.Password)
		c.Status(http.StatusUnauthorized)
		return
	}

	user, err := json.Marshal(auth)
	if err != nil {
		log.Printf("failed to marshal auth: %v\n", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	session := sessions.Default(c)
	session.Set("user", string(user))
	if err := session.Save(); err != nil {
		log.Printf("failed to save session: %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusFound, "/auth/hello")
}
