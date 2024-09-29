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

func (c *Controller) AuthLogin(ctx echo.Context) error {
	var req LoginRequest
	err := ctx.Bind(&req)
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

func (c *Controller) AuthRegister(ctx echo.Context) error {
	var req RegisterRequest
	err := ctx.Bind(&req)
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

func (c *Controller) AuthRefresh(ctx echo.Context) error {
	var req RefreshRequest
	err := ctx.Bind(&req)
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
