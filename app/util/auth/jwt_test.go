package auth

import (
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	expiresTime := time.Now().Local().Add(5 * time.Minute)

	token := Generate(&JWTPayload{
		UserId:   "123",
		UserName: "carl",
	}, expiresTime)
	t.Log(token)
}

func TestCheck(t *testing.T) {
	expiresTime := time.Now().Local().Add(5 * time.Minute)

	token := Generate(&JWTPayload{
		UserId:   "123",
		UserName: "carl",
	}, expiresTime)

	t.Log(token)
	t.Log(Check(token))
}
