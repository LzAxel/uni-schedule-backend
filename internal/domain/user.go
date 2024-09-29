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

type User struct {
	ID           ID
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
