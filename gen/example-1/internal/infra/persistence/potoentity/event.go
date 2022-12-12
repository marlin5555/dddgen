// Package potoentity Code generated, DO NOT EDIT.
package potoentity

import (
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/base"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/zero"
	"github.com/marlin5555/dddgen/gen/example-1/internal/infra/persistence"
	"github.com/marlin5555/dddgen/gen/example-1/internal/infra/persistence/po"
	"github.com/marlin5555/dddgen/gen/example-1/pkg/log"
)

// eventBusGetter impl entity.EventBusGetter
type eventBusGetter struct {
	value entity.EventBus
}

// GetEventBus return entity.EventBus
func (g eventBusGetter) GetEventBus() entity.EventBus {
	return g.value
}

// eventTypeGetter impl entity.EventTypeGetter
type eventTypeGetter struct {
	value entity.EventType
}

// GetEventType return entity.EventType
func (g eventTypeGetter) GetEventType() entity.EventType {
	return g.value
}

// ownershipGetter impl entity.OwnershipGetter
type ownershipGetter struct {
	value entity.Application
}

// GetOwnership return entity.Application
func (g ownershipGetter) GetOwnership() entity.Application {
	return g.value
}

// publicationGetter impl entity.PublicationGetter
type publicationGetter struct {
	value entity.Publication
}

// GetPublication return entity.Publication
func (g publicationGetter) GetPublication() entity.Publication {
	return g.value
}

// publisherGetter impl entity.PublisherGetter
type publisherGetter struct {
	value entity.Application
}

// GetPublisher return entity.Application
func (g publisherGetter) GetPublisher() entity.Application {
	return g.value
}

// subscriptionGetter impl entity.SubscriptionGetter
type subscriptionGetter struct {
	value entity.Subscription
}

// GetSubscription return entity.Subscription
func (g subscriptionGetter) GetSubscription() entity.Subscription {
	return g.value
}

// subscriberGetter impl entity.SubscriberGetter
type subscriberGetter struct {
	value entity.Application
}

// GetSubscriber return entity.Application
func (g subscriberGetter) GetSubscriber() entity.Application {
	return g.value
}

// eventBusEntity eventBus entity impl entity.EventBus
type eventBusEntity struct {
	base.CreatedAt
	base.Creator
	base.DeletedAt
	base.EventBusDescription
	base.EventBusID
	base.EventBusName
	base.EventBusParams
	base.UpdatedAt
	base.Updater
}

func (e *eventBusEntity) setCreatedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.CreatedAt); ok {
			e.CreatedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.CreatedAt", value)
	}
	// 零值处理
	e.CreatedAt = zero.CreatedAt{}
}
func (e *eventBusEntity) setCreator(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.Creator); ok {
			e.Creator = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.Creator", value)
	}
	// 零值处理
	e.Creator = zero.Creator{}
}
func (e *eventBusEntity) setDeletedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.DeletedAt); ok {
			e.DeletedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.DeletedAt", value)
	}
	// 零值处理
	e.DeletedAt = zero.DeletedAt{}
}
func (e *eventBusEntity) setEventBusDescription(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.EventBusDescription); ok {
			e.EventBusDescription = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.EventBusDescription", value)
	}
	// 零值处理
	e.EventBusDescription = zero.EventBusDescription{}
}
func (e *eventBusEntity) setEventBusID(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.EventBusID); ok {
			e.EventBusID = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.EventBusID", value)
	}
	// 零值处理
	e.EventBusID = zero.EventBusID{}
}
func (e *eventBusEntity) setEventBusName(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.EventBusName); ok {
			e.EventBusName = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.EventBusName", value)
	}
	// 零值处理
	e.EventBusName = zero.EventBusName{}
}
func (e *eventBusEntity) setEventBusParams(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.EventBusParams); ok {
			e.EventBusParams = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.EventBusParams", value)
	}
	// 零值处理
	e.EventBusParams = zero.EventBusParams{}
}
func (e *eventBusEntity) setUpdatedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.UpdatedAt); ok {
			e.UpdatedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.UpdatedAt", value)
	}
	// 零值处理
	e.UpdatedAt = zero.UpdatedAt{}
}
func (e *eventBusEntity) setUpdater(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.Updater); ok {
			e.Updater = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.Updater", value)
	}
	// 零值处理
	e.Updater = zero.Updater{}
}

