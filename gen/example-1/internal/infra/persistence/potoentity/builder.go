// Package potoentity Code generated, DO NOT EDIT.
package potoentity

import (
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/base"
	"github.com/marlin5555/dddgen/gen/example-1/internal/infra/persistence/po"
)

// Option EntityBuilder option
type Option func(builder *EntityBuilder)

// WithFields set fields
func WithFields(fields base.Fields) Option {
	return func(builder *EntityBuilder) {
		builder.checker.addFields(fields)
	}
}

// WithAccounts set accounts
func WithAccounts(pValue map[string]po.Account) Option {
	return func(builder *EntityBuilder) {
		builder.poAccounts = pValue
	}
}

// WithApplications set applications
func WithApplications(pValue map[string]po.Application) Option {
	return func(builder *EntityBuilder) {
		builder.poApplications = pValue
	}
}

// WithEventBuses set eventBuses
func WithEventBuses(pValue map[string]po.EventBus) Option {
	return func(builder *EntityBuilder) {
		builder.poEventBuses = pValue
	}
}

// WithEventTypes set eventTypes
func WithEventTypes(pValue map[string]po.EventType) Option {
	return func(builder *EntityBuilder) {
		builder.poEventTypes = pValue
	}
}

// WithPassports set passports
func WithPassports(pValue map[string]po.Passport) Option {
	return func(builder *EntityBuilder) {
		builder.poPassports = pValue
	}
}

// WithPublications set publications
func WithPublications(pValue map[string]po.Publication) Option {
	return func(builder *EntityBuilder) {
		builder.poPublications = pValue
	}
}

// WithSecrets set secrets
func WithSecrets(pValue map[string]po.Secret) Option {
	return func(builder *EntityBuilder) {
		builder.poSecrets = pValue
	}
}

// WithSubscriptions set subscriptions
func WithSubscriptions(pValue map[string]po.Subscription) Option {
	return func(builder *EntityBuilder) {
		builder.poSubscriptions = pValue
	}
}

// WithTechRelations set techRelations
func WithTechRelations(pValue map[string]po.TechRelation) Option {
	return func(builder *EntityBuilder) {
		builder.poTechRelations = pValue
	}
}

// EntityBuilder builder support entity.Account/entity.Application/entity.EventBus/entity.EventType/entity.Passport/entity.Publication/entity.Secret/entity.Subscription/entity.TechRelation/
type EntityBuilder struct {
	/** account **/
	poAccounts map[string]po.Account
	enAccounts map[string]*accountEntity
	/** application **/
	poApplications map[string]po.Application
	enApplications map[string]*applicationEntity
	/** eventBus **/
	poEventBuses map[string]po.EventBus
	enEventBuses map[string]*eventBusEntity
	/** eventType **/
	poEventTypes map[string]po.EventType
	enEventTypes map[string]*eventTypeEntity
	/** passport **/
	poPassports map[string]po.Passport
	enPassports map[string]*passportEntity
	/** publication **/
	poPublications map[string]po.Publication
	enPublications map[string]*publicationEntity
	/** secret **/
	poSecrets map[string]po.Secret
	enSecrets map[string]*secretEntity
	/** subscription **/
	poSubscriptions map[string]po.Subscription
	enSubscriptions map[string]*subscriptionEntity
	/** techRelation **/
	poTechRelations map[string]po.TechRelation
	enTechRelations map[string]*techRelationEntity
	checker         fieldChecker
}

// NewBuilder builder constructor
func NewBuilder(fields base.Fields, opts ...Option) *EntityBuilder {
	b := &EntityBuilder{
		poAccounts:      map[string]po.Account{},
		enAccounts:      map[string]*accountEntity{},
		poApplications:  map[string]po.Application{},
		enApplications:  map[string]*applicationEntity{},
		poEventBuses:    map[string]po.EventBus{},
		enEventBuses:    map[string]*eventBusEntity{},
		poEventTypes:    map[string]po.EventType{},
		enEventTypes:    map[string]*eventTypeEntity{},
		poPassports:     map[string]po.Passport{},
		enPassports:     map[string]*passportEntity{},
		poPublications:  map[string]po.Publication{},
		enPublications:  map[string]*publicationEntity{},
		poSecrets:       map[string]po.Secret{},
		enSecrets:       map[string]*secretEntity{},
		poSubscriptions: map[string]po.Subscription{},
		enSubscriptions: map[string]*subscriptionEntity{},
		poTechRelations: map[string]po.TechRelation{},
		enTechRelations: map[string]*techRelationEntity{},
	}
	b.checker = buildChecker(fields)
	for _, opt := range opts {
		opt(b)
	}
	return b
}

