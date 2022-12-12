// Package repo Code generated, DO NOT EDIT.
package repo

import (
	"context"

	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/base"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/req"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/zero"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/ports/infra/persistence"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/ports/repo"
	"github.com/marlin5555/dddgen/gen/example-1/internal/infra/persistence/po"
	"github.com/marlin5555/dddgen/gen/example-1/internal/infra/persistence/potoentity"
)

// EventRepo ...
type EventRepo struct {
	applicationDAO  persistence.ApplicationDAO
	eventBusDAO     persistence.EventBusDAO
	eventTypeDAO    persistence.EventTypeDAO
	publicationDAO  persistence.PublicationDAO
	subscriptionDAO persistence.SubscriptionDAO
}

// NewEventRepo ...
func NewEventRepo(applicationDAO persistence.ApplicationDAO, eventBusDAO persistence.EventBusDAO, eventTypeDAO persistence.EventTypeDAO, publicationDAO persistence.PublicationDAO, subscriptionDAO persistence.SubscriptionDAO) repo.EventRepository {
	return &EventRepo{
		applicationDAO:  applicationDAO,
		eventBusDAO:     eventBusDAO,
		eventTypeDAO:    eventTypeDAO,
		publicationDAO:  publicationDAO,
		subscriptionDAO: subscriptionDAO,
	}
}

// CreateEventBus create EventBus
func (r *EventRepo) CreateEventBus(ctx context.Context, req req.CreateEventBusReq) (string, error) {
	return r.eventBusDAO.Create(ctx, req)
}

// UpdateEventBus update EventBus
func (r *EventRepo) UpdateEventBus(ctx context.Context, req req.UpdateEventBusReq) error {
	return r.eventBusDAO.Update(ctx, req)
}

// DeleteEventBus delete EventBus
func (r *EventRepo) DeleteEventBus(ctx context.Context, req req.DeleteEventBusReq) error {
	return r.eventBusDAO.Delete(ctx, req)
}

// GetEventBuses general get EventBus method
func (r *EventRepo) GetEventBuses(ctx context.Context, rq req.GetEventBusesReq) (entity.EventBuses, uint32, error) {
	// 做主表查询，获取到主表记录
	result, total, err := r.eventBusDAO.Get(ctx, struct {
		req.GetEventBusesNoForeign
		base.PageOrderOperator // 分页查询信息
	}{
		GetEventBusesNoForeign: rq,
		PageOrderOperator:      rq,
	})
	// 关联表查询
	builder := potoentity.NewBuilder(rq,
		potoentity.WithEventBuses(base.ToStringKeyMap(result,
			func(i interface{}) string { return i.(base.EventBusID).GetEventBusID() })),
		potoentity.WithFields(r.eventBusDAO.BaseFields()))
	builder.Build()
	return builder.GetEventBuses(), total, err
}

func needEventBusIDs(req req.GetEventBusesPure) bool {
	if !base.IsZero(req.GetFuzzyEventBusName()) {
		return true
	}
	if !base.IsZero(req.GetEventBusID()) {
		return true
	}
	if !base.IsZero(req.GetEventBusIDs()) {
		return true
	}
	if !base.IsZero(req.GetEventBusName()) {
		return true
	}
	if !base.IsZero(req.GetEventBusNames()) {
		return true
	}
	return false
}

func getEventBusIDsByRef(ctx context.Context, dao persistence.EventBusDAO, pure req.GetEventBusesPure) (base.EventBusIDs, bool, error) {
	if !needEventBusIDs(pure) {
		return zero.EventBusIDs{}, false, nil
	}
	var eventBuses po.EventBuses
	var err error
	eventBuses, _, err = dao.Get(ctx, struct {
		zero.PageOrderOperator
		req.GetEventBusesPure
	}{GetEventBusesPure: pure})
	if len(eventBuses.GetEventBusIDs()) == 0 || err != nil {
		return zero.EventBusIDs{}, true, nil
	}
	return zero.EventBusIDs{Value: eventBuses.GetEventBusIDs()}, false, nil
}

func getEventBusesByIDsIfRelated(ctx context.Context, dao persistence.EventBusDAO, ids base.EventBusIDs, fields base.Fields) (po.EventBuses, error) {
	if !dao.IsRelated(fields) {
		return nil, nil
	}
	return getEventBusesByIDs(ctx, dao, ids)
}

