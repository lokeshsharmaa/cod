// Code generated by MockGen. DO NOT EDIT.
// Source: assignment_crossnokaye/cod/externalservice/clients/itemservice (interfaces: Client)

// Package mocks is a generated GoMock package.
package mocks

import (
	itemservice "assignment_crossnokaye/cod/externalservice/clients/itemservice"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CreateItem mocks base method.
func (m *MockClient) CreateItem(arg0 context.Context, arg1 *itemservice.CreatePayload) (*int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateItem", arg0, arg1)
	ret0, _ := ret[0].(*int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateItem indicates an expected call of CreateItem.
func (mr *MockClientMockRecorder) CreateItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItem", reflect.TypeOf((*MockClient)(nil).CreateItem), arg0, arg1)
}

// DeleteItem mocks base method.
func (m *MockClient) DeleteItem(arg0 context.Context, arg1 *itemservice.DeletePayload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteItem", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteItem indicates an expected call of DeleteItem.
func (mr *MockClientMockRecorder) DeleteItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteItem", reflect.TypeOf((*MockClient)(nil).DeleteItem), arg0, arg1)
}

// GetItem mocks base method.
func (m *MockClient) GetItem(arg0 context.Context, arg1 *itemservice.GetPayload) (*itemservice.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItem", arg0, arg1)
	ret0, _ := ret[0].(*itemservice.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItem indicates an expected call of GetItem.
func (mr *MockClientMockRecorder) GetItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItem", reflect.TypeOf((*MockClient)(nil).GetItem), arg0, arg1)
}

// UpdateItem mocks base method.
func (m *MockClient) UpdateItem(arg0 context.Context, arg1 *itemservice.UpdatePayload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItem", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateItem indicates an expected call of UpdateItem.
func (mr *MockClientMockRecorder) UpdateItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItem", reflect.TypeOf((*MockClient)(nil).UpdateItem), arg0, arg1)
}
