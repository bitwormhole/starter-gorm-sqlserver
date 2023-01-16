package sqlserver

import (
	"embed"

	startergorm "github.com/bitwormhole/starter-gorm"
	"github.com/bitwormhole/starter-gorm-sqlserver/gen/cfgsqlserver"
	"github.com/bitwormhole/starter-gorm-sqlserver/gen/cfgsqlserverdemo"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	theModuleName     = "github.com/bitwormhole/starter-gorm-sqlserver"
	theModuleVersion  = "v0.0.5"
	theModuleRevision = 5
	theModuleResPath  = "src/main/resources"
)

//go:embed src/main/resources
var theModuleResFS embed.FS

// DriverModule 导出[starter-gorm-sqlserver]模块
func DriverModule() application.Module {

	mb := &application.ModuleBuilder{}

	mb.Name(theModuleName + "#driver").Version(theModuleVersion).Revision(theModuleRevision)
	mb.OnMount(cfgsqlserver.ExportConfigForSQLServer)
	mb.Resources(collection.LoadEmbedResources(&theModuleResFS, theModuleResPath))
	mb.Dependency(startergorm.Module())

	return mb.Create()
}

// DemoModule 导出[#demo]模块
func DemoModule() application.Module {

	mb := &application.ModuleBuilder{}

	mb.Name(theModuleName + "#demo").Version(theModuleVersion).Revision(theModuleRevision)
	mb.OnMount(cfgsqlserverdemo.ExportConfigForSQLServerDemo)
	mb.Resources(collection.LoadEmbedResources(&theModuleResFS, theModuleResPath))
	mb.Dependency(DriverModule())

	return mb.Create()
}
