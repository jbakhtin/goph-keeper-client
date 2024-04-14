package models

import "time"

type Session struct {
	ID           *int
	AccessToken  string
	RefreshToken string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}
