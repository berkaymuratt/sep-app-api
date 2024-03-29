// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/berkaymuratt/sep-app-api/services (interfaces: SymptomsServiceI)

// Package services is a generated GoMock package.
package services

import (
	reflect "reflect"

	dtos "github.com/berkaymuratt/sep-app-api/dtos"
	models "github.com/berkaymuratt/sep-app-api/models"
	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockSymptomsServiceI is a mock of SymptomsServiceI interface.
type MockSymptomsServiceI struct {
	ctrl     *gomock.Controller
	recorder *MockSymptomsServiceIMockRecorder
}

// MockSymptomsServiceIMockRecorder is the mock recorder for MockSymptomsServiceI.
type MockSymptomsServiceIMockRecorder struct {
	mock *MockSymptomsServiceI
}

// NewMockSymptomsServiceI creates a new mock instance.
func NewMockSymptomsServiceI(ctrl *gomock.Controller) *MockSymptomsServiceI {
	mock := &MockSymptomsServiceI{ctrl: ctrl}
	mock.recorder = &MockSymptomsServiceIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSymptomsServiceI) EXPECT() *MockSymptomsServiceIMockRecorder {
	return m.recorder
}

// AddSymptom mocks base method.
func (m *MockSymptomsServiceI) AddSymptom(arg0 models.Symptom) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSymptom", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSymptom indicates an expected call of AddSymptom.
func (mr *MockSymptomsServiceIMockRecorder) AddSymptom(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSymptom", reflect.TypeOf((*MockSymptomsServiceI)(nil).AddSymptom), arg0)
}

// DeleteSymptom mocks base method.
func (m *MockSymptomsServiceI) DeleteSymptom(arg0 primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSymptom", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSymptom indicates an expected call of DeleteSymptom.
func (mr *MockSymptomsServiceIMockRecorder) DeleteSymptom(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSymptom", reflect.TypeOf((*MockSymptomsServiceI)(nil).DeleteSymptom), arg0)
}

// GetSymptomById mocks base method.
func (m *MockSymptomsServiceI) GetSymptomById(arg0 primitive.ObjectID) (*dtos.SymptomDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSymptomById", arg0)
	ret0, _ := ret[0].(*dtos.SymptomDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSymptomById indicates an expected call of GetSymptomById.
func (mr *MockSymptomsServiceIMockRecorder) GetSymptomById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSymptomById", reflect.TypeOf((*MockSymptomsServiceI)(nil).GetSymptomById), arg0)
}

// GetSymptoms mocks base method.
func (m *MockSymptomsServiceI) GetSymptoms() ([]*dtos.SymptomDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSymptoms")
	ret0, _ := ret[0].([]*dtos.SymptomDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSymptoms indicates an expected call of GetSymptoms.
func (mr *MockSymptomsServiceIMockRecorder) GetSymptoms() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSymptoms", reflect.TypeOf((*MockSymptomsServiceI)(nil).GetSymptoms))
}

// GetSymptomsByBodyPart mocks base method.
func (m *MockSymptomsServiceI) GetSymptomsByBodyPart(arg0 primitive.ObjectID) ([]*dtos.SymptomDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSymptomsByBodyPart", arg0)
	ret0, _ := ret[0].([]*dtos.SymptomDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSymptomsByBodyPart indicates an expected call of GetSymptomsByBodyPart.
func (mr *MockSymptomsServiceIMockRecorder) GetSymptomsByBodyPart(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSymptomsByBodyPart", reflect.TypeOf((*MockSymptomsServiceI)(nil).GetSymptomsByBodyPart), arg0)
}

// GetSymptomsByIds mocks base method.
func (m *MockSymptomsServiceI) GetSymptomsByIds(arg0 []primitive.ObjectID) ([]*dtos.SymptomDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSymptomsByIds", arg0)
	ret0, _ := ret[0].([]*dtos.SymptomDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSymptomsByIds indicates an expected call of GetSymptomsByIds.
func (mr *MockSymptomsServiceIMockRecorder) GetSymptomsByIds(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSymptomsByIds", reflect.TypeOf((*MockSymptomsServiceI)(nil).GetSymptomsByIds), arg0)
}

// UpdateSymptom mocks base method.
func (m *MockSymptomsServiceI) UpdateSymptom(arg0 primitive.ObjectID, arg1 models.Symptom) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSymptom", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSymptom indicates an expected call of UpdateSymptom.
func (mr *MockSymptomsServiceIMockRecorder) UpdateSymptom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSymptom", reflect.TypeOf((*MockSymptomsServiceI)(nil).UpdateSymptom), arg0, arg1)
}
