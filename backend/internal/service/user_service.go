package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/kerilOvs/backend/internal/models"
	"github.com/kerilOvs/backend/internal/storage"
)

type UserService struct {
	storage storage.UserStorage
}

func NewUserService(storage storage.UserStorage) *UserService {
	return &UserService{storage: storage}
}

func (s *UserService) CreateUser(id uuid.UUID, name, surname, email, password string) (*models.User, error) {
	if name == "" || surname == "" {
		return nil, errors.New("name and surname are required")
	}

	user := &models.User{
		ID:        id, // в зависимости от логики, или создаем, или присваиваем
		Name:      name,
		Surname:   surname,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}

	if err := s.storage.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserByID(id uuid.UUID) (*models.User, error) {
	user, err := s.storage.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Загружаем связанные данные
	docs, err := s.storage.GetUserDocs(id)
	if err != nil {
		return nil, err
	}

	// Преобразуем []*UserPhoto → []UserPhoto
	docVals := make([]models.UserDoc, len(docs))
	for i, p := range docs {
		if p != nil {
			docVals[i] = *p
		}
	}

	user.Documents = docVals

	return user, nil
}

func (s *UserService) UpdateUserProfile(id uuid.UUID, updates models.UserProfileUpdate) error {
	updateFields := make(map[string]interface{})

	if updates.Name != nil {
		if *updates.Name == "" {
			return errors.New("name cannot be empty")
		}
		updateFields["name"] = *updates.Name
	}

	if updates.Surname != nil {
		if *updates.Surname == "" {
			return errors.New("surname cannot be empty")
		}
		updateFields["surname"] = *updates.Surname
	}

	if updates.BirthDate != nil {
		updateFields["birth_date"] = *updates.BirthDate
	}

	if updates.Email != nil {
		updateFields["email"] = *updates.Email
	}

	if updates.Password != nil {
		updateFields["password"] = *updates.Password
	}

	if len(updateFields) > 0 {
		return s.storage.UpdateUser(id, updateFields)
	}

	return nil
}

func (s *UserService) AddUserDoc(userID uuid.UUID, docURL string) (*models.UserDoc, error) {
	if docURL == "" {
		return nil, errors.New("photo URL cannot be empty")
	}

	doc := &models.UserDoc{
		ID:     uuid.New(),
		UserID: userID,
		URL:    docURL,
	}

	if err := s.storage.AddDoc(doc); err != nil {
		return nil, err
	}

	return doc, nil
}

func (s *UserService) DeleteUser(id uuid.UUID) error {
	return s.storage.DeleteUser(id)
}

func (s *UserService) GetUserDocs(userID uuid.UUID) ([]*models.UserDoc, error) {
	return s.storage.GetUserDocs(userID)
}

func (s *UserService) RemoveUserDoc(userID, photoID uuid.UUID) error {

	return s.storage.RemoveDoc(userID, photoID)
}
