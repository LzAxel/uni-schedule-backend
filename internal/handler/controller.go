package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "uni-schedule-backend/docs"
	"uni-schedule-backend/internal/config"
	"uni-schedule-backend/internal/domain"
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

		schedule := v1api.Group("/schedule")
		{
			schedule.POST("", c.CreateSchedule, c.authMiddleware, c.requiredRolesMiddleware(domain.RoleAdmin))
			schedule.GET("/slug/:slug", c.GetScheduleBySlug)
			schedule.PATCH("/:id", c.UpdateSchedule, c.authMiddleware, c.requiredRolesMiddleware(domain.RoleAdmin, domain.RoleScheduleEditor))
			schedule.DELETE("/:id", c.DeleteSchedule, c.authMiddleware, c.requiredRolesMiddleware(domain.RoleAdmin))

			schedule.POST("/:schedule_id/slot", c.AddPairToSchedule, c.authMiddleware, c.requiredRolesMiddleware(domain.RoleAdmin))
			schedule.DELETE("/slot/:slot_id", c.DeleteSlotFromSchedule, c.authMiddleware, c.requiredRolesMiddleware(domain.RoleAdmin))
			schedule.PATCH("/slot/:slot_id", c.UpdateSlotInSchedule, c.authMiddleware, c.requiredRolesMiddleware(domain.RoleAdmin))
		}

		teacher := v1api.Group("/teacher")
		{
			teacher.POST("", c.CreateTeacher, c.authMiddleware, c.requiredRolesMiddleware(domain.RoleAdmin))
			teacher.PATCH("/:id", c.UpdateTeacher, c.authMiddleware, c.requiredRolesMiddleware(domain.RoleAdmin))
			teacher.DELETE("/:id", c.DeleteTeacher, c.authMiddleware, c.requiredRolesMiddleware(domain.RoleAdmin))
		}

		lesson := v1api.Group("/lesson")
		{
			lesson.POST("", c.CreateLesson, c.authMiddleware, c.requiredRolesMiddleware(domain.RoleAdmin))
			lesson.PATCH("/:id", c.UpdateLesson, c.authMiddleware, c.requiredRolesMiddleware(domain.RoleAdmin))
			lesson.DELETE("/:id", c.DeleteLesson, c.authMiddleware, c.requiredRolesMiddleware(domain.RoleAdmin))
		}
		v1api.GET("/swagger/*", echoSwagger.WrapHandler)
	}
}

func (c *Controller) Start() {
	c.init()
	err := c.server.Start(":" + config.GetConfig().AppConfig.Port)
	if err != nil {
		panic(fmt.Errorf("failed to start server: %v", err))
	}
}
