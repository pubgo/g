package xmigrate

import (
	"gopkg.in/gormigrate.v1"
)

var migrations []*gormigrate.Migration

func Registry(m *gormigrate.Migration) {
	migrations = append(migrations, m)
}

func Migrations() []*gormigrate.Migration {
	return migrations
}
