package middlewares

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"ushare/config"
	"ushare/helpers"
	"ushare/db"
)

type Claims struct {
	Mobile  string `json:"mobile"`
	Captcha string `json:"captcha"`
	jwt.StandardClaims
}

func GenerateToken(user *db.User) (string, error) {
	key := config.TokenKey
	expireTime := time.Now().UTC().Add(3 * time.Hour).Unix()

	claims := Claims{
		user.Mobile,
		user.Captcha,
		jwt.StandardClaims{
			ExpiresAt: expireTime,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(key))

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	key := config.TokenKey
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
		return false, helpers.InvalidToken
	}

	expireTime := claims.ExpiresAt
	timeNow := time.Now().UTC().Unix()
	if timeNow > expireTime {
		return false, helpers.ExpiredToken
	}

	return true, nil
}
