package mocks

type MockRowsScanner struct {
}

func (m *MockRowsScanner) Scan(dest ...interface{}) error {
	return nil
}

func (m *MockRowsScanner) Next() bool {
	return false
}

func (m *MockRowsScanner) Close() error {
	return nil
}

func (m *MockRowsScanner) Err() error {
	return nil
}
