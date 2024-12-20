package handler

import (
	"net/http"
	"uni-schedule-backend/internal/domain"

	"github.com/labstack/echo/v4"
)

// GetClass
// @Summary Get Class By ID
// @Description Get Class  By ID
// @Tags Class
// @ID classes-get-by-id
// @Produce  json
// @Param id path uint true "Class ID"
// @Success 200 {object} domain.Class
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /classes/{id} [get]
func (c *Controller) GetClass(ctx echo.Context) error {
	id, err := parseIDParam(ctx, "id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}
	class, err := c.Service.Class.GetByID(id)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, class)
}

// CreateClass
// @Summary Create Class
// @Description Create Class
// @Tags Class
// @ID classes-create
// @Produce  json
// @Param data body domain.CreateClassDTO true "Data"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /classes [post]
func (c *Controller) CreateClass(ctx echo.Context) error {
	var req domain.CreateClassDTO
	err := bindStruct(ctx, &req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}
	id, err := c.Service.Class.Create(req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, NewIDResponse(id))
}

// UpdateClass
// @Summary Update Class
// @Description Update Class
// @Tags Class
// @ID classes-update
// @Produce json
// @Param data body domain.UpdateClassDTO true "Data"
// @Param id path uint true "Class ID"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /classes/{id} [patch]
func (c *Controller) UpdateClass(ctx echo.Context) error {
	id, err := parseIDParam(ctx, "id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	var req domain.UpdateClassDTO
	err = bindStruct(ctx, &req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	err = c.Service.Class.Update(user.ID, id, req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, NewIDResponse(id))
}

// DeleteClass
// @Summary Delete Class
// @Description Delete Class
// @Tags Class
// @ID classes-delete
// @Produce  json
// @Param id path uint true "Class ID"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /classes/{id} [delete]
func (c *Controller) DeleteClass(ctx echo.Context) error {
	id, err := parseIDParam(ctx, "id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	err = c.Service.Class.Delete(user.ID, id)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, NewIDResponse(id))
}
