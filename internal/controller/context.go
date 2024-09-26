package controller

import (
	"github.com/labstack/echo/v4"
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain/user/model"
)

const (
	UserContextKey = "user"
)

func getUserFromContext(ctx echo.Context) (model.User, error) {
	user, ok := ctx.Get(UserContextKey).(model.User)
	if !ok {
		return model.User{}, apperror.NewServiceError("getUserFromContext: user not found in context", nil)
	}
	return user, nil
}

func addUserToContext(ctx echo.Context, user model.User) {
	ctx.Set(UserContextKey, user)
}
