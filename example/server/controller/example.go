package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pubgo/x/example/server/models"
	"github.com/pubgo/x/xdi"
)

func init() {
	var _ IExampleCtrl = (*ExampleCtrl)(nil)

	xdi.InitProvide(func(model models.IExampleModel) IExampleCtrl {
		return &ExampleCtrl{model: model}
	})
}

type IExampleCtrl interface {
	CreateOne(context *gin.Context)
	CreateMany(context *gin.Context)
	Delete(context *gin.Context)
	Update(context *gin.Context)
	GetOne(context *gin.Context)
	GetMany(context *gin.Context)
	Paginate(context *gin.Context)
	Clone() IExampleCtrl
}

var _ IExampleCtrl = (*ExampleCtrl)(nil)

type ExampleCtrl struct {
	model models.IExampleModel
}

func (t *ExampleCtrl) Clone() IExampleCtrl {
	return &(*t)
}

func (t *ExampleCtrl) CreateOne(context *gin.Context) {
	panic("implement me")
}

func (*ExampleCtrl) CreateMany(context *gin.Context) {
	panic("implement me")
}

func (*ExampleCtrl) Delete(context *gin.Context) {
	panic("implement me")
}

func (*ExampleCtrl) Update(context *gin.Context) {
	panic("implement me")
}

func (*ExampleCtrl) GetOne(context *gin.Context) {
	panic("implement me")
}

func (*ExampleCtrl) GetMany(context *gin.Context) {
	panic("implement me")
}

func (*ExampleCtrl) Paginate(context *gin.Context) {
	panic("implement me")
}
