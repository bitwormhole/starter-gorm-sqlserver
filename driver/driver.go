package driver

import (
	"errors"
	"strconv"
	"strings"

	"github.com/bitwormhole/starter-gorm/datasource"
	"github.com/bitwormhole/starter/markup"
	"gorm.io/driver/sqlserver"
	driver_pkg "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// SQLServerDriver 是SQLServer的starter-gorm驱动
type SQLServerDriver struct {
	markup.Component `class:"starter-gorm-driver-registry"`
}

func (inst *SQLServerDriver) _Impl() (datasource.Driver, datasource.DriverRegistry) {
	return inst, inst
}

// GetRegistration ...
func (inst *SQLServerDriver) GetRegistration() *datasource.DriverRegistration {
	return &datasource.DriverRegistration{
		Driver: inst,
	}
}

// Accept 用于确定是否支持给定的配置
func (inst *SQLServerDriver) Accept(cfg *datasource.Configuration) bool {
	name := cfg.Driver
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	return name == "sqlserver"
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

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &sqlserverDataSource{db: db}, nil
}

func (inst *SQLServerDriver) prepareForDefaultPort(cfg *datasource.Configuration) {
	const defport = 1433
	port := cfg.Port
	if port < 1 {
		port = defport
	}
	cfg.Port = port
}

////////////////////////////////////////////////////////////////////////////////

type sqlserverDataSource struct {
	db *gorm.DB
}

func (inst *sqlserverDataSource) _Impl() datasource.Source {
	return inst
}

func (inst *sqlserverDataSource) DB() (*gorm.DB, error) {
	return inst.db, nil
}

func (inst *sqlserverDataSource) Close() error {
	return nil
}