func getEventBusesByIDs(ctx context.Context, dao persistence.EventBusDAO, ids base.EventBusIDs) (po.EventBuses, error) {
	var eventBuses po.EventBuses
	var err error
	if eventBuses, _, err = dao.Get(ctx, struct {
		zero.FuzzyEventBusName
		zero.EventBusID
		base.EventBusIDs
		zero.EventBusName
		zero.EventBusNames
		zero.PageOrderOperator
	}{EventBusIDs: ids}); err != nil {
		return nil, err
	}
	return eventBuses, err
}

// CreateEventType create EventType
func (r *EventRepo) CreateEventType(ctx context.Context, req req.CreateEventTypeReq) (string, error) {
	return r.eventTypeDAO.Create(ctx, req)
}

// UpdateEventType update EventType
func (r *EventRepo) UpdateEventType(ctx context.Context, req req.UpdateEventTypeReq) error {
	return r.eventTypeDAO.Update(ctx, req)
}

// DeleteEventType delete EventType
func (r *EventRepo) DeleteEventType(ctx context.Context, req req.DeleteEventTypeReq) error {
	return r.eventTypeDAO.Delete(ctx, req)
}

// GetEventTypes general get EventType method
func (r *EventRepo) GetEventTypes(ctx context.Context, rq req.GetEventTypesReq) (entity.EventTypes, uint32, error) {
	var err error
	var isStop bool
	var mainAppIDs base.ApplicationIDs
	if mainAppIDs, isStop, err = getApplicationIDsByRef(ctx, r.applicationDAO, rq); isStop {
		return nil, 0, err
	}
	var eventBusIDs base.EventBusIDs
	if eventBusIDs, isStop, err = getEventBusIDsByRef(ctx, r.eventBusDAO, rq); isStop {
		return nil, 0, err
	}
	// 做主表查询，获取到主表记录
	result, total, err := r.eventTypeDAO.Get(ctx, struct {
		req.GetEventTypesNoForeign
		base.PageOrderOperator // 分页查询信息
		zero.OwnershipID
		base.OwnershipIDs
		zero.EventBusID
		base.EventBusIDs
	}{
		GetEventTypesNoForeign: rq,
		PageOrderOperator:      rq,
		OwnershipIDs:           zero.OwnershipIDs{Value: mainAppIDs.GetApplicationIDs()},
		EventBusIDs:            zero.EventBusIDs{Value: eventBusIDs.GetEventBusIDs()},
	})
	// 关联表查询
	var ownerships po.Applications
	if ownerships, err = getApplicationsByIDsIfRelated(ctx, r.applicationDAO, base.OwnershipIDsToApplicationIDs(result), rq); err != nil {
		return nil, 0, err
	}
	var eventBuses po.EventBuses
	if eventBuses, err = getEventBusesByIDsIfRelated(ctx, r.eventBusDAO, result, rq); err != nil {
		return nil, 0, err
	}
	builder := potoentity.NewBuilder(rq,
		potoentity.WithEventTypes(base.ToStringKeyMap(result,
			func(i interface{}) string { return i.(base.EventTypeID).GetEventTypeID() })),
		potoentity.WithApplications(base.ToStringKeyMap(ownerships,
			func(i interface{}) string { return i.(base.ApplicationID).GetApplicationID() })),
		potoentity.WithEventBuses(base.ToStringKeyMap(eventBuses,
			func(i interface{}) string { return i.(base.EventBusID).GetEventBusID() })),
		potoentity.WithFields(r.eventTypeDAO.BaseFields()))
	builder.Build()
	return builder.GetEventTypes(), total, err
}

func needEventTypeIDs(req req.GetEventTypesPure) bool {
	if !base.IsZero(req.GetEventBusID()) {
		return true
	}
	if !base.IsZero(req.GetEventBusIDs()) {
		return true
	}
	if !base.IsZero(req.GetFuzzyEventTypeName()) {
		return true
	}
	if !base.IsZero(req.GetEventTypeID()) {
		return true
	}
	if !base.IsZero(req.GetEventTypeIDs()) {
		return true
	}
	if !base.IsZero(req.GetOwnershipID()) {
		return true
	}
	if !base.IsZero(req.GetOwnershipIDs()) {
		return true
	}
	if !base.IsZero(req.GetEventTypeName()) {
		return true
	}
	if !base.IsZero(req.GetEventTypeNames()) {
		return true
	}
	if !base.IsZero(req.GetEventTypeStatus()) {
		return true
	}
	if !base.IsZero(req.GetEventTypeStatuses()) {
		return true
	}
	return false
}

