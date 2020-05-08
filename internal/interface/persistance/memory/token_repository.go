package memory

import (
	"dast-api/internal/domain/model"
	"sync"
	"time"
)

func NewTokenRepository() *TokenRepository {
	return &TokenRepository{
		mu:      &sync.Mutex{},
		durable: map[string]*model.LogIn{},
	}
}

type TokenRepository struct {
	mu      *sync.Mutex
	durable map[string]*model.LogIn
}

func (t TokenRepository) Create(token *model.LogIn, ttl time.Duration) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.durable[token.Token] = token
	return nil
}

func (t TokenRepository) Get(token string) (*model.LogIn, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	res, ok := t.durable[token]
	if !ok {
		return nil, nil
	}
	return res, nil
}
