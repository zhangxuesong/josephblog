package models

// Admin 管理员表
type Admin struct {
	ID       uint32 `gorm:"primaryKey;column:id;type:int(10) unsigned;not null"`
	Username string `gorm:"column:username;type:varchar(64);not null"` // 账号
	Password string `gorm:"column:password;type:varchar(64);not null"` // 密码
}