// Build build entity and relation
func (b *EntityBuilder) Build() {
	b.buildEntity()
	b.buildRelation()
}

// GetAccounts get Account entities
func (b *EntityBuilder) GetAccounts() entity.Accounts {
	var r entity.Accounts
	for _, e := range b.enAccounts {
		r = append(r, e)
	}
	return r
}

// GetApplications get Application entities
func (b *EntityBuilder) GetApplications() entity.Applications {
	var r entity.Applications
	for _, e := range b.enApplications {
		r = append(r, e)
	}
	return r
}

// GetEventBuses get EventBus entities
func (b *EntityBuilder) GetEventBuses() entity.EventBuses {
	var r entity.EventBuses
	for _, e := range b.enEventBuses {
		r = append(r, e)
	}
	return r
}

// GetEventTypes get EventType entities
func (b *EntityBuilder) GetEventTypes() entity.EventTypes {
	var r entity.EventTypes
	for _, e := range b.enEventTypes {
		r = append(r, e)
	}
	return r
}

// GetPassports get Passport entities
func (b *EntityBuilder) GetPassports() entity.Passports {
	var r entity.Passports
	for _, e := range b.enPassports {
		r = append(r, e)
	}
	return r
}

// GetPublications get Publication entities
func (b *EntityBuilder) GetPublications() entity.Publications {
	var r entity.Publications
	for _, e := range b.enPublications {
		r = append(r, e)
	}
	return r
}

// GetSecrets get Secret entities
func (b *EntityBuilder) GetSecrets() entity.Secrets {
	var r entity.Secrets
	for _, e := range b.enSecrets {
		r = append(r, e)
	}
	return r
}

// GetSubscriptions get Subscription entities
func (b *EntityBuilder) GetSubscriptions() entity.Subscriptions {
	var r entity.Subscriptions
	for _, e := range b.enSubscriptions {
		r = append(r, e)
	}
	return r
}

// GetTechRelations get TechRelation entities
func (b *EntityBuilder) GetTechRelations() entity.TechRelations {
	var r entity.TechRelations
	for _, e := range b.enTechRelations {
		r = append(r, e)
	}
	return r
}

// buildEntity build entity
func (b *EntityBuilder) buildEntity() {
	b.buildAccountEntity()
	b.buildApplicationEntity()
	b.buildEventBusEntity()
	b.buildEventTypeEntity()
	b.buildPassportEntity()
	b.buildPublicationEntity()
	b.buildSecretEntity()
	b.buildSubscriptionEntity()
	b.buildTechRelationEntity()
}
func (b *EntityBuilder) buildAccountEntity() {
	for k, p := range b.poAccounts {
		b.enAccounts[k] = buildAccount(b.checker, p)
	}
}
func (b *EntityBuilder) buildApplicationEntity() {
	for k, p := range b.poApplications {
		b.enApplications[k] = buildApplication(b.checker, p)
	}
}
func (b *EntityBuilder) buildEventBusEntity() {
	for k, p := range b.poEventBuses {
		b.enEventBuses[k] = buildEventBus(b.checker, p)
	}
}
func (b *EntityBuilder) buildEventTypeEntity() {
	for k, p := range b.poEventTypes {
		b.enEventTypes[k] = buildEventType(b.checker, p)
	}
}
func (b *EntityBuilder) buildPassportEntity() {
	for k, p := range b.poPassports {
		b.enPassports[k] = buildPassport(b.checker, p)
	}
}
func (b *EntityBuilder) buildPublicationEntity() {
	for k, p := range b.poPublications {
		b.enPublications[k] = buildPublication(b.checker, p)
	}
}
func (b *EntityBuilder) buildSecretEntity() {
	for k, p := range b.poSecrets {
		b.enSecrets[k] = buildSecret(b.checker, p)
	}
}
func (b *EntityBuilder) buildSubscriptionEntity() {
	for k, p := range b.poSubscriptions {
		b.enSubscriptions[k] = buildSubscription(b.checker, p)
	}
}
func (b *EntityBuilder) buildTechRelationEntity() {
	for k, p := range b.poTechRelations {
		b.enTechRelations[k] = buildTechRelation(b.checker, p)
	}
}

