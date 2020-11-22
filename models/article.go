package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Article 文章表
type Article struct {
	gorm.Model
	TagID     int    `gorm:"column:tag_id;type:int(10) unsigned;not null;index" json:"tag_id"`   // 标签ID
	Title     string `gorm:"column:title;type:varchar(100);not null" json:"title"`               // 标题
	Describe  string `gorm:"column:describe;type:varchar(255);not null" json:"describe"`         // 描述
	Content   string `gorm:"column:content;type:text;not null" json:"content"`                   // 内容
	Status    int8   `gorm:"column:status;type:tinyint(2) unsigned;not null" json:"status"`      // 状态 1正常 0不可见
	CreatedBy uint32 `gorm:"column:created_by;type:int(10) unsigned;not null" json:"created_by"` // 创建人
	DeletedBy int    `gorm:"column:deleted_by;type:int(11)" json:"deleted_by"`                   // 删除人
	Tag       Tag    `json:"tag"`
}

func (ac *Article) BeforeCreate(scope *gorm.Scope) error {
	ac.CreatedAt = time.Now()
	ac.UpdatedAt = time.Now()
	return nil
}

func (au *Article) BeforeUpdate(scope *gorm.Scope) error {
	au.UpdatedAt = time.Now()
	return nil
}
