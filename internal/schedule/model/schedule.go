package model

import (
	"uni-schedule-backend/internal/domain"
)

type Schedule struct {
	ID        domain.ID
	CreatorID domain.ID
	Name      string
	Slug      string
}

type ScheduleUpdate struct {
	Name      *string
	Slug      *string
	CreatorID *domain.ID
}
