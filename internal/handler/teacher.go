package handler

import (
	"net/http"
	"uni-schedule-backend/internal/domain"

	"github.com/labstack/echo/v4"
)

// CreateTeacher
// @Summary Create Teacher
// @Description Create a new teacher
// @Tags Teacher
// @ID teacher-create
// @Accept  json
// @Produce  json
// @Param data body domain.TeacherCreateDTO true "Data"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /teachers [post]
func (c *Controller) CreateTeacher(ctx echo.Context) error {
	var req domain.TeacherCreateDTO
	err := bindStruct(ctx, &req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	id, err := c.Service.Teacher.Create(req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, NewIDResponse(id))
}

// UpdateTeacher
// @Summary Update Teacher
// @Description Update teacher
// @Tags Teacher
// @ID teacher-update
// @Produce  json
// @Param id path uint true "Teacher ID"
// @Param data body domain.TeacherUpdateDTO true "Data"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /teachers/{id} [patch]
func (c *Controller) UpdateTeacher(ctx echo.Context) error {
	id, err := parseIDParam(ctx, "id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	var req domain.TeacherUpdateDTO
	err = bindStruct(ctx, &req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	err = c.Service.Teacher.Update(user.ID, id, req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, NewIDResponse(id))
}

// DeleteTeacher
// @Summary Delete Teacher
// @Description Delete teacher
// @Tags Teacher
// @ID teacher-delete
// @Produce  json
// @Param id path uint true "Teacher ID"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /teachers/{id} [delete]
func (c *Controller) DeleteTeacher(ctx echo.Context) error {
	id, err := parseIDParam(ctx, "id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	err = c.Service.Teacher.Delete(user.ID, id)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, NewIDResponse(id))
}

type GetTeahersResponse struct {
	Data       []domain.Teacher  `json:"data"`
	Pagination domain.Pagination `json:"pagination"`
}

// GetScheduleTeachers
// @Summary Get Schedule's Teachers
// @Description Get Schedule's Teachers
// @Tags Teacher
// @ID teacher-get-all
// @Produce  json
// @Param schedule_id path uint true "Schedule ID"
// @Param limit query uint false "Limit"
// @Param offset query uint false "Offset"
// @Success 200 {object} GetTeahersResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /schedules/{schedule_id}/teachers [get]
func (c *Controller) GetScheduleTeachers(ctx echo.Context) error {
	scheduleID, err := parseIDParam(ctx, "schedule_id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}
	limit, offset, err := getPagination(ctx)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	teachers, pagination, err := c.Service.Teacher.GetAll(scheduleID, limit, offset)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, GetTeahersResponse{
		Data:       teachers,
		Pagination: pagination,
	})
}

// GetTeacher
// @Summary Get Schedule's Teacher
// @Description Get Schedule's Teacher
// @Tags Teacher
// @ID teacher-get-id
// @Produce  json
// @Param id path uint true "Teacher ID"
// @Success 200 {object} domain.Teacher
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /teachers/{id} [get]
func (c *Controller) GetTeacher(ctx echo.Context) error {
	id, err := parseIDParam(ctx, "id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	teachers, err := c.Service.Teacher.GetByID(id)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, teachers)
}
