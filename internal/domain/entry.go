package domain

type ScheduleEntry struct {
	ID          uint64  `json:"id" db:"id"`
	ScheduleID  uint64  `json:"schedule_id" db:"schedule_id"`
	Day         Day     `json:"day" db:"day"`
	ClassNumber int     `json:"class_number" db:"class_number"`
	EvenClassID *uint64 `json:"even_class_id" db:"even_class_id"`
	OddClassID  *uint64 `json:"odd_class_id" db:"odd_class_id"`
	IsStatic    bool    `json:"is_static" db:"is_static"`
}

type CreateScheduleEntryDTO struct {
	Day         Day     `json:"day" binding:"required"`
	ScheduleID  uint64  `json:"schedule_id" binding:"required"`
	ClassNumber int     `json:"class_number" binding:"required"`
	EvenClassID *uint64 `json:"even_class_id,omitempty"`
	OddClassID  *uint64 `json:"odd_class_id,omitempty"`
	IsStatic    bool    `json:"is_static" binding:"required"`
}

type UpdateScheduleEntryDTO struct {
	Day         *Day    `json:"day,omitempty"`
	ClassNumber *int    `json:"class_number,omitempty"`
	EvenClassID *uint64 `json:"even_class_id,omitempty"`
	OddClassID  *uint64 `json:"odd_class_id,omitempty"`
	IsStatic    *bool   `json:"is_static,omitempty"`
}

type ScheduleEntryView struct {
	ID          uint64     `json:"id"`
	Day         Day        `json:"day"`
	ClassNumber int        `json:"class_number"`
	Even        *ClassView `json:"even,omitempty"`
	Odd         *ClassView `json:"odd,omitempty"`
	IsStatic    bool       `json:"is_static"`
}

func (s ScheduleEntry) ToView(even *ClassView, odd *ClassView) ScheduleEntryView {
	return ScheduleEntryView{
		ID:          s.ID,
		Day:         s.Day,
		ClassNumber: s.ClassNumber,
		Even:        even,
		Odd:         odd,
		IsStatic:    s.IsStatic,
	}
}
