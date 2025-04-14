package model

import "time"

// Role представляет роль пользователя
type Role string

const (
	RoleClient    Role = "client"
	RoleModerator Role = "moderator"
)

// User представляет пользователя системы
type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // не отдаем пароль в JSON
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

// IsValidRole проверяет, является ли роль допустимой
func IsValidRole(role Role) bool {
	return role == RoleClient || role == RoleModerator
}
