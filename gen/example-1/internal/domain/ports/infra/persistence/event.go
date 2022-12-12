// Package persistence Code generated, DO NOT EDIT.
package persistence

import (
	"context"

	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/req"
	"github.com/marlin5555/dddgen/gen/example-1/internal/infra/persistence/po"
)

// EventBusDAO EventBus dao
type EventBusDAO interface {
	Transaction
	FieldChecker
	Get(context.Context, req.GetEventBusesPO) (po.EventBuses, uint32, error)
	Create(context.Context, req.CreateEventBusReq) (string, error)
	Update(context.Context, req.UpdateEventBusReq) error
	Delete(context.Context, req.DeleteEventBusReq) error
}

// EventTypeDAO EventType dao
type EventTypeDAO interface {
	Transaction
	FieldChecker
	Get(context.Context, req.GetEventTypesPO) (po.EventTypes, uint32, error)
	Create(context.Context, req.CreateEventTypeReq) (string, error)
	Update(context.Context, req.UpdateEventTypeReq) error
	Delete(context.Context, req.DeleteEventTypeReq) error
}

// PublicationDAO Publication dao
type PublicationDAO interface {
	Transaction
	FieldChecker
	Get(context.Context, req.GetPublicationsPO) (po.Publications, uint32, error)
	Create(context.Context, req.CreatePublicationReq) (string, error)
	Update(context.Context, req.UpdatePublicationReq) error
	Delete(context.Context, req.DeletePublicationReq) error
}

// SubscriptionDAO Subscription dao
type SubscriptionDAO interface {
	Transaction
	FieldChecker
	Get(context.Context, req.GetSubscriptionsPO) (po.Subscriptions, uint32, error)
	Create(context.Context, req.CreateSubscriptionReq) (string, error)
	Update(context.Context, req.UpdateSubscriptionReq) error
	Delete(context.Context, req.DeleteSubscriptionReq) error
}