func getEventTypeIDsByRef(ctx context.Context, dao persistence.EventTypeDAO, pure req.GetEventTypesPure) (base.EventTypeIDs, bool, error) {
	if !needEventTypeIDs(pure) {
		return zero.EventTypeIDs{}, false, nil
	}
	var eventTypes po.EventTypes
	var err error
	eventTypes, _, err = dao.Get(ctx, struct {
		zero.PageOrderOperator
		req.GetEventTypesPure
	}{GetEventTypesPure: pure})
	if len(eventTypes.GetEventTypeIDs()) == 0 || err != nil {
		return zero.EventTypeIDs{}, true, nil
	}
	return zero.EventTypeIDs{Value: eventTypes.GetEventTypeIDs()}, false, nil
}

func getEventTypesByIDsIfRelated(ctx context.Context, dao persistence.EventTypeDAO, ids base.EventTypeIDs, fields base.Fields) (po.EventTypes, error) {
	if !dao.IsRelated(fields) {
		return nil, nil
	}
	return getEventTypesByIDs(ctx, dao, ids)
}

func getEventTypesByIDs(ctx context.Context, dao persistence.EventTypeDAO, ids base.EventTypeIDs) (po.EventTypes, error) {
	var eventTypes po.EventTypes
	var err error
	if eventTypes, _, err = dao.Get(ctx, struct {
		zero.EventBusID
		zero.EventBusIDs
		zero.FuzzyEventTypeName
		zero.EventTypeID
		base.EventTypeIDs
		zero.OwnershipID
		zero.OwnershipIDs
		zero.EventTypeName
		zero.EventTypeNames
		zero.EventTypeStatus
		zero.EventTypeStatuses
		zero.PageOrderOperator
	}{EventTypeIDs: ids}); err != nil {
		return nil, err
	}
	return eventTypes, err
}

// CreatePublication create Publication
func (r *EventRepo) CreatePublication(ctx context.Context, req req.CreatePublicationReq) (string, error) {
	return r.publicationDAO.Create(ctx, req)
}

// UpdatePublication update Publication
func (r *EventRepo) UpdatePublication(ctx context.Context, req req.UpdatePublicationReq) error {
	return r.publicationDAO.Update(ctx, req)
}

// DeletePublication delete Publication
func (r *EventRepo) DeletePublication(ctx context.Context, req req.DeletePublicationReq) error {
	return r.publicationDAO.Delete(ctx, req)
}

// GetPublications general get Publication method
func (r *EventRepo) GetPublications(ctx context.Context, rq req.GetPublicationsReq) (entity.Publications, uint32, error) {
	var err error
	var isStop bool
	var publisherIDs base.ApplicationIDs
	if publisherIDs, isStop, err = getApplicationIDsByRef(ctx, r.applicationDAO, rq); isStop {
		return nil, 0, err
	}
	var eventTypeIDs base.EventTypeIDs
	if eventTypeIDs, isStop, err = getEventTypeIDsByRef(ctx, r.eventTypeDAO, rq); isStop {
		return nil, 0, err
	}
	// 做主表查询，获取到主表记录
	result, total, err := r.publicationDAO.Get(ctx, struct {
		req.GetPublicationsNoForeign
		base.PageOrderOperator // 分页查询信息
		zero.PublisherID
		base.PublisherIDs
		zero.EventTypeID
		base.EventTypeIDs
	}{
		GetPublicationsNoForeign: rq,
		PageOrderOperator:        rq,
		PublisherIDs:             zero.PublisherIDs{Value: publisherIDs.GetApplicationIDs()},
		EventTypeIDs:             zero.EventTypeIDs{Value: eventTypeIDs.GetEventTypeIDs()},
	})
	// 关联表查询
	var publishers po.Applications
	if publishers, err = getApplicationsByIDsIfRelated(ctx, r.applicationDAO, base.PublisherIDsToApplicationIDs(result), rq); err != nil {
		return nil, 0, err
	}
	var eventTypes po.EventTypes
	if eventTypes, err = getEventTypesByIDsIfRelated(ctx, r.eventTypeDAO, result, rq); err != nil {
		return nil, 0, err
	}
	builder := potoentity.NewBuilder(rq,
		potoentity.WithPublications(base.ToStringKeyMap(result,
			func(i interface{}) string { return i.(base.PublicationID).GetPublicationID() })),
		potoentity.WithApplications(base.ToStringKeyMap(publishers,
			func(i interface{}) string { return i.(base.ApplicationID).GetApplicationID() })),
		potoentity.WithEventTypes(base.ToStringKeyMap(eventTypes,
			func(i interface{}) string { return i.(base.EventTypeID).GetEventTypeID() })),
		potoentity.WithFields(r.publicationDAO.BaseFields()))
	builder.Build()
	return builder.GetPublications(), total, err
}

