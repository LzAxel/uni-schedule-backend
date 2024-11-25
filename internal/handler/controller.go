package handler

import (
	"fmt"
	_ "uni-schedule-backend/docs"
	"uni-schedule-backend/internal/config"
	"uni-schedule-backend/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
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

		schedule := v1api.Group("/schedules")
		{
			schedule.GET("/slug/:slug", c.GetScheduleBySlug)
			schedule.GET("/my", c.GetMySchedules, c.authMiddleware)
			schedule.POST("", c.CreateSchedule, c.authMiddleware)
			schedule.PATCH("/:id", c.UpdateSchedule, c.authMiddleware)
			schedule.DELETE("/:id", c.DeleteSchedule, c.authMiddleware)

			schedule.GET("/:schedule_id/teachers", c.GetScheduleTeachers, c.authMiddleware)
			schedule.GET("/:schedule_id/subjects", c.GetScheduleSubjects, c.authMiddleware)
		}

		teacher := v1api.Group("/teachers")
		{
			teacher.POST("", c.CreateTeacher, c.authMiddleware)
			teacher.GET("/:id", c.GetTeacher, c.authMiddleware)
			teacher.PATCH("/:id", c.UpdateTeacher, c.authMiddleware)
			teacher.DELETE("/:id", c.DeleteTeacher, c.authMiddleware)
		}

		subjects := v1api.Group("/subjects")
		{
			subjects.GET("/:id", c.GetSubject, c.authMiddleware)
			subjects.POST("", c.CreateSubject, c.authMiddleware)
			subjects.PATCH("/:id", c.UpdateSubject, c.authMiddleware)
			subjects.DELETE("/:id", c.DeleteSubject, c.authMiddleware)
		}

		classes := v1api.Group("/classes")
		{
			classes.GET("/:id", c.GetClass, c.authMiddleware)
			classes.POST("", c.CreateClass, c.authMiddleware)
			classes.PATCH("/:id", c.UpdateClass, c.authMiddleware)
			classes.DELETE("/:id", c.DeleteClass, c.authMiddleware)
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
