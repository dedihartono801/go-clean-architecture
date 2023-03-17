package api

import (
	"log"

	"github.com/dedihartono801/go-clean-architecture/infrastructure/http"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/http/handler"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/identifier"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/repository"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/validator"
	"github.com/dedihartono801/go-clean-architecture/usecase/admin"
	"github.com/dedihartono801/go-clean-architecture/usecase/book"
	"github.com/dedihartono801/go-clean-architecture/usecase/product"
	"github.com/dedihartono801/go-clean-architecture/usecase/user"
	validatorv10 "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Execute(database *gorm.DB) {

	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())
	userRepository := repository.NewUserRepository(database)
	userService := user.NewUserService(userRepository, validator, identifier)
	userHandler := handler.NewUserHandler(userService)

	bookRepository := repository.NewBookRepository()
	bookService := book.NewService(bookRepository, validator)
	bookHandler := handler.NewBookHandler(bookService)

	adminRepository := repository.NewAdminRepository(database)
	adminService := admin.NewAdminService(adminRepository, validator, identifier)
	adminHandler := handler.NewAdminHandler(adminService)

	productService := product.NewProductService(validator, identifier)
	productHandler := handler.NewFilmHandler(productService)

	app := fiber.New()

	http.SetupRoutes(
		app,
		userHandler,
		bookHandler,
		adminHandler,
		productHandler,
	)

	if err := app.Listen(":5001"); err != nil {
		log.Fatalf("listen: %s", err)
	}
}
