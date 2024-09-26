package view

import "uni-schedule-backend/internal/domain"

type TeacherView struct {
	ID        domain.ID `json:"id"`
	ShortName string    `json:"short_name"`
	FullName  string    `json:"full_name"`
}
