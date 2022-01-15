package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/km2/go-login-example/model"
	"github.com/koron/go-dproxy"
)

type LoginMiddleware struct{}

func NewLoginMiddleware() *LoginMiddleware {
	return &LoginMiddleware{}
}

func (m *LoginMiddleware) Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		sessionUser := session.Get("user")
		if sessionUser == nil {
			log.Printf("user is not logged in\n")
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		authJSON, err := dproxy.New(sessionUser).String()
		if err != nil {
			log.Printf("failed to get session user as string: %v\n", err)
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		var auth model.Auth
		if err := json.Unmarshal([]byte(authJSON), &auth); err != nil {
			log.Printf("failed to unmarshal JSON: %v\n", err)
			c.Status(http.StatusInternalServerError)
			c.Abort()
			return
		}

		log.Printf("login successful: %s\n", auth.Email)

		c.Set("email", auth.Email)

		c.Next()
	}
}
