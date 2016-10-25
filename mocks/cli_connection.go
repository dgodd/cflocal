// Automatically generated by MockGen. DO NOT EDIT!
// Source: code.cloudfoundry.org/cli/plugin (interfaces: CliConnection)

package mocks

import (
	models "code.cloudfoundry.org/cli/plugin/models"
	gomock "github.com/golang/mock/gomock"
)

// Mock of CliConnection interface
type MockCliConnection struct {
	ctrl     *gomock.Controller
	recorder *_MockCliConnectionRecorder
}

// Recorder for MockCliConnection (not exported)
type _MockCliConnectionRecorder struct {
	mock *MockCliConnection
}

func NewMockCliConnection(ctrl *gomock.Controller) *MockCliConnection {
	mock := &MockCliConnection{ctrl: ctrl}
	mock.recorder = &_MockCliConnectionRecorder{mock}
	return mock
}

func (_m *MockCliConnection) EXPECT() *_MockCliConnectionRecorder {
	return _m.recorder
}

func (_m *MockCliConnection) AccessToken() (string, error) {
	ret := _m.ctrl.Call(_m, "AccessToken")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) AccessToken() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AccessToken")
}

func (_m *MockCliConnection) ApiEndpoint() (string, error) {
	ret := _m.ctrl.Call(_m, "ApiEndpoint")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) ApiEndpoint() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ApiEndpoint")
}

func (_m *MockCliConnection) ApiVersion() (string, error) {
	ret := _m.ctrl.Call(_m, "ApiVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) ApiVersion() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ApiVersion")
}

func (_m *MockCliConnection) CliCommand(_param0 ...string) ([]string, error) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CliCommand", _s...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) CliCommand(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CliCommand", arg0...)
}

func (_m *MockCliConnection) CliCommandWithoutTerminalOutput(_param0 ...string) ([]string, error) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CliCommandWithoutTerminalOutput", _s...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) CliCommandWithoutTerminalOutput(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CliCommandWithoutTerminalOutput", arg0...)
}

func (_m *MockCliConnection) DopplerEndpoint() (string, error) {
	ret := _m.ctrl.Call(_m, "DopplerEndpoint")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) DopplerEndpoint() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DopplerEndpoint")
}

func (_m *MockCliConnection) GetApp(_param0 string) (models.GetAppModel, error) {
	ret := _m.ctrl.Call(_m, "GetApp", _param0)
	ret0, _ := ret[0].(models.GetAppModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) GetApp(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetApp", arg0)
}

func (_m *MockCliConnection) GetApps() ([]models.GetAppsModel, error) {
	ret := _m.ctrl.Call(_m, "GetApps")
	ret0, _ := ret[0].([]models.GetAppsModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) GetApps() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetApps")
}

func (_m *MockCliConnection) GetCurrentOrg() (models.Organization, error) {
	ret := _m.ctrl.Call(_m, "GetCurrentOrg")
	ret0, _ := ret[0].(models.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) GetCurrentOrg() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCurrentOrg")
}

func (_m *MockCliConnection) GetCurrentSpace() (models.Space, error) {
	ret := _m.ctrl.Call(_m, "GetCurrentSpace")
	ret0, _ := ret[0].(models.Space)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) GetCurrentSpace() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCurrentSpace")
}

func (_m *MockCliConnection) GetOrg(_param0 string) (models.GetOrg_Model, error) {
	ret := _m.ctrl.Call(_m, "GetOrg", _param0)
	ret0, _ := ret[0].(models.GetOrg_Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) GetOrg(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOrg", arg0)
}

func (_m *MockCliConnection) GetOrgUsers(_param0 string, _param1 ...string) ([]models.GetOrgUsers_Model, error) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetOrgUsers", _s...)
	ret0, _ := ret[0].([]models.GetOrgUsers_Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) GetOrgUsers(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOrgUsers", _s...)
}

func (_m *MockCliConnection) GetOrgs() ([]models.GetOrgs_Model, error) {
	ret := _m.ctrl.Call(_m, "GetOrgs")
	ret0, _ := ret[0].([]models.GetOrgs_Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) GetOrgs() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOrgs")
}

func (_m *MockCliConnection) GetService(_param0 string) (models.GetService_Model, error) {
	ret := _m.ctrl.Call(_m, "GetService", _param0)
	ret0, _ := ret[0].(models.GetService_Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) GetService(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetService", arg0)
}

func (_m *MockCliConnection) GetServices() ([]models.GetServices_Model, error) {
	ret := _m.ctrl.Call(_m, "GetServices")
	ret0, _ := ret[0].([]models.GetServices_Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) GetServices() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetServices")
}

func (_m *MockCliConnection) GetSpace(_param0 string) (models.GetSpace_Model, error) {
	ret := _m.ctrl.Call(_m, "GetSpace", _param0)
	ret0, _ := ret[0].(models.GetSpace_Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) GetSpace(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSpace", arg0)
}

func (_m *MockCliConnection) GetSpaceUsers(_param0 string, _param1 string) ([]models.GetSpaceUsers_Model, error) {
	ret := _m.ctrl.Call(_m, "GetSpaceUsers", _param0, _param1)
	ret0, _ := ret[0].([]models.GetSpaceUsers_Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) GetSpaceUsers(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSpaceUsers", arg0, arg1)
}

func (_m *MockCliConnection) GetSpaces() ([]models.GetSpaces_Model, error) {
	ret := _m.ctrl.Call(_m, "GetSpaces")
	ret0, _ := ret[0].([]models.GetSpaces_Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) GetSpaces() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSpaces")
}

func (_m *MockCliConnection) HasAPIEndpoint() (bool, error) {
	ret := _m.ctrl.Call(_m, "HasAPIEndpoint")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) HasAPIEndpoint() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HasAPIEndpoint")
}

func (_m *MockCliConnection) HasOrganization() (bool, error) {
	ret := _m.ctrl.Call(_m, "HasOrganization")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) HasOrganization() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HasOrganization")
}

func (_m *MockCliConnection) HasSpace() (bool, error) {
	ret := _m.ctrl.Call(_m, "HasSpace")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) HasSpace() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HasSpace")
}

func (_m *MockCliConnection) IsLoggedIn() (bool, error) {
	ret := _m.ctrl.Call(_m, "IsLoggedIn")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) IsLoggedIn() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsLoggedIn")
}

func (_m *MockCliConnection) IsSSLDisabled() (bool, error) {
	ret := _m.ctrl.Call(_m, "IsSSLDisabled")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) IsSSLDisabled() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsSSLDisabled")
}

func (_m *MockCliConnection) LoggregatorEndpoint() (string, error) {
	ret := _m.ctrl.Call(_m, "LoggregatorEndpoint")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) LoggregatorEndpoint() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LoggregatorEndpoint")
}

func (_m *MockCliConnection) UserEmail() (string, error) {
	ret := _m.ctrl.Call(_m, "UserEmail")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) UserEmail() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UserEmail")
}

func (_m *MockCliConnection) UserGuid() (string, error) {
	ret := _m.ctrl.Call(_m, "UserGuid")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) UserGuid() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UserGuid")
}

func (_m *MockCliConnection) Username() (string, error) {
	ret := _m.ctrl.Call(_m, "Username")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCliConnectionRecorder) Username() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Username")
}