package main

import (
	"github.com/dedihartono801/go-clean-architecture/api"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/config"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/database"
	"github.com/gofiber/swagger"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @contact.name Dedi Hartono
// @contact.email dedihartono801@mail.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	swagger.New(swagger.Config{
		Title:        "Swagger API",
		DeepLinking:  false,
		DocExpansion: "none",
	})

	envConfig := config.SetupEnvFile()

	mysql := database.InitMysql(envConfig)
	//mongo := database.InitMongo(envConfig.MongoAddress, envConfig.DatabaseName)
	api.Execute(mysql)
}
