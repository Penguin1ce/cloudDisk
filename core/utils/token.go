package utils

import (
	"cloudDisk/core/define"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(id int, identity, name string) (string, error) {
	// id
	// identity
	// name
	now := time.Now()
	uc := define.UserClaims{
		Id:       id,
		Identity: identity,
		Name:     name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(240 * time.Hour)), // 令牌 240 小时后过期
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Subject:   identity,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	signedString, err := token.SignedString([]byte(define.JWTSecret))
	if err != nil {
		return "", err
	}
	return signedString, nil
}

// AnalyzeToken Token 解析保存
func AnalyzeToken(token string) (*define.UserClaims, error) {
	uc := new(define.UserClaims)
	jwtToken, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if !jwtToken.Valid {
		return uc, errors.New("invalid token")
	}
	return uc, nil
}
