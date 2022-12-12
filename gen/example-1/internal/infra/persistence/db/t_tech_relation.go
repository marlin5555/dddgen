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

// NewTechRelationDAO 构造方法
func NewTechRelationDAO(db *gorm.DB) persistence.TechRelationDAO {
	return &TechRelationDAO{
		db: db,
	}
}

// TechRelationDAO TechRelation 数据访问对象
type TechRelationDAO struct {
	db *gorm.DB
}

// Transaction 事务执行
func (e TechRelationDAO) Transaction(fun func(tx *gorm.DB) error) error {
	return e.db.Transaction(func(tx *gorm.DB) (err error) {
		return fun(tx)
	})
}

// IsRelated check fields is related EventBusDAO
func (e TechRelationDAO) IsRelated(fields base.Fields) bool {
	return checkFields(fields, po.TechRelation{}.TableName())
}

// BaseFields TechRelationDAO 相关的基本字段，如果未传入任何 fields, 由此进行兜底
func (e TechRelationDAO) BaseFields() base.Fields {
	return techRelationBaseFields{}
}

type techRelationBaseFields struct{}

// GetFieldNames return event set dao base fields
func (techRelationBaseFields) GetFieldNames() []string {
	return []string{
		fmt.Sprintf(COLUMN_PATTERN, po.TechRelation{}.TableName(), "id"),
	}
}

func (e TechRelationDAO) listBaseDB(ctx context.Context, r req.GetTechRelationsPO) *gorm.DB {
	eT := newTable(po.TechRelation{}.TableName())
	db := e.db.WithContext(ctx).Model(po.TechRelation{}).Where(eT.col("deleted_at").notDelete())
	mdb := mydb{db: db}

	for _, t := range []tuple{
		{predicate: eT.col("id").eq(), param: r.GetTechRelationID()},
		{predicate: eT.col("id").in(), param: r.GetTechRelationIDs()},
	} {
		mdb.appendWhere(t)
	}
	return mdb.db
}

// Get 通用 get 方法
func (e TechRelationDAO) Get(ctx context.Context, r req.GetTechRelationsPO) (po.TechRelations, uint32, error) {
	var total int64
	if err := e.listBaseDB(ctx, r).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var result []po.TechRelation
	if err := pageQuery(e.listBaseDB(ctx, r), r, &result); err != nil {
		return nil, 0, err
	}
	return result, uint32(total), nil

}
func (e TechRelationDAO) Update(ctx context.Context, r req.UpdateTechRelationReq) error {
	return e.db.WithContext(ctx).Updates(convertUpdateTechRelationReq(r)).Error
}

func (e TechRelationDAO) Delete(ctx context.Context, r req.DeleteTechRelationReq) error {
	return e.db.WithContext(ctx).Updates(convertDeleteTechRelationReq(r)).Error
}

func (e TechRelationDAO) Create(ctx context.Context, r req.CreateTechRelationReq) (string, error) {
	pobj := convertCreateTechRelationReq(r)
	if err := e.db.WithContext(ctx).Create(&pobj).Error; err != nil {
		return reflect.Zero(reflect.TypeOf(pobj.ID)).Interface().(string), err
	}
	return pobj.ID, nil
}

func convertUpdateTechRelationReq(r req.UpdateTechRelationReq) po.TechRelation {
	return po.TechRelation{
		Updater: r.GetRtx(),
	}
}

func convertDeleteTechRelationReq(r req.DeleteTechRelationReq) po.TechRelation {
	return po.TechRelation{
		ID:        r.GetTechRelationID(),
		DeletedAt: time.Now().Unix(),
		Updater:   r.GetRtx(),
	}
}

func convertCreateTechRelationReq(r req.CreateTechRelationReq) po.TechRelation {
	return po.TechRelation{
		ID:      uuid(),
		Creator: r.GetRtx(),
	}
}
