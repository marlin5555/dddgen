// Package repo Code generated, DO NOT EDIT.
package repo

import (
	"context"

	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/req"
)

// AccountRepository Account repo
type AccountRepository interface {
	// CreateAccount create Account
	CreateAccount(context.Context, req.CreateAccountReq) (string, error)
	// UpdateAccount update Account
	UpdateAccount(context.Context, req.UpdateAccountReq) error
	// DeleteAccount delete Account
	DeleteAccount(context.Context, req.DeleteAccountReq) error
	// GetAccounts general get Account method
	GetAccounts(context.Context, req.GetAccountsReq) (entity.Accounts, uint32, error)

	// CreatePassport create Passport
	CreatePassport(context.Context, req.CreatePassportReq) (string, error)
	// UpdatePassport update Passport
	UpdatePassport(context.Context, req.UpdatePassportReq) error
	// DeletePassport delete Passport
	DeletePassport(context.Context, req.DeletePassportReq) error
	// GetPassports general get Passport method
	GetPassports(context.Context, req.GetPassportsReq) (entity.Passports, uint32, error)

	// CreateSecret create Secret
	CreateSecret(context.Context, req.CreateSecretReq) (string, error)
	// UpdateSecret update Secret
	UpdateSecret(context.Context, req.UpdateSecretReq) error
	// DeleteSecret delete Secret
	DeleteSecret(context.Context, req.DeleteSecretReq) error
	// GetSecrets general get Secret method
	GetSecrets(context.Context, req.GetSecretsReq) (entity.Secrets, uint32, error)

	// CreateTechRelation create TechRelation
	CreateTechRelation(context.Context, req.CreateTechRelationReq) (string, error)
	// UpdateTechRelation update TechRelation
	UpdateTechRelation(context.Context, req.UpdateTechRelationReq) error
	// DeleteTechRelation delete TechRelation
	DeleteTechRelation(context.Context, req.DeleteTechRelationReq) error
	// GetTechRelations general get TechRelation method
	GetTechRelations(context.Context, req.GetTechRelationsReq) (entity.TechRelations, uint32, error)
}
