// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_sandbox is a generated GoMock package.
package mock_sandbox

import (
	sandbox "github.com/c3systems/c3-go/core/sandbox"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Play mocks base method
func (m *MockInterface) Play(config *sandbox.PlayConfig) ([]byte, error) {
	ret := m.ctrl.Call(m, "Play", config)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Play indicates an expected call of Play
func (mr *MockInterfaceMockRecorder) Play(config interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Play", reflect.TypeOf((*MockInterface)(nil).Play), config)
}
