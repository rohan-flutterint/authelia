// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/authelia/authelia/v4/internal/storage (interfaces: Provider)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	sql "database/sql"
	reflect "reflect"
	time "time"

	model "github.com/authelia/authelia/v4/internal/model"
	storage "github.com/authelia/authelia/v4/internal/storage"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockStorage is a mock of Provider interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// AppendAuthenticationLog mocks base method.
func (m *MockStorage) AppendAuthenticationLog(arg0 context.Context, arg1 model.AuthenticationAttempt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendAuthenticationLog", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendAuthenticationLog indicates an expected call of AppendAuthenticationLog.
func (mr *MockStorageMockRecorder) AppendAuthenticationLog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendAuthenticationLog", reflect.TypeOf((*MockStorage)(nil).AppendAuthenticationLog), arg0, arg1)
}

// BeginTX mocks base method.
func (m *MockStorage) BeginTX(arg0 context.Context) (context.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTX", arg0)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTX indicates an expected call of BeginTX.
func (mr *MockStorageMockRecorder) BeginTX(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTX", reflect.TypeOf((*MockStorage)(nil).BeginTX), arg0)
}

// Close mocks base method.
func (m *MockStorage) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStorageMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStorage)(nil).Close))
}

// Commit mocks base method.
func (m *MockStorage) Commit(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockStorageMockRecorder) Commit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockStorage)(nil).Commit), arg0)
}

// ConsumeIdentityVerification mocks base method.
func (m *MockStorage) ConsumeIdentityVerification(arg0 context.Context, arg1 string, arg2 model.NullIP) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsumeIdentityVerification", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConsumeIdentityVerification indicates an expected call of ConsumeIdentityVerification.
func (mr *MockStorageMockRecorder) ConsumeIdentityVerification(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumeIdentityVerification", reflect.TypeOf((*MockStorage)(nil).ConsumeIdentityVerification), arg0, arg1, arg2)
}

