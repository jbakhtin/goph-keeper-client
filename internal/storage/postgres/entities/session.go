package entities

import (
	"time"

	"github.com/jbakhtin/goph-keeper-client/internal/appmodules/auth/domain/models"
)

type Session struct {
	ID           *int       `json:"id,omitempty"`
	AccessToken  string     `json:"access_token,omitempty"`
	RefreshToken string     `json:"refresh_token,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}

func NewEntity(model models.Session) Session {
	return Session{
		ID:           nil,
		AccessToken:  model.AccessToken,
		RefreshToken: model.RefreshToken,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
	}
}

func (s *Session) ToModel() models.Session {
	return models.Session{
		ID:           s.ID,
		RefreshToken: s.RefreshToken,
		AccessToken:  s.AccessToken,
		CreatedAt:    s.CreatedAt,
		UpdatedAt:    s.UpdatedAt,
	}
}
