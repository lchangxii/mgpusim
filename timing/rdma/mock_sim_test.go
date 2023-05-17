// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/akita/akita/v3/sim (interfaces: Port,Engine)

package rdma

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sim "gitlab.com/akita/akita/v3/sim"
)

// MockPort is a mock of Port interface.
type MockPort struct {
	ctrl     *gomock.Controller
	recorder *MockPortMockRecorder
}

// MockPortMockRecorder is the mock recorder for MockPort.
type MockPortMockRecorder struct {
	mock *MockPort
}

// NewMockPort creates a new mock instance.
func NewMockPort(ctrl *gomock.Controller) *MockPort {
	mock := &MockPort{ctrl: ctrl}
	mock.recorder = &MockPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPort) EXPECT() *MockPortMockRecorder {
	return m.recorder
}

// AcceptHook mocks base method.
func (m *MockPort) AcceptHook(arg0 sim.Hook) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AcceptHook", arg0)
}

// AcceptHook indicates an expected call of AcceptHook.
func (mr *MockPortMockRecorder) AcceptHook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptHook", reflect.TypeOf((*MockPort)(nil).AcceptHook), arg0)
}

// CanSend mocks base method.
func (m *MockPort) CanSend() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanSend")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanSend indicates an expected call of CanSend.
func (mr *MockPortMockRecorder) CanSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanSend", reflect.TypeOf((*MockPort)(nil).CanSend))
}

// Component mocks base method.
func (m *MockPort) Component() sim.Component {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Component")
	ret0, _ := ret[0].(sim.Component)
	return ret0
}

// Component indicates an expected call of Component.
func (mr *MockPortMockRecorder) Component() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Component", reflect.TypeOf((*MockPort)(nil).Component))
}

// Hooks mocks base method.
func (m *MockPort) Hooks() []sim.Hook {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hooks")
	ret0, _ := ret[0].([]sim.Hook)
	return ret0
}

// Hooks indicates an expected call of Hooks.
func (mr *MockPortMockRecorder) Hooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hooks", reflect.TypeOf((*MockPort)(nil).Hooks))
}

// Name mocks base method.
func (m *MockPort) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockPortMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockPort)(nil).Name))
}

// NotifyAvailable mocks base method.
func (m *MockPort) NotifyAvailable(arg0 sim.VTimeInSec) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyAvailable", arg0)
}

// NotifyAvailable indicates an expected call of NotifyAvailable.
func (mr *MockPortMockRecorder) NotifyAvailable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyAvailable", reflect.TypeOf((*MockPort)(nil).NotifyAvailable), arg0)
}

// NumHooks mocks base method.
func (m *MockPort) NumHooks() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NumHooks")
	ret0, _ := ret[0].(int)
	return ret0
}

// NumHooks indicates an expected call of NumHooks.
func (mr *MockPortMockRecorder) NumHooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumHooks", reflect.TypeOf((*MockPort)(nil).NumHooks))
}

// Peek mocks base method.
func (m *MockPort) Peek() sim.Msg {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Peek")
	ret0, _ := ret[0].(sim.Msg)
	return ret0
}

// Peek indicates an expected call of Peek.
func (mr *MockPortMockRecorder) Peek() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Peek", reflect.TypeOf((*MockPort)(nil).Peek))
}

// Recv mocks base method.
func (m *MockPort) Recv(arg0 sim.Msg) *sim.SendError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv", arg0)
	ret0, _ := ret[0].(*sim.SendError)
	return ret0
}

// Recv indicates an expected call of Recv.
func (mr *MockPortMockRecorder) Recv(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockPort)(nil).Recv), arg0)
}

// Retrieve mocks base method.
func (m *MockPort) Retrieve(arg0 sim.VTimeInSec) sim.Msg {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Retrieve", arg0)
	ret0, _ := ret[0].(sim.Msg)
	return ret0
}

