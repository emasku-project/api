package domains

import (
	"api/internal/features/general/models"
)

type Setting struct {
	Id     uint   `json:"id"`
	UserId uint   `json:"user_id"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}

func (s *Setting) ToModel() *models.Setting {
	return &models.Setting{
		UserId: s.UserId,
		Key:    s.Key,
		Value:  s.Value,
	}
}

func FromSettingModel(m *models.Setting) *Setting {
	return &Setting{
		Id:     m.ID,
		UserId: m.UserId,
		Key:    m.Key,
		Value:  m.Value,
	}
}
