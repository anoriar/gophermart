// Code generated by MockGen. DO NOT EDIT.
// Source: password_service_interface.go

// Package mock_password_service is a generated GoMock package.
package mock_password_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPasswordServiceInterface is a mock of PasswordServiceInterface interface.
type MockPasswordServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPasswordServiceInterfaceMockRecorder
}

// MockPasswordServiceInterfaceMockRecorder is the mock recorder for MockPasswordServiceInterface.
type MockPasswordServiceInterfaceMockRecorder struct {
	mock *MockPasswordServiceInterface
}

// NewMockPasswordServiceInterface creates a new mock instance.
func NewMockPasswordServiceInterface(ctrl *gomock.Controller) *MockPasswordServiceInterface {
	mock := &MockPasswordServiceInterface{ctrl: ctrl}
	mock.recorder = &MockPasswordServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPasswordServiceInterface) EXPECT() *MockPasswordServiceInterfaceMockRecorder {
	return m.recorder
}

// GenerateHashedPassword mocks base method.
func (m *MockPasswordServiceInterface) GenerateHashedPassword(password string, salt []byte) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateHashedPassword", password, salt)
	ret0, _ := ret[0].(string)
	return ret0
}

// GenerateHashedPassword indicates an expected call of GenerateHashedPassword.
func (mr *MockPasswordServiceInterfaceMockRecorder) GenerateHashedPassword(password, salt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateHashedPassword", reflect.TypeOf((*MockPasswordServiceInterface)(nil).GenerateHashedPassword), password, salt)
}
