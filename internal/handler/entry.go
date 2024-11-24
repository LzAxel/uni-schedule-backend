package handler

import (
	"net/http"
	"uni-schedule-backend/internal/domain"

	"github.com/labstack/echo/v4"
)

// GetEntry
// @Summary Get Entry by ID
// @Description Get Entry by ID
// @Tags Entry
// @ID entry-get-by-id
// @Produce  json
// @Param id path uint true "Entry ID"
// @Success 200 {object} domain.ScheduleEntry
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /entries/{id} [get]
func (c *Controller) GetEntry(ctx echo.Context) error {
	id, err := parseIDParam(ctx, "id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}
	entry, err := c.Service.Entry.GetByID(id)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, entry)
}

// CreateEntry
// @Summary Create Entry
// @Description Create Entry
// @Tags Entry
// @ID entry-create
// @Produce  json
// @Param data body domain.CreateScheduleEntryDTO true "Data"
// @Success 201 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /entries [post]
func (c *Controller) CreateEntry(ctx echo.Context) error {
	var req domain.CreateScheduleEntryDTO
	err := bindStruct(ctx, &req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}
	id, err := c.Service.Entry.Create(req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, NewIDResponse(id))
}

// UpdateEntry
// @Summary Update Entry
// @Description Update Entry
// @Tags Entry
// @ID entry-update
// @Produce  json
// @Param id path uint true "Entry ID"
// @Param data body domain.UpdateScheduleEntryDTO true "Data"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /entries/{id} [patch]
func (c *Controller) UpdateEntry(ctx echo.Context) error {
	id, err := parseIDParam(ctx, "id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	var req domain.UpdateScheduleEntryDTO
	err = bindStruct(ctx, &req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	err = c.Service.Entry.Update(user.ID, id, req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, NewIDResponse(id))
}

// DeleteEntry
// @Summary Delete Entry
// @Description Delete Entry
// @Tags Entry
// @ID entry-delete
// @Produce  json
// @Param id path uint true "Entry ID"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /entries/{id} [delete]
func (c *Controller) DeleteEntry(ctx echo.Context) error {
	id, err := parseIDParam(ctx, "id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	err = c.Service.Entry.Delete(user.ID, id)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, NewIDResponse(id))
}
