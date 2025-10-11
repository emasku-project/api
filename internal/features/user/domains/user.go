package domains

import (
	"api/internal/features/user/models"
)

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) ToModel() *models.User {
	return &models.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}

func FromUserModel(m *models.User) *User {
	return &User{
		Id:       m.ID,
		Name:     m.Name,
		Email:    m.Email,
		Password: m.Password,
	}
}
