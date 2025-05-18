package storage

import (
	"github.com/google/uuid"
	"github.com/kerilOvs/backend/internal/models"
)

type UserStorage interface {
	// Основные операции с пользователем
	CreateUser(user *models.User) error
	GetUserByID(id uuid.UUID) (*models.User, error)
	UpdateUser(id uuid.UUID, updates map[string]interface{}) error
	DeleteUser(id uuid.UUID) error

	// Фото пользователя
	AddDoc(photo *models.UserDoc) error
	GetUserDocs(userID uuid.UUID) ([]*models.UserDoc, error)
	RemoveDoc(userID, photoID uuid.UUID) error
}
