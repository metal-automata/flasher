// Code generated by MockGen. DO NOT EDIT.
// Source: internal/model/model.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	context "context"
	os "os"
	reflect "reflect"

	common "github.com/bmc-toolbox/common"
	gomock "github.com/golang/mock/gomock"
	model "github.com/metal-toolbox/flasher/internal/model"
)

// MockDeviceQueryor is a mock of DeviceQueryor interface.
type MockDeviceQueryor struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceQueryorMockRecorder
}

// MockDeviceQueryorMockRecorder is the mock recorder for MockDeviceQueryor.
type MockDeviceQueryorMockRecorder struct {
	mock *MockDeviceQueryor
}

// NewMockDeviceQueryor creates a new mock instance.
func NewMockDeviceQueryor(ctrl *gomock.Controller) *MockDeviceQueryor {
	mock := &MockDeviceQueryor{ctrl: ctrl}
	mock.recorder = &MockDeviceQueryorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeviceQueryor) EXPECT() *MockDeviceQueryorMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockDeviceQueryor) Close(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDeviceQueryorMockRecorder) Close(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDeviceQueryor)(nil).Close), ctx)
}

// FirmwareInstall mocks base method.
func (m *MockDeviceQueryor) FirmwareInstall(ctx context.Context, componentSlug string, force bool, file *os.File) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FirmwareInstall", ctx, componentSlug, force, file)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FirmwareInstall indicates an expected call of FirmwareInstall.
func (mr *MockDeviceQueryorMockRecorder) FirmwareInstall(ctx, componentSlug, force, file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirmwareInstall", reflect.TypeOf((*MockDeviceQueryor)(nil).FirmwareInstall), ctx, componentSlug, force, file)
}

// FirmwareInstallStatus mocks base method.
func (m *MockDeviceQueryor) FirmwareInstallStatus(ctx context.Context, installVersion, componentSlug, bmcTaskID string) (model.ComponentFirmwareInstallStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FirmwareInstallStatus", ctx, installVersion, componentSlug, bmcTaskID)
	ret0, _ := ret[0].(model.ComponentFirmwareInstallStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FirmwareInstallStatus indicates an expected call of FirmwareInstallStatus.
func (mr *MockDeviceQueryorMockRecorder) FirmwareInstallStatus(ctx, installVersion, componentSlug, bmcTaskID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirmwareInstallStatus", reflect.TypeOf((*MockDeviceQueryor)(nil).FirmwareInstallStatus), ctx, installVersion, componentSlug, bmcTaskID)
}

// Inventory mocks base method.
func (m *MockDeviceQueryor) Inventory(ctx context.Context) (*common.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Inventory", ctx)
	ret0, _ := ret[0].(*common.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Inventory indicates an expected call of Inventory.
func (mr *MockDeviceQueryorMockRecorder) Inventory(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inventory", reflect.TypeOf((*MockDeviceQueryor)(nil).Inventory), ctx)
}

// Open mocks base method.
func (m *MockDeviceQueryor) Open(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open.
func (mr *MockDeviceQueryorMockRecorder) Open(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockDeviceQueryor)(nil).Open), ctx)
}

// PowerStatus mocks base method.
func (m *MockDeviceQueryor) PowerStatus(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerStatus", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PowerStatus indicates an expected call of PowerStatus.
func (mr *MockDeviceQueryorMockRecorder) PowerStatus(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerStatus", reflect.TypeOf((*MockDeviceQueryor)(nil).PowerStatus), ctx)
}

// ResetBMC mocks base method.
func (m *MockDeviceQueryor) ResetBMC(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetBMC", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetBMC indicates an expected call of ResetBMC.
func (mr *MockDeviceQueryorMockRecorder) ResetBMC(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetBMC", reflect.TypeOf((*MockDeviceQueryor)(nil).ResetBMC), ctx)
}

// SetPowerState mocks base method.
func (m *MockDeviceQueryor) SetPowerState(ctx context.Context, state string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPowerState", ctx, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPowerState indicates an expected call of SetPowerState.
func (mr *MockDeviceQueryorMockRecorder) SetPowerState(ctx, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPowerState", reflect.TypeOf((*MockDeviceQueryor)(nil).SetPowerState), ctx, state)
}