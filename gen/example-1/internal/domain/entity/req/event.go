// Package req Code generated, DO NOT EDIT.
package req

import "github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/base"

// get create update delete req for EventBus
type (
	// GetEventBusesNoForeign get EventBuses no foreign key
	GetEventBusesNoForeign interface {
		base.FuzzyEventBusName
		base.EventBusID
		base.EventBusIDs
		base.EventBusName
		base.EventBusNames
	}
	// GetEventBusesPure get EventBuses pure, 仅仅作用在单一实体表上
	GetEventBusesPure interface {
		base.FuzzyEventBusName
		base.EventBusID
		base.EventBusIDs
		base.EventBusName
		base.EventBusNames
	}
	// GetEventBusesPO get EventBuses po, 用于 persistence 上的 get 操作
	GetEventBusesPO interface {
		GetEventBusesPure
		base.PageOrderOperator
	}

	// GetEventBusesReq get EventBuses req, 用于对外 service 中使用
	GetEventBusesReq interface {
		GetEventBusesPure
		base.Fields
		base.PageOrderOperator
		base.Validator
	}

	DeleteEventBusReq interface {
		base.EventBusID
		base.Operator
		base.Validator
	}

	UpdateEventBusReq interface {
		base.EventBusID
		base.EventBusName
		base.EventBusParams
		base.Operator
		base.Validator
	}
	CreateEventBusReq interface {
		base.EventBusID
		base.EventBusName
		base.EventBusParams
		base.Operator
		base.Validator
	}
)

// get create update delete req for EventType
type (
	// GetEventTypesNoForeign get EventTypes no foreign key
	GetEventTypesNoForeign interface {
		base.FuzzyEventTypeName
		base.EventTypeID
		base.EventTypeIDs
		base.EventTypeName
		base.EventTypeNames
		base.EventTypeStatus
		base.EventTypeStatuses
	}
	// GetEventTypesPure get EventTypes pure, 仅仅作用在单一实体表上
	GetEventTypesPure interface {
		base.EventBusID
		base.EventBusIDs
		base.FuzzyEventTypeName
		base.EventTypeID
		base.EventTypeIDs
		base.OwnershipID
		base.OwnershipIDs
		base.EventTypeName
		base.EventTypeNames
		base.EventTypeStatus
		base.EventTypeStatuses
	}
	// GetEventTypesPO get EventTypes po, 用于 persistence 上的 get 操作
	GetEventTypesPO interface {
		GetEventTypesPure
		base.PageOrderOperator
	}

	// GetEventTypesReq get EventTypes req, 用于对外 service 中使用
	GetEventTypesReq interface {
		GetEventTypesPure

		// start. use for EventBus filter
		base.FuzzyEventBusName
		base.EventBusID
		base.EventBusIDs
		base.EventBusName
		base.EventBusNames
		// end.

		// start. use for Ownership filter
		base.FuzzyApplicationName
		base.FuzzyApplicationNickname
		base.ApplicationID
		base.ApplicationIDs
		base.ApplicationName
		base.ApplicationNames
		// end.

		base.Fields
		base.PageOrderOperator
		base.Validator
	}

	DeleteEventTypeReq interface {
		base.EventTypeID
		base.Operator
		base.Validator
	}

	UpdateEventTypeReq interface {
		base.EventBusID
		base.EventTypeID
		base.OwnershipID
		base.EventTypeName
		base.EventTypeStatus
		base.Operator
		base.Validator
	}
	CreateEventTypeReq interface {
		base.EventBusID
		base.EventTypeID
		base.OwnershipID
		base.EventTypeName
		base.EventTypeStatus
		base.Operator
		base.Validator
	}
)

