// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package products is a generated GoMock package.
package products

import (
	models "practice-app/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	krogo "github.com/krogertechnology/krogo/pkg/krogo"
)

// MockProductStore is a mock of ProductStore interface.
type MockProductStore struct {
	ctrl     *gomock.Controller
	recorder *MockProductStoreMockRecorder
}

// MockProductStoreMockRecorder is the mock recorder for MockProductStore.
type MockProductStoreMockRecorder struct {
	mock *MockProductStore
}

// NewMockProductStore creates a new mock instance.
func NewMockProductStore(ctrl *gomock.Controller) *MockProductStore {
	mock := &MockProductStore{ctrl: ctrl}
	mock.recorder = &MockProductStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductStore) EXPECT() *MockProductStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductStore) Create(ctx *krogo.Context, product *models.Product) (*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, product)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductStoreMockRecorder) Create(ctx, product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductStore)(nil).Create), ctx, product)
}

// GetAll mocks base method.
func (m *MockProductStore) GetAll(ctx *krogo.Context, params map[string]string) ([]models.ProductWithVariants, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, params)
	ret0, _ := ret[0].([]models.ProductWithVariants)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockProductStoreMockRecorder) GetAll(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockProductStore)(nil).GetAll), ctx, params)
}

// GetByID mocks base method.
func (m *MockProductStore) GetByID(ctx *krogo.Context, id string) (*models.ProductWithVariants, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*models.ProductWithVariants)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockProductStoreMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockProductStore)(nil).GetByID), ctx, id)
}