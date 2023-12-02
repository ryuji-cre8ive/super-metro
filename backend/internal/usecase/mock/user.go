// Code generated by MockGen. DO NOT EDIT.
// Source: user.go
//
// Generated by this command:
//
//	mockgen -source=user.go -destination=./mock/user.go
//
// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	reflect "reflect"

	v4 "github.com/labstack/echo/v4"
	domain "github.com/ryuji-cre8ive/super-metro/internal/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockUserUsecase is a mock of UserUsecase interface.
type MockUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUsecaseMockRecorder
}

// MockUserUsecaseMockRecorder is the mock recorder for MockUserUsecase.
type MockUserUsecaseMockRecorder struct {
	mock *MockUserUsecase
}

// NewMockUserUsecase creates a new mock instance.
func NewMockUserUsecase(ctrl *gomock.Controller) *MockUserUsecase {
	mock := &MockUserUsecase{ctrl: ctrl}
	mock.recorder = &MockUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUsecase) EXPECT() *MockUserUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserUsecase) Create(ctx v4.Context, email, userName, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, email, userName, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserUsecaseMockRecorder) Create(ctx, email, userName, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserUsecase)(nil).Create), ctx, email, userName, password)
}

// FindByEmail mocks base method.
func (m *MockUserUsecase) FindByEmail(ctx v4.Context, email string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", ctx, email)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockUserUsecaseMockRecorder) FindByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserUsecase)(nil).FindByEmail), ctx, email)
}

// TopUp mocks base method.
func (m *MockUserUsecase) TopUp(ctx v4.Context, id string, amount int) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TopUp", ctx, id, amount)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TopUp indicates an expected call of TopUp.
func (mr *MockUserUsecaseMockRecorder) TopUp(ctx, id, amount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TopUp", reflect.TypeOf((*MockUserUsecase)(nil).TopUp), ctx, id, amount)
}
