package handler

import (
	"net/http"
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"

	"github.com/labstack/echo/v4"
)

type CreateScheduleRequest struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

// CreateSchedule
// @Summary Create Schedule
// @Description Create a new schedule
// @Tags Schedule
// @ID schedule-create
// @Accept  json
// @Produce  json
// @Param data body CreateScheduleRequest true "Data"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /schedules [post]
func (c *Controller) CreateSchedule(ctx echo.Context) error {
	var req CreateScheduleRequest
	err := bindStruct(ctx, &req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}
	user, err := getUserFromContext(ctx)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	id, err := c.Service.Schedule.Create(domain.CreateScheduleDTO{
		UserID: user.ID,
		Slug:   req.Slug,
		Title:  req.Title,
	})
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, NewIDResponse(id))
}

// GetScheduleBySlug
// @Summary Get Schedule By Slug
// @Description Get schedule using slug
// @Tags Schedule
// @ID schedule-get-slug
// @Produce  json
// @Param slug path string true "Schedule Slug"
// @Success 200 {object} domain.ScheduleView
// @Failure 400 {object} ErrorResponse
// @Router /schedules/slug/{slug} [get]
func (c *Controller) GetScheduleBySlug(ctx echo.Context) error {
	slug := ctx.Param("slug")

	if slug == "" {
		return c.handleAppError(ctx, apperror.ErrInvalidSlug)
	}

	schedule, err := c.Service.Schedule.GetBySlug(slug)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, schedule)
}

type GetMySchedulesResponse struct {
	Data       []domain.Schedule `json:"data"`
	Pagination domain.Pagination `json:"pagination"`
}

// GetMySchedules
// @Summary Get Current User Schedules
// @Description Get Current User Schedules
// @Tags Schedule
// @ID schedule-get-my
// @Produce  json
// @Param limit query uint false "Limit"
// @Param offset query uint false "Offset"
// @Success 200 {object} GetMySchedulesResponse
// @Failure 400 {object} ErrorResponse
// @Router /schedules/my [get]
func (c *Controller) GetMySchedules(ctx echo.Context) error {
	user, err := getUserFromContext(ctx)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	limit, offset, err := getPagination(ctx)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	schedules, pagination, err := c.Service.Schedule.GetMy(user.ID, limit, offset)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, GetMySchedulesResponse{
		Data:       schedules,
		Pagination: pagination,
	})
}

type UpdateScheduleRequest struct {
	Slug  *string `json:"slug"`
	Title *string `json:"title"`
}

// UpdateSchedule
// @Summary Update Schedule
// @Description Update schedule
// @Tags Schedule
// @ID schedule-update
// @Accept  json
// @Produce  json
// @Param data body UpdateScheduleRequest true "Data"
// @Param id path uint true "Schedule ID"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /schedules/{id} [patch]
func (c *Controller) UpdateSchedule(ctx echo.Context) error {
	var req UpdateScheduleRequest
	err := bindStruct(ctx, &req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	id, err := parseIDParam(ctx, "id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	err = c.Service.Schedule.Update(user.ID, id, domain.UpdateScheduleDTO{
		Slug:  req.Slug,
		Title: req.Title,
	})
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, NewIDResponse(id))
}

// DeleteSchedule
// @Summary Delete Schedule
// @Description Delete schedule
// @Tags Schedule
// @ID schedule-delete
// @Produce  json
// @Param id path uint true "Schedule ID"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /schedules/{id} [delete]
func (c *Controller) DeleteSchedule(ctx echo.Context) error {
	id, err := parseIDParam(ctx, "id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	err = c.Service.Schedule.Delete(user.ID, id)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, NewIDResponse(id))
}
