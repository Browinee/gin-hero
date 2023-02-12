package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenExprireDuration = time.Hour * 2

var mySecret = []byte("jwt-key")

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenToken(user_id int64, username string) (string, error) {
	c := MyClaims{
		user_id,
		username,
		jwt.StandardClaims{ExpiresAt: time.Now().Add(TokenExprireDuration).Unix(), Issuer: "gin-hero"}}
	fmt.Printf("MyClaims %v\n", c)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	fmt.Printf("token %v\n", token)
	abc, _ := token.SignedString(mySecret)
	fmt.Printf("token.SignedString(mySecret)%v\n", abc)
	return abc, nil

}

func ParseToken(tokenString string) (*MyClaims, error) {
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
