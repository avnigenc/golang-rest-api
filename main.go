package main

import (
	"io"
	"log"
	"os"

	"github.com/avnigenc/go-api/controllers"
	"github.com/avnigenc/go-api/endpoint"
	"github.com/avnigenc/go-api/middlewares"
	"github.com/avnigenc/go-api/modules"
	"github.com/avnigenc/go-api/shared"

	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	logFile, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	var cfg shared.Config
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal("config error")
	}

	_ = os.Setenv(shared.ServerHost, cfg.ServerHost)
	_ = os.Setenv(shared.ServerPort, cfg.ServerPort)
	_ = os.Setenv(shared.DatabaseHost, cfg.DatabaseHost)
	_ = os.Setenv(shared.DatabasePort, cfg.DatabasePort)
	_ = os.Setenv(shared.JwtSecret, cfg.JwtSecret)
	_ = os.Setenv(shared.TokenIssuer, cfg.TokenIssuer)
	_ = os.Setenv(shared.TokenAudience, cfg.TokenAudience)
	_ = os.Setenv(shared.JwtExpireTime, cfg.JwtExpireTime.String())

	router := gin.Default()
	router.Use(gin.Recovery())

	modules.InitDB()

	// public group
	publicRoute := router.Group(endpoint.PREFIX)

	// Common
	publicRoute.GET(endpoint.ROOT, controllers.IndexController)
	publicRoute.GET(endpoint.HEALTH, controllers.HealthController)

	// Auth
	authRoute := publicRoute.Group(endpoint.AUTH)
	authRoute.POST(endpoint.LOGIN, controllers.LoginController)
	authRoute.POST(endpoint.REGISTER, controllers.RegisterController)

	// private group
	privateRoute := router.Group(endpoint.PREFIX, middlewares.TokenHandler)

	// Users
	usersRoute := privateRoute.Group(endpoint.USERS)
	usersRoute.GET(endpoint.ROOT, controllers.MeController)
	usersRoute.PUT(endpoint.ROOT, controllers.UpdateUserController)

	err = router.Run(":" + cfg.ServerPort)
	if err != nil {
		return
	}
}
