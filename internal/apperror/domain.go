package apperror

var (
	ErrInvalidLoginOrPassword     = New(ErrorTypeForbidden, "Invalid login or password", nil)
	ErrInvalidAuthorizationHeader = New(ErrorTypeUnauthorized, "Invalid authorization header", nil)
	ErrInvalidAccessToken         = New(ErrorTypeUnauthorized, "Invalid access token", nil)
	ErrInvalidRefreshToken        = New(ErrorTypeUnauthorized, "Invalid refresh token", nil)
	ErrAccessTokenIsExpired       = New(ErrorTypeUnauthorized, "Access token is expired", nil)
	ErrRefreshTokenIsExpired      = New(ErrorTypeUnauthorized, "Refresh token is expired", nil)

	ErrUserNotFound         = New(ErrorTypeNotFound, "User not found", nil)
	ErrUsernameAlreadyTaken = New(ErrorTypeConflict, "Username already taken", nil)
)
