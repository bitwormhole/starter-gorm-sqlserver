package sqlserver

import (
	"errors"
	"strconv"
	"strings"

	"github.com/bitwormhole/starter-gorm/datasource"
	driver_pkg "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Driver struct {
}

func (inst *Driver) Open(cfg *datasource.Configuration) (datasource.Source, error) {

	//	dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930/?database=gorm"
	//	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if cfg == nil {
		return nil, errors.New("config==nil")
	}

	dsnbuilder := &strings.Builder{}
	dsnbuilder.WriteString("sqlserver://")
	dsnbuilder.WriteString(cfg.Username)
	dsnbuilder.WriteString(":")
	dsnbuilder.WriteString(cfg.Password)
	dsnbuilder.WriteString("@")
	dsnbuilder.WriteString(cfg.Host)
	dsnbuilder.WriteString(":")
	dsnbuilder.WriteString(strconv.Itoa(cfg.Port))
	dsnbuilder.WriteString("?database=")
	dsnbuilder.WriteString(cfg.Database)
	dsn := dsnbuilder.String()

	dialector := driver_pkg.Open(dsn)
	if dialector == nil {
		return nil, errors.New("driver_sqlserver.Open() return nil")
	}

	gc := &gorm.Config{}

	builder := &datasource.GormDataSourceBuilder{}
	builder.Config1 = cfg
	builder.Config2 = gc
	builder.Dialector = dialector
	return builder.Create()
}
