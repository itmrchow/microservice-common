package token

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var (
	ErrEmptyToken   = errors.New("token is empty")
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// GenerateToken generates a JWT token for a user
func GenerateToken(userID string, secretKey string, issuer string, expireAt int) (tokenStr string, err error) {
	now := time.Now()

	registeredClaims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * time.Duration(expireAt))),
		Issuer:    issuer,
		Subject:   userID,
		Audience:  []string{userID},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, registeredClaims)

	tokenStr, err = token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return
}

// ValidateToken validates a JWT token
func ValidateToken(tokenStr string, secretKey string, issuer string) (userID string, err error) {

	if tokenStr == "" {
		return "", errors.New("token is empty")
	}

	// parse token
	tokenClaims, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	}, jwt.WithLeeway(5*time.Second))

	if err != nil {
		return
	}

	claims, ok := tokenClaims.Claims.(*jwt.RegisteredClaims)

	if !ok && !tokenClaims.Valid {
		return "", errors.New("token is invalid")
	}

	if claims.Subject == "" {
		return "", errors.New("token is invalid")
	}

	if claims.Issuer != issuer {
		return "", errors.New("token is invalid")
	}

	// 檢查 token 是否過期
	if claims.ExpiresAt.Before(time.Now()) {
		return "", errors.New("token has expired")
	}

	userID = claims.Subject

	return
}
