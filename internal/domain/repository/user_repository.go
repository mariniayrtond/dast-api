package repository

import "dast-api/internal/domain/model"

//go:generate moq -out user_repository_moq.go . UserRepository

type UserRepository interface {
	Save(user *model.User) error
	Get(id string) (*model.User, error)
	SearchByName(name string) (*model.User, error)
}
