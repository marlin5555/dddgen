// Package repo Code generated, DO NOT EDIT.
package repo

import (
	"context"

	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/base"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/req"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/zero"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/ports/infra/persistence"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/ports/repo"
	"github.com/marlin5555/dddgen/gen/example-1/internal/infra/persistence/po"
	"github.com/marlin5555/dddgen/gen/example-1/internal/infra/persistence/potoentity"
)

// ApplicationRepo ...
type ApplicationRepo struct {
	applicationDAO persistence.ApplicationDAO
}

// NewApplicationRepo ...
func NewApplicationRepo(applicationDAO persistence.ApplicationDAO) repo.ApplicationRepository {
	return &ApplicationRepo{
		applicationDAO: applicationDAO,
	}
}

// CreateApplication create Application
func (r *ApplicationRepo) CreateApplication(ctx context.Context, req req.CreateApplicationReq) (string, error) {
	return r.applicationDAO.Create(ctx, req)
}

// UpdateApplication update Application
func (r *ApplicationRepo) UpdateApplication(ctx context.Context, req req.UpdateApplicationReq) error {
	return r.applicationDAO.Update(ctx, req)
}

// DeleteApplication delete Application
func (r *ApplicationRepo) DeleteApplication(ctx context.Context, req req.DeleteApplicationReq) error {
	return r.applicationDAO.Delete(ctx, req)
}

// GetApplications general get Application method
func (r *ApplicationRepo) GetApplications(ctx context.Context, rq req.GetApplicationsReq) (entity.Applications, uint32, error) {
	// 做主表查询，获取到主表记录
	result, total, err := r.applicationDAO.Get(ctx, struct {
		req.GetApplicationsNoForeign
		base.PageOrderOperator // 分页查询信息
	}{
		GetApplicationsNoForeign: rq,
		PageOrderOperator:        rq,
	})
	// 关联表查询
	builder := potoentity.NewBuilder(rq,
		potoentity.WithApplications(base.ToStringKeyMap(result,
			func(i interface{}) string { return i.(base.ApplicationID).GetApplicationID() })),
		potoentity.WithFields(r.applicationDAO.BaseFields()))
	builder.Build()
	return builder.GetApplications(), total, err
}

func needApplicationIDs(req req.GetApplicationsPure) bool {
	if !base.IsZero(req.GetFuzzyApplicationName()) {
		return true
	}
	if !base.IsZero(req.GetFuzzyApplicationNickname()) {
		return true
	}
	if !base.IsZero(req.GetApplicationID()) {
		return true
	}
	if !base.IsZero(req.GetApplicationIDs()) {
		return true
	}
	if !base.IsZero(req.GetApplicationName()) {
		return true
	}
	if !base.IsZero(req.GetApplicationNames()) {
		return true
	}
	return false
}

func getApplicationIDsByRef(ctx context.Context, dao persistence.ApplicationDAO, pure req.GetApplicationsPure) (base.ApplicationIDs, bool, error) {
	if !needApplicationIDs(pure) {
		return zero.ApplicationIDs{}, false, nil
	}
	var applications po.Applications
	var err error
	applications, _, err = dao.Get(ctx, struct {
		zero.PageOrderOperator
		req.GetApplicationsPure
	}{GetApplicationsPure: pure})
	if len(applications.GetApplicationIDs()) == 0 || err != nil {
		return zero.ApplicationIDs{}, true, nil
	}
	return zero.ApplicationIDs{Value: applications.GetApplicationIDs()}, false, nil
}

func getApplicationsByIDsIfRelated(ctx context.Context, dao persistence.ApplicationDAO, ids base.ApplicationIDs, fields base.Fields) (po.Applications, error) {
	if !dao.IsRelated(fields) {
		return nil, nil
	}
	return getApplicationsByIDs(ctx, dao, ids)
}

func getApplicationsByIDs(ctx context.Context, dao persistence.ApplicationDAO, ids base.ApplicationIDs) (po.Applications, error) {
	var applications po.Applications
	var err error
	if applications, _, err = dao.Get(ctx, struct {
		zero.FuzzyApplicationName
		zero.FuzzyApplicationNickname
		zero.ApplicationID
		base.ApplicationIDs
		zero.ApplicationName
		zero.ApplicationNames
		zero.PageOrderOperator
	}{ApplicationIDs: ids}); err != nil {
		return nil, err
	}
	return applications, err
}
