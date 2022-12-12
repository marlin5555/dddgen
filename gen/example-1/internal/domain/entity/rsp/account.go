// Package rsp Code generated, DO NOT EDIT.
package rsp

import "github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity"

// AccountsRsp combination Rsp Total Accounts
type AccountsRsp struct {
	Rsp
	Total uint32
	entity.Accounts
}

// PassportsRsp combination Rsp Total Passports
type PassportsRsp struct {
	Rsp
	Total uint32
	entity.Passports
}

// SecretsRsp combination Rsp Total Secrets
type SecretsRsp struct {
	Rsp
	Total uint32
	entity.Secrets
}

// TechRelationsRsp combination Rsp Total TechRelations
type TechRelationsRsp struct {
	Rsp
	Total uint32
	entity.TechRelations
}
