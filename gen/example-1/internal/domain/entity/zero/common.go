// Package zero Code generated, DO NOT EDIT.
package zero

import (
	"strings"
	"time"
)

// CreatedAt interface has GetCreatedAt
type CreatedAt struct {
	CreatedAt time.Time
}

// GetCreatedAt return CreatedAt.CreatedAt impl base.CreatedAt
func (e CreatedAt) GetCreatedAt() time.Time {
	return e.CreatedAt
}

// Creator interface has GetCreator
type Creator struct {
	Creator string
}

// GetCreator return Creator.Creator impl base.Creator
func (e Creator) GetCreator() string {
	return e.Creator
}

// DeletedAt interface has GetDeletedAt
type DeletedAt struct {
	DeletedAt int64
}

// GetDeletedAt return DeletedAt.DeletedAt impl base.DeletedAt
func (e DeletedAt) GetDeletedAt() int64 {
	return e.DeletedAt
}

// UpdatedAt interface has GetUpdatedAt
type UpdatedAt struct {
	UpdatedAt time.Time
}

// GetUpdatedAt return UpdatedAt.UpdatedAt impl base.UpdatedAt
func (e UpdatedAt) GetUpdatedAt() time.Time {
	return e.UpdatedAt
}

// Updater interface has GetUpdater
type Updater struct {
	Updater string
}

// GetUpdater return Updater.Updater impl base.Updater
func (e Updater) GetUpdater() string {
	return e.Updater
}

// PageOrderOperator combination: PageQuery Order Operator
type PageOrderOperator struct {
	PageQuery
	Order
	Operator
}

// PageQuery 分页查询对象
type PageQuery struct {
	PageIndex int // 分页序号
	PageSize  int // 分页大小
}

// GetOffset 获取查询数据库的 offset
func (q PageQuery) GetOffset() int {
	return (q.PageIndex - 1) * q.PageSize
}

// GetLimit 获取查询数据库的 limit
func (q PageQuery) GetLimit() int {
	return q.PageSize
}

// Order 排序对象
type Order struct {
	SortedBy string
	Order    string
}

// OrderStr 返回排序语句
func (o Order) OrderStr() string {
	if !o.IsEmptyStr() {
		return strings.Join([]string{o.SortedBy, o.Order}, " ")
	}
	return ""
}

// IsEmptyStr 判断排序字段是否为空，并进行补充
func (o Order) IsEmptyStr() bool {
	if o.Order == "" {
		return true
	}

	if o.SortedBy != "" {
		return false
	}

	o.SortedBy = "desc"
	return false
}

// Operator 操作者
type Operator struct {
	Rtx string // 用户账户
}

func (o Operator) GetRtx() string {
	return o.Rtx
}
