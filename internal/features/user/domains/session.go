package domains

import (
	"api/internal/features/user/models"
)

type Session struct {
	Id     uint   `json:"id"`
	UserId uint   `json:"user_id"`
	Token  string `json:"token"`
}

func (s *Session) ToModel() *models.Session {
	return &models.Session{
		UserId: s.UserId,
		Token:  s.Token,
	}
}

func FromSessionModel(m *models.Session) *Session {
	return &Session{
		Id:     m.ID,
		UserId: m.UserId,
		Token:  m.Token,
	}
}
