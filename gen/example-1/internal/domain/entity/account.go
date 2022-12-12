// Package entity Code generated, DO NOT EDIT.
package entity

import "github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/base"

type (

	// Account entity
	Account interface {
		base.AccountID
		SecretGetter
		PassportGetter
		base.AccountName
		base.AccountNickname
		base.AccountDescription
		Trailer
	}
	// AccountGetter Account getter
	AccountGetter interface {
		GetAccount() Account
	}
	StudentGetter interface {
		GetStudent() Account
	}
	TeacherGetter interface {
		GetTeacher() Account
	}

	// Accounts account s
	Accounts []Account
)
type (

	// Passport entity
	Passport interface {
		base.PassportID
		AccountGetter
		base.PassportNationality
		base.PassportPostalAddress
		Trailer
	}
	// PassportGetter Passport getter
	PassportGetter interface {
		GetPassport() Passport
	}

	// Passports passport s
	Passports []Passport
)
type (

	// Secret entity
	Secret interface {
		base.SecretID
		AccountGetter
		base.SecretSecret
		Trailer
	}
	// SecretGetter Secret getter
	SecretGetter interface {
		GetSecret() Secret
	}

	// Secrets secret s
	Secrets []Secret
)
type (

	// TechRelation entity
	TechRelation interface {
		base.TechRelationID
		StudentGetter
		TeacherGetter
		Trailer
	}
	// TechRelationGetter TechRelation getter
	TechRelationGetter interface {
		GetTechRelation() TechRelation
	}

	// TechRelations techRelation s
	TechRelations []TechRelation
)
