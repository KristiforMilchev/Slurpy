package interfaces

type RowScanner interface {
	Scan(dest ...any) error
	Err() error
}
type RowsScanner interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error
	Err() error
}
