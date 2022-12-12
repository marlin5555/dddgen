// Package po Code generated, DO NOT EDIT.
package po

import (
	"time"

	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/base"
)

// EventBus table name: t_event_bus
type EventBus struct {
	CreatedAt   time.Time `gorm:"column:created_at;type:time;<-:create;autoCreateTime"`
	Creator     string    `gorm:"column:creator;type:varchar(255)"`
	DeletedAt   int64     `gorm:"column:deleted_at;default:0"`
	Description string    `gorm:"column:description;type:varchar(4096);default null"`
	ID          string    `gorm:"column:id;type:varchar(36);primaryKey;NOT NULL"`
	Name        string    `gorm:"column:name;type:varchar(255);NOT NULL;uniqueIndex:t_event_bus_u1"`
	Params      string    `gorm:"column:params;type:varchar(255)"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:time;autoUpdateTime"`
	Updater     string    `gorm:"column:updater;type:varchar(255)"`
}

type EventBuses []EventBus

func (p EventBuses) GetEventBusIDs() []string {
	return base.ToEventBusIDs(p)
}

// TableName EventBus impl schema.Tabler
func (p EventBus) TableName() string {
	return "t_event_bus"
}

// GetCreatedAt EventBus impl base.CreatedAt
func (p EventBus) GetCreatedAt() time.Time {
	return p.CreatedAt
}

// GetCreator EventBus impl base.Creator
func (p EventBus) GetCreator() string {
	return p.Creator
}

// GetDeletedAt EventBus impl base.DeletedAt
func (p EventBus) GetDeletedAt() int64 {
	return p.DeletedAt
}

// GetEventBusDescription EventBus impl base.Description
func (p EventBus) GetEventBusDescription() string {
	return p.Description
}

// GetEventBusID EventBus impl base.ID
func (p EventBus) GetEventBusID() string {
	return p.ID
}

// GetEventBusIDs EventBus impl base.IDs
func (p EventBus) GetEventBusIDs() []string {
	return []string{p.ID}
}

// GetEventBusName EventBus impl base.Name
func (p EventBus) GetEventBusName() string {
	return p.Name
}

// GetEventBusNames EventBus impl base.Names
func (p EventBus) GetEventBusNames() []string {
	return []string{p.Name}
}

// GetEventBusParams EventBus impl base.Params
func (p EventBus) GetEventBusParams() string {
	return p.Params
}

// GetUpdatedAt EventBus impl base.UpdatedAt
func (p EventBus) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

// GetUpdater EventBus impl base.Updater
func (p EventBus) GetUpdater() string {
	return p.Updater
}

// EventType table name: t_event_type
type EventType struct {
	CreatedAt   time.Time `gorm:"column:created_at;type:time;<-:create;autoCreateTime"`
	Creator     string    `gorm:"column:creator;type:varchar(255)"`
	DeletedAt   int64     `gorm:"column:deleted_at;default:0"`
	Description string    `gorm:"column:description;type:varchar(4096);default null"`
	EventBusID  string    `gorm:"column:event_bus_id;type:varchar(36);NOT NULL"`
	ID          string    `gorm:"column:id;type:varchar(36);primaryKey;NOT NULL"`
	MainAppID   string    `gorm:"column:main_app_id;type:varchar(36);NOT NULL"`
	Name        string    `gorm:"column:name;type:varchar(255);NOT NULL;uniqueIndex:t_event_type_u1"`
	Status      int       `gorm:"column:status;type:int;NOT NULL;default:0"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:time;autoUpdateTime"`
	Updater     string    `gorm:"column:updater;type:varchar(255)"`
}

type EventTypes []EventType

func (p EventTypes) GetEventTypeIDs() []string {
	return base.ToEventTypeIDs(p)
}
func (p EventTypes) GetEventBusIDs() []string {
	return base.ToEventBusIDs(p)
}
func (p EventTypes) GetOwnershipIDs() []string {
	return base.ToOwnershipIDs(p)
}

// TableName EventType impl schema.Tabler
func (p EventType) TableName() string {
	return "t_event_type"
}

// GetCreatedAt EventType impl base.CreatedAt
func (p EventType) GetCreatedAt() time.Time {
	return p.CreatedAt
}

