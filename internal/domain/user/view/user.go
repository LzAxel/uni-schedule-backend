package view

import (
	"time"
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/domain/user/model"
)

type UserView struct {
	ID        domain.ID  `json:"id"`
	Username  string     `json:"username"`
	Role      model.Role `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
}
