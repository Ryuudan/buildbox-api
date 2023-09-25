package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3"`
}

type CustomClaims struct {
	AccountID int       `json:"account_id"`
	UserID    int       `json:"user_id"`
	UserUUID  uuid.UUID `json:"user_uuid"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	jwt.StandardClaims
}

var ContextKeyClaims = contextKey("user")

type contextKey string
