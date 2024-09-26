package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"uni-schedule-backend/internal/config"
	"uni-schedule-backend/internal/service"
)

type Controller struct {
	Service *service.Service
	server  *echo.Echo
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		Service: service,
		server:  echo.New(),
	}
}

func (c *Controller) init() {
	c.initServer()
	c.initRoutes()
}

func (c *Controller) initServer() {
	cfg := config.GetConfig()

	c.server.Debug = cfg.AppConfig.IsDebug
	c.server.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORSWithConfig(middleware.DefaultCORSConfig),
	)
}

func (c *Controller) initRoutes() {
	v1api := c.server.Group("/api/v1")
	{
		auth := v1api.Group("/auth")
		{
			auth.POST("/login", c.AuthLogin)
			auth.POST("/register", c.AuthRegister)
			auth.POST("/refresh", c.AuthRefresh)
		}

		user := v1api.Group("/user")
		{
			user.POST("/test", c.UserTest, c.authMiddleware)
		}
	}
}

func (c *Controller) Start() {
	c.init()
	err := c.server.Start(":" + config.GetConfig().AppConfig.Port)
	if err != nil {
		panic(fmt.Errorf("failed to start server: %v", err))
	}
}
