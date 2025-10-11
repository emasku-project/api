package repositories

import (
	"api/internal/features/user/domains"
	"gorm.io/gorm"
)

type Session struct {
	db *gorm.DB
}

func NewSession(
	db *gorm.DB,
) *Session {
	return &Session{
		db: db,
	}
}

func (r *Session) Create(data domains.Session) (*domains.Session, error) {
	session := data.ToModel()
	if err := r.db.Create(&session).Error; err != nil {
		return nil, err
	} else {
		return domains.FromSessionModel(session), nil
	}
}

func (r *Session) CreateInTx(tx *gorm.DB, data domains.Session) (*domains.Session, error) {
	session := data.ToModel()
	if err := tx.Create(&session).Error; err != nil {
		return nil, err
	} else {
		return domains.FromSessionModel(session), nil
	}
}

func (r *Session) DeleteByToken(token string) error {
	return r.db.Where("token = ?", token).Unscoped().Delete(&domains.Session{}).Error
}
