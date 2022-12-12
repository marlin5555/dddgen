// Package base Code generated, DO NOT EDIT.
package base

import (
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/zero"
)

// base interface for Account
type (
	// AccountDescription interface has GetAccountDescription
	AccountDescription interface {
		GetAccountDescription() string
	}
	// FuzzyAccountName interface has GetFuzzyAccountName
	FuzzyAccountName interface {
		GetFuzzyAccountName() string
	}
	// FuzzyAccountNickname interface has GetFuzzyAccountNickname
	FuzzyAccountNickname interface {
		GetFuzzyAccountNickname() string
	}
	// AccountID interface has GetAccountID
	AccountID interface {
		GetAccountID() string
	}
	// AccountIDs interface has GetAccountIDs
	AccountIDs interface {
		GetAccountIDs() []string
	}
	// AccountName interface has GetAccountName
	AccountName interface {
		GetAccountName() string
	}
	// AccountNames interface has GetAccountNames
	AccountNames interface {
		GetAccountNames() []string
	}
	// AccountNickname interface has GetAccountNickname
	AccountNickname interface {
		GetAccountNickname() string
	}
)

// base interface for Passport
type (
	// FuzzyPassportPostalAddress interface has GetFuzzyPassportPostalAddress
	FuzzyPassportPostalAddress interface {
		GetFuzzyPassportPostalAddress() string
	}
	// PassportID interface has GetPassportID
	PassportID interface {
		GetPassportID() string
	}
	// PassportIDs interface has GetPassportIDs
	PassportIDs interface {
		GetPassportIDs() []string
	}
	// PassportNationality interface has GetPassportNationality
	PassportNationality interface {
		GetPassportNationality() string
	}
	// PassportNationalitys interface has GetPassportNationalitys
	PassportNationalitys interface {
		GetPassportNationalitys() []string
	}
	// PassportPostalAddress interface has GetPassportPostalAddress
	PassportPostalAddress interface {
		GetPassportPostalAddress() string
	}
)

// base interface for Secret
type (
	// SecretID interface has GetSecretID
	SecretID interface {
		GetSecretID() string
	}
	// SecretIDs interface has GetSecretIDs
	SecretIDs interface {
		GetSecretIDs() []string
	}
	// SecretSecret interface has GetSecretSecret
	SecretSecret interface {
		GetSecretSecret() string
	}
)

// base interface for TechRelation
type (
	// TechRelationID interface has GetTechRelationID
	TechRelationID interface {
		GetTechRelationID() string
	}
	// TechRelationIDs interface has GetTechRelationIDs
	TechRelationIDs interface {
		GetTechRelationIDs() []string
	}
	// StudentID interface has GetStudentID
	StudentID interface {
		GetStudentID() string
	}
	// StudentIDs interface has GetStudentIDs
	StudentIDs interface {
		GetStudentIDs() []string
	}
	// TeacherID interface has GetTeacherID
	TeacherID interface {
		GetTeacherID() string
	}
	// TeacherIDs interface has GetTeacherIDs
	TeacherIDs interface {
		GetTeacherIDs() []string
	}
)

// ToAccountIDs convert AccountID s to IDs
func ToAccountIDs[T AccountID](arrays []T) []string {
	return ToStringSet(arrays, func(i interface{}) string { return i.(T).GetAccountID() })
}

// ToStudentIDs convert StudentID s to []string
func ToStudentIDs[T StudentID](arrays []T) []string {
	return ToStringSet(arrays, func(i interface{}) string { return i.(T).GetStudentID() })
}

// StudentIDsToAccountIDs convert StudentIDs to AccountIDs
func StudentIDsToAccountIDs(i StudentIDs) AccountIDs {
	return zero.AccountIDs{Value: i.GetStudentIDs()}
}

// ToTeacherIDs convert TeacherID s to []string
func ToTeacherIDs[T TeacherID](arrays []T) []string {
	return ToStringSet(arrays, func(i interface{}) string { return i.(T).GetTeacherID() })
}

// TeacherIDsToAccountIDs convert TeacherIDs to AccountIDs
func TeacherIDsToAccountIDs(i TeacherIDs) AccountIDs {
	return zero.AccountIDs{Value: i.GetTeacherIDs()}
}

// ToPassportIDs convert PassportID s to IDs
func ToPassportIDs[T PassportID](arrays []T) []string {
	return ToStringSet(arrays, func(i interface{}) string { return i.(T).GetPassportID() })
}

// ToSecretIDs convert SecretID s to IDs
func ToSecretIDs[T SecretID](arrays []T) []string {
	return ToStringSet(arrays, func(i interface{}) string { return i.(T).GetSecretID() })
}

// ToTechRelationIDs convert TechRelationID s to IDs
func ToTechRelationIDs[T TechRelationID](arrays []T) []string {
	return ToStringSet(arrays, func(i interface{}) string { return i.(T).GetTechRelationID() })
}
