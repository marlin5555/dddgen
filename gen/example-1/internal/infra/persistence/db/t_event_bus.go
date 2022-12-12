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

// NewEventBusDAO 构造方法
func NewEventBusDAO(db *gorm.DB) persistence.EventBusDAO {
	return &EventBusDAO{
		db: db,
	}
}

// EventBusDAO EventBus 数据访问对象
type EventBusDAO struct {
	db *gorm.DB
}

// Transaction 事务执行
func (e EventBusDAO) Transaction(fun func(tx *gorm.DB) error) error {
	return e.db.Transaction(func(tx *gorm.DB) (err error) {
		return fun(tx)
	})
}

// IsRelated check fields is related EventBusDAO
func (e EventBusDAO) IsRelated(fields base.Fields) bool {
	return checkFields(fields, po.EventBus{}.TableName())
}

// BaseFields EventBusDAO 相关的基本字段，如果未传入任何 fields, 由此进行兜底
func (e EventBusDAO) BaseFields() base.Fields {
	return eventBusBaseFields{}
}

type eventBusBaseFields struct{}

// GetFieldNames return event set dao base fields
func (eventBusBaseFields) GetFieldNames() []string {
	return []string{
		fmt.Sprintf(COLUMN_PATTERN, po.EventBus{}.TableName(), "id"),
		fmt.Sprintf(COLUMN_PATTERN, po.EventBus{}.TableName(), "name"),
		fmt.Sprintf(COLUMN_PATTERN, po.EventBus{}.TableName(), "params"),
	}
}

func (e EventBusDAO) listBaseDB(ctx context.Context, r req.GetEventBusesPO) *gorm.DB {
	eT := newTable(po.EventBus{}.TableName())
	db := e.db.WithContext(ctx).Model(po.EventBus{}).Where(eT.col("deleted_at").notDelete())
	mdb := mydb{db: db}

	for _, t := range []tuple{
		{predicate: eT.col("name").like(), param: r.GetFuzzyEventBusName()},
		{predicate: eT.col("id").eq(), param: r.GetEventBusID()},
		{predicate: eT.col("id").in(), param: r.GetEventBusIDs()},
		{predicate: eT.col("name").eq(), param: r.GetEventBusName()},
		{predicate: eT.col("name").in(), param: r.GetEventBusNames()},
	} {
		mdb.appendWhere(t)
	}
	return mdb.db
}

// Get 通用 get 方法
func (e EventBusDAO) Get(ctx context.Context, r req.GetEventBusesPO) (po.EventBuses, uint32, error) {
	var total int64
	if err := e.listBaseDB(ctx, r).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var result []po.EventBus
	if err := pageQuery(e.listBaseDB(ctx, r), r, &result); err != nil {
		return nil, 0, err
	}
	return result, uint32(total), nil

}
func (e EventBusDAO) Update(ctx context.Context, r req.UpdateEventBusReq) error {
	return e.db.WithContext(ctx).Updates(convertUpdateEventBusReq(r)).Error
}

func (e EventBusDAO) Delete(ctx context.Context, r req.DeleteEventBusReq) error {
	return e.db.WithContext(ctx).Updates(convertDeleteEventBusReq(r)).Error
}

func (e EventBusDAO) Create(ctx context.Context, r req.CreateEventBusReq) (string, error) {
	pobj := convertCreateEventBusReq(r)
	if err := e.db.WithContext(ctx).Create(&pobj).Error; err != nil {
		return reflect.Zero(reflect.TypeOf(pobj.ID)).Interface().(string), err
	}
	return pobj.ID, nil
}

func convertUpdateEventBusReq(r req.UpdateEventBusReq) po.EventBus {
	return po.EventBus{
		ID:      r.GetEventBusID(),
		Name:    r.GetEventBusName(),
		Params:  r.GetEventBusParams(),
		Updater: r.GetRtx(),
	}
}

func convertDeleteEventBusReq(r req.DeleteEventBusReq) po.EventBus {
	return po.EventBus{
		ID:        r.GetEventBusID(),
		DeletedAt: time.Now().Unix(),
		Updater:   r.GetRtx(),
	}
}

func convertCreateEventBusReq(r req.CreateEventBusReq) po.EventBus {
	return po.EventBus{
		ID:      r.GetEventBusID(),
		Name:    r.GetEventBusName(),
		Params:  r.GetEventBusParams(),
		Creator: r.GetRtx(),
	}
}
