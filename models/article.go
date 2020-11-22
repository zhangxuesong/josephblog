package models

import "github.com/jinzhu/gorm"

// Article 文章表
type Article struct {
	gorm.Model
	TagID     int    `gorm:"column:tag_id;type:int(10) unsigned;not null;index"` // 标签ID
	Title     string `gorm:"column:title;type:varchar(100);not null"`            // 标题
	Describe  string `gorm:"column:describe;type:varchar(255);not null"`         // 描述
	Content   string `gorm:"column:content;type:text;not null"`                  // 内容
	Status    int8   `gorm:"column:status;type:tinyint(2) unsigned;not null"`    // 状态 1正常 0不可见
	CreatedBy uint32 `gorm:"column:created_by;type:int(10) unsigned;not null"`   // 创建人
	DeletedBy int    `gorm:"column:deleted_by;type:int(11)"`                     // 删除人
	Tag       Tag    `json:"tag"`
}
