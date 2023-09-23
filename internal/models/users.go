package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3"`
}

type CreateUserStruct struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=3"`
	FirstName   string `json:"first_name" validate:"required,min=1,max=200"`
	LastName    string `json:"last_name" validate:"required,min=1,max=200"`
	MiddleName  string `json:"middle_name" validate:"required,min=1,max=200"`
	Birthday    string `json:"birthday,omitempty" validate:"omitempty,datetime"`
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
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