func needPublicationIDs(req req.GetPublicationsPure) bool {
	if !base.IsZero(req.GetEventTypeID()) {
		return true
	}
	if !base.IsZero(req.GetEventTypeIDs()) {
		return true
	}
	if !base.IsZero(req.GetFuzzyPublicationEventTypeID()) {
		return true
	}
	if !base.IsZero(req.GetPublicationID()) {
		return true
	}
	if !base.IsZero(req.GetPublicationIDs()) {
		return true
	}
	if !base.IsZero(req.GetPublisherID()) {
		return true
	}
	if !base.IsZero(req.GetPublisherIDs()) {
		return true
	}
	if !base.IsZero(req.GetPublicationStatus()) {
		return true
	}
	if !base.IsZero(req.GetPublicationStatuses()) {
		return true
	}
	return false
}

func getPublicationIDsByRef(ctx context.Context, dao persistence.PublicationDAO, pure req.GetPublicationsPure) (base.PublicationIDs, bool, error) {
	if !needPublicationIDs(pure) {
		return zero.PublicationIDs{}, false, nil
	}
	var publications po.Publications
	var err error
	publications, _, err = dao.Get(ctx, struct {
		zero.PageOrderOperator
		req.GetPublicationsPure
	}{GetPublicationsPure: pure})
	if len(publications.GetPublicationIDs()) == 0 || err != nil {
		return zero.PublicationIDs{}, true, nil
	}
	return zero.PublicationIDs{Value: publications.GetPublicationIDs()}, false, nil
}

func getPublicationsByIDsIfRelated(ctx context.Context, dao persistence.PublicationDAO, ids base.PublicationIDs, fields base.Fields) (po.Publications, error) {
	if !dao.IsRelated(fields) {
		return nil, nil
	}
	return getPublicationsByIDs(ctx, dao, ids)
}

func getPublicationsByIDs(ctx context.Context, dao persistence.PublicationDAO, ids base.PublicationIDs) (po.Publications, error) {
	var publications po.Publications
	var err error
	if publications, _, err = dao.Get(ctx, struct {
		zero.EventTypeID
		zero.EventTypeIDs
		zero.FuzzyPublicationEventTypeID
		zero.PublicationID
		base.PublicationIDs
		zero.PublisherID
		zero.PublisherIDs
		zero.PublicationStatus
		zero.PublicationStatuses
		zero.PageOrderOperator
	}{PublicationIDs: ids}); err != nil {
		return nil, err
	}
	return publications, err
}

// CreateSubscription create Subscription
func (r *EventRepo) CreateSubscription(ctx context.Context, req req.CreateSubscriptionReq) (string, error) {
	return r.subscriptionDAO.Create(ctx, req)
}

// UpdateSubscription update Subscription
func (r *EventRepo) UpdateSubscription(ctx context.Context, req req.UpdateSubscriptionReq) error {
	return r.subscriptionDAO.Update(ctx, req)
}

// DeleteSubscription delete Subscription
func (r *EventRepo) DeleteSubscription(ctx context.Context, req req.DeleteSubscriptionReq) error {
	return r.subscriptionDAO.Delete(ctx, req)
}

