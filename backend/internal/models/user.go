package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Surname   string     `json:"surname"`
	BirthDate *time.Time `json:"birth_date,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Documents []UserDoc  `gorm:"foreignKey:UserID" json:"docs,omitempty"`
}

type UserDoc struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	URL    string    `json:"url"`
}

type UserProfileUpdate struct {
	Name      *string    `json:"name,omitempty"`
	Surname   *string    `json:"surname,omitempty"`
	BirthDate *time.Time `json:"birth_date,omitempty"`
	Email     *string    `json:"email"`
	Password  *string    `json:"password"`
}
