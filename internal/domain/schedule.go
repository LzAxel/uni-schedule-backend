package domain

import "time"

type Schedule struct {
	ID        uint64    `json:"id" db:"id"`
	UserID    uint64    `json:"user_id" db:"user_id"`
	Slug      string    `json:"slug" db:"slug"`
	Title     string    `json:"title" db:"title"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CreateScheduleDTO struct {
	UserID uint64 `json:"user_id" binding:"required"`
	Slug   string `json:"slug" binding:"required"`
	Title  string `json:"title" binding:"required"`
}

type UpdateScheduleDTO struct {
	Slug  *string `json:"slug,omitempty"`
	Title *string `json:"title,omitempty"`
}

type ScheduleView struct {
	ID      uint64                `json:"id"`
	UserID  uint64                `json:"user_id"`
	Slug    string                `json:"slug"`
	Title   string                `json:"title"`
	Entries DayGroupedClassesView `json:"entries"`
}

func (s Schedule) ToView(groupedClasses DayGroupedClassesView) ScheduleView {
	return ScheduleView{
		ID:      s.ID,
		UserID:  s.UserID,
		Slug:    s.Slug,
		Entries: groupedClasses,
		Title:   s.Title,
	}
}

type ScheduleGetAllFilters struct {
	UserID *uint64
}
