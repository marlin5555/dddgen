// Package repo Code generated, DO NOT EDIT.
package repo

import (
	"context"

	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/req"
)

// EventRepository Event repo
type EventRepository interface {
	// CreateEventBus create EventBus
	CreateEventBus(context.Context, req.CreateEventBusReq) (string, error)
	// UpdateEventBus update EventBus
	UpdateEventBus(context.Context, req.UpdateEventBusReq) error
	// DeleteEventBus delete EventBus
	DeleteEventBus(context.Context, req.DeleteEventBusReq) error
	// GetEventBuses general get EventBus method
	GetEventBuses(context.Context, req.GetEventBusesReq) (entity.EventBuses, uint32, error)

	// CreateEventType create EventType
	CreateEventType(context.Context, req.CreateEventTypeReq) (string, error)
	// UpdateEventType update EventType
	UpdateEventType(context.Context, req.UpdateEventTypeReq) error
	// DeleteEventType delete EventType
	DeleteEventType(context.Context, req.DeleteEventTypeReq) error
	// GetEventTypes general get EventType method
	GetEventTypes(context.Context, req.GetEventTypesReq) (entity.EventTypes, uint32, error)

	// CreatePublication create Publication
	CreatePublication(context.Context, req.CreatePublicationReq) (string, error)
	// UpdatePublication update Publication
	UpdatePublication(context.Context, req.UpdatePublicationReq) error
	// DeletePublication delete Publication
	DeletePublication(context.Context, req.DeletePublicationReq) error
	// GetPublications general get Publication method
	GetPublications(context.Context, req.GetPublicationsReq) (entity.Publications, uint32, error)

	// CreateSubscription create Subscription
	CreateSubscription(context.Context, req.CreateSubscriptionReq) (string, error)
	// UpdateSubscription update Subscription
	UpdateSubscription(context.Context, req.UpdateSubscriptionReq) error
	// DeleteSubscription delete Subscription
	DeleteSubscription(context.Context, req.DeleteSubscriptionReq) error
	// GetSubscriptions general get Subscription method
	GetSubscriptions(context.Context, req.GetSubscriptionsReq) (entity.Subscriptions, uint32, error)
}
