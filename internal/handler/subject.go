package handler

import (
	"net/http"
	"uni-schedule-backend/internal/domain"

	"github.com/labstack/echo/v4"
)

type GetSubjectsResponse struct {
	Data       []domain.Subject  `json:"data"`
	Pagination domain.Pagination `json:"pagination"`
}

// GetScheduleSubjects
// @Summary Get Schedule's Subjects
// @Description Get Schedule's Subjects
// @Tags Subjects
// @ID subjects-get-all-schedule
// @Produce  json
// @Param schedule_id path uint true "Schedule ID"
// @Param limit query uint false "Limit"
// @Param offset query uint false "Offset"
// @Success 200 {object} GetSubjectsResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /schedules/{schedule_id}/subjects [get]
func (c *Controller) GetScheduleSubjects(ctx echo.Context) error {
	scheduleID, err := parseIDParam(ctx, "schedule_id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}
	limit, offset, err := getPagination(ctx)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	subjects, pagination, err := c.Service.Subject.GetAll(scheduleID, limit, offset)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, GetSubjectsResponse{
		Data:       subjects,
		Pagination: pagination,
	})
}

// GetSubject
// @Summary Get Subject by ID
// @Description Get Subject by ID
// @Tags Subjects
// @ID subjects-get-by-id
// @Produce  json
// @Param id path uint true "Subject ID"
// @Success 200 {object} domain.Subject
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /subjects/{id} [get]
func (c *Controller) GetSubject(ctx echo.Context) error {
	id, err := parseIDParam(ctx, "id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}
	subject, err := c.Service.Subject.GetByID(id)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, subject)
}

// CreateSubject
// @Summary Create Subject
// @Description Create Subject
// @Tags Subjects
// @ID subjects-create
// @Produce  json
// @Param data body domain.CreateSubjectDTO true "Data"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /subjects [post]
func (c *Controller) CreateSubject(ctx echo.Context) error {
	var req domain.CreateSubjectDTO
	err := bindStruct(ctx, &req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	id, err := c.Service.Subject.Create(req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, NewIDResponse(id))
}

// UpdateSubject
// @Summary Update Subject
// @Description Update Subject
// @Tags Subjects
// @ID subjects-update
// @Produce  json
// @Param id path uint true "Subject ID"
// @Param data body domain.UpdateSubjectDTO true "Data"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /subjects/{id} [patch]
func (c *Controller) UpdateSubject(ctx echo.Context) error {
	id, err := parseIDParam(ctx, "id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	var req domain.UpdateSubjectDTO
	err = bindStruct(ctx, &req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	err = c.Service.Subject.Update(user.ID, id, req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, NewIDResponse(id))
}

// DeleteSubject
// @Summary Delete Subject
// @Description Delete Subject
// @Tags Subjects
// @ID subjects-delete
// @Produce  json
// @Param id path uint true "Subject ID"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /subjects/{id} [delete]
func (c *Controller) DeleteSubject(ctx echo.Context) error {
	id, err := parseIDParam(ctx, "id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		return c.handleAppError(ctx, err)
	}
	err = c.Service.Subject.Delete(user.ID, id)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, NewIDResponse(id))
}
