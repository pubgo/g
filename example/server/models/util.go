package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pubgo/g/xerror"
)

func Delete(db *gorm.DB, model interface{}, filter map[string]interface{}) (err error) {
	_sql := ""
	var _params = []interface{}{""}
	if filter != nil {
		for k, v := range filter {
			_sql += fmt.Sprintf("%s=?", k)
			_params = append(_params, v)
		}
	}
	_params[0] = _sql
	return db.Delete(model, _params...).Error
}

func Update(db *gorm.DB, data interface{}) (err error) {
	return db.Updates(data).Error
}

func FindOne(db *gorm.DB, filter map[string]interface{}, data interface{}, ) (err error) {
	_sql := ""
	var _params = []interface{}{""}
	if filter != nil {
		for k, v := range filter {
			_sql += fmt.Sprintf("%s=?", k)
			_params = append(_params, v)
		}
	}
	_params[0] = _sql
	return db.First(data, _params...).Error
}

func FindMany(db *gorm.DB, filter map[string]interface{}, data []interface{}) (err error) {
	_sql := ""
	var _params = []interface{}{""}
	if filter != nil {
		for k, v := range filter {
			_sql += fmt.Sprintf("%s=?", k)
			_params = append(_params, v)
		}
	}
	_params[0] = _sql
	return db.First(data, _params...).Error
}

func Paginate(db *gorm.DB, pageSize, pageIndex int64, filter map[string]interface{}, data []interface{}, ) (err error) {
	_sql := ""
	var _params []interface{}
	if filter != nil {
		for k, v := range filter {
			_sql += fmt.Sprintf("%s=?", k)
			_params = append(_params, v)
		}
	}

	if pageIndex < 1 {
		pageIndex = 1
	}
	return db.Where(_sql, _params...).Limit(pageSize).Offset(pageSize * (pageIndex - 1)).Find(data).Error
}

func CreateOne(db *gorm.DB, data interface{}) (err error) {
	return db.Create(data).Error
}
func CreateMany(db *gorm.DB, data ...interface{}) (err error) {
	if len(data) == 0 {
		return
	}

	tx := db.Begin()
	for _, dt := range data {
		xerror.Panic(tx.Create(dt).Error)
	}
	return tx.Commit().Error
}
