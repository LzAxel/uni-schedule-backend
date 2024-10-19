package domain

import (
	"time"
)

type Role uint

const (
	RoleAdmin Role = iota
	RoleScheduleEditor
	RoleStudent
)

func (r Role) String() string {
	switch r {
	case RoleAdmin:
		return "Admin"
	case RoleScheduleEditor:
		return "Editor"
	case RoleStudent:
		return "Student"
	default:
		return "Unknown"
	}
}

type User struct {
	ID           uint64
	Username     string
	PasswordHash string
	Role         Role
	CreatedAt    time.Time
}

type UserUpdate struct {
	Username *string
	Password *string
	Role     *Role
}

type UserUpdateDTO struct {
	Username     *string
	PasswordHash *string
	Role         *Role
}

type UserCreate struct {
	Username     string
	PasswordHash string
	Role         Role
	CreatedAt    time.Time
}
