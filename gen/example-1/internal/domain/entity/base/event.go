// Package base Code generated, DO NOT EDIT.
package base

// base interface for EventBus
type (
	// EventBusDescription interface has GetEventBusDescription
	EventBusDescription interface {
		GetEventBusDescription() string
	}
	// FuzzyEventBusName interface has GetFuzzyEventBusName
	FuzzyEventBusName interface {
		GetFuzzyEventBusName() string
	}
	// EventBusID interface has GetEventBusID
	EventBusID interface {
		GetEventBusID() string
	}
	// EventBusIDs interface has GetEventBusIDs
	EventBusIDs interface {
		GetEventBusIDs() []string
	}
	// EventBusName interface has GetEventBusName
	EventBusName interface {
		GetEventBusName() string
	}
	// EventBusNames interface has GetEventBusNames
	EventBusNames interface {
		GetEventBusNames() []string
	}
	// EventBusParams interface has GetEventBusParams
	EventBusParams interface {
		GetEventBusParams() string
	}
)

// base interface for EventType
type (
	// EventTypeDescription interface has GetEventTypeDescription
	EventTypeDescription interface {
		GetEventTypeDescription() string
	}
	// FuzzyEventTypeName interface has GetFuzzyEventTypeName
	FuzzyEventTypeName interface {
		GetFuzzyEventTypeName() string
	}
	// EventTypeID interface has GetEventTypeID
	EventTypeID interface {
		GetEventTypeID() string
	}
	// EventTypeIDs interface has GetEventTypeIDs
	EventTypeIDs interface {
		GetEventTypeIDs() []string
	}
	// OwnershipID interface has GetOwnershipID
	OwnershipID interface {
		GetOwnershipID() string
	}
	// OwnershipIDs interface has GetOwnershipIDs
	OwnershipIDs interface {
		GetOwnershipIDs() []string
	}
	// EventTypeName interface has GetEventTypeName
	EventTypeName interface {
		GetEventTypeName() string
	}
	// EventTypeNames interface has GetEventTypeNames
	EventTypeNames interface {
		GetEventTypeNames() []string
	}
	// EventTypeStatus interface has GetEventTypeStatus
	EventTypeStatus interface {
		GetEventTypeStatus() int
	}
	// EventTypeStatuses interface has GetEventTypeStatuses
	EventTypeStatuses interface {
		GetEventTypeStatuses() []int
	}
)

// base interface for Publication
type (
	// PublicationDescription interface has GetPublicationDescription
	PublicationDescription interface {
		GetPublicationDescription() string
	}
	// FuzzyPublicationEventTypeID interface has GetFuzzyPublicationEventTypeID
	FuzzyPublicationEventTypeID interface {
		GetFuzzyPublicationEventTypeID() string
	}
	// PublicationID interface has GetPublicationID
	PublicationID interface {
		GetPublicationID() string
	}
	// PublicationIDs interface has GetPublicationIDs
	PublicationIDs interface {
		GetPublicationIDs() []string
	}
	// PublisherID interface has GetPublisherID
	PublisherID interface {
		GetPublisherID() string
	}
	// PublisherIDs interface has GetPublisherIDs
	PublisherIDs interface {
		GetPublisherIDs() []string
	}
	// PublicationStatus interface has GetPublicationStatus
	PublicationStatus interface {
		GetPublicationStatus() int
	}
	// PublicationStatuses interface has GetPublicationStatuses
	PublicationStatuses interface {
		GetPublicationStatuses() []int
	}
)

// base interface for Subscription
type (
	// SubscriptionDescription interface has GetSubscriptionDescription
	SubscriptionDescription interface {
		GetSubscriptionDescription() string
	}
	// FuzzySubscriptionEventTypeID interface has GetFuzzySubscriptionEventTypeID
	FuzzySubscriptionEventTypeID interface {
		GetFuzzySubscriptionEventTypeID() string
	}
	// SubscriptionID interface has GetSubscriptionID
	SubscriptionID interface {
		GetSubscriptionID() string
	}
	// SubscriptionIDs interface has GetSubscriptionIDs
	SubscriptionIDs interface {
		GetSubscriptionIDs() []string
	}
	// SubscriptionStatus interface has GetSubscriptionStatus
	SubscriptionStatus interface {
		GetSubscriptionStatus() int
	}
	// SubscriptionStatuses interface has GetSubscriptionStatuses
	SubscriptionStatuses interface {
		GetSubscriptionStatuses() []int
	}
	// SubscriberID interface has GetSubscriberID
	SubscriberID interface {
		GetSubscriberID() string
	}
	// SubscriberIDs interface has GetSubscriberIDs
	SubscriberIDs interface {
		GetSubscriberIDs() []string
	}
)

// ToEventBusIDs convert EventBusID s to IDs
func ToEventBusIDs[T EventBusID](arrays []T) []string {
	return ToStringSet(arrays, func(i interface{}) string { return i.(T).GetEventBusID() })
}

// ToEventTypeIDs convert EventTypeID s to IDs
func ToEventTypeIDs[T EventTypeID](arrays []T) []string {
	return ToStringSet(arrays, func(i interface{}) string { return i.(T).GetEventTypeID() })
}

// ToPublicationIDs convert PublicationID s to IDs
func ToPublicationIDs[T PublicationID](arrays []T) []string {
	return ToStringSet(arrays, func(i interface{}) string { return i.(T).GetPublicationID() })
}

// ToSubscriptionIDs convert SubscriptionID s to IDs
func ToSubscriptionIDs[T SubscriptionID](arrays []T) []string {
	return ToStringSet(arrays, func(i interface{}) string { return i.(T).GetSubscriptionID() })
}
