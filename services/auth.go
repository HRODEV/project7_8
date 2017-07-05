package services

import (
	"encoding/base64"
)

type AuthService struct {
}

func (authService *AuthService) DecodeAuthorzationHeader(value string) string {
	decoded, _ := base64.URLEncoding.DecodeString(value)

	return string(decoded)
}
