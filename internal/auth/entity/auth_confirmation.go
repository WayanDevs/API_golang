package entity

import (
	"time"

	"github.com/google/uuid"
)

// AuthConfirmation represents the `auth.auth_confirmation` table.
type AuthConfirmation struct {
    ID                uint           `gorm:"primaryKey" json:"id"`
    AuthID            uuid.UUID      `gorm:"type:uuid;not null" json:"auth_id"`
    ConfirmationCode  string         `gorm:"type:varchar(255);not null" json:"confirmation_code"`
    ConfirmedBy       *string        `gorm:"type:varchar(255)" json:"confirmed_by"`
    ConfirmedAt       *time.Time     `gorm:"type:timestamptz" json:"confirmed_at"`
    Auth              Auth           `gorm:"foreignKey:AuthID;references:ID;constraint:OnDelete:CASCADE" json:"auth"`
}

// TableName overrides the default table name for AuthConfirmation
func (AuthConfirmation) TableName() string {
    return "auth.auth_confirmation"
}
