package http

import (
	_ "github.com/dedihartono801/go-clean-architecture/docs"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/http/handler"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
)

func SetupRoutes(
	app fiber.Router,
	userHandler handler.UserHandler,
	bookHandler handler.BookHandler,
	adminHandler handler.AdminHandler,
	productHandler handler.ProductHandler,
	skuHandler handler.SkuHandler,
	transactionHandler handler.TransactionHandler,

) {
	app.Get("/", monitor.New())
	app.Get("/docs/*", swagger.HandlerDefault)

	adminRoute := app.Group("/admin")
	adminRoute.Post("/login", adminHandler.Login)
	adminRoute.Post("/create", adminHandler.Create)
	adminRoute.Get("", middleware.AuthUser, adminHandler.Find)

	booksRoute := app.Group("/books", middleware.AuthUser)
	booksRoute.Get("", bookHandler.List)
	booksRoute.Get("/:id", bookHandler.Find)
	booksRoute.Put("/:id", bookHandler.Update)
	booksRoute.Post("", bookHandler.Create)
	booksRoute.Delete("/:id", bookHandler.Delete)

	usersRoute := app.Group("/users", middleware.AuthUser)
	usersRoute.Get("", userHandler.List)
	usersRoute.Get("/:id", userHandler.Find)
	usersRoute.Put("/:id", userHandler.Update)
	usersRoute.Post("", userHandler.Create)
	usersRoute.Delete("/:id", userHandler.Delete)

	skuRoute := app.Group("/sku", middleware.AuthUser)
	skuRoute.Get("", skuHandler.List)
	skuRoute.Post("", skuHandler.Create)

	productRoute := app.Group("/product", middleware.AuthUser)
	productRoute.Get("", productHandler.Product)

	app.Post("/checkout", middleware.AuthUser, transactionHandler.Checkout)
}
