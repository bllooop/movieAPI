package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	movieapi "movieapi"
	"movieapi/pkg/repository"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	salt     = "slfnoinrf90h390fnviofkl"
	tokenTTL = 12 * time.Hour
)

var jwtKey = []byte("secret_key")

type tokenClaims struct {
	jwt.StandardClaims
	UserRole string `json:"user_role"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user movieapi.User) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *AuthService) ParseToken(accesstok string) (string, error) {
	token, err := jwt.ParseWithClaims(accesstok, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return jwtKey, nil
	})
	if err != nil {
		return "", nil
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserRole, nil
}

func passHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) CreateToken(shopname, password string) (string, error) {
	user, err := s.repo.SignUser(shopname, password)
	if err != nil {
		return "", err
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["user_role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

/*

func (s *AuthService) CreateToken(shopname, password string) (string, error) {
	user, err := s.repo.SignUser(shopname, password)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		shop.Id,
	})
	return token.SignedString([]byte(signingKey))
} */
