// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/berkaymuratt/sep-app-api/services (interfaces: MiddlewareServiceI)

// Package services is a generated GoMock package.
package services

import (
	reflect "reflect"

	fiber "github.com/gofiber/fiber/v2"
	gomock "github.com/golang/mock/gomock"
)

// MockMiddlewareServiceI is a mock of MiddlewareServiceI interface.
type MockMiddlewareServiceI struct {
	ctrl     *gomock.Controller
	recorder *MockMiddlewareServiceIMockRecorder
}

// MockMiddlewareServiceIMockRecorder is the mock recorder for MockMiddlewareServiceI.
type MockMiddlewareServiceIMockRecorder struct {
	mock *MockMiddlewareServiceI
}

// NewMockMiddlewareServiceI creates a new mock instance.
func NewMockMiddlewareServiceI(ctrl *gomock.Controller) *MockMiddlewareServiceI {
	mock := &MockMiddlewareServiceI{ctrl: ctrl}
	mock.recorder = &MockMiddlewareServiceIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMiddlewareServiceI) EXPECT() *MockMiddlewareServiceIMockRecorder {
	return m.recorder
}

// CORSMiddleware mocks base method.
func (m *MockMiddlewareServiceI) CORSMiddleware() func(*fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CORSMiddleware")
	ret0, _ := ret[0].(func(*fiber.Ctx) error)
	return ret0
}

// CORSMiddleware indicates an expected call of CORSMiddleware.
func (mr *MockMiddlewareServiceIMockRecorder) CORSMiddleware() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CORSMiddleware", reflect.TypeOf((*MockMiddlewareServiceI)(nil).CORSMiddleware))
}

// Middleware mocks base method.
func (m *MockMiddlewareServiceI) Middleware(arg0 *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Middleware", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Middleware indicates an expected call of Middleware.
func (mr *MockMiddlewareServiceIMockRecorder) Middleware(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Middleware", reflect.TypeOf((*MockMiddlewareServiceI)(nil).Middleware), arg0)
}
