package interfaces

type Storage interface {
	New(dbPath *string, dbName *string, tables *string) bool
	Initialize()
	Query(sql *string, parameters *[]interface{}) (RowsScanner, error)
	QuerySingle(sql *string, parameters *[]interface{}) RowScanner
	Exec(sql *string, parameters *[]interface{}) error
	Open() bool
	Close() bool
}
