// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/berkaymuratt/sep-app-api/services (interfaces: JwtServiceI)

// Package services is a generated GoMock package.
package services

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockJwtServiceI is a mock of JwtServiceI interface.
type MockJwtServiceI struct {
	ctrl     *gomock.Controller
	recorder *MockJwtServiceIMockRecorder
}

// MockJwtServiceIMockRecorder is the mock recorder for MockJwtServiceI.
type MockJwtServiceIMockRecorder struct {
	mock *MockJwtServiceI
}

// NewMockJwtServiceI creates a new mock instance.
func NewMockJwtServiceI(ctrl *gomock.Controller) *MockJwtServiceI {
	mock := &MockJwtServiceI{ctrl: ctrl}
	mock.recorder = &MockJwtServiceIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJwtServiceI) EXPECT() *MockJwtServiceIMockRecorder {
	return m.recorder
}

// CheckJwt mocks base method.
func (m *MockJwtServiceI) CheckJwt(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckJwt", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckJwt indicates an expected call of CheckJwt.
func (mr *MockJwtServiceIMockRecorder) CheckJwt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckJwt", reflect.TypeOf((*MockJwtServiceI)(nil).CheckJwt), arg0)
}

// GenerateJwtToken mocks base method.
func (m *MockJwtServiceI) GenerateJwtToken(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateJwtToken", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateJwtToken indicates an expected call of GenerateJwtToken.
func (mr *MockJwtServiceIMockRecorder) GenerateJwtToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateJwtToken", reflect.TypeOf((*MockJwtServiceI)(nil).GenerateJwtToken), arg0)
}