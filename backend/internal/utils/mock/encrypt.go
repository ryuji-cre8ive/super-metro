// Code generated by MockGen. DO NOT EDIT.
// Source: encrypt.go
//
// Generated by this command:
//
//	mockgen -source=encrypt.go -destination=./mock/encrypt.go
//
// Package mock_utils is a generated GoMock package.
package mock_utils

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockEncrypt is a mock of Encrypt interface.
type MockEncrypt struct {
	ctrl     *gomock.Controller
	recorder *MockEncryptMockRecorder
}

// MockEncryptMockRecorder is the mock recorder for MockEncrypt.
type MockEncryptMockRecorder struct {
	mock *MockEncrypt
}

// NewMockEncrypt creates a new mock instance.
func NewMockEncrypt(ctrl *gomock.Controller) *MockEncrypt {
	mock := &MockEncrypt{ctrl: ctrl}
	mock.recorder = &MockEncryptMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncrypt) EXPECT() *MockEncryptMockRecorder {
	return m.recorder
}

// CheckHashPassword mocks base method.
func (m *MockEncrypt) CheckHashPassword(hashPassword, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckHashPassword", hashPassword, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckHashPassword indicates an expected call of CheckHashPassword.
func (mr *MockEncryptMockRecorder) CheckHashPassword(hashPassword, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckHashPassword", reflect.TypeOf((*MockEncrypt)(nil).CheckHashPassword), hashPassword, password)
}

// Decrypt mocks base method.
func (m *MockEncrypt) Decrypt(encryptedData string, key []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decrypt", encryptedData, key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decrypt indicates an expected call of Decrypt.
func (mr *MockEncryptMockRecorder) Decrypt(encryptedData, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decrypt", reflect.TypeOf((*MockEncrypt)(nil).Decrypt), encryptedData, key)
}

// Encrypt mocks base method.
func (m *MockEncrypt) Encrypt(plainText, key []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encrypt", plainText, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encrypt indicates an expected call of Encrypt.
func (mr *MockEncryptMockRecorder) Encrypt(plainText, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encrypt", reflect.TypeOf((*MockEncrypt)(nil).Encrypt), plainText, key)
}

// PasswordEncryptPasswordEncrypt mocks base method.
func (m *MockEncrypt) PasswordEncryptPasswordEncrypt(password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PasswordEncryptPasswordEncrypt", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PasswordEncryptPasswordEncrypt indicates an expected call of PasswordEncryptPasswordEncrypt.
func (mr *MockEncryptMockRecorder) PasswordEncryptPasswordEncrypt(password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PasswordEncryptPasswordEncrypt", reflect.TypeOf((*MockEncrypt)(nil).PasswordEncryptPasswordEncrypt), password)
}
