package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenExprireDuration = time.Hour * 2

var mySecret = []byte("jwt-key")

type MyClaims struct {
	UserID int64 `json:"user_id"`
	// Username string `json:"username"`
	jwt.StandardClaims
}

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return mySecret, nil
}
func GenTokenWithRefreshAndAccess(userID int64) (accessToken, refreshToken string, err error) {
	c := MyClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExprireDuration).Unix(),
			Issuer:    "gin-hero",
		}}
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
		Issuer:    "gin-hero",
	}).SignedString(mySecret)

	return
}

func GenToken(userID int64) (accessToken string, err error) {
	c := MyClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExprireDuration).Unix(),
			Issuer:    "gin-hero",
		}}
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)
	return
}
func ParseToken(tokenString string) (claims *MyClaims, err error) {
	var token *jwt.Token
	claims = new(MyClaims)
	// NOTE: parse tokenString to claims
	token, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		err = errors.New("invalid token")
	}
	return
}

func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	if _, err = jwt.Parse(rToken, keyFunc); err != nil {
		return
	}

	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)

	if v.Errors == jwt.ValidationErrorExpired {
		return GenTokenWithRefreshAndAccess(claims.UserID)
	}
	return
}