// buildEventBus build entity
func buildEventBus(checker fieldChecker, po po.EventBus) *eventBusEntity {
	e := &eventBusEntity{}
	t := newTable(persistence.T_EventBus)
	for _, tup := range []tuple{
		{pred: t.col("created_at").check(checker.hasColumn), attr: po, do: e.setCreatedAt},
		{pred: t.col("creator").check(checker.hasColumn), attr: po, do: e.setCreator},
		{pred: t.col("deleted_at").check(checker.hasColumn), attr: po, do: e.setDeletedAt},
		{pred: t.col("description").check(checker.hasColumn), attr: po, do: e.setEventBusDescription},
		{pred: t.col("id").check(checker.hasColumn), attr: po, do: e.setEventBusID},
		{pred: t.col("name").check(checker.hasColumn), attr: po, do: e.setEventBusName},
		{pred: t.col("params").check(checker.hasColumn), attr: po, do: e.setEventBusParams},
		{pred: t.col("updated_at").check(checker.hasColumn), attr: po, do: e.setUpdatedAt},
		{pred: t.col("updater").check(checker.hasColumn), attr: po, do: e.setUpdater},
	} {
		builderAppend(tup)
	}
	return e
}

// eventTypeEntity eventType entity impl entity.EventType
type eventTypeEntity struct {
	base.CreatedAt
	base.Creator
	base.DeletedAt
	base.EventTypeDescription
	eventBusGetter
	base.EventTypeID
	ownershipGetter
	base.EventTypeName
	base.EventTypeStatus
	base.UpdatedAt
	base.Updater
}

func (e *eventTypeEntity) setCreatedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.CreatedAt); ok {
			e.CreatedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.CreatedAt", value)
	}
	// 零值处理
	e.CreatedAt = zero.CreatedAt{}
}
func (e *eventTypeEntity) setCreator(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.Creator); ok {
			e.Creator = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.Creator", value)
	}
	// 零值处理
	e.Creator = zero.Creator{}
}
func (e *eventTypeEntity) setDeletedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.DeletedAt); ok {
			e.DeletedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.DeletedAt", value)
	}
	// 零值处理
	e.DeletedAt = zero.DeletedAt{}
}
func (e *eventTypeEntity) setEventTypeDescription(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.EventTypeDescription); ok {
			e.EventTypeDescription = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.EventTypeDescription", value)
	}
	// 零值处理
	e.EventTypeDescription = zero.EventTypeDescription{}
}
func (e *eventTypeEntity) setEventTypeID(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.EventTypeID); ok {
			e.EventTypeID = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.EventTypeID", value)
	}
	// 零值处理
	e.EventTypeID = zero.EventTypeID{}
}
func (e *eventTypeEntity) setEventTypeName(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.EventTypeName); ok {
			e.EventTypeName = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.EventTypeName", value)
	}
	// 零值处理
	e.EventTypeName = zero.EventTypeName{}
}
func (e *eventTypeEntity) setEventTypeStatus(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.EventTypeStatus); ok {
			e.EventTypeStatus = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.EventTypeStatus", value)
	}
	// 零值处理
	e.EventTypeStatus = zero.EventTypeStatus{}
}
func (e *eventTypeEntity) setUpdatedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.UpdatedAt); ok {
			e.UpdatedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.UpdatedAt", value)
	}
	// 零值处理
	e.UpdatedAt = zero.UpdatedAt{}
}
func (e *eventTypeEntity) setUpdater(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.Updater); ok {
			e.Updater = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.Updater", value)
	}
	// 零值处理
	e.Updater = zero.Updater{}
}

