package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	"os"
	"time"
)

type AuthTokenClaims struct {
	UserID int64  `json:"id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

const ExpiresTime = time.Hour * 1

func CreateToken(userId int64, email string) (string, error) {
	at := AuthTokenClaims{
		UserID: userId,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(ExpiresTime)),
		},
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, &at)

	return tkn.SignedString([]byte(os.Getenv("test")))
}

func GetClaimsByTokenString(tokenString string) (*AuthTokenClaims, error) {
	claims := AuthTokenClaims{}
	_, err1 := jwt.ParseWithClaims(tokenString, &claims,
		func(tkn *jwt.Token) (interface{}, error) {
			if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("not m format")
			}
			return []byte(os.Getenv("test")), nil
		})

	if err1 != nil {
		return nil, err1
	}

	return &claims, nil
}
