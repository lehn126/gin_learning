package auth

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var defaultSigningKey []byte = []byte("mySigningKey")

type JWTClaims struct {
	UserName string
	UserId   string
	jwt.RegisteredClaims
}

type JWTPayload struct {
	UserName string
	UserId   string
}

func NewJWTClaims(payload *JWTPayload, expiresTime time.Time) *JWTClaims {
	now := jwt.NewNumericDate(time.Now())
	return &JWTClaims{
		payload.UserName,
		payload.UserId,
		jwt.RegisteredClaims{
			IssuedAt:  now,
			NotBefore: now,
			ExpiresAt: jwt.NewNumericDate(expiresTime),
			// Issuer:    "test",
			// Subject:   "something",
			// ID:        "1",
			// Audience:  []string{"somebody_else"},
		},
	}
}

func getSigningKey() *[]byte {
	// TODO..., try to get signing key from config file or env
	return &defaultSigningKey
}

func Generate(payload *JWTPayload, expiresTime time.Time) string {
	if expiresTime.IsZero() {
		expiresTime = time.Now().Add(time.Hour)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, NewJWTClaims(payload, expiresTime))

	ss, err := token.SignedString(*getSigningKey()) // can't use pointer here
	if err != nil {
		panic(fmt.Sprintf("meet error when generate jwt token, error: %v", err))
	}
	return ss
}

func Check(token string) bool {
	jwtToken, err := jwt.ParseWithClaims(token, new(JWTClaims), func(token *jwt.Token) (interface{}, error) {
		return *getSigningKey(), nil // can't use pointer here
	})

	if claims, ok := jwtToken.Claims.(*JWTClaims); ok && jwtToken.Valid {
		log.Printf("jwt check passed: %v %v", claims.UserName, claims.RegisteredClaims.IssuedAt.Time)
		return true
	} else {
		/*
			jwt.ErrTokenMalformed
			jwt.ErrTokenSignatureInvalid
			jwt.ErrTokenExpired
			jwt.ErrTokenNotValidYet
		*/
		log.Printf("meet error when check jwt token, error: %v\n", err)
		return false
	}
}
