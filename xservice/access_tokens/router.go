package access_tokens

import (
	"github.com/gin-gonic/gin"
	"github.com/pubgo/x/models/token"
	"github.com/pubgo/x/xdi"
)

func init() {
	xdi.InitProvide(func(model token.IAccessTokens) IAccessTokenCtrl {
		return &accessTokenCtrl{model: model}
	})
}

// access_token
type IAccessTokenCtrl interface {
	Create(context *gin.Context)
	Delete(context *gin.Context)
	Update(context *gin.Context)
	GetOne(context *gin.Context)
	GetMany(context *gin.Context)
	Paginate(context *gin.Context)
	ResetToken(context *gin.Context)
}

type accessTokenCtrl struct {
	model token.IAccessTokens
}

func (t *accessTokenCtrl) ResetToken(context *gin.Context) {
	//	 {"description":"njnjnjss","scope":"artboard:read,repo","type":"oauth"}
}

func (t *accessTokenCtrl) Create(context *gin.Context) {
	// description
	// scope
	// type=oauth
}

func (t *accessTokenCtrl) Delete(context *gin.Context) {
	panic("implement me")
}

func (t *accessTokenCtrl) Update(context *gin.Context) {
	panic("implement me")
}

func (t *accessTokenCtrl) GetOne(context *gin.Context) {
	panic("implement me")
}

func (t *accessTokenCtrl) GetMany(context *gin.Context) {
	panic("implement me")
}

func (t *accessTokenCtrl) Paginate(context *gin.Context) {
	panic("implement me")
}
