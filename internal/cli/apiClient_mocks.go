// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/teserakt/c2/internal/cli (interfaces: APIClientFactory,C2Client)

// Package cli is a generated GoMock package.
package cli

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	cobra "github.com/spf13/cobra"
	pb "gitlab.com/teserakt/c2/pkg/pb"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockAPIClientFactory is a mock of APIClientFactory interface
type MockAPIClientFactory struct {
	ctrl     *gomock.Controller
	recorder *MockAPIClientFactoryMockRecorder
}

// MockAPIClientFactoryMockRecorder is the mock recorder for MockAPIClientFactory
type MockAPIClientFactoryMockRecorder struct {
	mock *MockAPIClientFactory
}

// NewMockAPIClientFactory creates a new mock instance
func NewMockAPIClientFactory(ctrl *gomock.Controller) *MockAPIClientFactory {
	mock := &MockAPIClientFactory{ctrl: ctrl}
	mock.recorder = &MockAPIClientFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPIClientFactory) EXPECT() *MockAPIClientFactoryMockRecorder {
	return m.recorder
}

// NewClient mocks base method
func (m *MockAPIClientFactory) NewClient(arg0 *cobra.Command) (C2Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewClient", arg0)
	ret0, _ := ret[0].(C2Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewClient indicates an expected call of NewClient
func (mr *MockAPIClientFactoryMockRecorder) NewClient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewClient", reflect.TypeOf((*MockAPIClientFactory)(nil).NewClient), arg0)
}

// MockC2Client is a mock of C2Client interface
type MockC2Client struct {
	ctrl     *gomock.Controller
	recorder *MockC2ClientMockRecorder
}

// MockC2ClientMockRecorder is the mock recorder for MockC2Client
type MockC2ClientMockRecorder struct {
	mock *MockC2Client
}

// NewMockC2Client creates a new mock instance
func NewMockC2Client(ctrl *gomock.Controller) *MockC2Client {
	mock := &MockC2Client{ctrl: ctrl}
	mock.recorder = &MockC2ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockC2Client) EXPECT() *MockC2ClientMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockC2Client) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockC2ClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockC2Client)(nil).Close))
}

