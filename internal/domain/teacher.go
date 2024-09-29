package domain

type Teacher struct {
	ID        ID
	ShortName string
	FullName  string
}

type TeacherUpdate struct {
	ShortName *string
	FullName  *string
}

type TeacherView struct {
	ID        ID     `json:"id"`
	ShortName string `json:"short_name"`
	FullName  string `json:"full_name"`
}
