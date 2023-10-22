package util

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/pramudyas69/Go-Grpc/server/internal/config"
	"time"
)

type CustomClaims struct {
	UserID string
	jwt.RegisteredClaims
}

func GenerateToken(cnf *config.Config, id string, exp int64) (string, error) {
	expired := time.Now().Add(time.Minute * time.Duration(exp))

	claims := &CustomClaims{
		UserID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expired),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(cnf.Token.Access_Token))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateToken(cnf *config.Config, token string) (*CustomClaims, error) {
	tokenStr, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cnf.Token.Access_Token), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tokenStr.Claims.(*CustomClaims); ok && tokenStr.Valid {
		return claims, nil
	}

	return nil, err
}