// CountClients mocks base method
func (m *MockC2Client) CountClients(arg0 context.Context, arg1 *pb.CountClientsRequest, arg2 ...grpc.CallOption) (*pb.CountClientsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CountClients", varargs...)
	ret0, _ := ret[0].(*pb.CountClientsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountClients indicates an expected call of CountClients
func (mr *MockC2ClientMockRecorder) CountClients(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountClients", reflect.TypeOf((*MockC2Client)(nil).CountClients), varargs...)
}

// CountClientsForTopic mocks base method
func (m *MockC2Client) CountClientsForTopic(arg0 context.Context, arg1 *pb.CountClientsForTopicRequest, arg2 ...grpc.CallOption) (*pb.CountClientsForTopicResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CountClientsForTopic", varargs...)
	ret0, _ := ret[0].(*pb.CountClientsForTopicResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountClientsForTopic indicates an expected call of CountClientsForTopic
func (mr *MockC2ClientMockRecorder) CountClientsForTopic(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountClientsForTopic", reflect.TypeOf((*MockC2Client)(nil).CountClientsForTopic), varargs...)
}

// CountTopics mocks base method
func (m *MockC2Client) CountTopics(arg0 context.Context, arg1 *pb.CountTopicsRequest, arg2 ...grpc.CallOption) (*pb.CountTopicsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CountTopics", varargs...)
	ret0, _ := ret[0].(*pb.CountTopicsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountTopics indicates an expected call of CountTopics
func (mr *MockC2ClientMockRecorder) CountTopics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountTopics", reflect.TypeOf((*MockC2Client)(nil).CountTopics), varargs...)
}

// CountTopicsForClient mocks base method
func (m *MockC2Client) CountTopicsForClient(arg0 context.Context, arg1 *pb.CountTopicsForClientRequest, arg2 ...grpc.CallOption) (*pb.CountTopicsForClientResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CountTopicsForClient", varargs...)
	ret0, _ := ret[0].(*pb.CountTopicsForClientResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountTopicsForClient indicates an expected call of CountTopicsForClient
func (mr *MockC2ClientMockRecorder) CountTopicsForClient(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountTopicsForClient", reflect.TypeOf((*MockC2Client)(nil).CountTopicsForClient), varargs...)
}

// GetClients mocks base method
func (m *MockC2Client) GetClients(arg0 context.Context, arg1 *pb.GetClientsRequest, arg2 ...grpc.CallOption) (*pb.GetClientsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetClients", varargs...)
	ret0, _ := ret[0].(*pb.GetClientsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClients indicates an expected call of GetClients
func (mr *MockC2ClientMockRecorder) GetClients(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClients", reflect.TypeOf((*MockC2Client)(nil).GetClients), varargs...)
}

// GetClientsForTopic mocks base method
func (m *MockC2Client) GetClientsForTopic(arg0 context.Context, arg1 *pb.GetClientsForTopicRequest, arg2 ...grpc.CallOption) (*pb.GetClientsForTopicResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetClientsForTopic", varargs...)
	ret0, _ := ret[0].(*pb.GetClientsForTopicResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientsForTopic indicates an expected call of GetClientsForTopic
func (mr *MockC2ClientMockRecorder) GetClientsForTopic(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientsForTopic", reflect.TypeOf((*MockC2Client)(nil).GetClientsForTopic), varargs...)
}

// GetTopics mocks base method
func (m *MockC2Client) GetTopics(arg0 context.Context, arg1 *pb.GetTopicsRequest, arg2 ...grpc.CallOption) (*pb.GetTopicsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTopics", varargs...)
	ret0, _ := ret[0].(*pb.GetTopicsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopics indicates an expected call of GetTopics
func (mr *MockC2ClientMockRecorder) GetTopics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopics", reflect.TypeOf((*MockC2Client)(nil).GetTopics), varargs...)
}

// GetTopicsForClient mocks base method
func (m *MockC2Client) GetTopicsForClient(arg0 context.Context, arg1 *pb.GetTopicsForClientRequest, arg2 ...grpc.CallOption) (*pb.GetTopicsForClientResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTopicsForClient", varargs...)
	ret0, _ := ret[0].(*pb.GetTopicsForClientResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopicsForClient indicates an expected call of GetTopicsForClient
func (mr *MockC2ClientMockRecorder) GetTopicsForClient(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopicsForClient", reflect.TypeOf((*MockC2Client)(nil).GetTopicsForClient), varargs...)
}

// NewClient mocks base method
func (m *MockC2Client) NewClient(arg0 context.Context, arg1 *pb.NewClientRequest, arg2 ...grpc.CallOption) (*pb.NewClientResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewClient", varargs...)
	ret0, _ := ret[0].(*pb.NewClientResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewClient indicates an expected call of NewClient
func (mr *MockC2ClientMockRecorder) NewClient(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewClient", reflect.TypeOf((*MockC2Client)(nil).NewClient), varargs...)
}

// NewClientKey mocks base method
func (m *MockC2Client) NewClientKey(arg0 context.Context, arg1 *pb.NewClientKeyRequest, arg2 ...grpc.CallOption) (*pb.NewClientKeyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewClientKey", varargs...)
	ret0, _ := ret[0].(*pb.NewClientKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewClientKey indicates an expected call of NewClientKey
func (mr *MockC2ClientMockRecorder) NewClientKey(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewClientKey", reflect.TypeOf((*MockC2Client)(nil).NewClientKey), varargs...)
}

// NewTopic mocks base method
func (m *MockC2Client) NewTopic(arg0 context.Context, arg1 *pb.NewTopicRequest, arg2 ...grpc.CallOption) (*pb.NewTopicResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewTopic", varargs...)
	ret0, _ := ret[0].(*pb.NewTopicResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewTopic indicates an expected call of NewTopic
func (mr *MockC2ClientMockRecorder) NewTopic(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTopic", reflect.TypeOf((*MockC2Client)(nil).NewTopic), varargs...)
}

// NewTopicClient mocks base method
func (m *MockC2Client) NewTopicClient(arg0 context.Context, arg1 *pb.NewTopicClientRequest, arg2 ...grpc.CallOption) (*pb.NewTopicClientResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewTopicClient", varargs...)
	ret0, _ := ret[0].(*pb.NewTopicClientResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewTopicClient indicates an expected call of NewTopicClient
func (mr *MockC2ClientMockRecorder) NewTopicClient(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTopicClient", reflect.TypeOf((*MockC2Client)(nil).NewTopicClient), varargs...)
}

// RemoveClient mocks base method
func (m *MockC2Client) RemoveClient(arg0 context.Context, arg1 *pb.RemoveClientRequest, arg2 ...grpc.CallOption) (*pb.RemoveClientResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveClient", varargs...)
	ret0, _ := ret[0].(*pb.RemoveClientResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveClient indicates an expected call of RemoveClient
func (mr *MockC2ClientMockRecorder) RemoveClient(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveClient", reflect.TypeOf((*MockC2Client)(nil).RemoveClient), varargs...)
}

// RemoveTopic mocks base method
func (m *MockC2Client) RemoveTopic(arg0 context.Context, arg1 *pb.RemoveTopicRequest, arg2 ...grpc.CallOption) (*pb.RemoveTopicResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveTopic", varargs...)
	ret0, _ := ret[0].(*pb.RemoveTopicResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveTopic indicates an expected call of RemoveTopic
func (mr *MockC2ClientMockRecorder) RemoveTopic(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTopic", reflect.TypeOf((*MockC2Client)(nil).RemoveTopic), varargs...)
}

// RemoveTopicClient mocks base method
func (m *MockC2Client) RemoveTopicClient(arg0 context.Context, arg1 *pb.RemoveTopicClientRequest, arg2 ...grpc.CallOption) (*pb.RemoveTopicClientResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveTopicClient", varargs...)
	ret0, _ := ret[0].(*pb.RemoveTopicClientResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveTopicClient indicates an expected call of RemoveTopicClient
func (mr *MockC2ClientMockRecorder) RemoveTopicClient(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTopicClient", reflect.TypeOf((*MockC2Client)(nil).RemoveTopicClient), varargs...)
}

// ResetClient mocks base method
func (m *MockC2Client) ResetClient(arg0 context.Context, arg1 *pb.ResetClientRequest, arg2 ...grpc.CallOption) (*pb.ResetClientResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResetClient", varargs...)
	ret0, _ := ret[0].(*pb.ResetClientResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetClient indicates an expected call of ResetClient
func (mr *MockC2ClientMockRecorder) ResetClient(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetClient", reflect.TypeOf((*MockC2Client)(nil).ResetClient), varargs...)
}

// SendMessage mocks base method
func (m *MockC2Client) SendMessage(arg0 context.Context, arg1 *pb.SendMessageRequest, arg2 ...grpc.CallOption) (*pb.SendMessageResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendMessage", varargs...)
	ret0, _ := ret[0].(*pb.SendMessageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessage indicates an expected call of SendMessage
func (mr *MockC2ClientMockRecorder) SendMessage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockC2Client)(nil).SendMessage), varargs...)
}

// SubscribeToEventStream mocks base method
func (m *MockC2Client) SubscribeToEventStream(arg0 context.Context, arg1 *pb.SubscribeToEventStreamRequest, arg2 ...grpc.CallOption) (pb.C2_SubscribeToEventStreamClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubscribeToEventStream", varargs...)
	ret0, _ := ret[0].(pb.C2_SubscribeToEventStreamClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeToEventStream indicates an expected call of SubscribeToEventStream
func (mr *MockC2ClientMockRecorder) SubscribeToEventStream(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToEventStream", reflect.TypeOf((*MockC2Client)(nil).SubscribeToEventStream), varargs...)
}
