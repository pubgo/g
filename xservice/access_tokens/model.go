package access_tokens

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"github.com/pubgo/x/xconfig"
	"github.com/pubgo/x/xconfig/xconfig_rds"
	"github.com/pubgo/x/xdi"
	"time"
)

func init() {
	xdi.InitProvide(func(rds *xconfig_rds.Rds, cfg *xconfig.Config) IAccessTokenModel {
		return &AccessToken{db: rds.GetRDS(cfg.Web.Db.Name)}
	})
}

type IAccessTokenModel interface {
	CreateOne(data AccessToken) (err error)
	CreateMany() (err error)
	Delete() (err error)
	Update() (err error)
	FindOne() (data AccessToken, err error)
	FindMany() (data []AccessToken, err error)
	Paginate() (data []AccessToken, err error)
}

type AccessToken struct {
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
	db            *gorm.DB
}

func (t *AccessToken) CreateOne(data AccessToken) (err error) {
	panic("implement me")
}

func (t *AccessToken) CreateMany() (err error) {
	panic("implement me")
}

func (t *AccessToken) Delete() (err error) {
	panic("implement me")
}

func (t *AccessToken) Update() (err error) {
	panic("implement me")
}

func (t *AccessToken) FindOne() (data AccessToken, err error) {
	panic("implement me")
}

func (t *AccessToken) FindMany() (data []AccessToken, err error) {
	panic("implement me")
}

func (t *AccessToken) Paginate() (data []AccessToken, err error) {
	panic("implement me")
}