// GetCreator EventType impl base.Creator
func (p EventType) GetCreator() string {
	return p.Creator
}

// GetDeletedAt EventType impl base.DeletedAt
func (p EventType) GetDeletedAt() int64 {
	return p.DeletedAt
}

// GetEventTypeDescription EventType impl base.Description
func (p EventType) GetEventTypeDescription() string {
	return p.Description
}

// GetEventBusID EventType impl base.EventBusID
func (p EventType) GetEventBusID() string {
	return p.EventBusID
}

// GetEventBusIDs EventType impl base.EventBusIDs
func (p EventType) GetEventBusIDs() []string {
	return []string{p.EventBusID}
}

// GetEventTypeID EventType impl base.ID
func (p EventType) GetEventTypeID() string {
	return p.ID
}

// GetEventTypeIDs EventType impl base.IDs
func (p EventType) GetEventTypeIDs() []string {
	return []string{p.ID}
}

// GetOwnershipID EventType impl base.MainAppID
func (p EventType) GetOwnershipID() string {
	return p.MainAppID
}

// GetOwnershipIDs EventType impl base.MainAppIDs
func (p EventType) GetOwnershipIDs() []string {
	return []string{p.MainAppID}
}

// GetEventTypeName EventType impl base.Name
func (p EventType) GetEventTypeName() string {
	return p.Name
}

// GetEventTypeNames EventType impl base.Names
func (p EventType) GetEventTypeNames() []string {
	return []string{p.Name}
}

// GetEventTypeStatus EventType impl base.Status
func (p EventType) GetEventTypeStatus() int {
	return p.Status
}

// GetEventTypeStatuses EventType impl base.Statuses
func (p EventType) GetEventTypeStatuses() []int {
	return []int{p.Status}
}

// GetUpdatedAt EventType impl base.UpdatedAt
func (p EventType) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

// GetUpdater EventType impl base.Updater
func (p EventType) GetUpdater() string {
	return p.Updater
}

// Publication table name: t_publication
type Publication struct {
	CreatedAt   time.Time `gorm:"column:created_at;type:time;<-:create;autoCreateTime"`
	Creator     string    `gorm:"column:creator;type:varchar(255)"`
	DeletedAt   int64     `gorm:"column:deleted_at;default:0"`
	Description string    `gorm:"column:description;type:varchar(4096);default null"`
	EventTypeID string    `gorm:"column:name;type:varchar(255);NOT NULL"`
	ID          string    `gorm:"column:id;type:varchar(36);primaryKey;NOT NULL"`
	PublisherID string    `gorm:"column:main_app_id;type:varchar(36);NOT NULL"`
	Status      int       `gorm:"column:status;type:int;NOT NULL;default:0"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:time;autoUpdateTime"`
	Updater     string    `gorm:"column:updater;type:varchar(255)"`
}

type Publications []Publication

func (p Publications) GetPublicationIDs() []string {
	return base.ToPublicationIDs(p)
}
func (p Publications) GetEventTypeIDs() []string {
	return base.ToEventTypeIDs(p)
}
func (p Publications) GetPublisherIDs() []string {
	return base.ToPublisherIDs(p)
}

// TableName Publication impl schema.Tabler
func (p Publication) TableName() string {
	return "t_publication"
}

// GetCreatedAt Publication impl base.CreatedAt
func (p Publication) GetCreatedAt() time.Time {
	return p.CreatedAt
}

// GetCreator Publication impl base.Creator
func (p Publication) GetCreator() string {
	return p.Creator
}

// GetDeletedAt Publication impl base.DeletedAt
func (p Publication) GetDeletedAt() int64 {
	return p.DeletedAt
}

// GetPublicationDescription Publication impl base.Description
func (p Publication) GetPublicationDescription() string {
	return p.Description
}

// GetEventTypeID Publication impl base.EventTypeID
func (p Publication) GetEventTypeID() string {
	return p.EventTypeID
}

// GetEventTypeIDs Publication impl base.EventTypeIDs
func (p Publication) GetEventTypeIDs() []string {
	return []string{p.EventTypeID}
}

// GetPublicationID Publication impl base.ID
func (p Publication) GetPublicationID() string {
	return p.ID
}

