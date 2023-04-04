package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"noir.github.com/aws_billing/internal/models"
	"noir.github.com/aws_billing/internal/routers/user"
)

//const DefaultPort = 8080

func initEnv() {
	models.DB = models.CreateDB()
}

func main() {
	initEnv()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
	}))
	r.POST("/api/auth/login", user.CheckUserStatus, user.ValidateLogin)
	r.POST("/api/auth/register", user.Registry)
	r.POST("/api/user/disable", user.RestoreToken, user.Disable)
	r.POST("/api/user/enable", user.RestoreToken, user.Enable)
	r.GET("/api/user/list", user.RestoreToken, user.List)
	//r.GET("api/auth/authorization", user.RestoreToken)
	r.GET("/api/billing/gcp", func(c *gin.Context) {

	})
	err := r.Run("0.0.0.0:80")
	if err != nil {
		panic(err)
	}
}
