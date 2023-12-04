// Code generated by MockGen. DO NOT EDIT.
// Source: companyHandler.go
//
// Generated by this command:
//
//	mockgen -source=companyHandler.go -destination=.mock/companyHandler_mock.go -package=handler
//
// Package handler is a generated GoMock package.
package handler

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "go.uber.org/mock/gomock"
)

// MockCompanyHandler is a mock of CompanyHandler interface.
type MockCompanyHandler struct {
	ctrl     *gomock.Controller
	recorder *MockCompanyHandlerMockRecorder
}

// MockCompanyHandlerMockRecorder is the mock recorder for MockCompanyHandler.
type MockCompanyHandlerMockRecorder struct {
	mock *MockCompanyHandler
}

// NewMockCompanyHandler creates a new mock instance.
func NewMockCompanyHandler(ctrl *gomock.Controller) *MockCompanyHandler {
	mock := &MockCompanyHandler{ctrl: ctrl}
	mock.recorder = &MockCompanyHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompanyHandler) EXPECT() *MockCompanyHandlerMockRecorder {
	return m.recorder
}

// AddCompany mocks base method.
func (m *MockCompanyHandler) AddCompany(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddCompany", c)
}

// AddCompany indicates an expected call of AddCompany.
func (mr *MockCompanyHandlerMockRecorder) AddCompany(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCompany", reflect.TypeOf((*MockCompanyHandler)(nil).AddCompany), c)
}

// ViewAllComapny mocks base method.
func (m *MockCompanyHandler) ViewAllComapny(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ViewAllComapny", c)
}

// ViewAllComapny indicates an expected call of ViewAllComapny.
func (mr *MockCompanyHandlerMockRecorder) ViewAllComapny(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewAllComapny", reflect.TypeOf((*MockCompanyHandler)(nil).ViewAllComapny), c)
}

// ViewCompanyByID mocks base method.
func (m *MockCompanyHandler) ViewCompanyByID(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ViewCompanyByID", c)
}

// ViewCompanyByID indicates an expected call of ViewCompanyByID.
func (mr *MockCompanyHandlerMockRecorder) ViewCompanyByID(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewCompanyByID", reflect.TypeOf((*MockCompanyHandler)(nil).ViewCompanyByID), c)
}