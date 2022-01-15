package main

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/km2/go-login-example/database/memory"
	"github.com/km2/go-login-example/handler"
	"github.com/km2/go-login-example/middleware"
)

var (
	CookieStoreKey = os.Getenv("COOKIT_STORE_KEY")
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/**/*")

	store := cookie.NewStore([]byte(CookieStoreKey))
	router.Use(sessions.Sessions("session", store))

	loginMiddleware := middleware.NewLoginMiddleware()

	loginHandler := handler.NewLoginHandler(memory.NewAuthMemoryDB())
	authHandler := handler.NewAuthHandler()

	authUserGroup := router.Group("/auth")
	authUserGroup.Use(loginMiddleware.Check())

	router.GET("/login", loginHandler.GetLogin)
	router.POST("/login", loginHandler.PostLogin)

	authUserGroup.GET("/hello", authHandler.GetHello)

	router.Run()
}