// buildEventType build entity
func buildEventType(checker fieldChecker, po po.EventType) *eventTypeEntity {
	e := &eventTypeEntity{}
	t := newTable(persistence.T_EventType)
	for _, tup := range []tuple{
		{pred: t.col("created_at").check(checker.hasColumn), attr: po, do: e.setCreatedAt},
		{pred: t.col("creator").check(checker.hasColumn), attr: po, do: e.setCreator},
		{pred: t.col("deleted_at").check(checker.hasColumn), attr: po, do: e.setDeletedAt},
		{pred: t.col("description").check(checker.hasColumn), attr: po, do: e.setEventTypeDescription},
		{pred: t.col("id").check(checker.hasColumn), attr: po, do: e.setEventTypeID},
		{pred: t.col("name").check(checker.hasColumn), attr: po, do: e.setEventTypeName},
		{pred: t.col("status").check(checker.hasColumn), attr: po, do: e.setEventTypeStatus},
		{pred: t.col("updated_at").check(checker.hasColumn), attr: po, do: e.setUpdatedAt},
		{pred: t.col("updater").check(checker.hasColumn), attr: po, do: e.setUpdater},
	} {
		builderAppend(tup)
	}
	return e
}

// publicationEntity publication entity impl entity.Publication
type publicationEntity struct {
	base.CreatedAt
	base.Creator
	base.DeletedAt
	base.PublicationDescription
	eventTypeGetter
	base.PublicationID
	publisherGetter
	base.PublicationStatus
	base.UpdatedAt
	base.Updater
}

func (e *publicationEntity) setCreatedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.CreatedAt); ok {
			e.CreatedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.CreatedAt", value)
	}
	// 零值处理
	e.CreatedAt = zero.CreatedAt{}
}
func (e *publicationEntity) setCreator(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.Creator); ok {
			e.Creator = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.Creator", value)
	}
	// 零值处理
	e.Creator = zero.Creator{}
}
func (e *publicationEntity) setDeletedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.DeletedAt); ok {
			e.DeletedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.DeletedAt", value)
	}
	// 零值处理
	e.DeletedAt = zero.DeletedAt{}
}
func (e *publicationEntity) setPublicationDescription(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.PublicationDescription); ok {
			e.PublicationDescription = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.PublicationDescription", value)
	}
	// 零值处理
	e.PublicationDescription = zero.PublicationDescription{}
}
func (e *publicationEntity) setPublicationID(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.PublicationID); ok {
			e.PublicationID = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.PublicationID", value)
	}
	// 零值处理
	e.PublicationID = zero.PublicationID{}
}
func (e *publicationEntity) setPublicationStatus(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.PublicationStatus); ok {
			e.PublicationStatus = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.PublicationStatus", value)
	}
	// 零值处理
	e.PublicationStatus = zero.PublicationStatus{}
}
func (e *publicationEntity) setUpdatedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.UpdatedAt); ok {
			e.UpdatedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.UpdatedAt", value)
	}
	// 零值处理
	e.UpdatedAt = zero.UpdatedAt{}
}
func (e *publicationEntity) setUpdater(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.Updater); ok {
			e.Updater = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.Updater", value)
	}
	// 零值处理
	e.Updater = zero.Updater{}
}

// buildPublication build entity
func buildPublication(checker fieldChecker, po po.Publication) *publicationEntity {
	e := &publicationEntity{}
	t := newTable(persistence.T_Publication)
	for _, tup := range []tuple{
		{pred: t.col("created_at").check(checker.hasColumn), attr: po, do: e.setCreatedAt},
		{pred: t.col("creator").check(checker.hasColumn), attr: po, do: e.setCreator},
		{pred: t.col("deleted_at").check(checker.hasColumn), attr: po, do: e.setDeletedAt},
		{pred: t.col("description").check(checker.hasColumn), attr: po, do: e.setPublicationDescription},
		{pred: t.col("id").check(checker.hasColumn), attr: po, do: e.setPublicationID},
		{pred: t.col("status").check(checker.hasColumn), attr: po, do: e.setPublicationStatus},
		{pred: t.col("updated_at").check(checker.hasColumn), attr: po, do: e.setUpdatedAt},
		{pred: t.col("updater").check(checker.hasColumn), attr: po, do: e.setUpdater},
	} {
		builderAppend(tup)
	}
	return e
}