// GetPublicationIDs Publication impl base.IDs
func (p Publication) GetPublicationIDs() []string {
	return []string{p.ID}
}

// GetPublisherID Publication impl base.PublisherID
func (p Publication) GetPublisherID() string {
	return p.PublisherID
}

// GetPublisherIDs Publication impl base.PublisherIDs
func (p Publication) GetPublisherIDs() []string {
	return []string{p.PublisherID}
}

// GetPublicationStatus Publication impl base.Status
func (p Publication) GetPublicationStatus() int {
	return p.Status
}

// GetPublicationStatuses Publication impl base.Statuses
func (p Publication) GetPublicationStatuses() []int {
	return []int{p.Status}
}

// GetUpdatedAt Publication impl base.UpdatedAt
func (p Publication) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

// GetUpdater Publication impl base.Updater
func (p Publication) GetUpdater() string {
	return p.Updater
}

// Subscription table name: t_subscription
type Subscription struct {
	CreatedAt    time.Time `gorm:"column:created_at;type:time;<-:create;autoCreateTime"`
	Creator      string    `gorm:"column:creator;type:varchar(255)"`
	DeletedAt    int64     `gorm:"column:deleted_at;default:0"`
	Description  string    `gorm:"column:description;type:varchar(4096);default null"`
	EventTypeID  string    `gorm:"column:name;type:varchar(255);NOT NULL"`
	ID           string    `gorm:"column:id;type:varchar(36);primaryKey;NOT NULL"`
	Status       int       `gorm:"column:status;type:int;NOT NULL;default:0"`
	SubscriberID string    `gorm:"column:main_app_id;type:varchar(36);NOT NULL"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:time;autoUpdateTime"`
	Updater      string    `gorm:"column:updater;type:varchar(255)"`
}

type Subscriptions []Subscription

func (p Subscriptions) GetSubscriptionIDs() []string {
	return base.ToSubscriptionIDs(p)
}
func (p Subscriptions) GetEventTypeIDs() []string {
	return base.ToEventTypeIDs(p)
}
func (p Subscriptions) GetSubscriberIDs() []string {
	return base.ToSubscriberIDs(p)
}

// TableName Subscription impl schema.Tabler
func (p Subscription) TableName() string {
	return "t_subscription"
}

// GetCreatedAt Subscription impl base.CreatedAt
func (p Subscription) GetCreatedAt() time.Time {
	return p.CreatedAt
}

// GetCreator Subscription impl base.Creator
func (p Subscription) GetCreator() string {
	return p.Creator
}

// GetDeletedAt Subscription impl base.DeletedAt
func (p Subscription) GetDeletedAt() int64 {
	return p.DeletedAt
}

// GetSubscriptionDescription Subscription impl base.Description
func (p Subscription) GetSubscriptionDescription() string {
	return p.Description
}

// GetEventTypeID Subscription impl base.EventTypeID
func (p Subscription) GetEventTypeID() string {
	return p.EventTypeID
}

// GetEventTypeIDs Subscription impl base.EventTypeIDs
func (p Subscription) GetEventTypeIDs() []string {
	return []string{p.EventTypeID}
}

// GetSubscriptionID Subscription impl base.ID
func (p Subscription) GetSubscriptionID() string {
	return p.ID
}

// GetSubscriptionIDs Subscription impl base.IDs
func (p Subscription) GetSubscriptionIDs() []string {
	return []string{p.ID}
}

// GetSubscriptionStatus Subscription impl base.Status
func (p Subscription) GetSubscriptionStatus() int {
	return p.Status
}

// GetSubscriptionStatuses Subscription impl base.Statuses
func (p Subscription) GetSubscriptionStatuses() []int {
	return []int{p.Status}
}

// GetSubscriberID Subscription impl base.SubscriberID
func (p Subscription) GetSubscriberID() string {
	return p.SubscriberID
}

// GetSubscriberIDs Subscription impl base.SubscriberIDs
func (p Subscription) GetSubscriberIDs() []string {
	return []string{p.SubscriberID}
}

// GetUpdatedAt Subscription impl base.UpdatedAt
func (p Subscription) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

// GetUpdater Subscription impl base.Updater
func (p Subscription) GetUpdater() string {
	return p.Updater
}
