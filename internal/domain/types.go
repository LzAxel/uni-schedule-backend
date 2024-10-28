package domain

type ClassType string

const (
	Lecture  ClassType = "lecture"
	Practice ClassType = "practice"
	Lab      ClassType = "lab"
	Combined ClassType = "combined"
)

type Day string

const (
	Monday    Day = "monday"
	Tuesday   Day = "tuesday"
	Wednesday Day = "wednesday"
	Thursday  Day = "thursday"
	Friday    Day = "friday"
	Saturday  Day = "saturday"
)
