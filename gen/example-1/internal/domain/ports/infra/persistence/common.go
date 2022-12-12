// Package persistence Code generated, DO NOT EDIT.
package persistence

import (
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/base"

	"gorm.io/gorm"
)

// Transaction 执行事务
type Transaction interface {
	Transaction(fun func(tx *gorm.DB) error) error
}

// FieldChecker own IsRelated BaseFields
type FieldChecker interface {
	IsRelated(fields base.Fields) bool
	BaseFields() base.Fields
}
