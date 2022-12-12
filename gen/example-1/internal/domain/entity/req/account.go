// Package req Code generated, DO NOT EDIT.
package req

import "github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/base"

// get create update delete req for Account
type (
	// GetAccountsNoForeign get Accounts no foreign key
	GetAccountsNoForeign interface {
		base.FuzzyAccountName
		base.FuzzyAccountNickname
		base.AccountID
		base.AccountIDs
		base.AccountName
		base.AccountNames
		base.AccountNickname
	}
	// GetAccountsPure get Accounts pure, 仅仅作用在单一实体表上
	GetAccountsPure interface {
		base.FuzzyAccountName
		base.FuzzyAccountNickname
		base.AccountID
		base.AccountIDs
		base.AccountName
		base.AccountNames
		base.AccountNickname
	}
	// GetAccountsPO get Accounts po, 用于 persistence 上的 get 操作
	GetAccountsPO interface {
		GetAccountsPure
		base.PageOrderOperator
	}

	// GetAccountsReq get Accounts req, 用于对外 service 中使用
	GetAccountsReq interface {
		GetAccountsPure
		base.Fields
		base.PageOrderOperator
		base.Validator
	}

	DeleteAccountReq interface {
		base.AccountID
		base.Operator
		base.Validator
	}

	UpdateAccountReq interface {
		base.AccountName
		base.AccountNickname
		base.Operator
		base.Validator
	}
	CreateAccountReq interface {
		base.AccountName
		base.AccountNickname
		base.Operator
		base.Validator
	}
)

// get create update delete req for Passport
type (
	// GetPassportsNoForeign get Passports no foreign key
	GetPassportsNoForeign interface {
		base.FuzzyPassportPostalAddress
		base.PassportID
		base.PassportIDs
		base.PassportNationality
		base.PassportNationalitys
	}
	// GetPassportsPure get Passports pure, 仅仅作用在单一实体表上
	GetPassportsPure interface {
		base.AccountID
		base.AccountIDs
		base.FuzzyPassportPostalAddress
		base.PassportID
		base.PassportIDs
		base.PassportNationality
		base.PassportNationalitys
	}
	// GetPassportsPO get Passports po, 用于 persistence 上的 get 操作
	GetPassportsPO interface {
		GetPassportsPure
		base.PageOrderOperator
	}

	// GetPassportsReq get Passports req, 用于对外 service 中使用
	GetPassportsReq interface {
		GetPassportsPure

		// start. use for Account filter
		base.FuzzyAccountName
		base.FuzzyAccountNickname
		base.AccountID
		base.AccountIDs
		base.AccountName
		base.AccountNames
		base.AccountNickname
		// end.

		base.Fields
		base.PageOrderOperator
		base.Validator
	}

	DeletePassportReq interface {
		base.PassportID
		base.Operator
		base.Validator
	}

	UpdatePassportReq interface {
		base.Operator
		base.Validator
	}
	CreatePassportReq interface {
		base.Operator
		base.Validator
	}
)

// get create update delete req for Secret
type (
	// GetSecretsNoForeign get Secrets no foreign key
	GetSecretsNoForeign interface {
		base.SecretID
		base.SecretIDs
	}
	// GetSecretsPure get Secrets pure, 仅仅作用在单一实体表上
	GetSecretsPure interface {
		base.AccountID
		base.SecretID
		base.SecretIDs
	}
	// GetSecretsPO get Secrets po, 用于 persistence 上的 get 操作
	GetSecretsPO interface {
		GetSecretsPure
		base.PageOrderOperator
	}

	// GetSecretsReq get Secrets req, 用于对外 service 中使用
	GetSecretsReq interface {
		GetSecretsPure

		// start. use for Account filter
		base.FuzzyAccountName
		base.FuzzyAccountNickname
		base.AccountID
		base.AccountIDs
		base.AccountName
		base.AccountNames
		base.AccountNickname
		// end.

		base.Fields
		base.PageOrderOperator
		base.Validator
	}

	DeleteSecretReq interface {
		base.SecretID
		base.Operator
		base.Validator
	}

	UpdateSecretReq interface {
		base.Operator
		base.Validator
	}
	CreateSecretReq interface {
		base.Operator
		base.Validator
	}
)

// get create update delete req for TechRelation
type (
	// GetTechRelationsNoForeign get TechRelations no foreign key
	GetTechRelationsNoForeign interface {
		base.TechRelationID
		base.TechRelationIDs
	}
	// GetTechRelationsPure get TechRelations pure, 仅仅作用在单一实体表上
	GetTechRelationsPure interface {
		base.TechRelationID
		base.TechRelationIDs
	}
	// GetTechRelationsPO get TechRelations po, 用于 persistence 上的 get 操作
	GetTechRelationsPO interface {
		GetTechRelationsPure
		base.PageOrderOperator
	}

	// GetTechRelationsReq get TechRelations req, 用于对外 service 中使用
	GetTechRelationsReq interface {
		GetTechRelationsPure

		// start. use for Student filter
		base.FuzzyAccountName
		base.FuzzyAccountNickname
		base.AccountID
		base.AccountIDs
		base.AccountName
		base.AccountNames
		base.AccountNickname
		// end.

		// start. use for Teacher filter
		base.FuzzyAccountName
		base.FuzzyAccountNickname
		base.AccountID
		base.AccountIDs
		base.AccountName
		base.AccountNames
		base.AccountNickname
		// end.

		base.Fields
		base.PageOrderOperator
		base.Validator
	}

	DeleteTechRelationReq interface {
		base.TechRelationID
		base.Operator
		base.Validator
	}

	UpdateTechRelationReq interface {
		base.Operator
		base.Validator
	}
	CreateTechRelationReq interface {
		base.Operator
		base.Validator
	}
)
