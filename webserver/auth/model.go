package auth

import "github.com/dgrijalva/jwt-go"

// JWTCustomClaims model for custom claim
type JWTCustomClaims struct {
	Email string `json:"email"`
	ID    int32  `json:"id"`
	jwt.StandardClaims
}

type UserPermissions struct {
	Permissions map[string]bool `json:"permissions"`
	Roles       map[string]bool `json:"roles"`
}
