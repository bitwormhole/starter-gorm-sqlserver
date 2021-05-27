package main

import (
	"fmt"

	"github.com/bitwormhole/starter-gorm-sqlserver/sqlserver"
)

func main() {
	fmt.Println("starter-gorm-sqlserver")
	dr := &sqlserver.Driver{}
	dr.Open(nil)
}
