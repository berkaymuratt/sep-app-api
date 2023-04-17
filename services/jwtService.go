package services

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtService struct {
	secretKey string
}

func NewJwtService() JwtService {
	return JwtService{
		secretKey: "secret",
	}
}

func (service JwtService) GenerateJwtToken(userId string) (string, error) {

	expiresAt := time.Now().Add(time.Hour * 1)

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
