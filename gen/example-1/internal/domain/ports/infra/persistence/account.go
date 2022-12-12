// Package persistence Code generated, DO NOT EDIT.
package persistence

import (
	"context"

	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/req"
	"github.com/marlin5555/dddgen/gen/example-1/internal/infra/persistence/po"
)

// AccountDAO Account dao
type AccountDAO interface {
	Transaction
	FieldChecker
	Get(context.Context, req.GetAccountsPO) (po.Accounts, uint32, error)
	Create(context.Context, req.CreateAccountReq) (string, error)
	Update(context.Context, req.UpdateAccountReq) error
	Delete(context.Context, req.DeleteAccountReq) error
}

// PassportDAO Passport dao
type PassportDAO interface {
	Transaction
	FieldChecker
	Get(context.Context, req.GetPassportsPO) (po.Passports, uint32, error)
	Create(context.Context, req.CreatePassportReq) (string, error)
	Update(context.Context, req.UpdatePassportReq) error
	Delete(context.Context, req.DeletePassportReq) error
}

// SecretDAO Secret dao
type SecretDAO interface {
	Transaction
	FieldChecker
	Get(context.Context, req.GetSecretsPO) (po.Secrets, uint32, error)
	Create(context.Context, req.CreateSecretReq) (string, error)
	Update(context.Context, req.UpdateSecretReq) error
	Delete(context.Context, req.DeleteSecretReq) error
}

// TechRelationDAO TechRelation dao
type TechRelationDAO interface {
	Transaction
	FieldChecker
	Get(context.Context, req.GetTechRelationsPO) (po.TechRelations, uint32, error)
	Create(context.Context, req.CreateTechRelationReq) (string, error)
	Update(context.Context, req.UpdateTechRelationReq) error
	Delete(context.Context, req.DeleteTechRelationReq) error
}
