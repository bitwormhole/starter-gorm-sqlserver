package sqlserver

import (
	"embed"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

//go:embed src/test/resources
var theTestResFS embed.FS

// ExportTestModule 【注意】仅供测试使用！
func ExportTestModule() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(myName + "#testing").Version(myVersion).Revision(myRevision)
	mb.Dependency(DriverModule())
	mb.Resources(collection.LoadEmbedResources(&theTestResFS, "src/test/resources"))

	return mb.Create()
}
