package driver

import (
	"errors"
	"strconv"
	"strings"

	"github.com/bitwormhole/starter-gorm/datasource"
	"github.com/bitwormhole/starter/markup"
	driver_pkg "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// SQLServerDriver 是SQLServer的starter-gorm驱动
type SQLServerDriver struct {
	markup.Component
}

func (inst *SQLServerDriver) _Impl() datasource.Driver {
	return inst
}

// Accept 用于确定是否支持给定的配置
func (inst *SQLServerDriver) Accept(cfg *datasource.Configuration) bool {
	name := cfg.Driver
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	return name == "sqlserver"
}

func (inst *SQLServerDriver) prepareForDefaultPort(cfg *datasource.Configuration) {
	const defport = 1433
	port := cfg.Port
	if port < 1 {
		port = defport
	}
	cfg.Port = port
}

// Open 打开数据源
func (inst *SQLServerDriver) Open(cfg *datasource.Configuration) (datasource.Source, error) {

	//	dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930/?database=gorm"
	//	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if cfg == nil {
		return nil, errors.New("config==nil")
	}

	inst.prepareForDefaultPort(cfg)

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
	builder.Config1 = *cfg
	builder.Config2 = *gc
	builder.Dialector = dialector
	return builder.Open()
}
