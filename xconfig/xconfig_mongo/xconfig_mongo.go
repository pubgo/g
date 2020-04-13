package xconfig_mongo

import (
	"context"
	"github.com/pubgo/g/pkg"
	"github.com/pubgo/g/xconfig"
	"github.com/pubgo/g/xdi"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"github.com/pubgo/g/xerror"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongo struct {
	_mongo map[string]*mongo.Client
}

func (t *Mongo) CloseMongo(name ...string) (err error) {
	defer xerror.RespErr(&err)

	c := t.GetMongo(name...)
	xerror.PanicM(c.Disconnect(context.Background()), "mongo close failed")

	return
}

// GetMongo get mongo instance
func (t *Mongo) GetMongo(name ...string) (c *mongo.Client) {
	_name := xconfig.DefaultName
	if len(name) > 0 {
		_name = name[0]
	}

	c = t._mongo[_name]
	xerror.PanicT(pkg.IsNone(c), "mongo instance %s is nil", _name)
	return
}

// GetDb get mongo database
func (t *Mongo) GetDb(name ...string) (db *mongo.Database, err error) {
	defer xerror.RespErr(&err)

	_db := ""
	_name := ""
	switch len(name) {
	case 0:
		_name = xconfig.DefaultName
		xerror.Panic(xdi.Invoke(func(config *xconfig.Config) {
			_cfg := config.Mongodb
			for _, cfg := range _cfg.Cfg {
				if _name == cfg.Name {
					_db = cfg.Database
					break
				}
			}
		}))
	case 1:
		_db = name[0]
	default:
		_name = name[0]
		_db = name[1]
	}

	xerror.PanicT(_db == "", "database name is nil")

	db = t.GetMongo(_name).Database(_db)
	return
}

func (t *Mongo) GetCol(name ...string) (col *mongo.Collection, err error) {
	defer xerror.RespErr(&err)

	xerror.PanicT(len(name) < 1, "mongo collection name is empty")
	db := xerror.PanicErr(t.GetDb(name[:len(name)-1]...)).(*mongo.Database)
	col = db.Collection(name[len(name)-1])

	return
}

func init() {
	xdi.InitProvide(func(cfg *xconfig.Config) *Mongo {
		defer xerror.Assert()

		// 加载配置
		_cfg := cfg.Mongodb

		if _cfg.Default == "" {

		}

		xerror.PanicT(_cfg.Default == "", "default name is nil")
		xerror.PanicT(len(_cfg.Cfg) == 0, "mongo config count is 0")
		_mongo := make(map[string]*mongo.Client, len(_cfg.Cfg))

		for _, cfg := range _cfg.Cfg {

			//认证参数设置，否则连不上
			opts := &options.ClientOptions{}
			opts.SetAuth(options.Credential{
				AuthMechanism: cfg.AuthMechanism,
				AuthSource:    cfg.AuthSource,
				Username:      cfg.Username,
				Password:      cfg.Password,
			})
			opts.ApplyURI(cfg.URL)

			_client := xerror.PanicErr(mongo.Connect(context.Background(), opts)).(*mongo.Client)
			ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
			xerror.PanicM(_client.Ping(ctx, readpref.Primary()), "mongo %s ping failed", cfg.AppName)
			_mongo[cfg.Name] = _client
		}
		_mongo[xconfig.DefaultName] = _mongo[_cfg.Default]
		return &Mongo{_mongo: _mongo}
	})
}

func containsKey(doc bson.Raw, key ...string) (b bool) {
	_, err := doc.LookupErr(key...)
	return err == nil
}
