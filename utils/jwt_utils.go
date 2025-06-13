package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWT struct {
	signingKey []byte
}

type CustomClaims struct {
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
	jwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{[]byte("jwt-secret-key")}
}

// ParseToken 解析 JWT Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// GenerateToken 生成 JWT Token
func (j *JWT) GenerateToken(userID, userName string) (string, error) {
	expireTime := time.Now().Add(24 * time.Hour)
	claims := CustomClaims{
		UserID:   userID,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "beego-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.signingKey)
}
