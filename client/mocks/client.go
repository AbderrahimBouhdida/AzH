// Copyright (C) 2022 Specter Ops, Inc.
//
// This file is part of AzureHound.
//
// AzureHound is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// AzureHound is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	json "encoding/json"
	reflect "reflect"

	azure "github.com/bloodhoundad/azurehound/models/azure"
	gomock "github.com/golang/mock/gomock"
)

// MockAzureClient is a mock of AzureClient interface.
type MockAzureClient struct {
	ctrl     *gomock.Controller
	recorder *MockAzureClientMockRecorder
}

// MockAzureClientMockRecorder is the mock recorder for MockAzureClient.
type MockAzureClientMockRecorder struct {
	mock *MockAzureClient
}

// NewMockAzureClient creates a new mock instance.
func NewMockAzureClient(ctrl *gomock.Controller) *MockAzureClient {
	mock := &MockAzureClient{ctrl: ctrl}
	mock.recorder = &MockAzureClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAzureClient) EXPECT() *MockAzureClientMockRecorder {
	return m.recorder
}

// GetAzureADApp mocks base method.
func (m *MockAzureClient) GetAzureADApp(arg0 context.Context, arg1 string, arg2 []string) (*azure.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureADApp", arg0, arg1, arg2)
	ret0, _ := ret[0].(*azure.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureADApp indicates an expected call of GetAzureADApp.
func (mr *MockAzureClientMockRecorder) GetAzureADApp(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureADApp", reflect.TypeOf((*MockAzureClient)(nil).GetAzureADApp), arg0, arg1, arg2)
}

// GetAzureADApps mocks base method.
func (m *MockAzureClient) GetAzureADApps(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 []string, arg6 int32, arg7 bool) (azure.ApplicationList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureADApps", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(azure.ApplicationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureADApps indicates an expected call of GetAzureADApps.
func (mr *MockAzureClientMockRecorder) GetAzureADApps(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureADApps", reflect.TypeOf((*MockAzureClient)(nil).GetAzureADApps), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// GetAzureADDirectoryObject mocks base method.
func (m *MockAzureClient) GetAzureADDirectoryObject(arg0 context.Context, arg1 string) (json.RawMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureADDirectoryObject", arg0, arg1)
	ret0, _ := ret[0].(json.RawMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureADDirectoryObject indicates an expected call of GetAzureADDirectoryObject.
func (mr *MockAzureClientMockRecorder) GetAzureADDirectoryObject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureADDirectoryObject", reflect.TypeOf((*MockAzureClient)(nil).GetAzureADDirectoryObject), arg0, arg1)
}

// GetAzureADGroup mocks base method.
func (m *MockAzureClient) GetAzureADGroup(arg0 context.Context, arg1 string, arg2 []string) (*azure.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureADGroup", arg0, arg1, arg2)
	ret0, _ := ret[0].(*azure.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureADGroup indicates an expected call of GetAzureADGroup.
func (mr *MockAzureClientMockRecorder) GetAzureADGroup(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureADGroup", reflect.TypeOf((*MockAzureClient)(nil).GetAzureADGroup), arg0, arg1, arg2)
}

// GetAzureADGroupOwners mocks base method.
func (m *MockAzureClient) GetAzureADGroupOwners(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 []string, arg6 int32, arg7 bool) (azure.DirectoryObjectList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureADGroupOwners", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(azure.DirectoryObjectList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureADGroupOwners indicates an expected call of GetAzureADGroupOwners.
func (mr *MockAzureClientMockRecorder) GetAzureADGroupOwners(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureADGroupOwners", reflect.TypeOf((*MockAzureClient)(nil).GetAzureADGroupOwners), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// GetAzureADGroups mocks base method.
func (m *MockAzureClient) GetAzureADGroups(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 []string, arg6 int32, arg7 bool) (azure.GroupList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureADGroups", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(azure.GroupList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureADGroups indicates an expected call of GetAzureADGroups.
func (mr *MockAzureClientMockRecorder) GetAzureADGroups(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureADGroups", reflect.TypeOf((*MockAzureClient)(nil).GetAzureADGroups), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// GetAzureADOrganization mocks base method.
func (m *MockAzureClient) GetAzureADOrganization(arg0 context.Context, arg1 []string) (*azure.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureADOrganization", arg0, arg1)
	ret0, _ := ret[0].(*azure.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureADOrganization indicates an expected call of GetAzureADOrganization.
func (mr *MockAzureClientMockRecorder) GetAzureADOrganization(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureADOrganization", reflect.TypeOf((*MockAzureClient)(nil).GetAzureADOrganization), arg0, arg1)
}

// GetAzureADRole mocks base method.
func (m *MockAzureClient) GetAzureADRole(arg0 context.Context, arg1 string, arg2 []string) (*azure.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureADRole", arg0, arg1, arg2)
	ret0, _ := ret[0].(*azure.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureADRole indicates an expected call of GetAzureADRole.
func (mr *MockAzureClientMockRecorder) GetAzureADRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureADRole", reflect.TypeOf((*MockAzureClient)(nil).GetAzureADRole), arg0, arg1, arg2)
}

// GetAzureADRoleAssignment mocks base method.
func (m *MockAzureClient) GetAzureADRoleAssignment(arg0 context.Context, arg1 string, arg2 []string) (*azure.UnifiedRoleAssignment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureADRoleAssignment", arg0, arg1, arg2)
	ret0, _ := ret[0].(*azure.UnifiedRoleAssignment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureADRoleAssignment indicates an expected call of GetAzureADRoleAssignment.
func (mr *MockAzureClientMockRecorder) GetAzureADRoleAssignment(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureADRoleAssignment", reflect.TypeOf((*MockAzureClient)(nil).GetAzureADRoleAssignment), arg0, arg1, arg2)
}

// GetAzureADRoleAssignments mocks base method.
func (m *MockAzureClient) GetAzureADRoleAssignments(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 []string, arg6 int32, arg7 bool) (azure.UnifiedRoleAssignmentList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureADRoleAssignments", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(azure.UnifiedRoleAssignmentList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureADRoleAssignments indicates an expected call of GetAzureADRoleAssignments.
func (mr *MockAzureClientMockRecorder) GetAzureADRoleAssignments(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureADRoleAssignments", reflect.TypeOf((*MockAzureClient)(nil).GetAzureADRoleAssignments), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// GetAzureADRoles mocks base method.
func (m *MockAzureClient) GetAzureADRoles(arg0 context.Context, arg1, arg2 string) (azure.RoleList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureADRoles", arg0, arg1, arg2)
	ret0, _ := ret[0].(azure.RoleList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureADRoles indicates an expected call of GetAzureADRoles.
func (mr *MockAzureClientMockRecorder) GetAzureADRoles(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureADRoles", reflect.TypeOf((*MockAzureClient)(nil).GetAzureADRoles), arg0, arg1, arg2)
}

// GetAzureADServicePrincipal mocks base method.
func (m *MockAzureClient) GetAzureADServicePrincipal(arg0 context.Context, arg1 string, arg2 []string) (*azure.ServicePrincipal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureADServicePrincipal", arg0, arg1, arg2)
	ret0, _ := ret[0].(*azure.ServicePrincipal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureADServicePrincipal indicates an expected call of GetAzureADServicePrincipal.
func (mr *MockAzureClientMockRecorder) GetAzureADServicePrincipal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureADServicePrincipal", reflect.TypeOf((*MockAzureClient)(nil).GetAzureADServicePrincipal), arg0, arg1, arg2)
}

// GetAzureADServicePrincipalOwners mocks base method.
func (m *MockAzureClient) GetAzureADServicePrincipalOwners(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 []string, arg6 int32, arg7 bool) (azure.DirectoryObjectList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureADServicePrincipalOwners", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(azure.DirectoryObjectList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureADServicePrincipalOwners indicates an expected call of GetAzureADServicePrincipalOwners.
func (mr *MockAzureClientMockRecorder) GetAzureADServicePrincipalOwners(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureADServicePrincipalOwners", reflect.TypeOf((*MockAzureClient)(nil).GetAzureADServicePrincipalOwners), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// GetAzureADServicePrincipals mocks base method.
func (m *MockAzureClient) GetAzureADServicePrincipals(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 []string, arg6 int32, arg7 bool) (azure.ServicePrincipalList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureADServicePrincipals", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(azure.ServicePrincipalList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureADServicePrincipals indicates an expected call of GetAzureADServicePrincipals.
func (mr *MockAzureClientMockRecorder) GetAzureADServicePrincipals(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureADServicePrincipals", reflect.TypeOf((*MockAzureClient)(nil).GetAzureADServicePrincipals), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// GetAzureADTenants mocks base method.
func (m *MockAzureClient) GetAzureADTenants(arg0 context.Context, arg1 bool) (azure.TenantList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureADTenants", arg0, arg1)
	ret0, _ := ret[0].(azure.TenantList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureADTenants indicates an expected call of GetAzureADTenants.
func (mr *MockAzureClientMockRecorder) GetAzureADTenants(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureADTenants", reflect.TypeOf((*MockAzureClient)(nil).GetAzureADTenants), arg0, arg1)
}

// GetAzureADUser mocks base method.
func (m *MockAzureClient) GetAzureADUser(arg0 context.Context, arg1 string, arg2 []string) (*azure.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureADUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(*azure.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureADUser indicates an expected call of GetAzureADUser.
func (mr *MockAzureClientMockRecorder) GetAzureADUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureADUser", reflect.TypeOf((*MockAzureClient)(nil).GetAzureADUser), arg0, arg1, arg2)
}

// GetAzureADUsers mocks base method.
func (m *MockAzureClient) GetAzureADUsers(arg0 context.Context, arg1, arg2, arg3 string, arg4 []string, arg5 int32, arg6 bool) (azure.UserList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureADUsers", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(azure.UserList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureADUsers indicates an expected call of GetAzureADUsers.
func (mr *MockAzureClientMockRecorder) GetAzureADUsers(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureADUsers", reflect.TypeOf((*MockAzureClient)(nil).GetAzureADUsers), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// GetAzureDevice mocks base method.
func (m *MockAzureClient) GetAzureDevice(arg0 context.Context, arg1 string, arg2 []string) (*azure.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureDevice", arg0, arg1, arg2)
	ret0, _ := ret[0].(*azure.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureDevice indicates an expected call of GetAzureDevice.
func (mr *MockAzureClientMockRecorder) GetAzureDevice(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureDevice", reflect.TypeOf((*MockAzureClient)(nil).GetAzureDevice), arg0, arg1, arg2)
}

// GetAzureDevices mocks base method.
func (m *MockAzureClient) GetAzureDevices(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 []string, arg6 int32, arg7 bool) (azure.DeviceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureDevices", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(azure.DeviceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureDevices indicates an expected call of GetAzureDevices.
func (mr *MockAzureClientMockRecorder) GetAzureDevices(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureDevices", reflect.TypeOf((*MockAzureClient)(nil).GetAzureDevices), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// GetAzureKeyVault mocks base method.
func (m *MockAzureClient) GetAzureKeyVault(arg0 context.Context, arg1, arg2, arg3 string) (*azure.KeyVault, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureKeyVault", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*azure.KeyVault)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureKeyVault indicates an expected call of GetAzureKeyVault.
func (mr *MockAzureClientMockRecorder) GetAzureKeyVault(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureKeyVault", reflect.TypeOf((*MockAzureClient)(nil).GetAzureKeyVault), arg0, arg1, arg2, arg3)
}

// GetAzureKeyVaults mocks base method.
func (m *MockAzureClient) GetAzureKeyVaults(arg0 context.Context, arg1 string, arg2 int32) (azure.KeyVaultList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureKeyVaults", arg0, arg1, arg2)
	ret0, _ := ret[0].(azure.KeyVaultList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureKeyVaults indicates an expected call of GetAzureKeyVaults.
func (mr *MockAzureClientMockRecorder) GetAzureKeyVaults(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureKeyVaults", reflect.TypeOf((*MockAzureClient)(nil).GetAzureKeyVaults), arg0, arg1, arg2)
}

// GetAzureManagementGroup mocks base method.
func (m *MockAzureClient) GetAzureManagementGroup(arg0 context.Context, arg1, arg2, arg3 string, arg4 bool) (*azure.ManagementGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureManagementGroup", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*azure.ManagementGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureManagementGroup indicates an expected call of GetAzureManagementGroup.
func (mr *MockAzureClientMockRecorder) GetAzureManagementGroup(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureManagementGroup", reflect.TypeOf((*MockAzureClient)(nil).GetAzureManagementGroup), arg0, arg1, arg2, arg3, arg4)
}

// GetAzureManagementGroups mocks base method.
func (m *MockAzureClient) GetAzureManagementGroups(arg0 context.Context) (azure.ManagementGroupList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureManagementGroups", arg0)
	ret0, _ := ret[0].(azure.ManagementGroupList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureManagementGroups indicates an expected call of GetAzureManagementGroups.
func (mr *MockAzureClientMockRecorder) GetAzureManagementGroups(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureManagementGroups", reflect.TypeOf((*MockAzureClient)(nil).GetAzureManagementGroups), arg0)
}

// GetAzureResourceGroup mocks base method.
func (m *MockAzureClient) GetAzureResourceGroup(arg0 context.Context, arg1, arg2 string) (*azure.ResourceGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureResourceGroup", arg0, arg1, arg2)
	ret0, _ := ret[0].(*azure.ResourceGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureResourceGroup indicates an expected call of GetAzureResourceGroup.
func (mr *MockAzureClientMockRecorder) GetAzureResourceGroup(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureResourceGroup", reflect.TypeOf((*MockAzureClient)(nil).GetAzureResourceGroup), arg0, arg1, arg2)
}

// GetAzureResourceGroups mocks base method.
func (m *MockAzureClient) GetAzureResourceGroups(arg0 context.Context, arg1, arg2 string, arg3 int32) (azure.ResourceGroupList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureResourceGroups", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(azure.ResourceGroupList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureResourceGroups indicates an expected call of GetAzureResourceGroups.
func (mr *MockAzureClientMockRecorder) GetAzureResourceGroups(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureResourceGroups", reflect.TypeOf((*MockAzureClient)(nil).GetAzureResourceGroups), arg0, arg1, arg2, arg3)
}

// GetAzureStorageAccount mocks base method.
func (m *MockAzureClient) GetAzureStorageAccount(arg0 context.Context, arg1, arg2, arg3, arg4 string) (*azure.StorageAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureStorageAccount", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*azure.StorageAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureStorageAccount indicates an expected call of GetAzureStorageAccount.
func (mr *MockAzureClientMockRecorder) GetAzureStorageAccount(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureStorageAccount", reflect.TypeOf((*MockAzureClient)(nil).GetAzureStorageAccount), arg0, arg1, arg2, arg3, arg4)
}

// GetAzureStorageAccounts mocks base method.
func (m *MockAzureClient) GetAzureStorageAccounts(arg0 context.Context, arg1 string, arg2 bool) (azure.StorageAccountList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureStorageAccounts", arg0, arg1, arg2)
	ret0, _ := ret[0].(azure.StorageAccountList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureStorageAccounts indicates an expected call of GetAzureStorageAccounts.
func (mr *MockAzureClientMockRecorder) GetAzureStorageAccounts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureStorageAccounts", reflect.TypeOf((*MockAzureClient)(nil).GetAzureStorageAccounts), arg0, arg1, arg2)
}

// GetAzureSubscription mocks base method.
func (m *MockAzureClient) GetAzureSubscription(arg0 context.Context, arg1 string) (*azure.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureSubscription", arg0, arg1)
	ret0, _ := ret[0].(*azure.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureSubscription indicates an expected call of GetAzureSubscription.
func (mr *MockAzureClientMockRecorder) GetAzureSubscription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureSubscription", reflect.TypeOf((*MockAzureClient)(nil).GetAzureSubscription), arg0, arg1)
}

// GetAzureSubscriptions mocks base method.
func (m *MockAzureClient) GetAzureSubscriptions(arg0 context.Context) (azure.SubscriptionList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureSubscriptions", arg0)
	ret0, _ := ret[0].(azure.SubscriptionList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureSubscriptions indicates an expected call of GetAzureSubscriptions.
func (mr *MockAzureClientMockRecorder) GetAzureSubscriptions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureSubscriptions", reflect.TypeOf((*MockAzureClient)(nil).GetAzureSubscriptions), arg0)
}

// GetAzureVirtualMachine mocks base method.
func (m *MockAzureClient) GetAzureVirtualMachine(arg0 context.Context, arg1, arg2, arg3, arg4 string) (*azure.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureVirtualMachine", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*azure.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureVirtualMachine indicates an expected call of GetAzureVirtualMachine.
func (mr *MockAzureClientMockRecorder) GetAzureVirtualMachine(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureVirtualMachine", reflect.TypeOf((*MockAzureClient)(nil).GetAzureVirtualMachine), arg0, arg1, arg2, arg3, arg4)
}

// GetAzureVirtualMachines mocks base method.
func (m *MockAzureClient) GetAzureVirtualMachines(arg0 context.Context, arg1 string, arg2 bool) (azure.VirtualMachineList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAzureVirtualMachines", arg0, arg1, arg2)
	ret0, _ := ret[0].(azure.VirtualMachineList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAzureVirtualMachines indicates an expected call of GetAzureVirtualMachines.
func (mr *MockAzureClientMockRecorder) GetAzureVirtualMachines(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAzureVirtualMachines", reflect.TypeOf((*MockAzureClient)(nil).GetAzureVirtualMachines), arg0, arg1, arg2)
}

// GetResourceRoleAssignments mocks base method.
func (m *MockAzureClient) GetResourceRoleAssignments(arg0 context.Context, arg1, arg2, arg3 string) (azure.RoleAssignmentList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceRoleAssignments", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(azure.RoleAssignmentList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceRoleAssignments indicates an expected call of GetResourceRoleAssignments.
func (mr *MockAzureClientMockRecorder) GetResourceRoleAssignments(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceRoleAssignments", reflect.TypeOf((*MockAzureClient)(nil).GetResourceRoleAssignments), arg0, arg1, arg2, arg3)
}

// GetRoleAssignmentsForResource mocks base method.
func (m *MockAzureClient) GetRoleAssignmentsForResource(arg0 context.Context, arg1, arg2 string) (azure.RoleAssignmentList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleAssignmentsForResource", arg0, arg1, arg2)
	ret0, _ := ret[0].(azure.RoleAssignmentList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleAssignmentsForResource indicates an expected call of GetRoleAssignmentsForResource.
func (mr *MockAzureClientMockRecorder) GetRoleAssignmentsForResource(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleAssignmentsForResource", reflect.TypeOf((*MockAzureClient)(nil).GetRoleAssignmentsForResource), arg0, arg1, arg2)
}

// ListAzureADAppMemberObjects mocks base method.
func (m *MockAzureClient) ListAzureADAppMemberObjects(arg0 context.Context, arg1 string, arg2 bool) <-chan azure.MemberObjectResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADAppMemberObjects", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan azure.MemberObjectResult)
	return ret0
}

// ListAzureADAppMemberObjects indicates an expected call of ListAzureADAppMemberObjects.
func (mr *MockAzureClientMockRecorder) ListAzureADAppMemberObjects(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADAppMemberObjects", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADAppMemberObjects), arg0, arg1, arg2)
}

// ListAzureADAppOwners mocks base method.
func (m *MockAzureClient) ListAzureADAppOwners(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 []string) <-chan azure.AppOwnerResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADAppOwners", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(<-chan azure.AppOwnerResult)
	return ret0
}

// ListAzureADAppOwners indicates an expected call of ListAzureADAppOwners.
func (mr *MockAzureClientMockRecorder) ListAzureADAppOwners(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADAppOwners", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADAppOwners), arg0, arg1, arg2, arg3, arg4, arg5)
}

// ListAzureADAppRoleAssignments mocks base method.
func (m *MockAzureClient) ListAzureADAppRoleAssignments(arg0 context.Context, arg1, arg2, arg3, arg4, arg5 string, arg6 []string) <-chan azure.AppRoleAssignmentResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADAppRoleAssignments", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(<-chan azure.AppRoleAssignmentResult)
	return ret0
}

// ListAzureADAppRoleAssignments indicates an expected call of ListAzureADAppRoleAssignments.
func (mr *MockAzureClientMockRecorder) ListAzureADAppRoleAssignments(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADAppRoleAssignments", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADAppRoleAssignments), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// ListAzureADApps mocks base method.
func (m *MockAzureClient) ListAzureADApps(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 []string) <-chan azure.ApplicationResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADApps", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(<-chan azure.ApplicationResult)
	return ret0
}

// ListAzureADApps indicates an expected call of ListAzureADApps.
func (mr *MockAzureClientMockRecorder) ListAzureADApps(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADApps", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADApps), arg0, arg1, arg2, arg3, arg4, arg5)
}

// ListAzureADGroupMembers mocks base method.
func (m *MockAzureClient) ListAzureADGroupMembers(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 []string) <-chan azure.MemberObjectResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADGroupMembers", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(<-chan azure.MemberObjectResult)
	return ret0
}

// ListAzureADGroupMembers indicates an expected call of ListAzureADGroupMembers.
func (mr *MockAzureClientMockRecorder) ListAzureADGroupMembers(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADGroupMembers", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADGroupMembers), arg0, arg1, arg2, arg3, arg4, arg5)
}

// ListAzureADGroupOwners mocks base method.
func (m *MockAzureClient) ListAzureADGroupOwners(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 []string) <-chan azure.GroupOwnerResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADGroupOwners", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(<-chan azure.GroupOwnerResult)
	return ret0
}

// ListAzureADGroupOwners indicates an expected call of ListAzureADGroupOwners.
func (mr *MockAzureClientMockRecorder) ListAzureADGroupOwners(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADGroupOwners", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADGroupOwners), arg0, arg1, arg2, arg3, arg4, arg5)
}

// ListAzureADGroups mocks base method.
func (m *MockAzureClient) ListAzureADGroups(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 []string) <-chan azure.GroupResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADGroups", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(<-chan azure.GroupResult)
	return ret0
}

// ListAzureADGroups indicates an expected call of ListAzureADGroups.
func (mr *MockAzureClientMockRecorder) ListAzureADGroups(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADGroups", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADGroups), arg0, arg1, arg2, arg3, arg4, arg5)
}

// ListAzureADRoleAssignments mocks base method.
func (m *MockAzureClient) ListAzureADRoleAssignments(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 []string) <-chan azure.UnifiedRoleAssignmentResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADRoleAssignments", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(<-chan azure.UnifiedRoleAssignmentResult)
	return ret0
}

// ListAzureADRoleAssignments indicates an expected call of ListAzureADRoleAssignments.
func (mr *MockAzureClientMockRecorder) ListAzureADRoleAssignments(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADRoleAssignments", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADRoleAssignments), arg0, arg1, arg2, arg3, arg4, arg5)
}

// ListAzureADRoles mocks base method.
func (m *MockAzureClient) ListAzureADRoles(arg0 context.Context, arg1, arg2 string) <-chan azure.RoleResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADRoles", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan azure.RoleResult)
	return ret0
}

// ListAzureADRoles indicates an expected call of ListAzureADRoles.
func (mr *MockAzureClientMockRecorder) ListAzureADRoles(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADRoles", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADRoles), arg0, arg1, arg2)
}

// ListAzureADServicePrincipalOwners mocks base method.
func (m *MockAzureClient) ListAzureADServicePrincipalOwners(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 []string) <-chan azure.ServicePrincipalOwnerResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADServicePrincipalOwners", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(<-chan azure.ServicePrincipalOwnerResult)
	return ret0
}

// ListAzureADServicePrincipalOwners indicates an expected call of ListAzureADServicePrincipalOwners.
func (mr *MockAzureClientMockRecorder) ListAzureADServicePrincipalOwners(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADServicePrincipalOwners", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADServicePrincipalOwners), arg0, arg1, arg2, arg3, arg4, arg5)
}

// ListAzureADServicePrincipals mocks base method.
func (m *MockAzureClient) ListAzureADServicePrincipals(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 []string) <-chan azure.ServicePrincipalResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADServicePrincipals", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(<-chan azure.ServicePrincipalResult)
	return ret0
}

// ListAzureADServicePrincipals indicates an expected call of ListAzureADServicePrincipals.
func (mr *MockAzureClientMockRecorder) ListAzureADServicePrincipals(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADServicePrincipals", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADServicePrincipals), arg0, arg1, arg2, arg3, arg4, arg5)
}

// ListAzureADTenants mocks base method.
func (m *MockAzureClient) ListAzureADTenants(arg0 context.Context, arg1 bool) <-chan azure.TenantResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADTenants", arg0, arg1)
	ret0, _ := ret[0].(<-chan azure.TenantResult)
	return ret0
}

// ListAzureADTenants indicates an expected call of ListAzureADTenants.
func (mr *MockAzureClientMockRecorder) ListAzureADTenants(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADTenants", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADTenants), arg0, arg1)
}

// ListAzureADUsers mocks base method.
func (m *MockAzureClient) ListAzureADUsers(arg0 context.Context, arg1, arg2, arg3 string, arg4 []string) <-chan azure.UserResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureADUsers", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(<-chan azure.UserResult)
	return ret0
}

// ListAzureADUsers indicates an expected call of ListAzureADUsers.
func (mr *MockAzureClientMockRecorder) ListAzureADUsers(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureADUsers", reflect.TypeOf((*MockAzureClient)(nil).ListAzureADUsers), arg0, arg1, arg2, arg3, arg4)
}

// ListAzureAutomationAccounts mocks base method.
func (m *MockAzureClient) ListAzureAutomationAccounts(arg0 context.Context, arg1 string, arg2 bool) <-chan azure.AutomationAccountResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureAutomationAccounts", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan azure.AutomationAccountResult)
	return ret0
}

// ListAzureAutomationAccounts indicates an expected call of ListAzureAutomationAccounts.
func (mr *MockAzureClientMockRecorder) ListAzureAutomationAccounts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureAutomationAccounts", reflect.TypeOf((*MockAzureClient)(nil).ListAzureAutomationAccounts), arg0, arg1, arg2)
}

// ListAzureDeviceRegisteredOwners mocks base method.
func (m *MockAzureClient) ListAzureDeviceRegisteredOwners(arg0 context.Context, arg1 string, arg2 bool) <-chan azure.DeviceRegisteredOwnerResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureDeviceRegisteredOwners", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan azure.DeviceRegisteredOwnerResult)
	return ret0
}

// ListAzureDeviceRegisteredOwners indicates an expected call of ListAzureDeviceRegisteredOwners.
func (mr *MockAzureClientMockRecorder) ListAzureDeviceRegisteredOwners(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureDeviceRegisteredOwners", reflect.TypeOf((*MockAzureClient)(nil).ListAzureDeviceRegisteredOwners), arg0, arg1, arg2)
}

// ListAzureDevices mocks base method.
func (m *MockAzureClient) ListAzureDevices(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 []string) <-chan azure.DeviceResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureDevices", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(<-chan azure.DeviceResult)
	return ret0
}

// ListAzureDevices indicates an expected call of ListAzureDevices.
func (mr *MockAzureClientMockRecorder) ListAzureDevices(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureDevices", reflect.TypeOf((*MockAzureClient)(nil).ListAzureDevices), arg0, arg1, arg2, arg3, arg4, arg5)
}

// ListAzureFunctionApps mocks base method.
func (m *MockAzureClient) ListAzureFunctionApps(arg0 context.Context, arg1 string, arg2 bool) <-chan azure.FunctionAppResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureFunctionApps", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan azure.FunctionAppResult)
	return ret0
}

// ListAzureFunctionApps indicates an expected call of ListAzureFunctionApps.
func (mr *MockAzureClientMockRecorder) ListAzureFunctionApps(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureFunctionApps", reflect.TypeOf((*MockAzureClient)(nil).ListAzureFunctionApps), arg0, arg1, arg2)
}

// ListAzureKeyVaults mocks base method.
func (m *MockAzureClient) ListAzureKeyVaults(arg0 context.Context, arg1 string, arg2 int32) <-chan azure.KeyVaultResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureKeyVaults", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan azure.KeyVaultResult)
	return ret0
}

// ListAzureKeyVaults indicates an expected call of ListAzureKeyVaults.
func (mr *MockAzureClientMockRecorder) ListAzureKeyVaults(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureKeyVaults", reflect.TypeOf((*MockAzureClient)(nil).ListAzureKeyVaults), arg0, arg1, arg2)
}

// ListAzureManagementGroupDescendants mocks base method.
func (m *MockAzureClient) ListAzureManagementGroupDescendants(arg0 context.Context, arg1 string) <-chan azure.DescendantInfoResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureManagementGroupDescendants", arg0, arg1)
	ret0, _ := ret[0].(<-chan azure.DescendantInfoResult)
	return ret0
}

// ListAzureManagementGroupDescendants indicates an expected call of ListAzureManagementGroupDescendants.
func (mr *MockAzureClientMockRecorder) ListAzureManagementGroupDescendants(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureManagementGroupDescendants", reflect.TypeOf((*MockAzureClient)(nil).ListAzureManagementGroupDescendants), arg0, arg1)
}

// ListAzureManagementGroups mocks base method.
func (m *MockAzureClient) ListAzureManagementGroups(arg0 context.Context) <-chan azure.ManagementGroupResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureManagementGroups", arg0)
	ret0, _ := ret[0].(<-chan azure.ManagementGroupResult)
	return ret0
}

// ListAzureManagementGroups indicates an expected call of ListAzureManagementGroups.
func (mr *MockAzureClientMockRecorder) ListAzureManagementGroups(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureManagementGroups", reflect.TypeOf((*MockAzureClient)(nil).ListAzureManagementGroups), arg0)
}

// ListAzureResourceGroups mocks base method.
func (m *MockAzureClient) ListAzureResourceGroups(arg0 context.Context, arg1, arg2 string) <-chan azure.ResourceGroupResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureResourceGroups", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan azure.ResourceGroupResult)
	return ret0
}

// ListAzureResourceGroups indicates an expected call of ListAzureResourceGroups.
func (mr *MockAzureClientMockRecorder) ListAzureResourceGroups(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureResourceGroups", reflect.TypeOf((*MockAzureClient)(nil).ListAzureResourceGroups), arg0, arg1, arg2)
}

// ListAzureStorageAccounts mocks base method.
func (m *MockAzureClient) ListAzureStorageAccounts(arg0 context.Context, arg1 string, arg2 bool) <-chan azure.StorageAccountResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureStorageAccounts", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan azure.StorageAccountResult)
	return ret0
}

// ListAzureStorageAccounts indicates an expected call of ListAzureStorageAccounts.
func (mr *MockAzureClientMockRecorder) ListAzureStorageAccounts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureStorageAccounts", reflect.TypeOf((*MockAzureClient)(nil).ListAzureStorageAccounts), arg0, arg1, arg2)
}

// ListAzureStorageContainers mocks base method.
func (m *MockAzureClient) ListAzureStorageContainers(arg0 context.Context, arg1, arg2, arg3 string, arg4 bool) <-chan azure.StorageContainerResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureStorageContainers", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(<-chan azure.StorageContainerResult)
	return ret0
}

// ListAzureStorageContainers indicates an expected call of ListAzureStorageContainers.
func (mr *MockAzureClientMockRecorder) ListAzureStorageContainers(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureStorageContainers", reflect.TypeOf((*MockAzureClient)(nil).ListAzureStorageContainers), arg0, arg1, arg2, arg3, arg4)
}

// ListAzureSubscriptions mocks base method.
func (m *MockAzureClient) ListAzureSubscriptions(arg0 context.Context) <-chan azure.SubscriptionResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureSubscriptions", arg0)
	ret0, _ := ret[0].(<-chan azure.SubscriptionResult)
	return ret0
}

// ListAzureSubscriptions indicates an expected call of ListAzureSubscriptions.
func (mr *MockAzureClientMockRecorder) ListAzureSubscriptions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureSubscriptions", reflect.TypeOf((*MockAzureClient)(nil).ListAzureSubscriptions), arg0)
}

// ListAzureVirtualMachines mocks base method.
func (m *MockAzureClient) ListAzureVirtualMachines(arg0 context.Context, arg1 string, arg2 bool) <-chan azure.VirtualMachineResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureVirtualMachines", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan azure.VirtualMachineResult)
	return ret0
}

// ListAzureVirtualMachines indicates an expected call of ListAzureVirtualMachines.
func (mr *MockAzureClientMockRecorder) ListAzureVirtualMachines(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureVirtualMachines", reflect.TypeOf((*MockAzureClient)(nil).ListAzureVirtualMachines), arg0, arg1, arg2)
}

// ListAzureWorkflows mocks base method.
func (m *MockAzureClient) ListAzureWorkflows(arg0 context.Context, arg1 string, arg2 bool) <-chan azure.WorkflowResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAzureWorkflows", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan azure.WorkflowResult)
	return ret0
}

// ListAzureWorkflows indicates an expected call of ListAzureWorkflows.
func (mr *MockAzureClientMockRecorder) ListAzureWorkflows(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAzureWorkflows", reflect.TypeOf((*MockAzureClient)(nil).ListAzureWorkflows), arg0, arg1, arg2)
}

// ListResourceRoleAssignments mocks base method.
func (m *MockAzureClient) ListResourceRoleAssignments(arg0 context.Context, arg1, arg2, arg3 string) <-chan azure.RoleAssignmentResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResourceRoleAssignments", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(<-chan azure.RoleAssignmentResult)
	return ret0
}

// ListResourceRoleAssignments indicates an expected call of ListResourceRoleAssignments.
func (mr *MockAzureClientMockRecorder) ListResourceRoleAssignments(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceRoleAssignments", reflect.TypeOf((*MockAzureClient)(nil).ListResourceRoleAssignments), arg0, arg1, arg2, arg3)
}

// ListRoleAssignmentsForResource mocks base method.
func (m *MockAzureClient) ListRoleAssignmentsForResource(arg0 context.Context, arg1, arg2 string) <-chan azure.RoleAssignmentResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoleAssignmentsForResource", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan azure.RoleAssignmentResult)
	return ret0
}

// ListRoleAssignmentsForResource indicates an expected call of ListRoleAssignmentsForResource.
func (mr *MockAzureClientMockRecorder) ListRoleAssignmentsForResource(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoleAssignmentsForResource", reflect.TypeOf((*MockAzureClient)(nil).ListRoleAssignmentsForResource), arg0, arg1, arg2)
}

// TenantInfo mocks base method.
func (m *MockAzureClient) TenantInfo() azure.Tenant {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantInfo")
	ret0, _ := ret[0].(azure.Tenant)
	return ret0
}

// TenantInfo indicates an expected call of TenantInfo.
func (mr *MockAzureClientMockRecorder) TenantInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantInfo", reflect.TypeOf((*MockAzureClient)(nil).TenantInfo))
}