// get create update delete req for Publication
type (
	// GetPublicationsNoForeign get Publications no foreign key
	GetPublicationsNoForeign interface {
		base.FuzzyPublicationEventTypeID
		base.PublicationID
		base.PublicationIDs
		base.PublicationStatus
		base.PublicationStatuses
	}
	// GetPublicationsPure get Publications pure, 仅仅作用在单一实体表上
	GetPublicationsPure interface {
		base.EventTypeID
		base.EventTypeIDs
		base.FuzzyPublicationEventTypeID
		base.PublicationID
		base.PublicationIDs
		base.PublisherID
		base.PublisherIDs
		base.PublicationStatus
		base.PublicationStatuses
	}
	// GetPublicationsPO get Publications po, 用于 persistence 上的 get 操作
	GetPublicationsPO interface {
		GetPublicationsPure
		base.PageOrderOperator
	}

	// GetPublicationsReq get Publications req, 用于对外 service 中使用
	GetPublicationsReq interface {
		GetPublicationsPure

		// start. use for EventType filter
		base.EventBusID
		base.EventBusIDs
		base.FuzzyEventTypeName
		base.EventTypeID
		base.EventTypeIDs
		base.OwnershipID
		base.OwnershipIDs
		base.EventTypeName
		base.EventTypeNames
		base.EventTypeStatus
		base.EventTypeStatuses
		// end.

		// start. use for Publisher filter
		base.FuzzyApplicationName
		base.FuzzyApplicationNickname
		base.ApplicationID
		base.ApplicationIDs
		base.ApplicationName
		base.ApplicationNames
		// end.

		base.Fields
		base.PageOrderOperator
		base.Validator
	}

	DeletePublicationReq interface {
		base.PublicationID
		base.Operator
		base.Validator
	}

	UpdatePublicationReq interface {
		base.EventTypeID
		base.PublicationID
		base.PublisherID
		base.PublicationStatus
		base.Operator
		base.Validator
	}
	CreatePublicationReq interface {
		base.EventTypeID
		base.PublicationID
		base.PublisherID
		base.PublicationStatus
		base.Operator
		base.Validator
	}
)

// get create update delete req for Subscription
type (
	// GetSubscriptionsNoForeign get Subscriptions no foreign key
	GetSubscriptionsNoForeign interface {
		base.FuzzySubscriptionEventTypeID
		base.SubscriptionID
		base.SubscriptionIDs
		base.SubscriptionStatus
		base.SubscriptionStatuses
	}
	// GetSubscriptionsPure get Subscriptions pure, 仅仅作用在单一实体表上
	GetSubscriptionsPure interface {
		base.EventTypeID
		base.EventTypeIDs
		base.FuzzySubscriptionEventTypeID
		base.SubscriptionID
		base.SubscriptionIDs
		base.SubscriptionStatus
		base.SubscriptionStatuses
		base.SubscriberID
		base.SubscriberIDs
	}
	// GetSubscriptionsPO get Subscriptions po, 用于 persistence 上的 get 操作
	GetSubscriptionsPO interface {
		GetSubscriptionsPure
		base.PageOrderOperator
	}

	// GetSubscriptionsReq get Subscriptions req, 用于对外 service 中使用
	GetSubscriptionsReq interface {
		GetSubscriptionsPure

		// start. use for EventType filter
		base.EventBusID
		base.EventBusIDs
		base.FuzzyEventTypeName
		base.EventTypeID
		base.EventTypeIDs
		base.OwnershipID
		base.OwnershipIDs
		base.EventTypeName
		base.EventTypeNames
		base.EventTypeStatus
		base.EventTypeStatuses
		// end.

		// start. use for Subscriber filter
		base.FuzzyApplicationName
		base.FuzzyApplicationNickname
		base.ApplicationID
		base.ApplicationIDs
		base.ApplicationName
		base.ApplicationNames
		// end.

		base.Fields
		base.PageOrderOperator
		base.Validator
	}

	DeleteSubscriptionReq interface {
		base.SubscriptionID
		base.Operator
		base.Validator
	}

	UpdateSubscriptionReq interface {
		base.EventTypeID
		base.SubscriptionID
		base.SubscriptionStatus
		base.SubscriberID
		base.Operator
		base.Validator
	}
	CreateSubscriptionReq interface {
		base.EventTypeID
		base.SubscriptionID
		base.SubscriptionStatus
		base.SubscriberID
		base.Operator
		base.Validator
	}
)