// Retrieve indicates an expected call of Retrieve.
func (mr *MockPortMockRecorder) Retrieve(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retrieve", reflect.TypeOf((*MockPort)(nil).Retrieve), arg0)
}

// Send mocks base method.
func (m *MockPort) Send(arg0 sim.Msg) *sim.SendError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(*sim.SendError)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockPortMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockPort)(nil).Send), arg0)
}

// SetConnection mocks base method.
func (m *MockPort) SetConnection(arg0 sim.Connection) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetConnection", arg0)
}

// SetConnection indicates an expected call of SetConnection.
func (mr *MockPortMockRecorder) SetConnection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConnection", reflect.TypeOf((*MockPort)(nil).SetConnection), arg0)
}

// MockEngine is a mock of Engine interface.
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine.
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance.
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// AcceptHook mocks base method.
func (m *MockEngine) AcceptHook(arg0 sim.Hook) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AcceptHook", arg0)
}

// AcceptHook indicates an expected call of AcceptHook.
func (mr *MockEngineMockRecorder) AcceptHook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptHook", reflect.TypeOf((*MockEngine)(nil).AcceptHook), arg0)
}

// Continue mocks base method.
func (m *MockEngine) Continue() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Continue")
}

// Continue indicates an expected call of Continue.
func (mr *MockEngineMockRecorder) Continue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Continue", reflect.TypeOf((*MockEngine)(nil).Continue))
}

// CurrentTime mocks base method.
func (m *MockEngine) CurrentTime() sim.VTimeInSec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentTime")
	ret0, _ := ret[0].(sim.VTimeInSec)
	return ret0
}

// CurrentTime indicates an expected call of CurrentTime.
func (mr *MockEngineMockRecorder) CurrentTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentTime", reflect.TypeOf((*MockEngine)(nil).CurrentTime))
}

// Finished mocks base method.
func (m *MockEngine) Finished() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Finished")
}

// Finished indicates an expected call of Finished.
func (mr *MockEngineMockRecorder) Finished() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finished", reflect.TypeOf((*MockEngine)(nil).Finished))
}

// Hooks mocks base method.
func (m *MockEngine) Hooks() []sim.Hook {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hooks")
	ret0, _ := ret[0].([]sim.Hook)
	return ret0
}

// Hooks indicates an expected call of Hooks.
func (mr *MockEngineMockRecorder) Hooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hooks", reflect.TypeOf((*MockEngine)(nil).Hooks))
}

// NumHooks mocks base method.
func (m *MockEngine) NumHooks() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NumHooks")
	ret0, _ := ret[0].(int)
	return ret0
}

// NumHooks indicates an expected call of NumHooks.
func (mr *MockEngineMockRecorder) NumHooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumHooks", reflect.TypeOf((*MockEngine)(nil).NumHooks))
}

// Pause mocks base method.
func (m *MockEngine) Pause() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Pause")
}

// Pause indicates an expected call of Pause.
func (mr *MockEngineMockRecorder) Pause() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pause", reflect.TypeOf((*MockEngine)(nil).Pause))
}

// RegisterSimulationEndHandler mocks base method.
func (m *MockEngine) RegisterSimulationEndHandler(arg0 sim.SimulationEndHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterSimulationEndHandler", arg0)
}

// RegisterSimulationEndHandler indicates an expected call of RegisterSimulationEndHandler.
func (mr *MockEngineMockRecorder) RegisterSimulationEndHandler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterSimulationEndHandler", reflect.TypeOf((*MockEngine)(nil).RegisterSimulationEndHandler), arg0)
}

// Run mocks base method.
func (m *MockEngine) Run() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run")
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockEngineMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockEngine)(nil).Run))
}

// Schedule mocks base method.
func (m *MockEngine) Schedule(arg0 sim.Event) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Schedule", arg0)
}

// Schedule indicates an expected call of Schedule.
func (mr *MockEngineMockRecorder) Schedule(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Schedule", reflect.TypeOf((*MockEngine)(nil).Schedule), arg0)
}
