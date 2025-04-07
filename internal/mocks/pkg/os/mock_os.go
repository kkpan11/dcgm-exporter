// Copyright (c) 2024, NVIDIA CORPORATION.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/NVIDIA/dcgm-exporter/internal/pkg/os (interfaces: OS)
//
// Generated by this command:
//
//	mockgen -destination=../../mocks/pkg/os/mock_os.go -package=os -copyright_file=../../../hack/header.txt . OS
//

// Package os is a generated GoMock package.
package os

import (
	os "os"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockOS is a mock of OS interface.
type MockOS struct {
	ctrl     *gomock.Controller
	recorder *MockOSMockRecorder
	isgomock struct{}
}

// MockOSMockRecorder is the mock recorder for MockOS.
type MockOSMockRecorder struct {
	mock *MockOS
}

// NewMockOS creates a new mock instance.
func NewMockOS(ctrl *gomock.Controller) *MockOS {
	mock := &MockOS{ctrl: ctrl}
	mock.recorder = &MockOSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOS) EXPECT() *MockOSMockRecorder {
	return m.recorder
}

// CreateTemp mocks base method.
func (m *MockOS) CreateTemp(dir, pattern string) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTemp", dir, pattern)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTemp indicates an expected call of CreateTemp.
func (mr *MockOSMockRecorder) CreateTemp(dir, pattern any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTemp", reflect.TypeOf((*MockOS)(nil).CreateTemp), dir, pattern)
}

// Exit mocks base method.
func (m *MockOS) Exit(code int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Exit", code)
}

// Exit indicates an expected call of Exit.
func (mr *MockOSMockRecorder) Exit(code any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exit", reflect.TypeOf((*MockOS)(nil).Exit), code)
}

// Getenv mocks base method.
func (m *MockOS) Getenv(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Getenv", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// Getenv indicates an expected call of Getenv.
func (mr *MockOSMockRecorder) Getenv(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Getenv", reflect.TypeOf((*MockOS)(nil).Getenv), key)
}

// Hostname mocks base method.
func (m *MockOS) Hostname() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hostname")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Hostname indicates an expected call of Hostname.
func (mr *MockOSMockRecorder) Hostname() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hostname", reflect.TypeOf((*MockOS)(nil).Hostname))
}

// IsNotExist mocks base method.
func (m *MockOS) IsNotExist(err error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNotExist", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNotExist indicates an expected call of IsNotExist.
func (mr *MockOSMockRecorder) IsNotExist(err any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNotExist", reflect.TypeOf((*MockOS)(nil).IsNotExist), err)
}

// MkdirTemp mocks base method.
func (m *MockOS) MkdirTemp(dir, pattern string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MkdirTemp", dir, pattern)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MkdirTemp indicates an expected call of MkdirTemp.
func (mr *MockOSMockRecorder) MkdirTemp(dir, pattern any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MkdirTemp", reflect.TypeOf((*MockOS)(nil).MkdirTemp), dir, pattern)
}

// Open mocks base method.
func (m *MockOS) Open(name string) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", name)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *MockOSMockRecorder) Open(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockOS)(nil).Open), name)
}

// ReadDir mocks base method.
func (m *MockOS) ReadDir(name string) ([]os.DirEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDir", name)
	ret0, _ := ret[0].([]os.DirEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDir indicates an expected call of ReadDir.
func (mr *MockOSMockRecorder) ReadDir(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDir", reflect.TypeOf((*MockOS)(nil).ReadDir), name)
}

// Remove mocks base method.
func (m *MockOS) Remove(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockOSMockRecorder) Remove(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockOS)(nil).Remove), name)
}

// RemoveAll mocks base method.
func (m *MockOS) RemoveAll(path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAll", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAll indicates an expected call of RemoveAll.
func (mr *MockOSMockRecorder) RemoveAll(path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAll", reflect.TypeOf((*MockOS)(nil).RemoveAll), path)
}

// Stat mocks base method.
func (m *MockOS) Stat(name string) (os.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stat", name)
	ret0, _ := ret[0].(os.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stat indicates an expected call of Stat.
func (mr *MockOSMockRecorder) Stat(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockOS)(nil).Stat), name)
}

// TempDir mocks base method.
func (m *MockOS) TempDir() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TempDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// TempDir indicates an expected call of TempDir.
func (mr *MockOSMockRecorder) TempDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TempDir", reflect.TypeOf((*MockOS)(nil).TempDir))
}
