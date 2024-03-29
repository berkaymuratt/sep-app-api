// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/berkaymuratt/sep-app-api/services (interfaces: BodyPartsServiceI)

// Package services is a generated GoMock package.
package services

import (
	reflect "reflect"

	models "github.com/berkaymuratt/sep-app-api/models"
	gomock "github.com/golang/mock/gomock"
)

// MockBodyPartsServiceI is a mock of BodyPartsServiceI interface.
type MockBodyPartsServiceI struct {
	ctrl     *gomock.Controller
	recorder *MockBodyPartsServiceIMockRecorder
}

// MockBodyPartsServiceIMockRecorder is the mock recorder for MockBodyPartsServiceI.
type MockBodyPartsServiceIMockRecorder struct {
	mock *MockBodyPartsServiceI
}

// NewMockBodyPartsServiceI creates a new mock instance.
func NewMockBodyPartsServiceI(ctrl *gomock.Controller) *MockBodyPartsServiceI {
	mock := &MockBodyPartsServiceI{ctrl: ctrl}
	mock.recorder = &MockBodyPartsServiceIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBodyPartsServiceI) EXPECT() *MockBodyPartsServiceIMockRecorder {
	return m.recorder
}

// GetBodyParts mocks base method.
func (m *MockBodyPartsServiceI) GetBodyParts() ([]*models.BodyPart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBodyParts")
	ret0, _ := ret[0].([]*models.BodyPart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBodyParts indicates an expected call of GetBodyParts.
func (mr *MockBodyPartsServiceIMockRecorder) GetBodyParts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBodyParts", reflect.TypeOf((*MockBodyPartsServiceI)(nil).GetBodyParts))
}
