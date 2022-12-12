// Package repo Code generated, DO NOT EDIT.
package repo

import (
	"context"

	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/req"
)

// ApplicationRepository Application repo
type ApplicationRepository interface {
	// CreateApplication create Application
	CreateApplication(context.Context, req.CreateApplicationReq) (string, error)
	// UpdateApplication update Application
	UpdateApplication(context.Context, req.UpdateApplicationReq) error
	// DeleteApplication delete Application
	DeleteApplication(context.Context, req.DeleteApplicationReq) error
	// GetApplications general get Application method
	GetApplications(context.Context, req.GetApplicationsReq) (entity.Applications, uint32, error)
}
