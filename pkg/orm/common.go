package orm

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

// 分页条件
type PageWhere struct {
	Where string
	Value []interface{}
}

// 分页参数返回
type IndexPage struct {
	Total uint64 `json:"total"`  //总数
	Page  uint64 `json:"page"  ` //页数
	Num   uint64 `json:"num"   ` //数量
}

// Create
func Create(value interface{}) error {
	return db.Create(value).Error
}

// Save
func Save(value interface{}) error {
	return db.Save(value).Error
}

// Updates
func Updates(where interface{}, value interface{}) error {
	return db.Model(where).Updates(value).Error
}

// Delete
func DeleteByModel(model interface{}) (count int64, err error) {
	db := db.Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// Delete
func DeleteByWhere(model, where interface{}) (count int64, err error) {
	db := db.Where(where).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// Delete
func DeleteByID(model interface{}, id uint) (count int64, err error) {
	db := db.Where("id=?", id).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// Delete
func DeleteByIDS(model interface{}, ids []uint) (count int64, err error) {
	db := db.Where("id in (?)", ids).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// First
func FirstByIDRelated(out interface{}, id uint, related interface{}) (notFound bool, err error) {
	err = db.First(out, id).Related(related).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}

// First
func FirstByID(out interface{}, id uint) (notFound bool, err error) {
	err = db.First(out, id).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}

// First
func First(where interface{}, out interface{}) (notFound bool, err error) {
	err = db.Where(where).First(out).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}

// Find
func Find(where interface{}, out interface{}, orders ...string) error {
	db := db.Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Find(out).Error
}

// Scan
func Scan(model, where interface{}, out interface{}) (notFound bool, err error) {
	err = db.Model(model).Where(where).Scan(out).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}

// ScanList
func ScanList(model, where interface{}, out interface{}, orders ...string) error {
	db := db.Model(model).Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Scan(out).Error
}

// GetPage
func GetPage(model, where interface{}, out interface{}, indexPage *IndexPage, order string, whereOrder ...PageWhere) error {
	db := db.Model(model).Where(where)
	if order != "" {
		db = db.Order(order).Order("id DESC")
	}
	if len(whereOrder) > 0 {
		for _, wo := range whereOrder {
			if wo.Where != "" {
				db = db.Where(wo.Where, wo.Value...)
			}
		}
	}
	err := db.Count(&indexPage.Total).Error
	if err != nil {
		return err
	}
	if indexPage.Total == 0 {
		return nil
	}
	return db.Offset((indexPage.Page - 1) * indexPage.Num).Limit(indexPage.Num).Find(out).Error
}

// GetPage
func GetPageProload(model, where interface{}, out interface{}, indexPage *IndexPage, preModel, order string, whereOrder ...PageWhere) error {
	db := db.Model(model)
	if preModel != "" {
		db = db.Preload(preModel)
	}
	db = db.Where(where)
	if order != "" {
		db = db.Order(order).Order("id DESC")
	}
	if len(whereOrder) > 0 {
		for _, wo := range whereOrder {
			if wo.Where != "" {
				db = db.Where(wo.Where, wo.Value...)
			}
		}
	}
	err := db.Count(&indexPage.Total).Error
	if err != nil {
		return err
	}
	if indexPage.Total == 0 {
		return nil
	}
	return db.Offset((indexPage.Page - 1) * indexPage.Num).Limit(indexPage.Num).Find(out).Error
}

// PluckList
func PluckList(model, where interface{}, fieldName string, out interface{}) error {
	return db.Model(model).Where(where).Pluck(fieldName, out).Error
}
