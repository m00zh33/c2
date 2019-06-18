// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/teserakt/c2/internal/protocols (interfaces: PubSubClient)

// Package protocols is a generated GoMock package.
package protocols

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPubSubClient is a mock of PubSubClient interface
type MockPubSubClient struct {
	ctrl     *gomock.Controller
	recorder *MockPubSubClientMockRecorder
}

// MockPubSubClientMockRecorder is the mock recorder for MockPubSubClient
type MockPubSubClientMockRecorder struct {
	mock *MockPubSubClient
}

// NewMockPubSubClient creates a new mock instance
func NewMockPubSubClient(ctrl *gomock.Controller) *MockPubSubClient {
	mock := &MockPubSubClient{ctrl: ctrl}
	mock.recorder = &MockPubSubClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPubSubClient) EXPECT() *MockPubSubClientMockRecorder {
	return m.recorder
}

// Connect mocks base method
func (m *MockPubSubClient) Connect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect
func (mr *MockPubSubClientMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockPubSubClient)(nil).Connect))
}

// Disconnect mocks base method
func (m *MockPubSubClient) Disconnect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disconnect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Disconnect indicates an expected call of Disconnect
func (mr *MockPubSubClientMockRecorder) Disconnect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockPubSubClient)(nil).Disconnect))
}

// Publish mocks base method
func (m *MockPubSubClient) Publish(arg0 context.Context, arg1 []byte, arg2 string, arg3 byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish
func (mr *MockPubSubClientMockRecorder) Publish(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockPubSubClient)(nil).Publish), arg0, arg1, arg2, arg3)
}

// SubscribeToTopic mocks base method
func (m *MockPubSubClient) SubscribeToTopic(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeToTopic", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubscribeToTopic indicates an expected call of SubscribeToTopic
func (mr *MockPubSubClientMockRecorder) SubscribeToTopic(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToTopic", reflect.TypeOf((*MockPubSubClient)(nil).SubscribeToTopic), arg0, arg1)
}

// SubscribeToTopics mocks base method
func (m *MockPubSubClient) SubscribeToTopics(arg0 context.Context, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeToTopics", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubscribeToTopics indicates an expected call of SubscribeToTopics
func (mr *MockPubSubClientMockRecorder) SubscribeToTopics(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToTopics", reflect.TypeOf((*MockPubSubClient)(nil).SubscribeToTopics), arg0, arg1)
}

// UnsubscribeFromTopic mocks base method
func (m *MockPubSubClient) UnsubscribeFromTopic(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsubscribeFromTopic", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsubscribeFromTopic indicates an expected call of UnsubscribeFromTopic
func (mr *MockPubSubClientMockRecorder) UnsubscribeFromTopic(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeFromTopic", reflect.TypeOf((*MockPubSubClient)(nil).UnsubscribeFromTopic), arg0, arg1)
}
