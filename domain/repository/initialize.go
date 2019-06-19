package repository

import (
	"github.com/gocraft/dbr"
	"github.com/kcwebapply/bm/infrastructure/db"
)

var conn *dbr.Connection
var sess *dbr.Session

func init() {
	conn = db.GetConnection()
	sess = conn.NewSession(nil)
	CreateCategoryTable()
	CreateCategoryTable()
}
