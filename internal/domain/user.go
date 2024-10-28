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
	ID           uint64    `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	PasswordHash string    `json:"password_hash" db:"password_hash"`
	Role         Role      `json:"role" db:"role"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type UserUpdateDTO struct {
	Username     *string `json:"username"`
	PasswordHash *string `json:"password_hash"`
	Role         *Role   `json:"role"`
}

type UserCreate struct {
	Username     string    `json:"username"`
	PasswordHash string    `json:"password_hash"`
	Role         Role      `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}

type UserView struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func (u User) ToView() UserView {
	return UserView{
		ID:       u.ID,
		Username: u.Username,
		Role:     u.Role.String(),
	}
}
