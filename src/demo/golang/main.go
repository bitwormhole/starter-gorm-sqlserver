package main

import (
	"github.com/bitwormhole/starter"

	starter_gorm_sqlserver "github.com/bitwormhole/starter-gorm-sqlserver"
)

func main() {
	i := starter.InitApp()
	i.UseMain(starter_gorm_sqlserver.DemoModule())
	i.Run()
}
