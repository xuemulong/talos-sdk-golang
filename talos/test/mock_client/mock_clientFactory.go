// Code generated by MockGen. DO NOT EDIT.
// Source: TalosClientFactory.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	consumer "github.com/XiaoMi/talos-sdk-golang/talos/thrift/consumer"
	message "github.com/XiaoMi/talos-sdk-golang/talos/thrift/message"
	topic "github.com/XiaoMi/talos-sdk-golang/talos/thrift/topic"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTalosClient is a mock of TalosClient interface
type MockTalosClient struct {
	ctrl     *gomock.Controller
	recorder *MockTalosClientMockRecorder
}

// MockTalosClientMockRecorder is the mock recorder for MockTalosClient
type MockTalosClientMockRecorder struct {
	mock *MockTalosClient
}

// NewMockTalosClient creates a new mock instance
func NewMockTalosClient(ctrl *gomock.Controller) *MockTalosClient {
	mock := &MockTalosClient{ctrl: ctrl}
	mock.recorder = &MockTalosClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTalosClient) EXPECT() *MockTalosClientMockRecorder {
	return m.recorder
}

// NewTopicClient mocks base method
func (m *MockTalosClient) NewTopicClient(url string) topic.TopicService {
	ret := m.ctrl.Call(m, "NewTopicClient", url)
	ret0, _ := ret[0].(topic.TopicService)
	return ret0
}

// NewTopicClient indicates an expected call of NewTopicClient
func (mr *MockTalosClientMockRecorder) NewTopicClient(url interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTopicClient", reflect.TypeOf((*MockTalosClient)(nil).NewTopicClient), url)
}

// NewMessageClient mocks base method
func (m *MockTalosClient) NewMessageClient(url string) message.MessageService {
	ret := m.ctrl.Call(m, "NewMessageClient", url)
	ret0, _ := ret[0].(message.MessageService)
	return ret0
}

// NewMessageClient indicates an expected call of NewMessageClient
func (mr *MockTalosClientMockRecorder) NewMessageClient(url interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewMessageClient", reflect.TypeOf((*MockTalosClient)(nil).NewMessageClient), url)
}

// NewConsumerClient mocks base method
func (m *MockTalosClient) NewConsumerClient(url string) consumer.ConsumerService {
	ret := m.ctrl.Call(m, "NewConsumerClient", url)
	ret0, _ := ret[0].(consumer.ConsumerService)
	return ret0
}

// NewConsumerClient indicates an expected call of NewConsumerClient
func (mr *MockTalosClientMockRecorder) NewConsumerClient(url interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewConsumerClient", reflect.TypeOf((*MockTalosClient)(nil).NewConsumerClient), url)
}

// NewTopicClientDefault mocks base method
func (m *MockTalosClient) NewTopicClientDefault() topic.TopicService {
	ret := m.ctrl.Call(m, "NewTopicClientDefault")
	ret0, _ := ret[0].(topic.TopicService)
	return ret0
}

// NewTopicClientDefault indicates an expected call of NewTopicClientDefault
func (mr *MockTalosClientMockRecorder) NewTopicClientDefault() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTopicClientDefault", reflect.TypeOf((*MockTalosClient)(nil).NewTopicClientDefault))
}

// NewMessageClientDefault mocks base method
func (m *MockTalosClient) NewMessageClientDefault() message.MessageService {
	ret := m.ctrl.Call(m, "NewMessageClientDefault")
	ret0, _ := ret[0].(message.MessageService)
	return ret0
}

// NewMessageClientDefault indicates an expected call of NewMessageClientDefault
func (mr *MockTalosClientMockRecorder) NewMessageClientDefault() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewMessageClientDefault", reflect.TypeOf((*MockTalosClient)(nil).NewMessageClientDefault))
}

// NewConsumerClientDefault mocks base method
func (m *MockTalosClient) NewConsumerClientDefault() consumer.ConsumerService {
	ret := m.ctrl.Call(m, "NewConsumerClientDefault")
	ret0, _ := ret[0].(consumer.ConsumerService)
	return ret0
}

// NewConsumerClientDefault indicates an expected call of NewConsumerClientDefault
func (mr *MockTalosClientMockRecorder) NewConsumerClientDefault() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewConsumerClientDefault", reflect.TypeOf((*MockTalosClient)(nil).NewConsumerClientDefault))
}
