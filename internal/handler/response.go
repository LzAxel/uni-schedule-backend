package handler

import (
	"strconv"
	"uni-schedule-backend/internal/apperror"

	"github.com/labstack/echo/v4"
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
		return 0, apperror.NewErrInvalidQueryParam(param)
	}
	return id, nil
}

func getPagination(ctx echo.Context) (limit uint64, offset uint64, err error) {
	limit = 25
	offset = 0

	limitStr := ctx.QueryParam("limit")
	if limitStr != "" {
		limit, err = strconv.ParseUint(limitStr, 10, 64)
		if err != nil {
			return 0, 0, apperror.NewErrInvalidQueryParam("limit")
		}
	}
	offsetStr := ctx.QueryParam("offset")
	if offsetStr != "" {
		offset, err = strconv.ParseUint(offsetStr, 10, 64)
		if err != nil {
			return 0, 0, apperror.NewErrInvalidQueryParam("offset")
		}
	}
	return limit, offset, nil
}
