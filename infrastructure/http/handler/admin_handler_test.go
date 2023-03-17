package handler

import (
	"testing"

	"github.com/dedihartono801/go-clean-architecture/infrastructure/identifier"
	repoMock "github.com/dedihartono801/go-clean-architecture/infrastructure/repository/mock"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/validator"
	usecaseMock "github.com/dedihartono801/go-clean-architecture/usecase/admin/mock"
	validatorv10 "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

func MockNewService(t *testing.T) usecaseMock.Service {

	repo := repoMock.NewMockAdminRepository()
	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())
	srv := usecaseMock.NewMockAdminService(repo, validator, identifier)
	return srv
}

func TestFind(t *testing.T) {
	srv := MockNewService(t)

	// Create a new Fiber app
	app := fiber.New()

	// Define test cases
	testCases := []struct {
		name       string
		id         string
		statusCode int
	}{
		{
			name:       "Found data",
			id:         "4d35bf38-8c50-4c85-8072-fd9794803a167",
			statusCode: 200,
		},
		{
			name:       "Not found data",
			id:         "8734JJHYD88",
			statusCode: 404,
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Define the getUserByID route
			app.Get("/admin", func(c *fiber.Ctx) error {
				// Set the adminID value in the context's locals map
				c.Locals("adminID", tc.id)
				handler := NewAdminHandler(srv)
				return handler.Find(c)
			})

			// Define a mock request for testing
			req := httptest.NewRequest(http.MethodGet, "/admin", nil)

			resp, err := app.Test(req, -1)

			// ensure that there are no errors
			assert.NoError(t, err)

			// ensure that the response status code is 200 OK
			assert.Equal(t, tc.statusCode, resp.StatusCode)
		})
	}

}
