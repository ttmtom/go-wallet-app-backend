// Code generated by MockGen. DO NOT EDIT.
// Source: user.go
//
// Generated by this command:
//
//	mockgen -source=user.go -destination=../../../test/mock/user.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	model "go-wallet-system/wallet_system/core/model"
	types "go-wallet-system/wallet_system/core/types"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
	isgomock struct{}
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRepository) Create(user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserRepositoryMockRecorder) Create(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), user)
}

// FindByID mocks base method.
func (m *MockUserRepository) FindByID(id string) *model.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*model.User)
	return ret0
}

// FindByID indicates an expected call of FindByID.
func (mr *MockUserRepositoryMockRecorder) FindByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUserRepository)(nil).FindByID), id)
}

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
	isgomock struct{}
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// UserInfo mocks base method.
func (m *MockUserService) UserInfo(name string) (*types.UserInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserInfoResponse", name)
	ret0, _ := ret[0].(*types.UserInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserInfo indicates an expected call of UserInfo.
func (mr *MockUserServiceMockRecorder) UserInfo(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserInfoResponse", reflect.TypeOf((*MockUserService)(nil).UserInfo), name)
}

// UserRegistration mocks base method.
func (m *MockUserService) UserRegistration(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserRegistration", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// UserRegistration indicates an expected call of UserRegistration.
func (mr *MockUserServiceMockRecorder) UserRegistration(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserRegistration", reflect.TypeOf((*MockUserService)(nil).UserRegistration), name)
}

// MockUserController is a mock of UserController interface.
type MockUserController struct {
	ctrl     *gomock.Controller
	recorder *MockUserControllerMockRecorder
	isgomock struct{}
}

// MockUserControllerMockRecorder is the mock recorder for MockUserController.
type MockUserControllerMockRecorder struct {
	mock *MockUserController
}

// NewMockUserController creates a new mock instance.
func NewMockUserController(ctrl *gomock.Controller) *MockUserController {
	mock := &MockUserController{ctrl: ctrl}
	mock.recorder = &MockUserControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserController) EXPECT() *MockUserControllerMockRecorder {
	return m.recorder
}

// GetUserInfo mocks base method.
func (m *MockUserController) GetUserInfo(name string) (*types.UserInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfo", name)
	ret0, _ := ret[0].(*types.UserInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfo indicates an expected call of GetUserInfo.
func (mr *MockUserControllerMockRecorder) GetUserInfo(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfo", reflect.TypeOf((*MockUserController)(nil).GetUserInfo), name)
}

// UserRegister mocks base method.
func (m *MockUserController) UserRegister(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserRegister", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// UserRegister indicates an expected call of UserRegister.
func (mr *MockUserControllerMockRecorder) UserRegister(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserRegister", reflect.TypeOf((*MockUserController)(nil).UserRegister), name)
}
