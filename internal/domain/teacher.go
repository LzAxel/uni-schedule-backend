package domain

type Teacher struct {
	ID        uint64
	ShortName string
	FullName  string
}

type TeacherUpdate struct {
	ShortName *string
	FullName  *string
}

type TeacherCreate struct {
	ShortName string
	FullName  string
}

type TeacherView struct {
	ID        uint64 `json:"id"`
	ShortName string `json:"short_name"`
	FullName  string `json:"full_name"`
}
