package domain

type Teacher struct {
	ID         uint64 `db:"id" json:"id"`
	FirstName  string `db:"first_name" json:"first_name"`
	LastName   string `db:"last_name" json:"last_name"`
	Surname    string `db:"surname" json:"surname"`
	ScheduleID uint64 `db:"schedule_id" json:"schedule_id" `
}

type TeacherUpdateDTO struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Surname   *string `json:"surname,omitempty"`
}

type TeacherCreateDTO struct {
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	ScheduleID uint64 `json:"schedule_id" binding:"required"`
}

type TeacherView struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Surname   string `json:"surname"`
}

func (t *Teacher) ToView() TeacherView {
	return TeacherView{
		ID:        t.ID,
		FirstName: t.FirstName,
		LastName:  t.LastName,
		Surname:   t.Surname,
	}
}
