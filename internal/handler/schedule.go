package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"
)

type CreateScheduleRequest struct {
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	CreatorID uint64 `json:"creator_id"`
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
// @Router /schedule [post]
func (c *Controller) CreateSchedule(ctx echo.Context) error {
	var req CreateScheduleRequest
	err := bindStruct(ctx, &req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	id, err := c.Service.Schedule.CreateSchedule(domain.ScheduleCreate{
		Name:      req.Name,
		Slug:      req.Slug,
		CreatorID: req.CreatorID,
	})

	return ctx.JSON(http.StatusCreated, NewIDResponse(id))
}

type AddSlotToScheduleRequest struct {
	Weekday          time.Weekday `json:"weekday"`
	Number           uint         `json:"number"`
	IsAlternating    bool         `json:"is_alternating"`
	EvenWeekLessonID uint64       `json:"even_week_lesson_id"`
	OddWeekLessonID  uint64       `json:"odd_week_lesson_id"`
}

// AddPairToSchedule
// @Summary Add Pair Slot To Schedule
// @Description Add a new pair slot to schedule
// @Tags Schedule
// @ID schedule-add-pair
// @Accept  json
// @Produce  json
// @Param schedule_id path uint true "Schedule ID"
// @Param data body AddSlotToScheduleRequest true "Data"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /schedule/{schedule_id}/slot [post]
func (c *Controller) AddPairToSchedule(ctx echo.Context) error {
	var req AddSlotToScheduleRequest

	scheduleID, err := parseIDParam(ctx, "schedule_id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	err = bindStruct(ctx, &req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	id, err := c.Service.Schedule.CreateSlot(domain.ScheduleSlotCreate{
		ScheduleID:       scheduleID,
		Weekday:          req.Weekday,
		Number:           req.Number,
		IsAlternating:    req.IsAlternating,
		EvenWeekLessonID: req.EvenWeekLessonID,
		OddWeekLessonID:  req.OddWeekLessonID,
	})

	return ctx.JSON(http.StatusCreated, NewIDResponse(id))
}

// DeleteSlotFromSchedule
// @Summary Delete Pair Slot From Schedule
// @Description Delete pair slot from schedule
// @Tags Schedule
// @ID schedule-delete-pair
// @Produce  json
// @Param slot_id path uint true "Slot ID"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /schedule/slot/{slot_id} [delete]
func (c *Controller) DeleteSlotFromSchedule(ctx echo.Context) error {
	slotID, err := parseIDParam(ctx, "slot_id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	err = c.Service.Schedule.DeleteSlot(slotID)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, NewIDResponse(slotID))
}

type UpdateSlotInScheduleRequest struct {
	Weekday          *time.Weekday `json:"weekday"`
	Number           *uint         `json:"number"`
	IsAlternating    *bool         `json:"is_alternating"`
	EvenWeekLessonID *uint64       `json:"even_week_lesson_id"`
	OddWeekLessonID  *uint64       `json:"odd_week_lesson_id"`
}

// UpdateSlotInSchedule
// @Summary Update Pair Slot In Schedule
// @Description Update pair slot in schedule. To clear weekLessonID - set it to 0.
// @Tags Schedule
// @ID schedule-update-pair
// @Accept  json
// @Produce  json
// @Param slot_id path uint true "Slot ID"
// @Param data body UpdateSlotInScheduleRequest true "Data"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /schedule/slot/{slot_id} [patch]
func (c *Controller) UpdateSlotInSchedule(ctx echo.Context) error {
	slotID, err := parseIDParam(ctx, "slot_id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	var req UpdateSlotInScheduleRequest

	err = bindStruct(ctx, &req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	err = c.Service.Schedule.UpdateSlot(slotID, domain.ScheduleSlotUpdate{
		Weekday:          req.Weekday,
		Number:           req.Number,
		IsAlternating:    req.IsAlternating,
		EvenWeekLessonID: req.EvenWeekLessonID,
		OddWeekLessonID:  req.OddWeekLessonID,
	})
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, NewIDResponse(slotID))
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
// @Router /schedule/slug/{slug} [get]
func (c *Controller) GetScheduleBySlug(ctx echo.Context) error {
	slug := ctx.Param("slug")

	if slug == "" {
		return c.handleAppError(ctx, apperror.ErrInvalidSlug)
	}

	schedule, err := c.Service.Schedule.GetScheduleBySlug(slug)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, schedule)
}

type UpdateScheduleRequest struct {
	Name      *string `json:"name"`
	Slug      *string `json:"slug"`
	CreatorID *uint64 `json:"creator_id"`
}

// UpdateSchedule
// @Summary Update Schedule
// @Description Update schedule
// @Tags Schedule
// @ID schedule-update
// @Accept  json
// @Produce  json
// @Param id path uint true "Schedule ID"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /schedule/{id} [patch]
func (c *Controller) UpdateSchedule(ctx echo.Context) error {
	id, err := parseIDParam(ctx, "id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	var req UpdateScheduleRequest
	err = bindStruct(ctx, &req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	err = c.Service.Schedule.UpdateSchedule(id, domain.ScheduleUpdate{
		Name:      req.Name,
		Slug:      req.Slug,
		CreatorID: req.CreatorID,
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
// @Router /schedule/{id} [delete]
func (c *Controller) DeleteSchedule(ctx echo.Context) error {
	id, err := parseIDParam(ctx, "id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	err = c.Service.Schedule.DeleteSchedule(id)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, NewIDResponse(id))
}
