package mocks

import "errors"

type MockRowScanner struct {
}

func (m *MockRowScanner) Scan(dest ...any) error {

	if len(dest) > 0 {
		*dest[0].(*int) = 1
	}
	return nil
}

func (m *MockRowScanner) Err() error {
	return errors.New("no rows found")
}
