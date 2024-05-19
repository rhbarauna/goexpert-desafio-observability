package mocks

import (
	"github.com/stretchr/testify/mock"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var _ trace.Span = (*MockSpan)(nil)

type MockSpan struct{}

func (m *MockSpan) End(options ...trace.SpanEndOption)                  {}
func (m *MockSpan) AddEvent(name string, options ...trace.EventOption)  {}
func (m *MockSpan) IsRecording() bool                                   { return false }
func (m *MockSpan) RecordError(err error, options ...trace.EventOption) {}
func (m *MockSpan) SpanContext() trace.SpanContext                      { return trace.SpanContext{} }
func (m *MockSpan) SetStatus(code codes.Code, description string)       {}
func (m *MockSpan) SetName(name string)                                 {}
func (m *MockSpan) SetAttributes(kv ...attribute.KeyValue)              {}
func (m *MockSpan) TracerProvider() trace.TracerProvider                { return nil }
func (m *MockSpan) AddLink(link trace.Link)                             {}
func (m *MockSpan) span()                                               {}

var _ trace.TracerProvider = (*TracerProviderMock)(nil)

type TracerProviderMock struct {
	mock.Mock
}

func NewTracerProviderMock() *TracerProviderMock {
	return &TracerProviderMock{}
}

func (tp TracerProviderMock) Tracer(name string, options ...trace.TracerOption) trace.Tracer {
	return nil
}
func (tp TracerProviderMock) tracerProvider() {}
