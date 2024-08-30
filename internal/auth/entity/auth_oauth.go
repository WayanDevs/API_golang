package entity

import (
	"time"

	"github.com/google/uuid"
)

// AuthOAuth represents the `auth.auth_oauth` table.
type AuthOAuth struct {
    ID         uint           `gorm:"primaryKey" json:"id"`
    AuthID     uuid.UUID      `gorm:"type:uuid;not null" json:"auth_id"`
    Provider   string         `gorm:"type:varchar(255);not null" json:"provider"`
    ProviderID string         `gorm:"type:varchar(255);not null" json:"provider_id"`
    CreatedAt  time.Time      `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"created_at"`
    Auth       Auth           `gorm:"foreignKey:AuthID;references:ID;constraint:OnDelete:CASCADE" json:"auth"`
}

// TableName overrides the default table name for AuthOAuth
func (AuthOAuth) TableName() string {
    return "auth.auth_oauth"
}
