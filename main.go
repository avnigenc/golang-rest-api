package main

import (
	"github.com/avnigenc/go-api/controllers"
	"github.com/avnigenc/go-api/middlewares"
	"github.com/avnigenc/go-api/modules"
	"github.com/avnigenc/go-api/shared"
	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	"io"
	"log"
	"os"
)

func main() {
	logFile, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	var cfg shared.Config
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal("config error")
	}

	// TODO: Refactor handle errors
	_ = os.Setenv("ServerHost", cfg.ServerHost)
	_ = os.Setenv("ServerPort", cfg.ServerPort)
	_ = os.Setenv("DatabaseHost", cfg.DatabaseHost)
	_ = os.Setenv("DatabasePort", cfg.DatabasePort)
	_ = os.Setenv("JwtSecret", cfg.JwtSecret)
	_ = os.Setenv("JwtExpireTime", cfg.JwtExpireTime.String())

	router := gin.Default()
	router.Use(gin.Recovery())

	modules.InitDB()

	publicRoute := router.Group("/api")
	publicRoute.GET("", controllers.IndexController)

	// Auth
	authRoute := publicRoute.Group("/auth")
	authRoute.POST("login", controllers.LoginController)
	authRoute.POST("register", controllers.RegisterController)

	privateRoute := router.Group("/api", middlewares.TokenHandler)
	// Users
	usersRoute := privateRoute.Group("users")
	usersRoute.GET("", controllers.MeController)
	usersRoute.PUT("", controllers.UpdateUserController)

	err = router.Run(":" + cfg.ServerPort)
	if err != nil {
		return 
	}
}

