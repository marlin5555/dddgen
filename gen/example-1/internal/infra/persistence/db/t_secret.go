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

// NewSecretDAO 构造方法
func NewSecretDAO(db *gorm.DB) persistence.SecretDAO {
	return &SecretDAO{
		db: db,
	}
}

// SecretDAO Secret 数据访问对象
type SecretDAO struct {
	db *gorm.DB
}

// Transaction 事务执行
func (e SecretDAO) Transaction(fun func(tx *gorm.DB) error) error {
	return e.db.Transaction(func(tx *gorm.DB) (err error) {
		return fun(tx)
	})
}

// IsRelated check fields is related EventBusDAO
func (e SecretDAO) IsRelated(fields base.Fields) bool {
	return checkFields(fields, po.Secret{}.TableName())
}

// BaseFields SecretDAO 相关的基本字段，如果未传入任何 fields, 由此进行兜底
func (e SecretDAO) BaseFields() base.Fields {
	return secretBaseFields{}
}

type secretBaseFields struct{}

// GetFieldNames return event set dao base fields
func (secretBaseFields) GetFieldNames() []string {
	return []string{
		fmt.Sprintf(COLUMN_PATTERN, po.Secret{}.TableName(), "id"),
	}
}

func (e SecretDAO) listBaseDB(ctx context.Context, r req.GetSecretsPO) *gorm.DB {
	eT := newTable(po.Secret{}.TableName())
	db := e.db.WithContext(ctx).Model(po.Secret{}).Where(eT.col("deleted_at").notDelete())
	mdb := mydb{db: db}

	for _, t := range []tuple{
		{predicate: eT.col("account_id").eq(), param: r.GetAccountID()},
		{predicate: eT.col("id").eq(), param: r.GetSecretID()},
		{predicate: eT.col("id").in(), param: r.GetSecretIDs()},
	} {
		mdb.appendWhere(t)
	}
	return mdb.db
}

// Get 通用 get 方法
func (e SecretDAO) Get(ctx context.Context, r req.GetSecretsPO) (po.Secrets, uint32, error) {
	var total int64
	if err := e.listBaseDB(ctx, r).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var result []po.Secret
	if err := pageQuery(e.listBaseDB(ctx, r), r, &result); err != nil {
		return nil, 0, err
	}
	return result, uint32(total), nil

}
func (e SecretDAO) Update(ctx context.Context, r req.UpdateSecretReq) error {
	return e.db.WithContext(ctx).Updates(convertUpdateSecretReq(r)).Error
}

func (e SecretDAO) Delete(ctx context.Context, r req.DeleteSecretReq) error {
	return e.db.WithContext(ctx).Updates(convertDeleteSecretReq(r)).Error
}

func (e SecretDAO) Create(ctx context.Context, r req.CreateSecretReq) (string, error) {
	pobj := convertCreateSecretReq(r)
	if err := e.db.WithContext(ctx).Create(&pobj).Error; err != nil {
		return reflect.Zero(reflect.TypeOf(pobj.ID)).Interface().(string), err
	}
	return pobj.ID, nil
}

func convertUpdateSecretReq(r req.UpdateSecretReq) po.Secret {
	return po.Secret{
		Updater: r.GetRtx(),
	}
}

func convertDeleteSecretReq(r req.DeleteSecretReq) po.Secret {
	return po.Secret{
		ID:        r.GetSecretID(),
		DeletedAt: time.Now().Unix(),
		Updater:   r.GetRtx(),
	}
}

func convertCreateSecretReq(r req.CreateSecretReq) po.Secret {
	return po.Secret{
		ID:      uuid(),
		Creator: r.GetRtx(),
	}
}
