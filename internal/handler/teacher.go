package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"uni-schedule-backend/internal/domain"
)

type CreateTeacherRequest struct {
	ShortName string `json:"short_name"`
	FullName  string `json:"full_name"`
}

// CreateTeacher
// @Summary Create Teacher
// @Description Create a new teacher
// @Tags Teacher
// @ID teacher-create
// @Accept  json
// @Produce  json
// @Param data body CreateTeacherRequest true "Data"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /teacher [post]
func (c *Controller) CreateTeacher(ctx echo.Context) error {
	var req CreateTeacherRequest
	err := bindStruct(ctx, &req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	id, err := c.Service.Teacher.Create(domain.TeacherCreate{
		ShortName: req.ShortName,
		FullName:  req.FullName,
	})
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, NewIDResponse(id))
}

type UpdateTeacherRequest struct {
	ShortName *string `json:"short_name"`
	FullName  *string `json:"full_name"`
}

// UpdateTeacher
// @Summary Update Teacher
// @Description Update teacher
// @Tags Teacher
// @ID teacher-update
// @Produce  json
// @Param id path uint true "Teacher ID"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /teacher/{id} [patch]
func (c *Controller) UpdateTeacher(ctx echo.Context) error {
	id, err := parseIDParam(ctx, "id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	var req UpdateTeacherRequest
	err = bindStruct(ctx, &req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	err = c.Service.Teacher.Update(id, domain.TeacherUpdate{
		ShortName: req.ShortName,
		FullName:  req.FullName,
	})
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
// @Router /teacher/{id} [delete]
func (c *Controller) DeleteTeacher(ctx echo.Context) error {
	id, err := parseIDParam(ctx, "id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	err = c.Service.Teacher.Delete(id)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, NewIDResponse(id))
}
