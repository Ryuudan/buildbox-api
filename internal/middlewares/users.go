package middlewares

import (
	"github.com/Pyakz/buildbox-api/internal/services"
)

type UserMiddleware struct {
	userService services.UserService
}

func NewUserMiddleware(userService services.UserService) *UserMiddleware {
	return &UserMiddleware{
		userService: userService,
	}
}
