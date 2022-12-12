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

// NewPassportDAO 构造方法
func NewPassportDAO(db *gorm.DB) persistence.PassportDAO {
	return &PassportDAO{
		db: db,
	}
}

// PassportDAO Passport 数据访问对象
type PassportDAO struct {
	db *gorm.DB
}

// Transaction 事务执行
func (e PassportDAO) Transaction(fun func(tx *gorm.DB) error) error {
	return e.db.Transaction(func(tx *gorm.DB) (err error) {
		return fun(tx)
	})
}

// IsRelated check fields is related EventBusDAO
func (e PassportDAO) IsRelated(fields base.Fields) bool {
	return checkFields(fields, po.Passport{}.TableName())
}

// BaseFields PassportDAO 相关的基本字段，如果未传入任何 fields, 由此进行兜底
func (e PassportDAO) BaseFields() base.Fields {
	return passportBaseFields{}
}

type passportBaseFields struct{}

// GetFieldNames return event set dao base fields
func (passportBaseFields) GetFieldNames() []string {
	return []string{
		fmt.Sprintf(COLUMN_PATTERN, po.Passport{}.TableName(), "id"),
	}
}

func (e PassportDAO) listBaseDB(ctx context.Context, r req.GetPassportsPO) *gorm.DB {
	eT := newTable(po.Passport{}.TableName())
	db := e.db.WithContext(ctx).Model(po.Passport{}).Where(eT.col("deleted_at").notDelete())
	mdb := mydb{db: db}

	for _, t := range []tuple{
		{predicate: eT.col("account_id").eq(), param: r.GetAccountID()},
		{predicate: eT.col("account_id").in(), param: r.GetAccountID()},
		{predicate: eT.col("postal_address").like(), param: r.GetFuzzyPassportPostalAddress()},
		{predicate: eT.col("id").eq(), param: r.GetPassportID()},
		{predicate: eT.col("id").in(), param: r.GetPassportIDs()},
		{predicate: eT.col("nationality").eq(), param: r.GetPassportNationality()},
		{predicate: eT.col("nationality").in(), param: r.GetPassportNationalitys()},
	} {
		mdb.appendWhere(t)
	}
	return mdb.db
}

// Get 通用 get 方法
func (e PassportDAO) Get(ctx context.Context, r req.GetPassportsPO) (po.Passports, uint32, error) {
	var total int64
	if err := e.listBaseDB(ctx, r).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var result []po.Passport
	if err := pageQuery(e.listBaseDB(ctx, r), r, &result); err != nil {
		return nil, 0, err
	}
	return result, uint32(total), nil

}
func (e PassportDAO) Update(ctx context.Context, r req.UpdatePassportReq) error {
	return e.db.WithContext(ctx).Updates(convertUpdatePassportReq(r)).Error
}

func (e PassportDAO) Delete(ctx context.Context, r req.DeletePassportReq) error {
	return e.db.WithContext(ctx).Updates(convertDeletePassportReq(r)).Error
}

func (e PassportDAO) Create(ctx context.Context, r req.CreatePassportReq) (string, error) {
	pobj := convertCreatePassportReq(r)
	if err := e.db.WithContext(ctx).Create(&pobj).Error; err != nil {
		return reflect.Zero(reflect.TypeOf(pobj.ID)).Interface().(string), err
	}
	return pobj.ID, nil
}

func convertUpdatePassportReq(r req.UpdatePassportReq) po.Passport {
	return po.Passport{
		Updater: r.GetRtx(),
	}
}

func convertDeletePassportReq(r req.DeletePassportReq) po.Passport {
	return po.Passport{
		ID:        r.GetPassportID(),
		DeletedAt: time.Now().Unix(),
		Updater:   r.GetRtx(),
	}
}

func convertCreatePassportReq(r req.CreatePassportReq) po.Passport {
	return po.Passport{
		ID:      uuid(),
		Creator: r.GetRtx(),
	}
}