// GetSubscriptions general get Subscription method
func (r *EventRepo) GetSubscriptions(ctx context.Context, rq req.GetSubscriptionsReq) (entity.Subscriptions, uint32, error) {
	var err error
	var isStop bool
	var subscriberIDs base.ApplicationIDs
	if subscriberIDs, isStop, err = getApplicationIDsByRef(ctx, r.applicationDAO, rq); isStop {
		return nil, 0, err
	}
	var eventTypeIDs base.EventTypeIDs
	if eventTypeIDs, isStop, err = getEventTypeIDsByRef(ctx, r.eventTypeDAO, rq); isStop {
		return nil, 0, err
	}
	// 做主表查询，获取到主表记录
	result, total, err := r.subscriptionDAO.Get(ctx, struct {
		req.GetSubscriptionsNoForeign
		base.PageOrderOperator // 分页查询信息
		zero.SubscriberID
		base.SubscriberIDs
		zero.EventTypeID
		base.EventTypeIDs
	}{
		GetSubscriptionsNoForeign: rq,
		PageOrderOperator:         rq,
		SubscriberIDs:             zero.SubscriberIDs{Value: subscriberIDs.GetApplicationIDs()},
		EventTypeIDs:              zero.EventTypeIDs{Value: eventTypeIDs.GetEventTypeIDs()},
	})
	// 关联表查询
	var subscribers po.Applications
	if subscribers, err = getApplicationsByIDsIfRelated(ctx, r.applicationDAO, base.SubscriberIDsToApplicationIDs(result), rq); err != nil {
		return nil, 0, err
	}
	var eventTypes po.EventTypes
	if eventTypes, err = getEventTypesByIDsIfRelated(ctx, r.eventTypeDAO, result, rq); err != nil {
		return nil, 0, err
	}
	builder := potoentity.NewBuilder(rq,
		potoentity.WithSubscriptions(base.ToStringKeyMap(result,
			func(i interface{}) string { return i.(base.SubscriptionID).GetSubscriptionID() })),
		potoentity.WithApplications(base.ToStringKeyMap(subscribers,
			func(i interface{}) string { return i.(base.ApplicationID).GetApplicationID() })),
		potoentity.WithEventTypes(base.ToStringKeyMap(eventTypes,
			func(i interface{}) string { return i.(base.EventTypeID).GetEventTypeID() })),
		potoentity.WithFields(r.subscriptionDAO.BaseFields()))
	builder.Build()
	return builder.GetSubscriptions(), total, err
}

func needSubscriptionIDs(req req.GetSubscriptionsPure) bool {
	if !base.IsZero(req.GetEventTypeID()) {
		return true
	}
	if !base.IsZero(req.GetEventTypeIDs()) {
		return true
	}
	if !base.IsZero(req.GetFuzzySubscriptionEventTypeID()) {
		return true
	}
	if !base.IsZero(req.GetSubscriptionID()) {
		return true
	}
	if !base.IsZero(req.GetSubscriptionIDs()) {
		return true
	}
	if !base.IsZero(req.GetSubscriptionStatus()) {
		return true
	}
	if !base.IsZero(req.GetSubscriptionStatuses()) {
		return true
	}
	if !base.IsZero(req.GetSubscriberID()) {
		return true
	}
	if !base.IsZero(req.GetSubscriberIDs()) {
		return true
	}
	return false
}

func getSubscriptionIDsByRef(ctx context.Context, dao persistence.SubscriptionDAO, pure req.GetSubscriptionsPure) (base.SubscriptionIDs, bool, error) {
	if !needSubscriptionIDs(pure) {
		return zero.SubscriptionIDs{}, false, nil
	}
	var subscriptions po.Subscriptions
	var err error
	subscriptions, _, err = dao.Get(ctx, struct {
		zero.PageOrderOperator
		req.GetSubscriptionsPure
	}{GetSubscriptionsPure: pure})
	if len(subscriptions.GetSubscriptionIDs()) == 0 || err != nil {
		return zero.SubscriptionIDs{}, true, nil
	}
	return zero.SubscriptionIDs{Value: subscriptions.GetSubscriptionIDs()}, false, nil
}

func getSubscriptionsByIDsIfRelated(ctx context.Context, dao persistence.SubscriptionDAO, ids base.SubscriptionIDs, fields base.Fields) (po.Subscriptions, error) {
	if !dao.IsRelated(fields) {
		return nil, nil
	}
	return getSubscriptionsByIDs(ctx, dao, ids)
}

func getSubscriptionsByIDs(ctx context.Context, dao persistence.SubscriptionDAO, ids base.SubscriptionIDs) (po.Subscriptions, error) {
	var subscriptions po.Subscriptions
	var err error
	if subscriptions, _, err = dao.Get(ctx, struct {
		zero.EventTypeID
		zero.EventTypeIDs
		zero.FuzzySubscriptionEventTypeID
		zero.SubscriptionID
		base.SubscriptionIDs
		zero.SubscriptionStatus
		zero.SubscriptionStatuses
		zero.SubscriberID
		zero.SubscriberIDs
		zero.PageOrderOperator
	}{SubscriptionIDs: ids}); err != nil {
		return nil, err
	}
	return subscriptions, err
}
