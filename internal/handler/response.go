package handler

import (
	"github.com/labstack/echo/v4"
	"strconv"
	"uni-schedule-backend/internal/apperror"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type IDResponse struct {
	ID uint64 `json:"id"`
}

func NewIDResponse(id uint64) IDResponse {
	return IDResponse{
		ID: id,
	}
}

func bindStruct(ctx echo.Context, i any) error {
	err := ctx.Bind(&i)
	if err != nil {
		return apperror.ErrInvalidInputBody
	}
	return nil
}

func parseIDParam(ctx echo.Context, param string) (uint64, error) {
	idStr := ctx.Param(param)
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, apperror.NewErrInvalidIDParam(param)
	}
	return id, nil
}
