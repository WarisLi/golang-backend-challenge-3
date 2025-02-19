package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockBeefRepository struct {
	mock.Mock
}

func (m *MockBeefRepository) GetData() ([]byte, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Error(1)
}
