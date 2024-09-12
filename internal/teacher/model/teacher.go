package model

import "uni-schedule-backend/internal/domain"

type Teacher struct {
	ID        domain.ID
	ShortName string
	FullName  string
}

type TeacherUpdate struct {
	ShortName *string
	FullName  *string
}
