package handler

import "github.com/labstack/echo/v4"

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthLogin
// @Summary Login
// @Description Get access and refresh token pair by username and password
// @Tags Auth
// @ID auth-login
// @Accept  json
// @Produce  json
// @Param data body LoginRequest true "Data"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} ErrorResponse
// @Router /auth/login [post]
func (c *Controller) AuthLogin(ctx echo.Context) error {
	var req LoginRequest
	err := bindStruct(ctx, &req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	tokenPair, err := c.Service.Auth.Login(req.Username, req.Password)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(200, AuthResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	})
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthRegister
// @Summary Register
// @Description Create an account and get access and refresh token pair
// @Tags Auth
// @ID auth-register
// @Accept  json
// @Produce  json
// @Param data body RegisterRequest true "Data"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} ErrorResponse
// @Router /auth/register [post]
func (c *Controller) AuthRegister(ctx echo.Context) error {
	var req RegisterRequest
	err := bindStruct(ctx, &req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	tokenPair, err := c.Service.Auth.Register(req.Username, req.Password)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(200, AuthResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	})
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// AuthRefresh
// @Summary Refresh Tokens
// @Description Generate a new access and refresh token pair using the refresh token
// @Tags Auth
// @ID auth-refresh
// @Accept  json
// @Produce  json
// @Param data body RefreshRequest true "Data"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} ErrorResponse
// @Router /auth/refresh [post]
func (c *Controller) AuthRefresh(ctx echo.Context) error {
	var req RefreshRequest
	err := bindStruct(ctx, &req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	tokenPair, err := c.Service.Auth.RefreshToken(req.RefreshToken)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(200, AuthResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	})
}
