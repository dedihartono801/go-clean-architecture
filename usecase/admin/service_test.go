package admin

import (
	"testing"

	"github.com/dedihartono801/go-clean-architecture/domain"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/identifier"
	repoMock "github.com/dedihartono801/go-clean-architecture/infrastructure/repository/mock"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/validator"
	validatorv10 "github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	repo := repoMock.NewMockAdminRepository()
	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())
	srv := NewAdminService(repo, validator, identifier)

	expected := &domain.Admin{
		ID:       "4d35bf38-8c50-4c85-8072-fd9794803a16",
		Name:     "diding",
		Email:    "diding@gmail.com",
		Password: "56334b8232e95fb59b0fc93f2bc0d5c1fdbf5f120d91ac9f5d4c9db14544e007dd163cba5af3de3f027a6d47280f1407c19a5c1b8fc8ca10a4d7ef431341f135",
	}
	repo.Create(expected)

	// Define test cases
	testCases := []struct {
		name     string
		id       string
		expected *domain.Admin
		wantErr  bool
	}{
		{
			name:     "Found data",
			id:       expected.ID,
			expected: expected,
			wantErr:  false,
		},
		{
			name:     "Not found data",
			id:       "8734JJHYD88",
			expected: nil,
			wantErr:  true,
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Test retrieving an existing admin
			actual, err := srv.Find(tc.id)
			assert.Equal(t, tc.expected, actual, "Expected and actual data should be equal")

			assert.Equal(t, tc.wantErr, err != nil, "Expected error and actual error should be equal")
		})
	}

}
