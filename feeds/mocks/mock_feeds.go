// Code generated by MockGen. DO NOT EDIT.
// Source: feeds/interfaces.go

// Package mock_feeds is a generated GoMock package.
package mock_feeds

import (
	reflect "reflect"

	types "github.com/ethereum/go-ethereum/core/types"
	contracts "github.com/forta-network/forta-node/contracts"
	domain "github.com/forta-network/forta-node/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockBlockFeed is a mock of BlockFeed interface.
type MockBlockFeed struct {
	ctrl     *gomock.Controller
	recorder *MockBlockFeedMockRecorder
}

// MockBlockFeedMockRecorder is the mock recorder for MockBlockFeed.
type MockBlockFeedMockRecorder struct {
	mock *MockBlockFeed
}

// NewMockBlockFeed creates a new mock instance.
func NewMockBlockFeed(ctrl *gomock.Controller) *MockBlockFeed {
	mock := &MockBlockFeed{ctrl: ctrl}
	mock.recorder = &MockBlockFeedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlockFeed) EXPECT() *MockBlockFeedMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockBlockFeed) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockBlockFeedMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockBlockFeed)(nil).Start))
}

// Subscribe mocks base method.
func (m *MockBlockFeed) Subscribe(handler func(*domain.BlockEvent) error) <-chan error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", handler)
	ret0, _ := ret[0].(<-chan error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockBlockFeedMockRecorder) Subscribe(handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockBlockFeed)(nil).Subscribe), handler)
}

// MockTransactionFeed is a mock of TransactionFeed interface.
type MockTransactionFeed struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionFeedMockRecorder
}

// MockTransactionFeedMockRecorder is the mock recorder for MockTransactionFeed.
type MockTransactionFeedMockRecorder struct {
	mock *MockTransactionFeed
}

// NewMockTransactionFeed creates a new mock instance.
func NewMockTransactionFeed(ctrl *gomock.Controller) *MockTransactionFeed {
	mock := &MockTransactionFeed{ctrl: ctrl}
	mock.recorder = &MockTransactionFeedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionFeed) EXPECT() *MockTransactionFeedMockRecorder {
	return m.recorder
}

// ForEachTransaction mocks base method.
func (m *MockTransactionFeed) ForEachTransaction(blockHandler func(*domain.BlockEvent) error, txHandler func(*domain.TransactionEvent) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEachTransaction", blockHandler, txHandler)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEachTransaction indicates an expected call of ForEachTransaction.
func (mr *MockTransactionFeedMockRecorder) ForEachTransaction(blockHandler, txHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEachTransaction", reflect.TypeOf((*MockTransactionFeed)(nil).ForEachTransaction), blockHandler, txHandler)
}

// MockLogFeed is a mock of LogFeed interface.
type MockLogFeed struct {
	ctrl     *gomock.Controller
	recorder *MockLogFeedMockRecorder
}

// MockLogFeedMockRecorder is the mock recorder for MockLogFeed.
type MockLogFeedMockRecorder struct {
	mock *MockLogFeed
}

// NewMockLogFeed creates a new mock instance.
func NewMockLogFeed(ctrl *gomock.Controller) *MockLogFeed {
	mock := &MockLogFeed{ctrl: ctrl}
	mock.recorder = &MockLogFeedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogFeed) EXPECT() *MockLogFeedMockRecorder {
	return m.recorder
}

// ForEachLog mocks base method.
func (m *MockLogFeed) ForEachLog(blockHandler func(*domain.Block) error, handler func(types.Log) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEachLog", blockHandler, handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEachLog indicates an expected call of ForEachLog.
func (mr *MockLogFeedMockRecorder) ForEachLog(blockHandler, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEachLog", reflect.TypeOf((*MockLogFeed)(nil).ForEachLog), blockHandler, handler)
}

// MockAlertFeed is a mock of AlertFeed interface.
type MockAlertFeed struct {
	ctrl     *gomock.Controller
	recorder *MockAlertFeedMockRecorder
}

// MockAlertFeedMockRecorder is the mock recorder for MockAlertFeed.
type MockAlertFeedMockRecorder struct {
	mock *MockAlertFeed
}

// NewMockAlertFeed creates a new mock instance.
func NewMockAlertFeed(ctrl *gomock.Controller) *MockAlertFeed {
	mock := &MockAlertFeed{ctrl: ctrl}
	mock.recorder = &MockAlertFeedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertFeed) EXPECT() *MockAlertFeedMockRecorder {
	return m.recorder
}

// ForEachAlert mocks base method.
func (m *MockAlertFeed) ForEachAlert(blockHandler func(*domain.Block) error, handler func(types.Log, *contracts.AlertsAlertBatch) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEachAlert", blockHandler, handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEachAlert indicates an expected call of ForEachAlert.
func (mr *MockAlertFeedMockRecorder) ForEachAlert(blockHandler, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEachAlert", reflect.TypeOf((*MockAlertFeed)(nil).ForEachAlert), blockHandler, handler)
}
