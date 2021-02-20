package models

import (
	"github.com/pubgo/x/xconfig"
	"github.com/pubgo/x/xconfig/xconfig_rds"
	"github.com/pubgo/x/xdi"
	"xorm.io/xorm"
)

func init() {
	var _ IExampleModel = (*Example)(nil)

	xdi.InitProvide(func(rds *xconfig_rds.Rds, cfg *xconfig.Config) IExampleModel {
		return &Example{db: rds.GetRDS(cfg.Web.Db.Name)}
	})
}

type IExampleModel interface {
	CreateOne(data Example) (err error)
	CreateMany() (err error)
	Delete() (err error)
	Update() (err error)
	FindOne() (data Example, err error)
	FindMany() (data []Example, err error)
}

type Example struct {
	db *xorm.Engine
}

func (Example) CreateOne(data Example) (err error) {
	panic("implement me")
}

func (Example) CreateMany() (err error) {
	panic("implement me")
}

func (Example) Delete() (err error) {
	panic("implement me")
}

func (Example) Update() (err error) {
	panic("implement me")
}

func (Example) FindOne() (data Example, err error) {
	panic("implement me")
}

func (Example) FindMany() (data []Example, err error) {
	panic("implement me")
}
