// Package rsp Code generated, DO NOT EDIT.
package rsp

import "github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity"

// EventBusesRsp combination Rsp Total EventBuses
type EventBusesRsp struct {
	Rsp
	Total uint32
	entity.EventBuses
}

// EventTypesRsp combination Rsp Total EventTypes
type EventTypesRsp struct {
	Rsp
	Total uint32
	entity.EventTypes
}

// PublicationsRsp combination Rsp Total Publications
type PublicationsRsp struct {
	Rsp
	Total uint32
	entity.Publications
}

// SubscriptionsRsp combination Rsp Total Subscriptions
type SubscriptionsRsp struct {
	Rsp
	Total uint32
	entity.Subscriptions
}
