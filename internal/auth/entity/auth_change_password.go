package entity

import (
	"time"

	"github.com/google/uuid"
)

// AuthChangePassword represents the `auth.auth_change_password` table.
type AuthChangePassword struct {
    ID               uint           `gorm:"primaryKey" json:"id"`
    AuthID           uuid.UUID      `gorm:"type:uuid;not null" json:"auth_id"`
    NewPasswordHash  string         `gorm:"type:varchar(255);not null" json:"new_password_hash"`
    CreatedAt        time.Time      `gorm:"type:timestamptz;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
    ConfirmedAt      *time.Time     `gorm:"type:timestamptz" json:"confirmed_at"`
    Auth             Auth           `gorm:"foreignKey:AuthID;references:ID;constraint:OnDelete:CASCADE" json:"auth"`
}

// TableName overrides the default table name for AuthChangePassword
func (AuthChangePassword) TableName() string {
    return "auth.auth_change_password"
}
