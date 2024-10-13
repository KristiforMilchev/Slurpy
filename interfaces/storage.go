package interfaces

import "database/sql"

type Storage interface {
	New(dbPath *string, dbName *string, tables *string) bool
	Initialize()
	Query(sql *string, parameters *[]interface{}) (*sql.Rows, error)
	QuerySingle(sql *string, parameters *[]interface{}) *sql.Row
	Exec(sql *string, parameters *[]interface{}) error
	Open() bool
	Close()
}
