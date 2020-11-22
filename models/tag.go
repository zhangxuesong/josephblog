package models

import (
	"github.com/jinzhu/gorm"
	"github.com/zhangxuesong/josephblog/pkg/orm"
	"time"
)

// Tag 标签表
type Tag struct {
	orm.Model
	Name      string `gorm:"column:name;type:varchar(100);not null" json:"name"`                 // 标签名称
	CreatedBy uint32 `gorm:"column:created_by;type:int(10) unsigned;not null" json:"created_by"` // 创建人
	UpdatedBy int    `gorm:"column:updated_by;type:int(10) unsigned;not null" json:"updated_by"` // 修改人
	DeletedBy int    `gorm:"column:deleted_by;type:int(10) unsigned" json:"deleted_by"`          // 删除人
}

func (tc *Tag) BeforeCreate(scope *gorm.Scope) error {
	tc.CreatedAt = time.Now()
	tc.UpdatedAt = time.Now()
	return nil
}

func (tu *Tag) BeforeUpdate(scope *gorm.Scope) error {
	tu.UpdatedAt = time.Now()
	return nil
}
