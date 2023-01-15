package cfgsqlserver

import "github.com/bitwormhole/starter/application"

// ExportConfigForSQLServer ...
func ExportConfigForSQLServer(cb application.ConfigBuilder) error {
	// return nil
	return autoGenConfig(cb)
}
