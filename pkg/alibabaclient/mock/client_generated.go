// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	gomock "github.com/golang/mock/gomock"
)

// MockAPI is a mock of API interface.
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI.
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance.
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// DescribeAvailableZoneByInstanceType mocks base method.
func (m *MockAPI) DescribeAvailableZoneByInstanceType(arg0 string) (*ecs.DescribeAvailableResourceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeAvailableZoneByInstanceType", arg0)
	ret0, _ := ret[0].(*ecs.DescribeAvailableResourceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAvailableZoneByInstanceType indicates an expected call of DescribeAvailableZoneByInstanceType.
func (mr *MockAPIMockRecorder) DescribeAvailableZoneByInstanceType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAvailableZoneByInstanceType", reflect.TypeOf((*MockAPI)(nil).DescribeAvailableZoneByInstanceType), arg0)
}

// GetAvailableZonesByInstanceType mocks base method.
func (m *MockAPI) GetAvailableZonesByInstanceType(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableZonesByInstanceType", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableZonesByInstanceType indicates an expected call of GetAvailableZonesByInstanceType.
func (mr *MockAPIMockRecorder) GetAvailableZonesByInstanceType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableZonesByInstanceType", reflect.TypeOf((*MockAPI)(nil).GetAvailableZonesByInstanceType), arg0)
}