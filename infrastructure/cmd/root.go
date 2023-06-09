package cmd

import (
	"os"

	"github.com/dedihartono801/go-clean-architecture/infrastructure/identifier"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/repository"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/validator"
	pvalidator "github.com/go-playground/validator/v10"

	"github.com/dedihartono801/go-clean-architecture/usecase/user"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func Execute(database *gorm.DB) {
	// dependency
	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(pvalidator.New())
	userRepository := repository.NewUserRepository(database)
	userService := user.NewUserService(userRepository, validator, identifier)

	// userCmd
	userCmd := NewUserCmd()
	userCmd.AddCommand(NewCreateUserCmd(userService))
	userCmd.AddCommand(NewUpdateUserCmd(userService))

	// rootCmd
	rootCmd := NewRootCmd()
	rootCmd.AddCommand(userCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "go-clean-arch",
		Short: "This application exemplifies the use of clean architecture using go language",
	}
	return rootCmd
}
