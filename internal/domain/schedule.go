package domain

type Schedule struct {
	ID     uint64 `json:"id" db:"id"`
	UserID uint64 `json:"user_id" db:"user_id"`
	Slug   string `json:"slug" db:"slug"`
}

type CreateScheduleDTO struct {
	UserID uint64 `json:"user_id" binding:"required"`
	Slug   string `json:"slug" binding:"required"`
}

type UpdateScheduleDTO struct {
	Slug *string `json:"slug,omitempty"`
}

type ScheduleView struct {
	ID      uint64              `json:"id"`
	UserID  uint64              `json:"user_id"`
	Slug    string              `json:"slug"`
	Entries []ScheduleEntryView `json:"entries"`
}

func (s Schedule) ToView(entries []ScheduleEntryView) ScheduleView {
	return ScheduleView{
		ID:      s.ID,
		UserID:  s.UserID,
		Slug:    s.Slug,
		Entries: entries,
	}
}

type ScheduleGetAllFilters struct {
	UserID *uint64
}
