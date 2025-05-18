package postgres

import (
	"errors"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/kerilOvs/backend/internal/models"
)

type UserPostgresStorage struct {
	db *gorm.DB
}

func NewUserPostgresStorage(db *gorm.DB) *UserPostgresStorage {
	return &UserPostgresStorage{db: db}
}

func (s *UserPostgresStorage) CreateUser(user *models.User) error {
	return s.db.Create(user).Error
}

func (s *UserPostgresStorage) GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserPostgresStorage) UpdateUser(id uuid.UUID, updates map[string]interface{}) error {
	return s.db.Model(&models.User{}).Where("id = ?", id).Updates(updates).Error
}

func (s *UserPostgresStorage) DeleteUser(id uuid.UUID) error {
	return s.db.Where("id = ?", id).Delete(&models.User{}).Error
}

func (s *UserPostgresStorage) AddDoc(doc *models.UserDoc) error {
	return s.db.Create(doc).Error
}

func (s *UserPostgresStorage) GetUserDocs(userID uuid.UUID) ([]*models.UserDoc, error) {
	var docs []*models.UserDoc
	err := s.db.Where("user_id = ?", userID).Find(&docs).Error
	return docs, err
}

func (s *UserPostgresStorage) RemoveDoc(userID, docID uuid.UUID) error {
	return s.db.Where("id = ? AND user_id = ?", docID, userID).Delete(&models.UserDoc{}).Error
}

func (s *UserPostgresStorage) UpdateUserAbout(id uuid.UUID, about string) error {
	return s.db.Model(&models.User{}).Where("id = ?", id).Update("about_myself", about).Error
}

func (s *UserPostgresStorage) UpdateUserName(id uuid.UUID, name string) error {
	return s.db.Model(&models.User{}).Where("id = ?", id).Update("name", name).Error
}

/* пример простой обновлять
func (s *UserPostgresStorage) UpdateUserSurname(id uuid.UUID, surname string) error {
	return s.db.Model(&models.User{}).Where("id = ?", id).Update("surname", surname).Error
}*/
