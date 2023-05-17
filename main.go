package main

import (
	"context"
	"log"

	"github.com/dedihartono801/go-clean-architecture/api"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/config"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/database"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/worker"
	"github.com/gofiber/swagger"
	"github.com/hibiken/asynq"
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
	redisOpt := asynq.RedisClientOpt{
		Addr: envConfig.RedisAddress,
	}
	ctx := context.Background()
	taskDistributor := worker.NewRedisTaskDistributor(redisOpt, ctx)

	mysql := database.InitMysql(envConfig)
	go runWorkerServer(*envConfig, redisOpt)
	api.Execute(mysql, taskDistributor)
}

func runWorkerServer(config config.Config, redisOpt asynq.RedisClientOpt) {
	taskProcessor := worker.NewServer(redisOpt)
	err := taskProcessor.Start()
	if err != nil {
		log.Fatalf("failed to start worker")
	}
	log.Println("start worker")
}
