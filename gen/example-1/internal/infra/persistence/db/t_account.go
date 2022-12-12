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

// NewAccountDAO 构造方法
func NewAccountDAO(db *gorm.DB) persistence.AccountDAO {
	return &AccountDAO{
		db: db,
	}
}

// AccountDAO Account 数据访问对象
type AccountDAO struct {
	db *gorm.DB
}

// Transaction 事务执行
func (e AccountDAO) Transaction(fun func(tx *gorm.DB) error) error {
	return e.db.Transaction(func(tx *gorm.DB) (err error) {
		return fun(tx)
	})
}

// IsRelated check fields is related EventBusDAO
func (e AccountDAO) IsRelated(fields base.Fields) bool {
	return checkFields(fields, po.Account{}.TableName())
}

// BaseFields AccountDAO 相关的基本字段，如果未传入任何 fields, 由此进行兜底
func (e AccountDAO) BaseFields() base.Fields {
	return accountBaseFields{}
}

type accountBaseFields struct{}

// GetFieldNames return event set dao base fields
func (accountBaseFields) GetFieldNames() []string {
	return []string{
		fmt.Sprintf(COLUMN_PATTERN, po.Account{}.TableName(), "id"),
		fmt.Sprintf(COLUMN_PATTERN, po.Account{}.TableName(), "name"),
		fmt.Sprintf(COLUMN_PATTERN, po.Account{}.TableName(), "nickname"),
	}
}

func (e AccountDAO) listBaseDB(ctx context.Context, r req.GetAccountsPO) *gorm.DB {
	eT := newTable(po.Account{}.TableName())
	db := e.db.WithContext(ctx).Model(po.Account{}).Where(eT.col("deleted_at").notDelete())
	mdb := mydb{db: db}

	for _, t := range []tuple{
		{predicate: eT.col("name").like(), param: r.GetFuzzyAccountName()},
		{predicate: eT.col("nickname").like(), param: r.GetFuzzyAccountNickname()},
		{predicate: eT.col("id").eq(), param: r.GetAccountID()},
		{predicate: eT.col("id").in(), param: r.GetAccountIDs()},
		{predicate: eT.col("name").eq(), param: r.GetAccountName()},
		{predicate: eT.col("name").in(), param: r.GetAccountNames()},
		{predicate: eT.col("nickname").eq(), param: r.GetAccountNickname()},
	} {
		mdb.appendWhere(t)
	}
	return mdb.db
}

// Get 通用 get 方法
func (e AccountDAO) Get(ctx context.Context, r req.GetAccountsPO) (po.Accounts, uint32, error) {
	var total int64
	if err := e.listBaseDB(ctx, r).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var result []po.Account
	if err := pageQuery(e.listBaseDB(ctx, r), r, &result); err != nil {
		return nil, 0, err
	}
	return result, uint32(total), nil

}
func (e AccountDAO) Update(ctx context.Context, r req.UpdateAccountReq) error {
	return e.db.WithContext(ctx).Updates(convertUpdateAccountReq(r)).Error
}

func (e AccountDAO) Delete(ctx context.Context, r req.DeleteAccountReq) error {
	return e.db.WithContext(ctx).Updates(convertDeleteAccountReq(r)).Error
}

func (e AccountDAO) Create(ctx context.Context, r req.CreateAccountReq) (string, error) {
	pobj := convertCreateAccountReq(r)
	if err := e.db.WithContext(ctx).Create(&pobj).Error; err != nil {
		return reflect.Zero(reflect.TypeOf(pobj.ID)).Interface().(string), err
	}
	return pobj.ID, nil
}

func convertUpdateAccountReq(r req.UpdateAccountReq) po.Account {
	return po.Account{
		Name:     r.GetAccountName(),
		Nickname: r.GetAccountNickname(),
		Updater:  r.GetRtx(),
	}
}

func convertDeleteAccountReq(r req.DeleteAccountReq) po.Account {
	return po.Account{
		ID:        r.GetAccountID(),
		DeletedAt: time.Now().Unix(),
		Updater:   r.GetRtx(),
	}
}

func convertCreateAccountReq(r req.CreateAccountReq) po.Account {
	return po.Account{
		ID:       uuid(),
		Name:     r.GetAccountName(),
		Nickname: r.GetAccountNickname(),
		Creator:  r.GetRtx(),
	}
}
