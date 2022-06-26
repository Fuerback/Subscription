// Code generated by MockGen. DO NOT EDIT.
// Source: ./core/domain/product.go

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	domain "github.com/Fuerback/subscription/core/domain"
	dto "github.com/Fuerback/subscription/core/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockProductService is a mock of ProductService interface.
type MockProductService struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceMockRecorder
}

// MockProductServiceMockRecorder is the mock recorder for MockProductService.
type MockProductServiceMockRecorder struct {
	mock *MockProductService
}

// NewMockProductService creates a new mock instance.
func NewMockProductService(ctrl *gomock.Controller) *MockProductService {
	mock := &MockProductService{ctrl: ctrl}
	mock.recorder = &MockProductServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductService) EXPECT() *MockProductServiceMockRecorder {
	return m.recorder
}

// Fetch mocks base method.
func (m *MockProductService) Fetch(response http.ResponseWriter, request *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Fetch", response, request)
}

// Fetch indicates an expected call of Fetch.
func (mr *MockProductServiceMockRecorder) Fetch(response, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockProductService)(nil).Fetch), response, request)
}

// FetchOne mocks base method.
func (m *MockProductService) FetchOne(response http.ResponseWriter, request *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FetchOne", response, request)
}

// FetchOne indicates an expected call of FetchOne.
func (mr *MockProductServiceMockRecorder) FetchOne(response, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchOne", reflect.TypeOf((*MockProductService)(nil).FetchOne), response, request)
}

// MockProductUseCase is a mock of ProductUseCase interface.
type MockProductUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockProductUseCaseMockRecorder
}

// MockProductUseCaseMockRecorder is the mock recorder for MockProductUseCase.
type MockProductUseCaseMockRecorder struct {
	mock *MockProductUseCase
}

// NewMockProductUseCase creates a new mock instance.
func NewMockProductUseCase(ctrl *gomock.Controller) *MockProductUseCase {
	mock := &MockProductUseCase{ctrl: ctrl}
	mock.recorder = &MockProductUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductUseCase) EXPECT() *MockProductUseCaseMockRecorder {
	return m.recorder
}

// Fetch mocks base method.
func (m *MockProductUseCase) Fetch(paginationRequest *dto.PaginationRequestParms) (*domain.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", paginationRequest)
	ret0, _ := ret[0].(*domain.Pagination)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockProductUseCaseMockRecorder) Fetch(paginationRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockProductUseCase)(nil).Fetch), paginationRequest)
}

// FetchOne mocks base method.
func (m *MockProductUseCase) FetchOne(id string) (*domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchOne", id)
	ret0, _ := ret[0].(*domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchOne indicates an expected call of FetchOne.
func (mr *MockProductUseCaseMockRecorder) FetchOne(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchOne", reflect.TypeOf((*MockProductUseCase)(nil).FetchOne), id)
}

// MockProductRepository is a mock of ProductRepository interface.
type MockProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepositoryMockRecorder
}

// MockProductRepositoryMockRecorder is the mock recorder for MockProductRepository.
type MockProductRepositoryMockRecorder struct {
	mock *MockProductRepository
}

// NewMockProductRepository creates a new mock instance.
func NewMockProductRepository(ctrl *gomock.Controller) *MockProductRepository {
	mock := &MockProductRepository{ctrl: ctrl}
	mock.recorder = &MockProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepository) EXPECT() *MockProductRepositoryMockRecorder {
	return m.recorder
}

// Fetch mocks base method.
func (m *MockProductRepository) Fetch(paginationRequest *dto.PaginationRequestParms) (*domain.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", paginationRequest)
	ret0, _ := ret[0].(*domain.Pagination)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockProductRepositoryMockRecorder) Fetch(paginationRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockProductRepository)(nil).Fetch), paginationRequest)
}

// FetchOne mocks base method.
func (m *MockProductRepository) FetchOne(id string) (*domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchOne", id)
	ret0, _ := ret[0].(*domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchOne indicates an expected call of FetchOne.
func (mr *MockProductRepositoryMockRecorder) FetchOne(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchOne", reflect.TypeOf((*MockProductRepository)(nil).FetchOne), id)
}
