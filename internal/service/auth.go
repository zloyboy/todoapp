package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	todo "github.com/zloyboy/todoapp"
	"github.com/zloyboy/todoapp/internal/reposit"
)

const (
	salt      = "qwerty123"
	signInKey = "qwertyuiop"
	tokenTTL  = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

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

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signInKey))
}

func (s *AuthService) generatePasswordHash(pswd string) string {
	hash := sha1.New()
	hash.Write([]byte(pswd))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
