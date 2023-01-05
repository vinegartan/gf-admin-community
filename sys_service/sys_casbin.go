// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package sys_service

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	ICasbin interface {
		InstallHook(userType sys_enum.UserType, hookFunc sys_model.CasbinHookFunc) int64
		UnInstallHook(savedHookId int64)
		CleanAllHook()
		Check() error
		Enforcer() *casbin.Enforcer
		Middleware(r *ghttp.Request)
		AddRoleForUserInDomain(userName string, roleName string, domain string) (bool, error)
		DeleteRoleForUserInDomain(userName, roleName string, domain string) (bool, error)
		DeleteRolesForUser(userName string, domain string) (bool, error)
		AddPermissionForUser(roleName, path, method string) (bool, error)
		AddPermissionsForUser(roleName string, path []string) (bool, error)
		DeletePermissionForUser(roleName, path, method string) (bool, error)
		DeletePermissionsForUser(roleName string) (bool, error)
		EnforceCheck(userName, path, role, method interface{}) (bool, error)
	}
)

var (
	localCasbin ICasbin
)

func Casbin() ICasbin {
	if localCasbin == nil {
		panic("implement not found for interface ICasbin, forgot register?")
	}
	return localCasbin
}

func RegisterCasbin(i ICasbin) {
	localCasbin = i
}
