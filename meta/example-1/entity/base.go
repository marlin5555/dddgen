package entity

import "time"

// Trailer field combination
type Trailer struct {
	Creator   string    `gorm:"column:creator;type:varchar(255)" json:"creator"`                        // 创建人
	Updater   string    `gorm:"column:updater;type:varchar(255)" json:"updater"`                        // 更新人
	DeletedAt int64     `gorm:"column:deleted_at;default:0" json:"deleted_at"`                          // 是否已删除，0: 未删除，1: 已删除
	CreatedAt time.Time `gorm:"column:created_at;type:time;<-:create;autoCreateTime" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:time;autoUpdateTime" json:"updated_at"`           // 修改时间
}
