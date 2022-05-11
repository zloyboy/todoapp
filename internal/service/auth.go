package service

import (
	"crypto/sha1"
	"fmt"

	todo "github.com/zloyboy/todoapp"
	"github.com/zloyboy/todoapp/internal/reposit"
)

const salt = "qwerty123"

type AuthService struct {
	repo reposit.Authorization
}

func NewAuthService(repo reposit.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(pswd string) string {
	hash := sha1.New()
	hash.Write([]byte(pswd))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
