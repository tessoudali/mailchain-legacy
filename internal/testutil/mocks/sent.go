// Copyright 2019 Finobo
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
// Source: sent.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockSent is a mock of Sent interface
type MockSent struct {
	ctrl     *gomock.Controller
	recorder *MockSentMockRecorder
}

// MockSentMockRecorder is the mock recorder for MockSent
type MockSentMockRecorder struct {
	mock *MockSent
}

// NewMockSent creates a new mock instance
func NewMockSent(ctrl *gomock.Controller) *MockSent {
	mock := &MockSent{ctrl: ctrl}
	mock.recorder = &MockSentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSent) EXPECT() *MockSentMockRecorder {
	return m.recorder
}

// PutMessage mocks base method
func (m *MockSent) PutMessage(path string, msg io.Reader, headers map[string]string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutMessage", path, msg, headers)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutMessage indicates an expected call of PutMessage
func (mr *MockSentMockRecorder) PutMessage(path, msg, headers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutMessage", reflect.TypeOf((*MockSent)(nil).PutMessage), path, msg, headers)
}