package main

import (
	"github.com/dedihartono801/go-clean-architecture/infrastructure/cmd"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/config"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/database"
)

func main() {
	envConfig := config.SetupEnvFile()
	db := database.InitMysql(envConfig)
	cmd.Execute(db)
}
