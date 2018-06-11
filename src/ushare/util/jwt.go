package util

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"ushare/config"
)

type Claims struct {
	Mobile string `json:"mobile"`
	Code   string `json:"code"`
	jwt.StandardClaims
}

func GenerateToken(mobile string, code string) (string, error) {
	key := config.Conf.Read("token", "key")
	expireTime := time.Now().UTC().Add(3 * time.Hour).Unix()

	claims := Claims{
		mobile,
		code,
		jwt.StandardClaims{
			ExpiresAt: expireTime,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(key))

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	key := config.Conf.Read("token", "key")
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func ValidateToken(token string) (bool, interface{}) {
	claims, err := ParseToken(token)
	if err != nil {
		return false, InvalidToken
	}

	expireTime := claims.ExpiresAt
	timeNow := time.Now().UTC().Unix()
	if timeNow > expireTime {
		return false, ExpiredToken
	}

	return true, nil
}
