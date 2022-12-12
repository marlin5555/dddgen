// Package persistence Code generated, DO NOT EDIT.
package persistence

import (
	"context"

	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/req"
	"github.com/marlin5555/dddgen/gen/example-1/internal/infra/persistence/po"
)

// ApplicationDAO Application dao
type ApplicationDAO interface {
	Transaction
	FieldChecker
	Get(context.Context, req.GetApplicationsPO) (po.Applications, uint32, error)
	Create(context.Context, req.CreateApplicationReq) (string, error)
	Update(context.Context, req.UpdateApplicationReq) error
	Delete(context.Context, req.DeleteApplicationReq) error
}
