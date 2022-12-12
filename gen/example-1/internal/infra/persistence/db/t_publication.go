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

// NewPublicationDAO 构造方法
func NewPublicationDAO(db *gorm.DB) persistence.PublicationDAO {
	return &PublicationDAO{
		db: db,
	}
}

// PublicationDAO Publication 数据访问对象
type PublicationDAO struct {
	db *gorm.DB
}

// Transaction 事务执行
func (e PublicationDAO) Transaction(fun func(tx *gorm.DB) error) error {
	return e.db.Transaction(func(tx *gorm.DB) (err error) {
		return fun(tx)
	})
}

// IsRelated check fields is related EventBusDAO
func (e PublicationDAO) IsRelated(fields base.Fields) bool {
	return checkFields(fields, po.Publication{}.TableName())
}

// BaseFields PublicationDAO 相关的基本字段，如果未传入任何 fields, 由此进行兜底
func (e PublicationDAO) BaseFields() base.Fields {
	return publicationBaseFields{}
}

type publicationBaseFields struct{}

// GetFieldNames return event set dao base fields
func (publicationBaseFields) GetFieldNames() []string {
	return []string{
		fmt.Sprintf(COLUMN_PATTERN, po.Publication{}.TableName(), "name"),
		fmt.Sprintf(COLUMN_PATTERN, po.Publication{}.TableName(), "id"),
		fmt.Sprintf(COLUMN_PATTERN, po.Publication{}.TableName(), "status"),
	}
}

func (e PublicationDAO) listBaseDB(ctx context.Context, r req.GetPublicationsPO) *gorm.DB {
	eT := newTable(po.Publication{}.TableName())
	db := e.db.WithContext(ctx).Model(po.Publication{}).Where(eT.col("deleted_at").notDelete())
	mdb := mydb{db: db}

	for _, t := range []tuple{
		{predicate: eT.col("name").eq(), param: r.GetEventTypeID()},
		{predicate: eT.col("name").in(), param: r.GetEventTypeID()},
		{predicate: eT.col("name").like(), param: r.GetFuzzyPublicationEventTypeID()},
		{predicate: eT.col("id").eq(), param: r.GetPublicationID()},
		{predicate: eT.col("id").in(), param: r.GetPublicationIDs()},
		{predicate: eT.col("main_app_id").eq(), param: r.GetPublisherID()},
		{predicate: eT.col("main_app_id").in(), param: r.GetPublisherID()},
		{predicate: eT.col("status").eq(), param: r.GetPublicationStatus()},
		{predicate: eT.col("status").in(), param: r.GetPublicationStatuses()},
	} {
		mdb.appendWhere(t)
	}
	return mdb.db
}

// Get 通用 get 方法
func (e PublicationDAO) Get(ctx context.Context, r req.GetPublicationsPO) (po.Publications, uint32, error) {
	var total int64
	if err := e.listBaseDB(ctx, r).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var result []po.Publication
	if err := pageQuery(e.listBaseDB(ctx, r), r, &result); err != nil {
		return nil, 0, err
	}
	return result, uint32(total), nil

}
func (e PublicationDAO) Update(ctx context.Context, r req.UpdatePublicationReq) error {
	return e.db.WithContext(ctx).Updates(convertUpdatePublicationReq(r)).Error
}

func (e PublicationDAO) Delete(ctx context.Context, r req.DeletePublicationReq) error {
	return e.db.WithContext(ctx).Updates(convertDeletePublicationReq(r)).Error
}

func (e PublicationDAO) Create(ctx context.Context, r req.CreatePublicationReq) (string, error) {
	pobj := convertCreatePublicationReq(r)
	if err := e.db.WithContext(ctx).Create(&pobj).Error; err != nil {
		return reflect.Zero(reflect.TypeOf(pobj.ID)).Interface().(string), err
	}
	return pobj.ID, nil
}

func convertUpdatePublicationReq(r req.UpdatePublicationReq) po.Publication {
	return po.Publication{
		EventTypeID: r.GetEventTypeID(),
		ID:          r.GetPublicationID(),
		PublisherID: r.GetPublisherID(),
		Status:      r.GetPublicationStatus(),
		Updater:     r.GetRtx(),
	}
}

func convertDeletePublicationReq(r req.DeletePublicationReq) po.Publication {
	return po.Publication{
		ID:        r.GetPublicationID(),
		DeletedAt: time.Now().Unix(),
		Updater:   r.GetRtx(),
	}
}

func convertCreatePublicationReq(r req.CreatePublicationReq) po.Publication {
	return po.Publication{
		EventTypeID: r.GetEventTypeID(),
		ID:          r.GetPublicationID(),
		PublisherID: r.GetPublisherID(),
		Status:      r.GetPublicationStatus(),
		Creator:     r.GetRtx(),
	}
}
