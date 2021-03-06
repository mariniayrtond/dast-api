package usecase

import (
	"crypto/sha256"
	"dast-api/internal/config"
	"dast-api/internal/domain/model"
	"dast-api/internal/domain/repository"
	"dast-api/pkg/uid"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserUseCase interface {
	RegisterUser(string, string, string) (*model.User, error)
	LogIn(string, string) (string, error)
	AlreadyLogIn(string, string) error
	GetUser(id string) (*model.User, error)
}

func NewUserUseCase(uRepo repository.UserRepository, tokenRepo repository.TokenRepository) UserUseCase {
	return userCRUDImpl{
		uRepo:     uRepo,
		tokenRepo: tokenRepo,
	}
}

type userCRUDImpl struct {
	uRepo     repository.UserRepository
	tokenRepo repository.TokenRepository
}

func (u userCRUDImpl) AlreadyLogIn(username string, token string) error {
	t, err := u.tokenRepo.Get(token)
	if t == nil || err != nil {
		return errors.New("unauthorized")
	}

	user, err := u.uRepo.Get(t.UserID)
	if err != nil {
		return errors.New("unauthorized")
	}

	if user.Name != username {
		return errors.New("unauthorized")
	}

	return nil
}

func (u userCRUDImpl) LogIn(name string, password string) (string, error) {
	user, err := u.uRepo.SearchByName(name)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", fmt.Errorf("user %s does not exists, you have to create it before login", name)
	}

	if errHashing := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); errHashing != nil {
		return "", errors.New("user/password don't exist or is not matching")
	}

	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s-%s-%s-%s-%d", user.Name, user.Password, user.Email, user.ID, time.Now().YearDay())))

	tokenLogIn := model.NewLogIn(user.ID, fmt.Sprintf("%x", h.Sum(nil)))

	possibleToken, err := u.tokenRepo.Get(tokenLogIn.Token)
	if err != nil {
		return "", errors.New("user/password don't exist or is not matching")
	}

	if possibleToken != nil {
		return possibleToken.Token, nil
	}

	if err := u.tokenRepo.Create(tokenLogIn, config.TTLForToken); err != nil {
		return "", errors.New(fmt.Sprintf("error generating token for user:%s", name))
	}

	return tokenLogIn.Token, nil
}

func (u userCRUDImpl) RegisterUser(name string, email string, password string) (*model.User, error) {
	id, err := uid.GenerateUUID()
	if err != nil {
		return nil, err
	}

	user := model.NewUser(id, name, email)
	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(encryptedPass)
	errInsert := u.uRepo.Save(user)

	return user, errInsert
}

func (u userCRUDImpl) GetUser(id string) (*model.User, error) {
	return u.GetUser(id)
}
