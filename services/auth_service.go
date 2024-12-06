package services

import (
	"errors"
	"time"

	"go-mongodb-api/config"
	"go-mongodb-api/repositories"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = config.GetConfig("JWT_SECRET")
var tokenExpiryStr = config.GetConfig("TOKEN_EXPIRY")

type AuthService struct {
	Repo *repositories.UserRepository
}

func NewAuthService(repo *repositories.UserRepository, secret string, expiry time.Duration) *AuthService {
	return &AuthService{
		Repo: repo,
	}
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.Repo.FetchByEmail(username)
	if err != nil {
		return "", err
	}

	if user.Password != password {
		return "", errors.New("invalid username or password")
	}

	token, err := s.generateToken(user.ID.Hex(), user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) generateToken(userID, email string) (string, error) {
	tokenExpiry, err := time.ParseDuration(tokenExpiryStr)
	if err != nil {
		return "", errors.New("invalid token expiry duration")
	}

	claims := jwt.MapClaims{
		"userID": userID,
		"email":  email,
		"exp":    time.Now().Add(tokenExpiry).Unix(),
		"iat":    time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *AuthService) ValidateToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, errors.New("invalid token")
}