// subscriptionEntity subscription entity impl entity.Subscription
type subscriptionEntity struct {
	base.CreatedAt
	base.Creator
	base.DeletedAt
	base.SubscriptionDescription
	eventTypeGetter
	base.SubscriptionID
	base.SubscriptionStatus
	subscriberGetter
	base.UpdatedAt
	base.Updater
}

func (e *subscriptionEntity) setCreatedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.CreatedAt); ok {
			e.CreatedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.CreatedAt", value)
	}
	// 零值处理
	e.CreatedAt = zero.CreatedAt{}
}
func (e *subscriptionEntity) setCreator(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.Creator); ok {
			e.Creator = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.Creator", value)
	}
	// 零值处理
	e.Creator = zero.Creator{}
}
func (e *subscriptionEntity) setDeletedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.DeletedAt); ok {
			e.DeletedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.DeletedAt", value)
	}
	// 零值处理
	e.DeletedAt = zero.DeletedAt{}
}
func (e *subscriptionEntity) setSubscriptionDescription(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.SubscriptionDescription); ok {
			e.SubscriptionDescription = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.SubscriptionDescription", value)
	}
	// 零值处理
	e.SubscriptionDescription = zero.SubscriptionDescription{}
}
func (e *subscriptionEntity) setSubscriptionID(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.SubscriptionID); ok {
			e.SubscriptionID = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.SubscriptionID", value)
	}
	// 零值处理
	e.SubscriptionID = zero.SubscriptionID{}
}
func (e *subscriptionEntity) setSubscriptionStatus(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.SubscriptionStatus); ok {
			e.SubscriptionStatus = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.SubscriptionStatus", value)
	}
	// 零值处理
	e.SubscriptionStatus = zero.SubscriptionStatus{}
}
func (e *subscriptionEntity) setUpdatedAt(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.UpdatedAt); ok {
			e.UpdatedAt = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.UpdatedAt", value)
	}
	// 零值处理
	e.UpdatedAt = zero.UpdatedAt{}
}
func (e *subscriptionEntity) setUpdater(pred bool, value interface{}) {
	if pred {
		if w, ok := value.(base.Updater); ok {
			e.Updater = w
			return
		}
		log.InfofWithFuncName("value[%+v] not a base.Updater", value)
	}
	// 零值处理
	e.Updater = zero.Updater{}
}

// buildSubscription build entity
func buildSubscription(checker fieldChecker, po po.Subscription) *subscriptionEntity {
	e := &subscriptionEntity{}
	t := newTable(persistence.T_Subscription)
	for _, tup := range []tuple{
		{pred: t.col("created_at").check(checker.hasColumn), attr: po, do: e.setCreatedAt},
		{pred: t.col("creator").check(checker.hasColumn), attr: po, do: e.setCreator},
		{pred: t.col("deleted_at").check(checker.hasColumn), attr: po, do: e.setDeletedAt},
		{pred: t.col("description").check(checker.hasColumn), attr: po, do: e.setSubscriptionDescription},
		{pred: t.col("id").check(checker.hasColumn), attr: po, do: e.setSubscriptionID},
		{pred: t.col("status").check(checker.hasColumn), attr: po, do: e.setSubscriptionStatus},
		{pred: t.col("updated_at").check(checker.hasColumn), attr: po, do: e.setUpdatedAt},
		{pred: t.col("updater").check(checker.hasColumn), attr: po, do: e.setUpdater},
	} {
		builderAppend(tup)
	}
	return e
}