// DeactivateOAuth2Session mocks base method.
func (m *MockStorage) DeactivateOAuth2Session(arg0 context.Context, arg1 storage.OAuth2SessionType, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeactivateOAuth2Session", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeactivateOAuth2Session indicates an expected call of DeactivateOAuth2Session.
func (mr *MockStorageMockRecorder) DeactivateOAuth2Session(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeactivateOAuth2Session", reflect.TypeOf((*MockStorage)(nil).DeactivateOAuth2Session), arg0, arg1, arg2)
}

// DeactivateOAuth2SessionByRequestID mocks base method.
func (m *MockStorage) DeactivateOAuth2SessionByRequestID(arg0 context.Context, arg1 storage.OAuth2SessionType, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeactivateOAuth2SessionByRequestID", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeactivateOAuth2SessionByRequestID indicates an expected call of DeactivateOAuth2SessionByRequestID.
func (mr *MockStorageMockRecorder) DeactivateOAuth2SessionByRequestID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeactivateOAuth2SessionByRequestID", reflect.TypeOf((*MockStorage)(nil).DeactivateOAuth2SessionByRequestID), arg0, arg1, arg2)
}

// DeletePreferredDuoDevice mocks base method.
func (m *MockStorage) DeletePreferredDuoDevice(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePreferredDuoDevice", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePreferredDuoDevice indicates an expected call of DeletePreferredDuoDevice.
func (mr *MockStorageMockRecorder) DeletePreferredDuoDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePreferredDuoDevice", reflect.TypeOf((*MockStorage)(nil).DeletePreferredDuoDevice), arg0, arg1)
}

// DeleteTOTPConfiguration mocks base method.
func (m *MockStorage) DeleteTOTPConfiguration(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTOTPConfiguration", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTOTPConfiguration indicates an expected call of DeleteTOTPConfiguration.
func (mr *MockStorageMockRecorder) DeleteTOTPConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTOTPConfiguration", reflect.TypeOf((*MockStorage)(nil).DeleteTOTPConfiguration), arg0, arg1)
}

// DeleteWebauthnDevice mocks base method.
func (m *MockStorage) DeleteWebauthnDevice(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWebauthnDevice", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWebauthnDevice indicates an expected call of DeleteWebauthnDevice.
func (mr *MockStorageMockRecorder) DeleteWebauthnDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWebauthnDevice", reflect.TypeOf((*MockStorage)(nil).DeleteWebauthnDevice), arg0, arg1)
}

// DeleteWebauthnDeviceByUsername mocks base method.
func (m *MockStorage) DeleteWebauthnDeviceByUsername(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWebauthnDeviceByUsername", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWebauthnDeviceByUsername indicates an expected call of DeleteWebauthnDeviceByUsername.
func (mr *MockStorageMockRecorder) DeleteWebauthnDeviceByUsername(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWebauthnDeviceByUsername", reflect.TypeOf((*MockStorage)(nil).DeleteWebauthnDeviceByUsername), arg0, arg1, arg2)
}

// FindIdentityVerification mocks base method.
func (m *MockStorage) FindIdentityVerification(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindIdentityVerification", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindIdentityVerification indicates an expected call of FindIdentityVerification.
func (mr *MockStorageMockRecorder) FindIdentityVerification(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindIdentityVerification", reflect.TypeOf((*MockStorage)(nil).FindIdentityVerification), arg0, arg1)
}

// LoadAuthenticationLogs mocks base method.
func (m *MockStorage) LoadAuthenticationLogs(arg0 context.Context, arg1 string, arg2 time.Time, arg3, arg4 int) ([]model.AuthenticationAttempt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadAuthenticationLogs", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]model.AuthenticationAttempt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadAuthenticationLogs indicates an expected call of LoadAuthenticationLogs.
func (mr *MockStorageMockRecorder) LoadAuthenticationLogs(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadAuthenticationLogs", reflect.TypeOf((*MockStorage)(nil).LoadAuthenticationLogs), arg0, arg1, arg2, arg3, arg4)
}

// LoadOAuth2BlacklistedJTI mocks base method.
func (m *MockStorage) LoadOAuth2BlacklistedJTI(arg0 context.Context, arg1 string) (*model.OAuth2BlacklistedJTI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadOAuth2BlacklistedJTI", arg0, arg1)
	ret0, _ := ret[0].(*model.OAuth2BlacklistedJTI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadOAuth2BlacklistedJTI indicates an expected call of LoadOAuth2BlacklistedJTI.
func (mr *MockStorageMockRecorder) LoadOAuth2BlacklistedJTI(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadOAuth2BlacklistedJTI", reflect.TypeOf((*MockStorage)(nil).LoadOAuth2BlacklistedJTI), arg0, arg1)
}

// LoadOAuth2ConsentPreConfigurations mocks base method.
func (m *MockStorage) LoadOAuth2ConsentPreConfigurations(arg0 context.Context, arg1 string, arg2 uuid.UUID) (*storage.ConsentPreConfigRows, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadOAuth2ConsentPreConfigurations", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storage.ConsentPreConfigRows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadOAuth2ConsentPreConfigurations indicates an expected call of LoadOAuth2ConsentPreConfigurations.
func (mr *MockStorageMockRecorder) LoadOAuth2ConsentPreConfigurations(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadOAuth2ConsentPreConfigurations", reflect.TypeOf((*MockStorage)(nil).LoadOAuth2ConsentPreConfigurations), arg0, arg1, arg2)
}

// LoadOAuth2ConsentSessionByChallengeID mocks base method.
func (m *MockStorage) LoadOAuth2ConsentSessionByChallengeID(arg0 context.Context, arg1 uuid.UUID) (*model.OAuth2ConsentSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadOAuth2ConsentSessionByChallengeID", arg0, arg1)
	ret0, _ := ret[0].(*model.OAuth2ConsentSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadOAuth2ConsentSessionByChallengeID indicates an expected call of LoadOAuth2ConsentSessionByChallengeID.
func (mr *MockStorageMockRecorder) LoadOAuth2ConsentSessionByChallengeID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadOAuth2ConsentSessionByChallengeID", reflect.TypeOf((*MockStorage)(nil).LoadOAuth2ConsentSessionByChallengeID), arg0, arg1)
}

// LoadOAuth2Session mocks base method.
func (m *MockStorage) LoadOAuth2Session(arg0 context.Context, arg1 storage.OAuth2SessionType, arg2 string) (*model.OAuth2Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadOAuth2Session", arg0, arg1, arg2)
	ret0, _ := ret[0].(*model.OAuth2Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadOAuth2Session indicates an expected call of LoadOAuth2Session.
func (mr *MockStorageMockRecorder) LoadOAuth2Session(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadOAuth2Session", reflect.TypeOf((*MockStorage)(nil).LoadOAuth2Session), arg0, arg1, arg2)
}

// LoadPreferred2FAMethod mocks base method.
func (m *MockStorage) LoadPreferred2FAMethod(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadPreferred2FAMethod", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadPreferred2FAMethod indicates an expected call of LoadPreferred2FAMethod.
func (mr *MockStorageMockRecorder) LoadPreferred2FAMethod(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadPreferred2FAMethod", reflect.TypeOf((*MockStorage)(nil).LoadPreferred2FAMethod), arg0, arg1)
}

// LoadPreferredDuoDevice mocks base method.
func (m *MockStorage) LoadPreferredDuoDevice(arg0 context.Context, arg1 string) (*model.DuoDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadPreferredDuoDevice", arg0, arg1)
	ret0, _ := ret[0].(*model.DuoDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadPreferredDuoDevice indicates an expected call of LoadPreferredDuoDevice.
func (mr *MockStorageMockRecorder) LoadPreferredDuoDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadPreferredDuoDevice", reflect.TypeOf((*MockStorage)(nil).LoadPreferredDuoDevice), arg0, arg1)
}

// LoadTOTPConfiguration mocks base method.
func (m *MockStorage) LoadTOTPConfiguration(arg0 context.Context, arg1 string) (*model.TOTPConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadTOTPConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*model.TOTPConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadTOTPConfiguration indicates an expected call of LoadTOTPConfiguration.
func (mr *MockStorageMockRecorder) LoadTOTPConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadTOTPConfiguration", reflect.TypeOf((*MockStorage)(nil).LoadTOTPConfiguration), arg0, arg1)
}

// LoadTOTPConfigurations mocks base method.
func (m *MockStorage) LoadTOTPConfigurations(arg0 context.Context, arg1, arg2 int) ([]model.TOTPConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadTOTPConfigurations", arg0, arg1, arg2)
	ret0, _ := ret[0].([]model.TOTPConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadTOTPConfigurations indicates an expected call of LoadTOTPConfigurations.
func (mr *MockStorageMockRecorder) LoadTOTPConfigurations(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadTOTPConfigurations", reflect.TypeOf((*MockStorage)(nil).LoadTOTPConfigurations), arg0, arg1, arg2)
}

// LoadUserInfo mocks base method.
func (m *MockStorage) LoadUserInfo(arg0 context.Context, arg1 string) (model.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadUserInfo", arg0, arg1)
	ret0, _ := ret[0].(model.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadUserInfo indicates an expected call of LoadUserInfo.
func (mr *MockStorageMockRecorder) LoadUserInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadUserInfo", reflect.TypeOf((*MockStorage)(nil).LoadUserInfo), arg0, arg1)
}

// LoadUserOpaqueIdentifier mocks base method.
func (m *MockStorage) LoadUserOpaqueIdentifier(arg0 context.Context, arg1 uuid.UUID) (*model.UserOpaqueIdentifier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadUserOpaqueIdentifier", arg0, arg1)
	ret0, _ := ret[0].(*model.UserOpaqueIdentifier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadUserOpaqueIdentifier indicates an expected call of LoadUserOpaqueIdentifier.
func (mr *MockStorageMockRecorder) LoadUserOpaqueIdentifier(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadUserOpaqueIdentifier", reflect.TypeOf((*MockStorage)(nil).LoadUserOpaqueIdentifier), arg0, arg1)
}

// LoadUserOpaqueIdentifierBySignature mocks base method.
func (m *MockStorage) LoadUserOpaqueIdentifierBySignature(arg0 context.Context, arg1, arg2, arg3 string) (*model.UserOpaqueIdentifier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadUserOpaqueIdentifierBySignature", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*model.UserOpaqueIdentifier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadUserOpaqueIdentifierBySignature indicates an expected call of LoadUserOpaqueIdentifierBySignature.
func (mr *MockStorageMockRecorder) LoadUserOpaqueIdentifierBySignature(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadUserOpaqueIdentifierBySignature", reflect.TypeOf((*MockStorage)(nil).LoadUserOpaqueIdentifierBySignature), arg0, arg1, arg2, arg3)
}

// LoadUserOpaqueIdentifiers mocks base method.
func (m *MockStorage) LoadUserOpaqueIdentifiers(arg0 context.Context) ([]model.UserOpaqueIdentifier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadUserOpaqueIdentifiers", arg0)
	ret0, _ := ret[0].([]model.UserOpaqueIdentifier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadUserOpaqueIdentifiers indicates an expected call of LoadUserOpaqueIdentifiers.
func (mr *MockStorageMockRecorder) LoadUserOpaqueIdentifiers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadUserOpaqueIdentifiers", reflect.TypeOf((*MockStorage)(nil).LoadUserOpaqueIdentifiers), arg0)
}

// LoadWebauthnDeviceByID mocks base method.
func (m *MockStorage) LoadWebauthnDeviceByID(arg0 context.Context, arg1 int) (*model.WebauthnDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadWebauthnDeviceByID", arg0, arg1)
	ret0, _ := ret[0].(*model.WebauthnDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadWebauthnDeviceByID indicates an expected call of LoadWebauthnDeviceByID.
func (mr *MockStorageMockRecorder) LoadWebauthnDeviceByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadWebauthnDeviceByID", reflect.TypeOf((*MockStorage)(nil).LoadWebauthnDeviceByID), arg0, arg1)
}

// LoadWebauthnDevices mocks base method.
func (m *MockStorage) LoadWebauthnDevices(arg0 context.Context, arg1, arg2 int) ([]model.WebauthnDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadWebauthnDevices", arg0, arg1, arg2)
	ret0, _ := ret[0].([]model.WebauthnDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadWebauthnDevices indicates an expected call of LoadWebauthnDevices.
func (mr *MockStorageMockRecorder) LoadWebauthnDevices(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadWebauthnDevices", reflect.TypeOf((*MockStorage)(nil).LoadWebauthnDevices), arg0, arg1, arg2)
}

// LoadWebauthnDevicesByUsername mocks base method.
func (m *MockStorage) LoadWebauthnDevicesByUsername(arg0 context.Context, arg1, arg2 string) ([]model.WebauthnDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadWebauthnDevicesByUsername", arg0, arg1, arg2)
	ret0, _ := ret[0].([]model.WebauthnDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadWebauthnDevicesByUsername indicates an expected call of LoadWebauthnDevicesByUsername.
func (mr *MockStorageMockRecorder) LoadWebauthnDevicesByUsername(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadWebauthnDevicesByUsername", reflect.TypeOf((*MockStorage)(nil).LoadWebauthnDevicesByUsername), arg0, arg1, arg2)
}

// RevokeOAuth2Session mocks base method.
func (m *MockStorage) RevokeOAuth2Session(arg0 context.Context, arg1 storage.OAuth2SessionType, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeOAuth2Session", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeOAuth2Session indicates an expected call of RevokeOAuth2Session.
func (mr *MockStorageMockRecorder) RevokeOAuth2Session(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeOAuth2Session", reflect.TypeOf((*MockStorage)(nil).RevokeOAuth2Session), arg0, arg1, arg2)
}

// RevokeOAuth2SessionByRequestID mocks base method.
func (m *MockStorage) RevokeOAuth2SessionByRequestID(arg0 context.Context, arg1 storage.OAuth2SessionType, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeOAuth2SessionByRequestID", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeOAuth2SessionByRequestID indicates an expected call of RevokeOAuth2SessionByRequestID.
func (mr *MockStorageMockRecorder) RevokeOAuth2SessionByRequestID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeOAuth2SessionByRequestID", reflect.TypeOf((*MockStorage)(nil).RevokeOAuth2SessionByRequestID), arg0, arg1, arg2)
}

// Rollback mocks base method.
func (m *MockStorage) Rollback(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockStorageMockRecorder) Rollback(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockStorage)(nil).Rollback), arg0)
}

// SaveIdentityVerification mocks base method.
func (m *MockStorage) SaveIdentityVerification(arg0 context.Context, arg1 model.IdentityVerification) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveIdentityVerification", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveIdentityVerification indicates an expected call of SaveIdentityVerification.
func (mr *MockStorageMockRecorder) SaveIdentityVerification(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveIdentityVerification", reflect.TypeOf((*MockStorage)(nil).SaveIdentityVerification), arg0, arg1)
}

// SaveOAuth2BlacklistedJTI mocks base method.
func (m *MockStorage) SaveOAuth2BlacklistedJTI(arg0 context.Context, arg1 model.OAuth2BlacklistedJTI) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveOAuth2BlacklistedJTI", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveOAuth2BlacklistedJTI indicates an expected call of SaveOAuth2BlacklistedJTI.
func (mr *MockStorageMockRecorder) SaveOAuth2BlacklistedJTI(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveOAuth2BlacklistedJTI", reflect.TypeOf((*MockStorage)(nil).SaveOAuth2BlacklistedJTI), arg0, arg1)
}

// SaveOAuth2ConsentPreConfiguration mocks base method.
func (m *MockStorage) SaveOAuth2ConsentPreConfiguration(arg0 context.Context, arg1 model.OAuth2ConsentPreConfig) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveOAuth2ConsentPreConfiguration", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveOAuth2ConsentPreConfiguration indicates an expected call of SaveOAuth2ConsentPreConfiguration.
func (mr *MockStorageMockRecorder) SaveOAuth2ConsentPreConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveOAuth2ConsentPreConfiguration", reflect.TypeOf((*MockStorage)(nil).SaveOAuth2ConsentPreConfiguration), arg0, arg1)
}

// SaveOAuth2ConsentSession mocks base method.
func (m *MockStorage) SaveOAuth2ConsentSession(arg0 context.Context, arg1 model.OAuth2ConsentSession) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveOAuth2ConsentSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveOAuth2ConsentSession indicates an expected call of SaveOAuth2ConsentSession.
func (mr *MockStorageMockRecorder) SaveOAuth2ConsentSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveOAuth2ConsentSession", reflect.TypeOf((*MockStorage)(nil).SaveOAuth2ConsentSession), arg0, arg1)
}

// SaveOAuth2ConsentSessionGranted mocks base method.
func (m *MockStorage) SaveOAuth2ConsentSessionGranted(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveOAuth2ConsentSessionGranted", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveOAuth2ConsentSessionGranted indicates an expected call of SaveOAuth2ConsentSessionGranted.
func (mr *MockStorageMockRecorder) SaveOAuth2ConsentSessionGranted(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveOAuth2ConsentSessionGranted", reflect.TypeOf((*MockStorage)(nil).SaveOAuth2ConsentSessionGranted), arg0, arg1)
}

// SaveOAuth2ConsentSessionResponse mocks base method.
func (m *MockStorage) SaveOAuth2ConsentSessionResponse(arg0 context.Context, arg1 model.OAuth2ConsentSession, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveOAuth2ConsentSessionResponse", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveOAuth2ConsentSessionResponse indicates an expected call of SaveOAuth2ConsentSessionResponse.
func (mr *MockStorageMockRecorder) SaveOAuth2ConsentSessionResponse(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveOAuth2ConsentSessionResponse", reflect.TypeOf((*MockStorage)(nil).SaveOAuth2ConsentSessionResponse), arg0, arg1, arg2)
}

// SaveOAuth2ConsentSessionSubject mocks base method.
func (m *MockStorage) SaveOAuth2ConsentSessionSubject(arg0 context.Context, arg1 model.OAuth2ConsentSession) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveOAuth2ConsentSessionSubject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveOAuth2ConsentSessionSubject indicates an expected call of SaveOAuth2ConsentSessionSubject.
func (mr *MockStorageMockRecorder) SaveOAuth2ConsentSessionSubject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveOAuth2ConsentSessionSubject", reflect.TypeOf((*MockStorage)(nil).SaveOAuth2ConsentSessionSubject), arg0, arg1)
}

// SaveOAuth2Session mocks base method.
func (m *MockStorage) SaveOAuth2Session(arg0 context.Context, arg1 storage.OAuth2SessionType, arg2 model.OAuth2Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveOAuth2Session", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveOAuth2Session indicates an expected call of SaveOAuth2Session.
func (mr *MockStorageMockRecorder) SaveOAuth2Session(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveOAuth2Session", reflect.TypeOf((*MockStorage)(nil).SaveOAuth2Session), arg0, arg1, arg2)
}

// SavePreferred2FAMethod mocks base method.
func (m *MockStorage) SavePreferred2FAMethod(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePreferred2FAMethod", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SavePreferred2FAMethod indicates an expected call of SavePreferred2FAMethod.
func (mr *MockStorageMockRecorder) SavePreferred2FAMethod(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePreferred2FAMethod", reflect.TypeOf((*MockStorage)(nil).SavePreferred2FAMethod), arg0, arg1, arg2)
}

// SavePreferredDuoDevice mocks base method.
func (m *MockStorage) SavePreferredDuoDevice(arg0 context.Context, arg1 model.DuoDevice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePreferredDuoDevice", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SavePreferredDuoDevice indicates an expected call of SavePreferredDuoDevice.
func (mr *MockStorageMockRecorder) SavePreferredDuoDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePreferredDuoDevice", reflect.TypeOf((*MockStorage)(nil).SavePreferredDuoDevice), arg0, arg1)
}

// SaveTOTPConfiguration mocks base method.
func (m *MockStorage) SaveTOTPConfiguration(arg0 context.Context, arg1 model.TOTPConfiguration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTOTPConfiguration", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTOTPConfiguration indicates an expected call of SaveTOTPConfiguration.
func (mr *MockStorageMockRecorder) SaveTOTPConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTOTPConfiguration", reflect.TypeOf((*MockStorage)(nil).SaveTOTPConfiguration), arg0, arg1)
}

// SaveUserOpaqueIdentifier mocks base method.
func (m *MockStorage) SaveUserOpaqueIdentifier(arg0 context.Context, arg1 model.UserOpaqueIdentifier) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUserOpaqueIdentifier", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUserOpaqueIdentifier indicates an expected call of SaveUserOpaqueIdentifier.
func (mr *MockStorageMockRecorder) SaveUserOpaqueIdentifier(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUserOpaqueIdentifier", reflect.TypeOf((*MockStorage)(nil).SaveUserOpaqueIdentifier), arg0, arg1)
}

// SaveWebauthnDevice mocks base method.
func (m *MockStorage) SaveWebauthnDevice(arg0 context.Context, arg1 model.WebauthnDevice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveWebauthnDevice", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveWebauthnDevice indicates an expected call of SaveWebauthnDevice.
func (mr *MockStorageMockRecorder) SaveWebauthnDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveWebauthnDevice", reflect.TypeOf((*MockStorage)(nil).SaveWebauthnDevice), arg0, arg1)
}

// SchemaEncryptionChangeKey mocks base method.
func (m *MockStorage) SchemaEncryptionChangeKey(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaEncryptionChangeKey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SchemaEncryptionChangeKey indicates an expected call of SchemaEncryptionChangeKey.
func (mr *MockStorageMockRecorder) SchemaEncryptionChangeKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaEncryptionChangeKey", reflect.TypeOf((*MockStorage)(nil).SchemaEncryptionChangeKey), arg0, arg1)
}

// SchemaEncryptionCheckKey mocks base method.
func (m *MockStorage) SchemaEncryptionCheckKey(arg0 context.Context, arg1 bool) (storage.EncryptionValidationResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaEncryptionCheckKey", arg0, arg1)
	ret0, _ := ret[0].(storage.EncryptionValidationResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SchemaEncryptionCheckKey indicates an expected call of SchemaEncryptionCheckKey.
func (mr *MockStorageMockRecorder) SchemaEncryptionCheckKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaEncryptionCheckKey", reflect.TypeOf((*MockStorage)(nil).SchemaEncryptionCheckKey), arg0, arg1)
}

// SchemaLatestVersion mocks base method.
func (m *MockStorage) SchemaLatestVersion() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaLatestVersion")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SchemaLatestVersion indicates an expected call of SchemaLatestVersion.
func (mr *MockStorageMockRecorder) SchemaLatestVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaLatestVersion", reflect.TypeOf((*MockStorage)(nil).SchemaLatestVersion))
}

// SchemaMigrate mocks base method.
func (m *MockStorage) SchemaMigrate(arg0 context.Context, arg1 bool, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaMigrate", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SchemaMigrate indicates an expected call of SchemaMigrate.
func (mr *MockStorageMockRecorder) SchemaMigrate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaMigrate", reflect.TypeOf((*MockStorage)(nil).SchemaMigrate), arg0, arg1, arg2)
}

// SchemaMigrationHistory mocks base method.
func (m *MockStorage) SchemaMigrationHistory(arg0 context.Context) ([]model.Migration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaMigrationHistory", arg0)
	ret0, _ := ret[0].([]model.Migration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SchemaMigrationHistory indicates an expected call of SchemaMigrationHistory.
func (mr *MockStorageMockRecorder) SchemaMigrationHistory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaMigrationHistory", reflect.TypeOf((*MockStorage)(nil).SchemaMigrationHistory), arg0)
}

// SchemaMigrationsDown mocks base method.
func (m *MockStorage) SchemaMigrationsDown(arg0 context.Context, arg1 int) ([]model.SchemaMigration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaMigrationsDown", arg0, arg1)
	ret0, _ := ret[0].([]model.SchemaMigration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SchemaMigrationsDown indicates an expected call of SchemaMigrationsDown.
func (mr *MockStorageMockRecorder) SchemaMigrationsDown(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaMigrationsDown", reflect.TypeOf((*MockStorage)(nil).SchemaMigrationsDown), arg0, arg1)
}

// SchemaMigrationsUp mocks base method.
func (m *MockStorage) SchemaMigrationsUp(arg0 context.Context, arg1 int) ([]model.SchemaMigration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaMigrationsUp", arg0, arg1)
	ret0, _ := ret[0].([]model.SchemaMigration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SchemaMigrationsUp indicates an expected call of SchemaMigrationsUp.
func (mr *MockStorageMockRecorder) SchemaMigrationsUp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaMigrationsUp", reflect.TypeOf((*MockStorage)(nil).SchemaMigrationsUp), arg0, arg1)
}

// SchemaTables mocks base method.
func (m *MockStorage) SchemaTables(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaTables", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SchemaTables indicates an expected call of SchemaTables.
func (mr *MockStorageMockRecorder) SchemaTables(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaTables", reflect.TypeOf((*MockStorage)(nil).SchemaTables), arg0)
}

// SchemaVersion mocks base method.
func (m *MockStorage) SchemaVersion(arg0 context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaVersion", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SchemaVersion indicates an expected call of SchemaVersion.
func (mr *MockStorageMockRecorder) SchemaVersion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaVersion", reflect.TypeOf((*MockStorage)(nil).SchemaVersion), arg0)
}

// StartupCheck mocks base method.
func (m *MockStorage) StartupCheck() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartupCheck")
	ret0, _ := ret[0].(error)
	return ret0
}

// StartupCheck indicates an expected call of StartupCheck.
func (mr *MockStorageMockRecorder) StartupCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartupCheck", reflect.TypeOf((*MockStorage)(nil).StartupCheck))
}

// UpdateTOTPConfigurationSignIn mocks base method.
func (m *MockStorage) UpdateTOTPConfigurationSignIn(arg0 context.Context, arg1 int, arg2 sql.NullTime) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTOTPConfigurationSignIn", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTOTPConfigurationSignIn indicates an expected call of UpdateTOTPConfigurationSignIn.
func (mr *MockStorageMockRecorder) UpdateTOTPConfigurationSignIn(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTOTPConfigurationSignIn", reflect.TypeOf((*MockStorage)(nil).UpdateTOTPConfigurationSignIn), arg0, arg1, arg2)
}

// UpdateWebauthnDeviceDescription mocks base method.
func (m *MockStorage) UpdateWebauthnDeviceDescription(arg0 context.Context, arg1 string, arg2 int, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWebauthnDeviceDescription", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWebauthnDeviceDescription indicates an expected call of UpdateWebauthnDeviceDescription.
func (mr *MockStorageMockRecorder) UpdateWebauthnDeviceDescription(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWebauthnDeviceDescription", reflect.TypeOf((*MockStorage)(nil).UpdateWebauthnDeviceDescription), arg0, arg1, arg2, arg3)
}

// UpdateWebauthnDeviceSignIn mocks base method.
func (m *MockStorage) UpdateWebauthnDeviceSignIn(arg0 context.Context, arg1 model.WebauthnDevice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWebauthnDeviceSignIn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWebauthnDeviceSignIn indicates an expected call of UpdateWebauthnDeviceSignIn.
func (mr *MockStorageMockRecorder) UpdateWebauthnDeviceSignIn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWebauthnDeviceSignIn", reflect.TypeOf((*MockStorage)(nil).UpdateWebauthnDeviceSignIn), arg0, arg1)
}
