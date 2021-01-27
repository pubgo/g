package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/pubgo/x/xdi"
)

func init() {
	xdi.InitProvide(func() IAuthCtrl {
		return &authCtrl{}
	})
}

// auth
type IAuthCtrl interface {
	Register(context *gin.Context)
	CreatePassword(context *gin.Context)
	UpdatePassword(context *gin.Context)
	Login(context *gin.Context)
	Logout(context *gin.Context)
	Verify(context *gin.Context)
}

type authCtrl struct {
}

func (t *authCtrl) Register(context *gin.Context) {
	panic("implement me")
}

func (t *authCtrl) CreatePassword(context *gin.Context) {
	panic("implement me")
}

func (t *authCtrl) UpdatePassword(context *gin.Context) {
	panic("implement me")
}

func (t *authCtrl) Login(context *gin.Context) {
	panic("implement me")
}

func (t *authCtrl) Logout(context *gin.Context) {
	panic("implement me")
}

func (t *authCtrl) Verify(context *gin.Context) {
	panic("implement me")
}
