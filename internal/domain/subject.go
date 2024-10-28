package domain

type Subject struct {
	ID         uint64 `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	ScheduleID uint64 `json:"schedule_id" db:"schedule_id"`
}

type CreateSubjectDTO struct {
	Name       string `json:"name" binding:"required"`
	ScheduleID uint64 `json:"schedule_id" binding:"required"`
}

type UpdateSubjectDTO struct {
	Name *string `json:"name,omitempty"`
}

type SubjectView struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func (s *Subject) ToView() SubjectView {
	return SubjectView{
		ID:   s.ID,
		Name: s.Name,
	}
}
