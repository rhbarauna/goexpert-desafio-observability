// Code generated by MockGen. DO NOT EDIT.
// Source: go.opentelemetry.io/otel/trace (interfaces: Tracer)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	trace "go.opentelemetry.io/otel/trace"
)

// MockTracer is a mock of Tracer interface.
type MockTracer struct {
	ctrl     *gomock.Controller
	recorder *MockTracerMockRecorder
}

// MockTracerMockRecorder is the mock recorder for MockTracer.
type MockTracerMockRecorder struct {
	mock *MockTracer
}

// NewMockTracer creates a new mock instance.
func NewMockTracer(ctrl *gomock.Controller) *MockTracer {
	mock := &MockTracer{ctrl: ctrl}
	mock.recorder = &MockTracerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTracer) EXPECT() *MockTracerMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockTracer) Start(arg0 context.Context, arg1 string, arg2 ...trace.SpanStartOption) (context.Context, trace.Span) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Start", varargs...)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(trace.Span)
	return ret0, ret1
}

// Start indicates an expected call of Start.
func (mr *MockTracerMockRecorder) Start(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockTracer)(nil).Start), varargs...)
}

// tracer mocks base method.
func (m *MockTracer) tracer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "tracer")
}

// tracer indicates an expected call of tracer.
func (mr *MockTracerMockRecorder) tracer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "tracer", reflect.TypeOf((*MockTracer)(nil).tracer))
}
