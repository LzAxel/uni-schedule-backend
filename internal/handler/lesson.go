package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"uni-schedule-backend/internal/domain"
)

type CreateLessonRequest struct {
	Name       string `json:"name"`
	Location   string `json:"location"`
	TeacherID  uint64 `json:"teacher_id"`
	LessonType uint64 `json:"lesson_type"`
}

// CreateLesson
// @Summary Create New Lesson
// @Description Create a new lesson
// @Tags Lesson
// @ID lesson-create
// @Accept  json
// @Produce  json
// @Param data body CreateLessonRequest true "Data"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /lesson [post]
func (c *Controller) CreateLesson(ctx echo.Context) error {
	var req CreateLessonRequest
	err := bindStruct(ctx, &req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	id, err := c.Service.Lesson.Create(domain.LessonCreate{
		Name:       req.Name,
		Location:   req.Location,
		TeacherID:  req.TeacherID,
		LessonType: domain.LessonType(req.LessonType),
	})
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, NewIDResponse(id))
}

type UpdateLessonRequest struct {
	Name       *string `json:"name"`
	Location   *string `json:"location"`
	TeacherID  *uint64 `json:"teacher_id"`
	LessonType *uint64 `json:"lesson_type"`
}

// UpdateLesson
// @Summary Update Lesson
// @Description Update the lesson by id
// @Tags Lesson
// @ID lesson-update
// @Accept  json
// @Produce  json
// @Param data body UpdateLessonRequest true "Data"
// @Param id path uint true "Lesson ID"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /lesson/{id} [patch]
func (c *Controller) UpdateLesson(ctx echo.Context) error {
	id, err := parseIDParam(ctx, "id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	var req UpdateLessonRequest
	err = bindStruct(ctx, &req)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	err = c.Service.Lesson.Update(id, domain.LessonUpdate{
		Name:       req.Name,
		Location:   req.Location,
		TeacherID:  req.TeacherID,
		LessonType: (*domain.LessonType)(req.LessonType),
	})
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, NewIDResponse(id))
}

// DeleteLesson
// @Summary Delete Lesson
// @Description Delete the lesson by id
// @Tags Lesson
// @ID lesson-delete
// @Produce  json
// @Param id path uint true "Lesson ID"
// @Success 200 {object} IDResponse
// @Failure 400 {object} ErrorResponse
// @Security Bearer
// @Router /lesson/{id} [delete]
func (c *Controller) DeleteLesson(ctx echo.Context) error {
	id, err := parseIDParam(ctx, "id")
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	err = c.Service.Lesson.Delete(id)
	if err != nil {
		return c.handleAppError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, NewIDResponse(id))
}
