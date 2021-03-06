// Code generated by MockGen. DO NOT EDIT.
// Source: exporter.go

// Package mock_main is a generated GoMock package.
package mock_main

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockExportInterface is a mock of ExportInterface interface
type MockExportInterface struct {
	ctrl     *gomock.Controller
	recorder *MockExportInterfaceMockRecorder
}

// MockExportInterfaceMockRecorder is the mock recorder for MockExportInterface
type MockExportInterfaceMockRecorder struct {
	mock *MockExportInterface
}

// NewMockExportInterface creates a new mock instance
func NewMockExportInterface(ctrl *gomock.Controller) *MockExportInterface {
	mock := &MockExportInterface{ctrl: ctrl}
	mock.recorder = &MockExportInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExportInterface) EXPECT() *MockExportInterfaceMockRecorder {
	return m.recorder
}

// Format mocks base method
func (m *MockExportInterface) Format() MockFormatInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Format")
	ret0, _ := ret[0].(MockFormatInterface)
	return ret0
}

// Format indicates an expected call of Format
func (mr *MockExportInterfaceMockRecorder) Format() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Format", reflect.TypeOf((*MockExportInterface)(nil).Format))
}

// GetFileName mocks base method
func (m *MockExportInterface) GetFileName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetFileName indicates an expected call of GetFileName
func (mr *MockExportInterfaceMockRecorder) GetFileName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileName", reflect.TypeOf((*MockExportInterface)(nil).GetFileName))
}

// GetSource mocks base method
func (m *MockExportInterface) GetSource() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSource")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSource indicates an expected call of GetSource
func (mr *MockExportInterfaceMockRecorder) GetSource() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSource", reflect.TypeOf((*MockExportInterface)(nil).GetSource))
}
