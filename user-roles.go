package authority

import (
	"github.com/google/uuid"
)

// UserRole represents the relationship between users and roles
type UserRole struct {
	ID     uint
	UserID uuid.UUID
	RoleID uint
}

// TableName sets the table name
func (u UserRole) TableName() string {
	return tablePrefix + "user_roles"
}
