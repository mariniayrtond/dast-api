package memory

import (
	"dast-api/internal/domain/model"
	"sync"
)

func NewUserRepository() *UserRepository {
	return &UserRepository{
		mu:      &sync.Mutex{},
		durable: map[string]*model.User{},
	}
}

type UserRepository struct {
	mu      *sync.Mutex
	durable map[string]*model.User
}

func (u UserRepository) Save(user *model.User) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.durable[user.ID] = user
	return nil
}

func (u UserRepository) Get(id string) (*model.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	user, ok := u.durable[id]
	if !ok {
		return nil, nil
	}
	return user, nil
}

func (u UserRepository) SearchByName(name string) (*model.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	for _, user := range u.durable {
		if user.Name == name {
			return user, nil
		}
	}

	return nil, nil
}
