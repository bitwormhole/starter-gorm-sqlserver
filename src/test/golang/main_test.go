package golang

import (
	"testing"

	sqlserver "github.com/bitwormhole/starter-gorm-sqlserver"
	"github.com/bitwormhole/starter/tests"
)

func TestSQLServerDriver(t *testing.T) {

	i := tests.Starter(t)
	i.Use(sqlserver.ExportTestModule())
	rt := i.RunTest()

	rt.Exit()
}
