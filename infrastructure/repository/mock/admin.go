package mock

import (
	"errors"

	"github.com/dedihartono801/go-clean-architecture/domain"
)

// Define the MockAdminRepository interface
type AdminRepository interface {
	Find(id string) (*domain.Admin, error)
	Create(admin *domain.Admin) error
	FindByEmail(email string) (*domain.Admin, error)
}

type mockAdminRepository struct {
	admin map[string]*domain.Admin
}

func NewMockAdminRepository() AdminRepository {
	return &mockAdminRepository{
		admin: make(map[string]*domain.Admin),
	}

}

// Create a new admin
func (m *mockAdminRepository) Create(admin *domain.Admin) error {
	if _, ok := m.admin[admin.ID]; ok {
		return errors.New("user already exists")
	}
	m.admin[admin.ID] = admin
	return nil
}

// // Update an existing user
// func (m *mockAdminRepository) Update(user *domain.User) error {
// 	if _, ok := m.users[user.ID]; !ok {
// 		return errors.New("user not found")
// 	}
// 	user.UpdatedAt = time.Now()
// 	m.users[user.ID] = user
// 	return nil
// }

// Delete an existing user
// func (m *mockAdminRepository) Delete(id string) error {
// 	if _, ok := m.users[id]; !ok {
// 		return errors.New("user not found")
// 	}
// 	delete(m.users, id)
// 	return nil
// }

// FindByID finds a user by ID
func (m *mockAdminRepository) Find(id string) (*domain.Admin, error) {
	if admin, ok := m.admin[id]; ok {
		return admin, nil
	}
	return nil, errors.New("admin not found")
}

// FindByEmail finds a user by email
func (m *mockAdminRepository) FindByEmail(email string) (*domain.Admin, error) {
	for _, admin := range m.admin {
		if admin.Email == email {
			return admin, nil
		}
	}
	return nil, errors.New("admin not found")
}
