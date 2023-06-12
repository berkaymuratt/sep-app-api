package services

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

//go:generate mockgen -destination=../mocks/service/mockJwtService.go -package=services github.com/berkaymuratt/sep-app-api/services JwtServiceI
type JwtServiceI interface {
	GenerateJwtToken(userId string) (string, error)
	CheckJwt(jwtToken string) (string, error)
}

type JwtService struct {
	JwtServiceI
	secretKey string
}

func NewJwtService() JwtService {
	return JwtService{
		secretKey: "secret",
	}
}

func (service JwtService) GenerateJwtToken(userId string) (string, error) {

	expiresAt := time.Now().Add(time.Hour * 24)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    userId,
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	})

	return claims.SignedString([]byte(service.secretKey))
}

func (service JwtService) CheckJwt(jwtToken string) (string, error) {

	token, err := jwt.ParseWithClaims(jwtToken, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(service.secretKey), nil
	})

	if err != nil {
		return "", err
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	return claims.Issuer, nil
}
