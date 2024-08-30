package entity

import (
	"time"

	"github.com/google/uuid"
)

// AuthChangeEmail represents the `auth.auth_change_email` table.
type AuthChangeEmail struct {
    ID             uint           `gorm:"primaryKey" json:"id"`
    AuthID         uuid.UUID      `gorm:"type:uuid;not null" json:"auth_id"`
    NewEmail       string         `gorm:"type:varchar(255);not null" json:"new_email"`
    CreatedAt      time.Time      `gorm:"type:timestamptz;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
    ConfirmedAt    *time.Time     `gorm:"type:timestamptz" json:"confirmed_at"`
    Auth           Auth           `gorm:"foreignKey:AuthID;references:ID;constraint:OnDelete:CASCADE" json:"auth"`
}

// TableName overrides the default table name for AuthChangeEmail
func (AuthChangeEmail) TableName() string {
    return "auth.auth_change_email"
}
