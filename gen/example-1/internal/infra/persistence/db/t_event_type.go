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

// NewEventTypeDAO 构造方法
func NewEventTypeDAO(db *gorm.DB) persistence.EventTypeDAO {
	return &EventTypeDAO{
		db: db,
	}
}

// EventTypeDAO EventType 数据访问对象
type EventTypeDAO struct {
	db *gorm.DB
}

// Transaction 事务执行
func (e EventTypeDAO) Transaction(fun func(tx *gorm.DB) error) error {
	return e.db.Transaction(func(tx *gorm.DB) (err error) {
		return fun(tx)
	})
}

// IsRelated check fields is related EventBusDAO
func (e EventTypeDAO) IsRelated(fields base.Fields) bool {
	return checkFields(fields, po.EventType{}.TableName())
}

// BaseFields EventTypeDAO 相关的基本字段，如果未传入任何 fields, 由此进行兜底
func (e EventTypeDAO) BaseFields() base.Fields {
	return eventTypeBaseFields{}
}

type eventTypeBaseFields struct{}

// GetFieldNames return event set dao base fields
func (eventTypeBaseFields) GetFieldNames() []string {
	return []string{
		fmt.Sprintf(COLUMN_PATTERN, po.EventType{}.TableName(), "id"),
		fmt.Sprintf(COLUMN_PATTERN, po.EventType{}.TableName(), "name"),
		fmt.Sprintf(COLUMN_PATTERN, po.EventType{}.TableName(), "status"),
	}
}

func (e EventTypeDAO) listBaseDB(ctx context.Context, r req.GetEventTypesPO) *gorm.DB {
	eT := newTable(po.EventType{}.TableName())
	db := e.db.WithContext(ctx).Model(po.EventType{}).Where(eT.col("deleted_at").notDelete())
	mdb := mydb{db: db}

	for _, t := range []tuple{
		{predicate: eT.col("event_bus_id").eq(), param: r.GetEventBusID()},
		{predicate: eT.col("event_bus_id").in(), param: r.GetEventBusID()},
		{predicate: eT.col("name").like(), param: r.GetFuzzyEventTypeName()},
		{predicate: eT.col("id").eq(), param: r.GetEventTypeID()},
		{predicate: eT.col("id").in(), param: r.GetEventTypeIDs()},
		{predicate: eT.col("main_app_id").eq(), param: r.GetOwnershipID()},
		{predicate: eT.col("main_app_id").in(), param: r.GetOwnershipID()},
		{predicate: eT.col("name").eq(), param: r.GetEventTypeName()},
		{predicate: eT.col("name").in(), param: r.GetEventTypeNames()},
		{predicate: eT.col("status").eq(), param: r.GetEventTypeStatus()},
		{predicate: eT.col("status").in(), param: r.GetEventTypeStatuses()},
	} {
		mdb.appendWhere(t)
	}
	return mdb.db
}

// Get 通用 get 方法
func (e EventTypeDAO) Get(ctx context.Context, r req.GetEventTypesPO) (po.EventTypes, uint32, error) {
	var total int64
	if err := e.listBaseDB(ctx, r).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var result []po.EventType
	if err := pageQuery(e.listBaseDB(ctx, r), r, &result); err != nil {
		return nil, 0, err
	}
	return result, uint32(total), nil

}
func (e EventTypeDAO) Update(ctx context.Context, r req.UpdateEventTypeReq) error {
	return e.db.WithContext(ctx).Updates(convertUpdateEventTypeReq(r)).Error
}

func (e EventTypeDAO) Delete(ctx context.Context, r req.DeleteEventTypeReq) error {
	return e.db.WithContext(ctx).Updates(convertDeleteEventTypeReq(r)).Error
}

func (e EventTypeDAO) Create(ctx context.Context, r req.CreateEventTypeReq) (string, error) {
	pobj := convertCreateEventTypeReq(r)
	if err := e.db.WithContext(ctx).Create(&pobj).Error; err != nil {
		return reflect.Zero(reflect.TypeOf(pobj.ID)).Interface().(string), err
	}
	return pobj.ID, nil
}

func convertUpdateEventTypeReq(r req.UpdateEventTypeReq) po.EventType {
	return po.EventType{
		EventBusID: r.GetEventBusID(),
		ID:         r.GetEventTypeID(),
		MainAppID:  r.GetOwnershipID(),
		Name:       r.GetEventTypeName(),
		Status:     r.GetEventTypeStatus(),
		Updater:    r.GetRtx(),
	}
}

func convertDeleteEventTypeReq(r req.DeleteEventTypeReq) po.EventType {
	return po.EventType{
		ID:        r.GetEventTypeID(),
		DeletedAt: time.Now().Unix(),
		Updater:   r.GetRtx(),
	}
}

func convertCreateEventTypeReq(r req.CreateEventTypeReq) po.EventType {
	return po.EventType{
		EventBusID: r.GetEventBusID(),
		ID:         r.GetEventTypeID(),
		MainAppID:  r.GetOwnershipID(),
		Name:       r.GetEventTypeName(),
		Status:     r.GetEventTypeStatus(),
		Creator:    r.GetRtx(),
	}
}
