// Package entity Code generated, DO NOT EDIT.
package entity

import "github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/base"

type (

	// EventBus entity
	EventBus interface {
		base.EventBusID
		base.EventBusName
		base.EventBusParams
		base.EventBusDescription
		Trailer
	}
	// EventBusGetter EventBus getter
	EventBusGetter interface {
		GetEventBus() EventBus
	}

	// EventBuses eventBus s
	EventBuses []EventBus
)
type (

	// EventType entity
	EventType interface {
		base.EventTypeID
		base.EventTypeName
		OwnershipGetter
		EventBusGetter
		base.EventTypeStatus
		base.EventTypeDescription
		Trailer
	}
	// EventTypeGetter EventType getter
	EventTypeGetter interface {
		GetEventType() EventType
	}

	// EventTypes eventType s
	EventTypes []EventType
)
type (

	// Publication entity
	Publication interface {
		base.PublicationID
		PublisherGetter
		EventTypeGetter
		base.PublicationStatus
		base.PublicationDescription
		Trailer
	}
	// PublicationGetter Publication getter
	PublicationGetter interface {
		GetPublication() Publication
	}

	// Publications publication s
	Publications []Publication
)
type (

	// Subscription entity
	Subscription interface {
		base.SubscriptionID
		SubscriberGetter
		EventTypeGetter
		base.SubscriptionStatus
		base.SubscriptionDescription
		Trailer
	}
	// SubscriptionGetter Subscription getter
	SubscriptionGetter interface {
		GetSubscription() Subscription
	}

	// Subscriptions subscription s
	Subscriptions []Subscription
)
