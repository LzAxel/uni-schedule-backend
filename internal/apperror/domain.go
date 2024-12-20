package apperror

import (
	"fmt"
	"strings"
	"uni-schedule-backend/internal/domain"
)

var (
	ErrInvalidInputBody = New(ErrorTypeBadRequest, "Invalid input body", nil)

	ErrInvalidLoginOrPassword     = New(ErrorTypeForbidden, "Invalid login or password", nil)
	ErrInvalidAuthorizationHeader = New(ErrorTypeUnauthorized, "Invalid authorization header", nil)
	ErrInvalidAccessToken         = New(ErrorTypeUnauthorized, "Invalid access token", nil)
	ErrInvalidRefreshToken        = New(ErrorTypeBadRequest, "Invalid refresh token", nil)
	ErrAccessTokenIsExpired       = New(ErrorTypeUnauthorized, "Access token is expired", nil)
	ErrRefreshTokenIsExpired      = New(ErrorTypeUnauthorized, "Refresh token is expired", nil)

	ErrUserNotFound         = New(ErrorTypeNotFound, "User not found", nil)
	ErrUsernameAlreadyTaken = New(ErrorTypeConflict, "Username already taken", nil)

	ErrSlugAlreadyInUse = New(ErrorTypeConflict, "Slug already in use", nil)
	ErrInvalidSlug      = New(ErrorTypeBadRequest, "Invalid slug", nil)
	ErrInvalidTitle     = New(ErrorTypeBadRequest, "Invalid title", nil)

	ErrDontHavePermission = New(ErrorTypeForbidden, "You don't have permission", nil)

	ErrScheduleNotFound     = New(ErrorTypeNotFound, "Schedule not found", nil)
	ErrEmptyShortName       = New(ErrorTypeBadRequest, "Short name is empty", nil)
	ErrEmptyScheduleSlug    = New(ErrorTypeBadRequest, "Schedule slug is empty", nil)
	ErrEmptyScheduleName    = New(ErrorTypeBadRequest, "Schedule name is empty", nil)
	ErrInvalidWeekdayNumber = New(ErrorTypeBadRequest, "Invalid weekday number. Should be from 0 to 6", nil)

	ErrClassWithSamePositionAlreadyExists = New(ErrorTypeConflict, "Class with same position already exists", nil)
	ErrSingleClassAlreadySet              = New(ErrorTypeConflict, "Single class already set", nil)
	ErrCannotSetSeveralSingleClasses      = New(ErrorTypeConflict, "Cannot set several single classes", nil)
	ErrCannotSetSameClassesPositions      = New(ErrorTypeConflict, "Cannot set same classes positions", nil)
	ErrClassesAlreadySet                  = New(ErrorTypeConflict, "Classes already set", nil)
	ErrClassOnlyEvenOrSingleClass         = New(ErrorTypeConflict, "Can set only even(odd) or single class at a time", nil)
)

func NewErrUserShoutHaveRole(role ...domain.Role) error {
	formattedRoles := make([]string, 0, len(role))
	for _, r := range role {
		formattedRoles = append(formattedRoles, r.String())
	}
	return New(ErrorTypeForbidden, fmt.Sprintf("User should have one of roles: %s", strings.Join(formattedRoles, ", ")), nil)
}

func NewErrInvalidQueryParam(name string) error {
	return New(ErrorTypeBadRequest, fmt.Sprintf("Invalid query param: %s", name), nil)
}
