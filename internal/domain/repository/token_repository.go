package repository

import (
	"dast-api/internal/domain/model"
	"time"
)

type TokenRepository interface {
	Create(token *model.LogIn, ttl time.Duration) error
	Get(token string) (*model.LogIn, error)
}
