// Package entity Code generated, DO NOT EDIT.
package entity

import "github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/base"

type (

	// Application entity
	Application interface {
		base.ApplicationID
		base.ApplicationName
		base.ApplicationNickname
		base.ApplicationDescription
		Trailer
	}
	// ApplicationGetter Application getter
	ApplicationGetter interface {
		GetApplication() Application
	}
	OwnershipGetter interface {
		GetOwnership() Application
	}
	SubscriberGetter interface {
		GetSubscriber() Application
	}
	PublisherGetter interface {
		GetPublisher() Application
	}

	// Applications application s
	Applications []Application
)
