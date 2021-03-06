// Code generated by MockGen. DO NOT EDIT.
// Source: ./load_balancer.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/outscale-dev/cluster-api-provider-outscale.git/api/v1beta1"
	osc "github.com/outscale/osc-sdk-go/v2"
)

// MockOscLoadBalancerInterface is a mock of OscLoadBalancerInterface interface.
type MockOscLoadBalancerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockOscLoadBalancerInterfaceMockRecorder
}

// MockOscLoadBalancerInterfaceMockRecorder is the mock recorder for MockOscLoadBalancerInterface.
type MockOscLoadBalancerInterfaceMockRecorder struct {
	mock *MockOscLoadBalancerInterface
}

// NewMockOscLoadBalancerInterface creates a new mock instance.
func NewMockOscLoadBalancerInterface(ctrl *gomock.Controller) *MockOscLoadBalancerInterface {
	mock := &MockOscLoadBalancerInterface{ctrl: ctrl}
	mock.recorder = &MockOscLoadBalancerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOscLoadBalancerInterface) EXPECT() *MockOscLoadBalancerInterfaceMockRecorder {
	return m.recorder
}

// ConfigureHealthCheck mocks base method.
func (m *MockOscLoadBalancerInterface) ConfigureHealthCheck(spec *v1beta1.OscLoadBalancer) (*osc.LoadBalancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigureHealthCheck", spec)
	ret0, _ := ret[0].(*osc.LoadBalancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigureHealthCheck indicates an expected call of ConfigureHealthCheck.
func (mr *MockOscLoadBalancerInterfaceMockRecorder) ConfigureHealthCheck(spec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureHealthCheck", reflect.TypeOf((*MockOscLoadBalancerInterface)(nil).ConfigureHealthCheck), spec)
}

// CreateLoadBalancer mocks base method.
func (m *MockOscLoadBalancerInterface) CreateLoadBalancer(spec *v1beta1.OscLoadBalancer, subnetId, securityGroupId string) (*osc.LoadBalancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLoadBalancer", spec, subnetId, securityGroupId)
	ret0, _ := ret[0].(*osc.LoadBalancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLoadBalancer indicates an expected call of CreateLoadBalancer.
func (mr *MockOscLoadBalancerInterfaceMockRecorder) CreateLoadBalancer(spec, subnetId, securityGroupId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLoadBalancer", reflect.TypeOf((*MockOscLoadBalancerInterface)(nil).CreateLoadBalancer), spec, subnetId, securityGroupId)
}

// DeleteLoadBalancer mocks base method.
func (m *MockOscLoadBalancerInterface) DeleteLoadBalancer(spec *v1beta1.OscLoadBalancer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLoadBalancer", spec)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLoadBalancer indicates an expected call of DeleteLoadBalancer.
func (mr *MockOscLoadBalancerInterfaceMockRecorder) DeleteLoadBalancer(spec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLoadBalancer", reflect.TypeOf((*MockOscLoadBalancerInterface)(nil).DeleteLoadBalancer), spec)
}

// GetLoadBalancer mocks base method.
func (m *MockOscLoadBalancerInterface) GetLoadBalancer(spec *v1beta1.OscLoadBalancer) (*osc.LoadBalancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoadBalancer", spec)
	ret0, _ := ret[0].(*osc.LoadBalancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoadBalancer indicates an expected call of GetLoadBalancer.
func (mr *MockOscLoadBalancerInterfaceMockRecorder) GetLoadBalancer(spec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoadBalancer", reflect.TypeOf((*MockOscLoadBalancerInterface)(nil).GetLoadBalancer), spec)
}
