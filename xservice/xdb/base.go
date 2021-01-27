package xdb

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pubgo/x/xerror"
	"github.com/pubgo/schema"
	"strings"
)

func SearchObject(db *gorm.DB, wheres map[string]interface{}, out interface{}) error {
	if out == nil {
		return errors.New("param[outArr] is empty")
	}

	if len(wheres) > 0 {
		db = db.Where(wheres)
	}

	err := db.Find(out).Error
	if err != nil {
		return err
	}

	return nil
}

func SearchObjectByIn(db *gorm.DB, wheres map[string]interface{},
	ins map[string]interface{}, out interface{}) error {

	if out == nil {
		return errors.New("param[outArr] is empty")
	}

	if len(wheres) > 0 {
		db = db.Where(wheres)
	}

	if len(ins) > 0 {
		for key, value := range ins {
			db = db.Where(key, value)
		}
	}

	err := db.Find(out).Error
	if err != nil {
		return err
	}

	return nil
}

func SearchObjectByOrder(db *gorm.DB, wheres map[string]interface{}, ins map[string]interface{},
	orders string, limit, offset int, out interface{}) error {
	if out == nil {
		return errors.New("param[outArr] is empty")
	}

	if len(wheres) > 0 {
		db = db.Where(wheres)
	}

	if len(ins) > 0 {
		for key, value := range ins {
			db = db.Where(key, value)
		}
	}

	if orders != "" {
		orderList := strings.Split(orders, ",")
		if len(orderList) > 1 {
			for _, itemOrder := range orderList {
				db = db.Order(itemOrder)
			}
		} else {
			db = db.Order(orders)
		}
	}

	if limit > 0 {
		db = db.Limit(limit)
	}

	if offset > 0 {
		db = db.Offset(offset)
	}

	err := db.Find(out).Error
	if err != nil {
		return err
	}
	return nil
}

// updates maybe use gorm.Expr()
func UpdateObject(db *gorm.DB, wheres, updates map[string]interface{}) error {
	if len(wheres) == 0 {
		return errors.New("param[wheres] length is zero")
	}
	if len(updates) == 0 {
		return errors.New("param[updates] length is zero")
	}

	err := db.Where(wheres).UpdateColumns(updates).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateObjectIns(db *gorm.DB, wheres, ins, updates map[string]interface{}) error {
	if len(wheres) == 0 && len(ins) == 0 {
		return errors.New("param[wheres and ins] length is zero")
	}

	if len(updates) == 0 {
		return errors.New("param[updates] length is zero")
	}

	if len(wheres) > 0 {
		db = db.Where(wheres)
	}

	if len(ins) > 0 {
		for key, value := range ins {
			db = db.Where(key, value)
		}
	}

	err := db.UpdateColumns(updates).Error
	if err != nil {
		return err
	}

	return nil
}

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

func TableNames(db *sql.DB) ([]string, error) {
	return schema.TableNames(db)
}

func ViewNames(db *sql.DB) ([]string, error) {
	return schema.ViewNames(db)
}

type TbCol struct {
}

func TableCols(db *sql.DB, tbNames ...string) (map[string]TbCol, error) {
	if len(tbNames) == 0 {
		ds, err := schema.Tables(db)
	} else {
		for _, tb := range tbNames {
			ds, err := schema.Table(db, tb)
			if err != nil {
				return nil, err
			}

			for _, d := range ds {
				d.Name()
				d.DatabaseTypeName()
				d.DecimalSize()
				d.Length()
				d.Nullable()
				d.ScanType()
			}
		}
	}
}
