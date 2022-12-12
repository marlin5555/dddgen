// Package base Code generated, DO NOT EDIT.
package base

import (
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/zero"
)

// base interface for Application
type (
	// ApplicationDescription interface has GetApplicationDescription
	ApplicationDescription interface {
		GetApplicationDescription() string
	}
	// FuzzyApplicationName interface has GetFuzzyApplicationName
	FuzzyApplicationName interface {
		GetFuzzyApplicationName() string
	}
	// FuzzyApplicationNickname interface has GetFuzzyApplicationNickname
	FuzzyApplicationNickname interface {
		GetFuzzyApplicationNickname() string
	}
	// ApplicationID interface has GetApplicationID
	ApplicationID interface {
		GetApplicationID() string
	}
	// ApplicationIDs interface has GetApplicationIDs
	ApplicationIDs interface {
		GetApplicationIDs() []string
	}
	// ApplicationName interface has GetApplicationName
	ApplicationName interface {
		GetApplicationName() string
	}
	// ApplicationNames interface has GetApplicationNames
	ApplicationNames interface {
		GetApplicationNames() []string
	}
	// ApplicationNickname interface has GetApplicationNickname
	ApplicationNickname interface {
		GetApplicationNickname() string
	}
)

// ToApplicationIDs convert ApplicationID s to IDs
func ToApplicationIDs[T ApplicationID](arrays []T) []string {
	return ToStringSet(arrays, func(i interface{}) string { return i.(T).GetApplicationID() })
}

// ToOwnershipIDs convert OwnershipID s to []string
func ToOwnershipIDs[T OwnershipID](arrays []T) []string {
	return ToStringSet(arrays, func(i interface{}) string { return i.(T).GetOwnershipID() })
}

// OwnershipIDsToApplicationIDs convert OwnershipIDs to ApplicationIDs
func OwnershipIDsToApplicationIDs(i OwnershipIDs) ApplicationIDs {
	return zero.ApplicationIDs{Value: i.GetOwnershipIDs()}
}

// ToSubscriberIDs convert SubscriberID s to []string
func ToSubscriberIDs[T SubscriberID](arrays []T) []string {
	return ToStringSet(arrays, func(i interface{}) string { return i.(T).GetSubscriberID() })
}

// SubscriberIDsToApplicationIDs convert SubscriberIDs to ApplicationIDs
func SubscriberIDsToApplicationIDs(i SubscriberIDs) ApplicationIDs {
	return zero.ApplicationIDs{Value: i.GetSubscriberIDs()}
}

// ToPublisherIDs convert PublisherID s to []string
func ToPublisherIDs[T PublisherID](arrays []T) []string {
	return ToStringSet(arrays, func(i interface{}) string { return i.(T).GetPublisherID() })
}

// PublisherIDsToApplicationIDs convert PublisherIDs to ApplicationIDs
func PublisherIDsToApplicationIDs(i PublisherIDs) ApplicationIDs {
	return zero.ApplicationIDs{Value: i.GetPublisherIDs()}
}
