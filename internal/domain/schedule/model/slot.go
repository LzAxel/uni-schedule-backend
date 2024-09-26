package model

import (
	"time"
	"uni-schedule-backend/internal/domain"
)

type ScheduleSlot struct {
	ID               domain.ID
	ScheduleID       domain.ID
	Weekday          time.Weekday
	Number           uint
	IsAlternating    bool
	EvenWeekLessonID *domain.ID
	OddWeekLessonID  *domain.ID
}

type ScheduleSlotUpdate struct {
	ScheduleID       *domain.ID
	Weekday          *time.Weekday
	Number           *uint
	IsAlternating    *bool
	EvenWeekLessonID *domain.ID
	OddWeekLessonID  *domain.ID
}
