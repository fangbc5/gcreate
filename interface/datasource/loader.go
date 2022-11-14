package datasource

import (
	"gcreate/db"
)

type Db interface {
	Conn()
	Close()
	GetTables(names ...string) []db.Table
}
