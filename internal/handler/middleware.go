package handler

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"slices"
	"strings"
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"
)

func (c *Controller) requiredRolesMiddleware(roles ...domain.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			user, err := getUserFromContext(ctx)
			if err != nil {
				return c.handleAppError(ctx, err)
			}

			if !slices.Contains(roles, user.Role) {
				return c.handleAppError(ctx, apperror.NewErrUserShoutHaveRole(roles...))
			}

			return next(ctx)
		}
	}
}

func (c *Controller) authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authorization := ctx.Request().Header.Get("Authorization")
		authorizationSplitted := strings.Split(authorization, " ")

		if len(authorizationSplitted) != 2 {
			return c.handleAppError(ctx, apperror.ErrInvalidAuthorizationHeader)
		}
		if authorizationSplitted[0] != "Bearer" {
			return c.handleAppError(ctx, apperror.ErrInvalidAuthorizationHeader)
		}

		user, err := c.Service.Auth.GetUserFromAccessToken(authorizationSplitted[1])
		if err != nil {
			return c.handleAppError(ctx, err)
		}

		addUserToContext(ctx, user)

		return next(ctx)
	}
}

func appErrorTypeToCode(_type int) int {
	switch _type {
	case apperror.ErrorTypeDatabase:
		return http.StatusInternalServerError
	case apperror.ErrorTypeNotFound:
		return http.StatusNotFound
	case apperror.ErrorTypeConflict:
		return http.StatusConflict
	case apperror.ErrorTypeForbidden:
		return http.StatusForbidden
	case apperror.ErrorTypeUnauthorized:
		return http.StatusUnauthorized
	case apperror.ErrorTypeBadRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func (c *Controller) handleAppError(ctx echo.Context, err error) error {
	if apperror.IsAppError(err) {
		var appErr *apperror.AppError
		errors.As(err, &appErr)
		status := appErrorTypeToCode(appErr.Type)

		return ctx.JSON(status, ErrorResponse{
			Error: appErr.Message,
		})
	}

	fmt.Println("Unexpected error:", err)
	return ctx.JSON(http.StatusInternalServerError, ErrorResponse{
		Error: "Internal server error",
	})

}
