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

// NewSubscriptionDAO 构造方法
func NewSubscriptionDAO(db *gorm.DB) persistence.SubscriptionDAO {
	return &SubscriptionDAO{
		db: db,
	}
}

// SubscriptionDAO Subscription 数据访问对象
type SubscriptionDAO struct {
	db *gorm.DB
}

// Transaction 事务执行
func (e SubscriptionDAO) Transaction(fun func(tx *gorm.DB) error) error {
	return e.db.Transaction(func(tx *gorm.DB) (err error) {
		return fun(tx)
	})
}

// IsRelated check fields is related EventBusDAO
func (e SubscriptionDAO) IsRelated(fields base.Fields) bool {
	return checkFields(fields, po.Subscription{}.TableName())
}

// BaseFields SubscriptionDAO 相关的基本字段，如果未传入任何 fields, 由此进行兜底
func (e SubscriptionDAO) BaseFields() base.Fields {
	return subscriptionBaseFields{}
}

type subscriptionBaseFields struct{}

// GetFieldNames return event set dao base fields
func (subscriptionBaseFields) GetFieldNames() []string {
	return []string{
		fmt.Sprintf(COLUMN_PATTERN, po.Subscription{}.TableName(), "name"),
		fmt.Sprintf(COLUMN_PATTERN, po.Subscription{}.TableName(), "id"),
		fmt.Sprintf(COLUMN_PATTERN, po.Subscription{}.TableName(), "status"),
	}
}

func (e SubscriptionDAO) listBaseDB(ctx context.Context, r req.GetSubscriptionsPO) *gorm.DB {
	eT := newTable(po.Subscription{}.TableName())
	db := e.db.WithContext(ctx).Model(po.Subscription{}).Where(eT.col("deleted_at").notDelete())
	mdb := mydb{db: db}

	for _, t := range []tuple{
		{predicate: eT.col("name").eq(), param: r.GetEventTypeID()},
		{predicate: eT.col("name").in(), param: r.GetEventTypeID()},
		{predicate: eT.col("name").like(), param: r.GetFuzzySubscriptionEventTypeID()},
		{predicate: eT.col("id").eq(), param: r.GetSubscriptionID()},
		{predicate: eT.col("id").in(), param: r.GetSubscriptionIDs()},
		{predicate: eT.col("status").eq(), param: r.GetSubscriptionStatus()},
		{predicate: eT.col("status").in(), param: r.GetSubscriptionStatuses()},
		{predicate: eT.col("main_app_id").eq(), param: r.GetSubscriberID()},
		{predicate: eT.col("main_app_id").in(), param: r.GetSubscriberID()},
	} {
		mdb.appendWhere(t)
	}
	return mdb.db
}

// Get 通用 get 方法
func (e SubscriptionDAO) Get(ctx context.Context, r req.GetSubscriptionsPO) (po.Subscriptions, uint32, error) {
	var total int64
	if err := e.listBaseDB(ctx, r).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var result []po.Subscription
	if err := pageQuery(e.listBaseDB(ctx, r), r, &result); err != nil {
		return nil, 0, err
	}
	return result, uint32(total), nil

}
func (e SubscriptionDAO) Update(ctx context.Context, r req.UpdateSubscriptionReq) error {
	return e.db.WithContext(ctx).Updates(convertUpdateSubscriptionReq(r)).Error
}

func (e SubscriptionDAO) Delete(ctx context.Context, r req.DeleteSubscriptionReq) error {
	return e.db.WithContext(ctx).Updates(convertDeleteSubscriptionReq(r)).Error
}

func (e SubscriptionDAO) Create(ctx context.Context, r req.CreateSubscriptionReq) (string, error) {
	pobj := convertCreateSubscriptionReq(r)
	if err := e.db.WithContext(ctx).Create(&pobj).Error; err != nil {
		return reflect.Zero(reflect.TypeOf(pobj.ID)).Interface().(string), err
	}
	return pobj.ID, nil
}

func convertUpdateSubscriptionReq(r req.UpdateSubscriptionReq) po.Subscription {
	return po.Subscription{
		EventTypeID:  r.GetEventTypeID(),
		ID:           r.GetSubscriptionID(),
		Status:       r.GetSubscriptionStatus(),
		SubscriberID: r.GetSubscriberID(),
		Updater:      r.GetRtx(),
	}
}

func convertDeleteSubscriptionReq(r req.DeleteSubscriptionReq) po.Subscription {
	return po.Subscription{
		ID:        r.GetSubscriptionID(),
		DeletedAt: time.Now().Unix(),
		Updater:   r.GetRtx(),
	}
}

func convertCreateSubscriptionReq(r req.CreateSubscriptionReq) po.Subscription {
	return po.Subscription{
		EventTypeID:  r.GetEventTypeID(),
		ID:           r.GetSubscriptionID(),
		Status:       r.GetSubscriptionStatus(),
		SubscriberID: r.GetSubscriberID(),
		Creator:      r.GetRtx(),
	}
}
