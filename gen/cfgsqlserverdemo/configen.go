package cfgsqlserverdemo

import "github.com/bitwormhole/starter/application"

// ExportConfigForSQLServerDemo ...
func ExportConfigForSQLServerDemo(cb application.ConfigBuilder) error {
	// return nil
	return autoGenConfig(cb)
}
