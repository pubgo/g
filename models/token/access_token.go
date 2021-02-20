package token

import (
	"database/sql"
	"github.com/pubgo/x/xconfig"
	"github.com/pubgo/x/xconfig/xconfig_rds"
	"github.com/pubgo/x/xdi"
	"time"
)

func init() {
	xdi.InitProvide(func(rds *xconfig_rds.Rds, cfg *xconfig.Config) IAccessTokens {
		return &AccessTokens{}
	})
}

type IAccessTokens interface {
	CreateOne(data AccessTokens) (err error)
	CreateMany() (err error)
	Delete() (err error)
	Update() (err error)
	FindOne() (data AccessTokens, err error)
	FindMany() (data []AccessTokens, err error)
	Paginate() (data []AccessTokens, err error)
}

type AccessTokens struct {
	ID            int          `json:"id"`
	Type          string       `json:"type"`
	UserID        int          `json:"user_id"`
	Description   string       `json:"description"`
	LastUsedAt    sql.NullTime `json:"last_used_at"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	SpaceID       int          `json:"space_id"`
	Scope         string       `json:"scope"`
	ApplicationID string       `json:"application_id"`
	Token         string       `json:"token"`
	Application   string       `json:"application"`
	Serializer    string       `json:"_serializer"`
}

func (t *AccessTokens) CreateOne(data AccessTokens) (err error) {
	panic("implement me")
}

func (t *AccessTokens) CreateMany() (err error) {
	panic("implement me")
}

func (t *AccessTokens) Delete() (err error) {
	panic("implement me")
}

func (t *AccessTokens) Update() (err error) {
	panic("implement me")
}

func (t *AccessTokens) FindOne() (data AccessTokens, err error) {
	panic("implement me")
}

func (t *AccessTokens) FindMany() (data []AccessTokens, err error) {
	panic("implement me")
}

func (t *AccessTokens) Paginate() (data []AccessTokens, err error) {
	panic("implement me")
}
