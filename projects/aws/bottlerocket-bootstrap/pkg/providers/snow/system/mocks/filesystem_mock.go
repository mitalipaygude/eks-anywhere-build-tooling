// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/eks-anywhere-build-tooling/bottlerocket-bootstrap/pkg/providers/snow/system (interfaces: FileSystem)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFileSystem is a mock of FileSystem interface.
type MockFileSystem struct {
	ctrl     *gomock.Controller
	recorder *MockFileSystemMockRecorder
}

// MockFileSystemMockRecorder is the mock recorder for MockFileSystem.
type MockFileSystemMockRecorder struct {
	mock *MockFileSystem
}

// NewMockFileSystem creates a new mock instance.
func NewMockFileSystem(ctrl *gomock.Controller) *MockFileSystem {
	mock := &MockFileSystem{ctrl: ctrl}
	mock.recorder = &MockFileSystemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileSystem) EXPECT() *MockFileSystemMockRecorder {
	return m.recorder
}

// MountVolume mocks base method.
func (m *MockFileSystem) MountVolume(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MountVolume", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MountVolume indicates an expected call of MountVolume.
func (mr *MockFileSystemMockRecorder) MountVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MountVolume", reflect.TypeOf((*MockFileSystem)(nil).MountVolume), arg0, arg1)
}

// Partition mocks base method.
func (m *MockFileSystem) Partition(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Partition", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Partition indicates an expected call of Partition.
func (mr *MockFileSystemMockRecorder) Partition(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Partition", reflect.TypeOf((*MockFileSystem)(nil).Partition), arg0)
}
