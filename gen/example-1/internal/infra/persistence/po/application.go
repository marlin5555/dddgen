// Package po Code generated, DO NOT EDIT.
package po

import (
	"time"

	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/base"
)

// Application table name: t_application
type Application struct {
	CreatedAt   time.Time `gorm:"column:created_at;type:time;<-:create;autoCreateTime"`
	Creator     string    `gorm:"column:creator;type:varchar(255)"`
	DeletedAt   int64     `gorm:"column:deleted_at;default:0"`
	Description string    `gorm:"column:description;type:varchar(4096);default null"`
	ID          string    `gorm:"column:id;type:varchar(36);primaryKey;NOT NULL"`
	Name        string    `gorm:"column:name;type:varchar(255);NOT NULL;uniqueIndex:t_application_u1"`
	Nickname    string    `gorm:"column:nickname;type:varchar(255);"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:time;autoUpdateTime"`
	Updater     string    `gorm:"column:updater;type:varchar(255)"`
}

type Applications []Application

func (p Applications) GetApplicationIDs() []string {
	return base.ToApplicationIDs(p)
}

// TableName Application impl schema.Tabler
func (p Application) TableName() string {
	return "t_application"
}

// GetCreatedAt Application impl base.CreatedAt
func (p Application) GetCreatedAt() time.Time {
	return p.CreatedAt
}

// GetCreator Application impl base.Creator
func (p Application) GetCreator() string {
	return p.Creator
}

// GetDeletedAt Application impl base.DeletedAt
func (p Application) GetDeletedAt() int64 {
	return p.DeletedAt
}

// GetApplicationDescription Application impl base.Description
func (p Application) GetApplicationDescription() string {
	return p.Description
}

// GetApplicationID Application impl base.ID
func (p Application) GetApplicationID() string {
	return p.ID
}

// GetApplicationIDs Application impl base.IDs
func (p Application) GetApplicationIDs() []string {
	return []string{p.ID}
}

// GetApplicationName Application impl base.Name
func (p Application) GetApplicationName() string {
	return p.Name
}

// GetApplicationNames Application impl base.Names
func (p Application) GetApplicationNames() []string {
	return []string{p.Name}
}

// GetApplicationNickname Application impl base.Nickname
func (p Application) GetApplicationNickname() string {
	return p.Nickname
}

// GetUpdatedAt Application impl base.UpdatedAt
func (p Application) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

// GetUpdater Application impl base.Updater
func (p Application) GetUpdater() string {
	return p.Updater
}
