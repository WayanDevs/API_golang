package entity

import (
	"time"

	"github.com/google/uuid"
)

// Auth represents the `auth.auth` table.
type Auth struct {
    ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
    Email       string         `gorm:"type:varchar(255);unique;not null" json:"email"`
    Hash        string         `gorm:"type:varchar(255);not null" json:"hash"`
    Token       string         `gorm:"type:char(5);not null" json:"token"`
    CreatedAt   time.Time      `gorm:"type:timestamptz;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt   time.Time      `gorm:"type:timestamptz;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
    DeletedAt   *time.Time     `gorm:"type:timestamptz" json:"deleted_at"`

    // Relationships
    Confirmations []AuthConfirmation `gorm:"foreignKey:AuthID"`
    OAuths        []AuthOAuth        `gorm:"foreignKey:AuthID"`
    ChangeEmails  []AuthChangeEmail  `gorm:"foreignKey:AuthID"`
    ChangePasswords []AuthChangePassword `gorm:"foreignKey:AuthID"`
}

// TableName overrides the default table name for Auth
func (Auth) TableName() string {
    return "auth.auth"
}
