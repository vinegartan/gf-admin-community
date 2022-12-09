// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure for table sys_user.
type SysUser struct {
	Id        int64       `json:"id"        description:""`
	Username  string      `json:"username"  description:"账号"`
	Password  string      `json:"password"  description:"密码"`
	State     int         `json:"state"     description:"状态：0未激活、1正常、-1封号、-2异常、-3已注销"`
	Type      int         `json:"type"      description:"用户类型，0匿名，1用户，2微商，4商户、8广告主、16服务商、32运营中心"`
	Mobile    string      `json:"mobile"    description:"手机号"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" description:""`
}