// buildRelation build relation
func (b *EntityBuilder) buildRelation() {
	b.buildAccountRelation()
	b.buildApplicationRelation()
	b.buildEventBusRelation()
	b.buildEventTypeRelation()
	b.buildPassportRelation()
	b.buildPublicationRelation()
	b.buildSecretRelation()
	b.buildSubscriptionRelation()
	b.buildTechRelationRelation()
}
func (b *EntityBuilder) buildAccountRelation() {

	// 进行 Account - reference 关联
	for k, e := range b.enAccounts {
		if a, ok := b.enSecrets[b.poAccounts[k].ID]; ok {
			e.secretGetter = secretGetter{value: a}
		}
		if a, ok := b.enPassports[b.poAccounts[k].ID]; ok {
			e.passportGetter = passportGetter{value: a}
		}
	}
}
func (b *EntityBuilder) buildApplicationRelation() {

}
func (b *EntityBuilder) buildEventBusRelation() {

}
func (b *EntityBuilder) buildEventTypeRelation() {

	// 进行 EventType - reference 关联
	for k, e := range b.enEventTypes {
		if a, ok := b.enEventBuses[b.poEventTypes[k].EventBusID]; ok {
			e.eventBusGetter = eventBusGetter{value: a}
		}
		if a, ok := b.enApplications[b.poEventTypes[k].MainAppID]; ok {
			e.ownershipGetter = ownershipGetter{value: a}
		}
	}
}
func (b *EntityBuilder) buildPassportRelation() {

	// 进行 Passport - reference 关联
	for k, e := range b.enPassports {
		if a, ok := b.enAccounts[b.poPassports[k].AccountID]; ok {
			e.accountGetter = accountGetter{value: a}
		}
	}
}
func (b *EntityBuilder) buildPublicationRelation() {

	// 进行 Publication - reference 关联
	for k, e := range b.enPublications {
		if a, ok := b.enEventTypes[b.poPublications[k].EventTypeID]; ok {
			e.eventTypeGetter = eventTypeGetter{value: a}
		}
		if a, ok := b.enApplications[b.poPublications[k].PublisherID]; ok {
			e.publisherGetter = publisherGetter{value: a}
		}
	}
}
func (b *EntityBuilder) buildSecretRelation() {

	// 进行 Secret - reference 关联
	for k, e := range b.enSecrets {
		if a, ok := b.enAccounts[b.poSecrets[k].AccountID]; ok {
			e.accountGetter = accountGetter{value: a}
		}
	}
}
func (b *EntityBuilder) buildSubscriptionRelation() {

	// 进行 Subscription - reference 关联
	for k, e := range b.enSubscriptions {
		if a, ok := b.enEventTypes[b.poSubscriptions[k].EventTypeID]; ok {
			e.eventTypeGetter = eventTypeGetter{value: a}
		}
		if a, ok := b.enApplications[b.poSubscriptions[k].SubscriberID]; ok {
			e.subscriberGetter = subscriberGetter{value: a}
		}
	}
}
func (b *EntityBuilder) buildTechRelationRelation() {

	// 进行 TechRelation - reference 关联
	for k, e := range b.enTechRelations {
		if a, ok := b.enAccounts[b.poTechRelations[k].StudentID]; ok {
			e.studentGetter = studentGetter{value: a}
		}
		if a, ok := b.enAccounts[b.poTechRelations[k].TeacherID]; ok {
			e.teacherGetter = teacherGetter{value: a}
		}
	}
}
