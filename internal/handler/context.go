package handler

import (
	"github.com/labstack/echo/v4"
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"
)

const (
	UserContextKey = "user"
)

func getUserFromContext(ctx echo.Context) (domain.User, error) {
	user, ok := ctx.Get(UserContextKey).(domain.User)
	if !ok {
		return domain.User{}, apperror.NewServiceError("getUserFromContext: user not found in context", nil)
	}
	return user, nil
}

func addUserToContext(ctx echo.Context, user domain.User) {
	ctx.Set(UserContextKey, user)
}
