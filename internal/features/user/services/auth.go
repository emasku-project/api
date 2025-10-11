package services

import (
	"errors"
	"net/http"

	"api/internal/features/user/domains"
	"api/internal/features/user/dto/requests"
	"api/internal/features/user/dto/responses"
	"api/internal/features/user/repositories"
	"api/pkg/http/failure"
	"api/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Auth struct {
	userRepo    *repositories.User
	sessionRepo *repositories.Session
	db          *gorm.DB
}

func NewAuth(
	userRepo *repositories.User,
	sessionRepo *repositories.Session,
	db *gorm.DB,
) *Auth {
	return &Auth{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		db:          db,
	}
}

func (s *Auth) Login(req requests.Login) (*responses.Login, *failure.App) {
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, failure.New(
				http.StatusBadRequest,
				"Alamat email atau kata sandi tidak valid!",
				err,
			)
		}
		return nil, failure.NewInternal(err)
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	); err != nil {
		return nil, failure.New(
			http.StatusBadRequest,
			"Alamat email atau kata sandi tidak valid!",
			err,
		)
	}

	if session, err := s.sessionRepo.Create(
		domains.Session{
			UserId: user.Id,
			Token:  uuid.NewString(),
		},
	); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.Login{
			Token: session.Token,
		}, nil
	}
}

func (s *Auth) Register(req requests.Register) (*responses.Register, *failure.App) {
	var customErr *failure.App = nil
	var token string
	if err := s.db.Transaction(
		func(tx *gorm.DB) error {
			password, err := bcrypt.GenerateFromPassword(
				[]byte(req.Password),
				bcrypt.DefaultCost,
			)
			if err != nil {
				return err
			}

			user, err := s.userRepo.CreateInTx(
				tx,
				domains.User{
					Name:     req.Name,
					Email:    req.Email,
					Password: string(password),
				},
			)
			if err != nil {
				if errors.Is(err, gorm.ErrDuplicatedKey) {
					customErr = failure.New(
						http.StatusBadRequest,
						"Alamat email sudah terdaftar!",
						err,
					)
				}
				return err
			}

			if session, err := s.sessionRepo.CreateInTx(
				tx, domains.Session{
					UserId: user.Id,
					Token:  uuid.NewString(),
				},
			); err != nil {
				return err
			} else {
				token = session.Token
			}

			return nil
		},
	); err != nil {
		if customErr != nil {
			return nil, customErr
		}
		return nil, failure.NewInternal(err)
	}

	return &responses.Register{
		Token: token,
	}, nil
}

func (s *Auth) Logout(c *gin.Context) (*responses.Logout, *failure.App) {
	if session, err := utils.GetAuthenticatedSession(c); err != nil {
		return nil, failure.NewUnauthorized()
	} else {
		if err := s.sessionRepo.DeleteByToken(session.Token); err != nil {
			return nil, failure.NewInternal(err)
		}

		return &responses.Logout{
			Message: "oke",
		}, nil
	}
}
