// Code generated by MockGen. DO NOT EDIT.
// Source: app/interfaces/api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	usecases "github.com/itiroinazawa/golangapi/app/usecases"
	reflect "reflect"
)

// MockInteractor is a mock of Interactor interface.
type MockInteractor struct {
	ctrl     *gomock.Controller
	recorder *MockInteractorMockRecorder
}

// MockInteractorMockRecorder is the mock recorder for MockInteractor.
type MockInteractorMockRecorder struct {
	mock *MockInteractor
}

// NewMockInteractor creates a new mock instance.
func NewMockInteractor(ctrl *gomock.Controller) *MockInteractor {
	mock := &MockInteractor{ctrl: ctrl}
	mock.recorder = &MockInteractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInteractor) EXPECT() *MockInteractorMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockInteractor) CreateAccount(documentNumber string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", documentNumber)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockInteractorMockRecorder) CreateAccount(documentNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockInteractor)(nil).CreateAccount), documentNumber)
}

// CreateTransaction mocks base method.
func (m *MockInteractor) CreateTransaction(accountID, operationTypeID int, amount float32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", accountID, operationTypeID, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockInteractorMockRecorder) CreateTransaction(accountID, operationTypeID, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockInteractor)(nil).CreateTransaction), accountID, operationTypeID, amount)
}

// GetAccount mocks base method.
func (m *MockInteractor) GetAccount(id int) (usecases.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", id)
	ret0, _ := ret[0].(usecases.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockInteractorMockRecorder) GetAccount(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockInteractor)(nil).GetAccount), id)
}

// GetTransactionsByAccountID mocks base method.
func (m *MockInteractor) GetTransactionsByAccountID(accountID int) ([]usecases.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionsByAccountID", accountID)
	ret0, _ := ret[0].([]usecases.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionsByAccountID indicates an expected call of GetTransactionsByAccountID.
func (mr *MockInteractorMockRecorder) GetTransactionsByAccountID(accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionsByAccountID", reflect.TypeOf((*MockInteractor)(nil).GetTransactionsByAccountID), accountID)
}
