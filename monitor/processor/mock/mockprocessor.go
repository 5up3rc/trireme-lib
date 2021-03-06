// Code generated by MockGen. DO NOT EDIT.
// Source: monitor/processor/interfaces.go

// Package mockprocessor is a generated GoMock package.
package mockprocessor

import (
	context "context"
	reflect "reflect"

	common "github.com/aporeto-inc/trireme-lib/common"
	gomock "github.com/golang/mock/gomock"
)

// MockProcessor is a mock of Processor interface
// nolint
type MockProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockProcessorMockRecorder
}

// MockProcessorMockRecorder is the mock recorder for MockProcessor
// nolint
type MockProcessorMockRecorder struct {
	mock *MockProcessor
}

// NewMockProcessor creates a new mock instance
// nolint
func NewMockProcessor(ctrl *gomock.Controller) *MockProcessor {
	mock := &MockProcessor{ctrl: ctrl}
	mock.recorder = &MockProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
// nolint
func (m *MockProcessor) EXPECT() *MockProcessorMockRecorder {
	return m.recorder
}

// Start mocks base method
// nolint
func (m *MockProcessor) Start(ctx context.Context, eventInfo *common.EventInfo) error {
	ret := m.ctrl.Call(m, "Start", ctx, eventInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
// nolint
func (mr *MockProcessorMockRecorder) Start(ctx, eventInfo interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockProcessor)(nil).Start), ctx, eventInfo)
}

// Stop mocks base method
// nolint
func (m *MockProcessor) Stop(ctx context.Context, eventInfo *common.EventInfo) error {
	ret := m.ctrl.Call(m, "Stop", ctx, eventInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
// nolint
func (mr *MockProcessorMockRecorder) Stop(ctx, eventInfo interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockProcessor)(nil).Stop), ctx, eventInfo)
}

// Create mocks base method
// nolint
func (m *MockProcessor) Create(ctx context.Context, eventInfo *common.EventInfo) error {
	ret := m.ctrl.Call(m, "Create", ctx, eventInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
// nolint
func (mr *MockProcessorMockRecorder) Create(ctx, eventInfo interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProcessor)(nil).Create), ctx, eventInfo)
}

// Destroy mocks base method
// nolint
func (m *MockProcessor) Destroy(ctx context.Context, eventInfo *common.EventInfo) error {
	ret := m.ctrl.Call(m, "Destroy", ctx, eventInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy
// nolint
func (mr *MockProcessorMockRecorder) Destroy(ctx, eventInfo interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockProcessor)(nil).Destroy), ctx, eventInfo)
}

// Pause mocks base method
// nolint
func (m *MockProcessor) Pause(ctx context.Context, eventInfo *common.EventInfo) error {
	ret := m.ctrl.Call(m, "Pause", ctx, eventInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Pause indicates an expected call of Pause
// nolint
func (mr *MockProcessorMockRecorder) Pause(ctx, eventInfo interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pause", reflect.TypeOf((*MockProcessor)(nil).Pause), ctx, eventInfo)
}

// ReSync mocks base method
// nolint
func (m *MockProcessor) ReSync(ctx context.Context, EventInfo *common.EventInfo) error {
	ret := m.ctrl.Call(m, "ReSync", ctx, EventInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReSync indicates an expected call of ReSync
// nolint
func (mr *MockProcessorMockRecorder) ReSync(ctx, EventInfo interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReSync", reflect.TypeOf((*MockProcessor)(nil).ReSync), ctx, EventInfo)
}
