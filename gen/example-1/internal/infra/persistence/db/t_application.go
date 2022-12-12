// Package db Code generated, DO NOT EDIT.
package db

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/base"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/entity/req"
	"github.com/marlin5555/dddgen/gen/example-1/internal/domain/ports/infra/persistence"
	"github.com/marlin5555/dddgen/gen/example-1/internal/infra/persistence/po"

	"gorm.io/gorm"
)

// NewApplicationDAO 构造方法
func NewApplicationDAO(db *gorm.DB) persistence.ApplicationDAO {
	return &ApplicationDAO{
		db: db,
	}
}

// ApplicationDAO Application 数据访问对象
type ApplicationDAO struct {
	db *gorm.DB
}

// Transaction 事务执行
func (e ApplicationDAO) Transaction(fun func(tx *gorm.DB) error) error {
	return e.db.Transaction(func(tx *gorm.DB) (err error) {
		return fun(tx)
	})
}

// IsRelated check fields is related EventBusDAO
func (e ApplicationDAO) IsRelated(fields base.Fields) bool {
	return checkFields(fields, po.Application{}.TableName())
}

// BaseFields ApplicationDAO 相关的基本字段，如果未传入任何 fields, 由此进行兜底
func (e ApplicationDAO) BaseFields() base.Fields {
	return applicationBaseFields{}
}

type applicationBaseFields struct{}

// GetFieldNames return event set dao base fields
func (applicationBaseFields) GetFieldNames() []string {
	return []string{
		fmt.Sprintf(COLUMN_PATTERN, po.Application{}.TableName(), "id"),
		fmt.Sprintf(COLUMN_PATTERN, po.Application{}.TableName(), "name"),
		fmt.Sprintf(COLUMN_PATTERN, po.Application{}.TableName(), "nickname"),
	}
}

func (e ApplicationDAO) listBaseDB(ctx context.Context, r req.GetApplicationsPO) *gorm.DB {
	eT := newTable(po.Application{}.TableName())
	db := e.db.WithContext(ctx).Model(po.Application{}).Where(eT.col("deleted_at").notDelete())
	mdb := mydb{db: db}

	for _, t := range []tuple{
		{predicate: eT.col("name").like(), param: r.GetFuzzyApplicationName()},
		{predicate: eT.col("nickname").like(), param: r.GetFuzzyApplicationNickname()},
		{predicate: eT.col("id").eq(), param: r.GetApplicationID()},
		{predicate: eT.col("id").in(), param: r.GetApplicationIDs()},
		{predicate: eT.col("name").eq(), param: r.GetApplicationName()},
		{predicate: eT.col("name").in(), param: r.GetApplicationNames()},
	} {
		mdb.appendWhere(t)
	}
	return mdb.db
}

// Get 通用 get 方法
func (e ApplicationDAO) Get(ctx context.Context, r req.GetApplicationsPO) (po.Applications, uint32, error) {
	var total int64
	if err := e.listBaseDB(ctx, r).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var result []po.Application
	if err := pageQuery(e.listBaseDB(ctx, r), r, &result); err != nil {
		return nil, 0, err
	}
	return result, uint32(total), nil

}
func (e ApplicationDAO) Update(ctx context.Context, r req.UpdateApplicationReq) error {
	return e.db.WithContext(ctx).Updates(convertUpdateApplicationReq(r)).Error
}

func (e ApplicationDAO) Delete(ctx context.Context, r req.DeleteApplicationReq) error {
	return e.db.WithContext(ctx).Updates(convertDeleteApplicationReq(r)).Error
}

func (e ApplicationDAO) Create(ctx context.Context, r req.CreateApplicationReq) (string, error) {
	pobj := convertCreateApplicationReq(r)
	if err := e.db.WithContext(ctx).Create(&pobj).Error; err != nil {
		return reflect.Zero(reflect.TypeOf(pobj.ID)).Interface().(string), err
	}
	return pobj.ID, nil
}

func convertUpdateApplicationReq(r req.UpdateApplicationReq) po.Application {
	return po.Application{
		ID:       r.GetApplicationID(),
		Name:     r.GetApplicationName(),
		Nickname: r.GetApplicationNickname(),
		Updater:  r.GetRtx(),
	}
}

func convertDeleteApplicationReq(r req.DeleteApplicationReq) po.Application {
	return po.Application{
		ID:        r.GetApplicationID(),
		DeletedAt: time.Now().Unix(),
		Updater:   r.GetRtx(),
	}
}

func convertCreateApplicationReq(r req.CreateApplicationReq) po.Application {
	return po.Application{
		ID:       r.GetApplicationID(),
		Name:     r.GetApplicationName(),
		Nickname: r.GetApplicationNickname(),
		Creator:  r.GetRtx(),
	}
}
